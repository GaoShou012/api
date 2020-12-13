package stream_redis_stream

import (
	"context"
	"fmt"
	"framework/class/logger"
	"framework/class/stream"
	"github.com/go-redis/redis/v8"
	"sync"
)

var _ stream.Stream = &redisStream{}

type redisStream struct {
	redisClient *redis.Client
	opts        *Options
}

func (s *redisStream) Init() error {
	s.redisClient = s.opts.redisClient
	if s.opts.logger == nil {
		s.opts.logger = &iLogger{}
	}
	return nil
}

func (s *redisStream) Connect(dns string) error {
	client, err := connect(dns)
	if err != nil {
		return err
	}
	s.redisClient = client
	return nil
}

func (s *redisStream) Push(topic string, message []byte) (string, error) {
	values := make(map[string]interface{})
	values["Payload"] = message
	xAddArgs := &redis.XAddArgs{
		Stream:       topic,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "*",
		Values:       values,
	}
	msgId, err := s.redisClient.XAdd(context.TODO(), xAddArgs).Result()
	return msgId, err
}

func (s *redisStream) Pull(topic string, lastMessageId string, count uint64) ([]stream.Event, error) {
	readArgs := &redis.XReadArgs{
		Streams: []string{topic, lastMessageId},
		Count:   int64(count),
		Block:   -1,
	}
	res, err := s.redisClient.XRead(context.TODO(), readArgs).Result()
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
				s.redisClient.XDel(context.TODO(), topic, message.ID)
				s.opts.logger.Log(logger.ErrorLevel, "Assert payload is failed")
				continue
			}

			evt := &redisEvent{
				id:          message.ID,
				streamName:  topic,
				message:     []byte(payload),
				redisClient: s.redisClient,
			}
			events = append(events, evt)
		}
	}

	return events, nil
}

func (s *redisStream) Subscribe(topic string, handler stream.Handler) (stream.Subscriber, error) {
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
				res, err := s.redisClient.XRead(context.TODO(), xReadArgs).Result()
				if err != nil {
					s.opts.logger.Log(logger.ErrorLevel, err)
					continue
				}
				for _, theStream := range res {
					for _, message := range theStream.Messages {
						evt := &redisEvent{
							id:          message.ID,
							streamName:  theStream.Stream,
							message:     nil,
							redisClient: s.redisClient,
							err:         nil,
						}

						payload, ok := message.Values["payload"].(string)
						if !ok {
							evt.err = fmt.Errorf("Assert Type field is failed\n")
							if err := handler(evt); err != nil {
								s.opts.logger.Log(logger.ErrorLevel, err)
							}
							continue
						}
						evt.message = []byte(payload)

						if err := handler(evt); err != nil {
							s.opts.logger.Log(logger.ErrorLevel, err)
						}
					}
				}
			}
		}
	}()

	return sub, nil
}
