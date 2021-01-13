package stream_redis_stream_v8

import (
	"context"
	"fmt"
	"framework/class/logger"
	"framework/class/stream"
	"framework/env"
	"github.com/go-redis/redis/v8"
	"sync"
)

var _ stream.Stream = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Connect(dns string) error {
	client, err := connect(dns)
	if err != nil {
		return err
	}
	p.opts.redisClient = client
	return nil
}

func (p *plugin) Delete(topic string) error {
	_, err := p.opts.redisClient.Del(context.TODO(), topic).Result()
	return err
}

func (p *plugin) Push(topic string, message []byte) (string, error) {
	values := make(map[string]interface{})
	values["Payload"] = message
	xAddArgs := &redis.XAddArgs{
		Stream:       topic,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "*",
		Values:       values,
	}
	msgId, err := p.opts.redisClient.XAdd(context.TODO(), xAddArgs).Result()
	return msgId, err
}

func (p *plugin) Pull(topic string, lastMessageId string, count uint64) ([]stream.Event, error) {
	readArgs := &redis.XReadArgs{
		Streams: []string{topic, lastMessageId},
		Count:   int64(count),
		Block:   -1,
	}
	res, err := p.opts.redisClient.XRead(context.TODO(), readArgs).Result()
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
				p.opts.redisClient.XDel(context.TODO(), topic, message.ID)
				p.opts.logger.Log(logger.ErrorLevel, "Assert payload is failed")
				continue
			}

			evt := &Event{
				id:          message.ID,
				streamName:  topic,
				message:     []byte(payload),
				redisClient: p.opts.redisClient,
			}
			events = append(events, evt)
		}
	}

	return events, nil
}

func (p *plugin) RevPull(topic string, lastMessageId string, count uint64) ([]stream.Event, error) {
	res, err := p.opts.redisClient.XRevRangeN(context.TODO(), topic, lastMessageId, "-", int64(count)+1).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var events []stream.Event

	for _, message := range res {
		payload, ok := message.Values["Payload"].(string)
		if !ok {
			p.opts.redisClient.XDel(context.TODO(), topic, message.ID)
			env.Logger.Error("Assert payload is failed")
			continue
		}

		evt := &Event{
			id:          message.ID,
			streamName:  topic,
			message:     []byte(payload),
			redisClient: p.opts.redisClient,
		}
		events = append(events, evt)
	}

	// 如果第一条数据等于目标数据，去掉第一条数据
	// 如果第一条数据不等于目标数据，去掉最后一条数据，保证最大条目数相等
	if len(events) > 0 {
		if events[0].Id() == lastMessageId {
			events = events[1:]
		} else {
			if len(events) > int(count) {
				events = events[0 : len(events)-2]
			}
		}
	}

	return events, nil
}

func (p *plugin) PullById(topic string, messageId string) (stream.Event, error) {
	res, err := p.opts.redisClient.XRange(context.TODO(), topic, messageId, messageId).Result()
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
		redisClient: p.opts.redisClient,
		err:         nil,
	}
	return evt, nil
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
				res, err := s.opts.redisClient.XRead(context.TODO(), xReadArgs).Result()
				if err != nil {
					s.opts.logger.Log(logger.ErrorLevel, err)
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
