package queue_redis

import (
	"cs/class/queue"
	"cs/meta"
)

var _ queue.Queue = &plugin{}

type plugin struct {
	queueLength uint64
}

func (p *plugin) Init() error {
	return nil
}

func (p *plugin) SetLenMax(name string, num uint64) error {
	p.queueLength = num
	return nil
}
func (p *plugin) GetLenMax(name string) (uint64, error) {
	panic("implement me")
}

/*
	获取队列头位会话信息
*/
func (p *plugin) GetTheFirstSession(name string) (*queue.Item, error) {
	panic("implement me")
}

func (p *plugin) Join(name string, session meta.Session) error {
	panic("implement me")
}

func (p *plugin) Leave(name string, session meta.Session) error {
	panic("implement me")
}

func (p *plugin) Len(name string) (uint64, error) {
	panic("implement me")
}

func (p *plugin) GetOffset(name string, session meta.Session) (uint64, error) {
	panic("implement me")
}
