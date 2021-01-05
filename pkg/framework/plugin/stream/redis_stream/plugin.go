package stream_redis_stream

import (
	"context"
	"fmt"
	"framework/class/stream"
	"framework/env"
	"github.com/go-redis/redis"
	"sync"
)

var _ stream.Stream = &plugin{}

type plugin struct {
	opts *Options
}

func (s *plugin) Init() error {
	return nil
}

func (s *plugin) Connect(dns string) error {
	client, err := connect(dns)
	if err != nil {
		return err
	}
	s.opts.redisClient = client
	return nil
}

func (s *plugin) GetById(topic string, messageId string) (stream.Event, error) {
	res, err := s.opts.redisClient.XRange(topic, messageId, messageId).Result()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	event := res[0]
	message := event.Values["Payload"].(string)
	evt := &Event{
		id:          event.ID,
		streamName:  topic,
		message:     []byte(message),
		redisClient: s.opts.redisClient,
		err:         nil,
	}
	return evt, nil
}

func (s *plugin) Push(topic string, message []byte) (string, error) {
	values := make(map[string]interface{})
	values["Payload"] = message
	xAddArgs := &redis.XAddArgs{
		Stream:       topic,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "*",
		Values:       values,
	}
	msgId, err := s.opts.redisClient.XAdd(xAddArgs).Result()
	return msgId, err
}

func (s *plugin) Pull(topic string, lastMessageId string, count uint64) ([]stream.Event, error) {
	readArgs := &redis.XReadArgs{
		Streams: []string{topic, lastMessageId},
		Count:   int64(count),
		Block:   -1,
	}
	res, err := s.opts.redisClient.XRead(readArgs).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var events []stream.Event

	for _, row := range res {
		for _, message := range row.Messages {
			payload, ok := message.Values["Payload"].(string)
			if !ok {
				s.opts.redisClient.XDel(topic, message.ID)
				env.Logger.Error("Assert payload is failed")
				continue
			}

			evt := &Event{
				id:          message.ID,
				streamName:  topic,
				message:     []byte(payload),
				redisClient: s.opts.redisClient,
			}
			events = append(events, evt)
		}
	}

	return events, nil
}

func (s *plugin) Subscribe(topic string, handler stream.Handler) (stream.Subscriber, error) {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	sub := &subscriber{}
	sub.setupUnSubscribe(func() {
		cancel()
		wg.Wait()
	})

	go func() {
		wg.Add(1)
		defer wg.Done()

		xReadArgs := &redis.XReadArgs{
			Streams: []string{topic, "0"},
			Count:   1,
			Block:   0,
		}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				res, err := s.opts.redisClient.XRead(xReadArgs).Result()
				if err != nil {
					env.Logger.Error(err)
					continue
				}
				for _, theStream := range res {
					for _, message := range theStream.Messages {
						evt := &Event{
							id:          message.ID,
							streamName:  theStream.Stream,
							message:     nil,
							redisClient: s.opts.redisClient,
							err:         nil,
						}

						payload, ok := message.Values["payload"].(string)
						if !ok {
							evt.err = fmt.Errorf("Assert Type field is failed\n")
							if err := handler(evt); err != nil {
								env.Logger.Error(err)
							}
							continue
						}
						evt.message = []byte(payload)

						if err := handler(evt); err != nil {
							env.Logger.Error(err)
						}
					}
				}
			}
		}
	}()

	return sub, nil
}
