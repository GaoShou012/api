package client_event_redis_stream

import (
	"api/cs"
	"encoding/json"
	"fmt"
)

type plugin struct {
	opts *Options
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) Push(uuid string, event *cs.Event) (string, error) {
	j, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	topic := fmt.Sprintf("client:event:stream:%s", uuid)
	return p.opts.Stream.Push(topic, j)
}
