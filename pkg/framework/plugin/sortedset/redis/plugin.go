package redis

import (
	"framework/class/sortedset"
	"github.com/go-redis/redis"
	"im/class/channel"
)

var _ sortedset.Sortedset = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Exists(topic string) (bool, error) {
	num,err := p.opts.redisClient.ZCard(topic).Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	return true, nil
}

func (p *plugin) Len(topic string) int64 {
	num,err := p.opts.redisClient.ZCard(topic).Result()
	if err != nil {
		panic(err)
	}
	return num
}

func (p *plugin) Find(topic string, page int64, pageSize int64) ([]sortedset.Item, error) {
	members,err := p.opts.redisClient.ZRange(topic,page,page*pageSize).Result()
	if err != nil {
		return nil,err
	}


	i := 0
	items := make([]sortedset.Item, len(members))
	for v, k := range members {
		evt := &item{
			iKey:k,
			iVal:v,
		}
		items[i] = evt
		i++
	}
	return items,nil
}

func (p *plugin) SetItem(topic string, key string, val float64) error {
	_,err := p.opts.redisClient.ZAdd(topic,redis.Z{
		Score:  val,
		Member: key,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) GetOffset(topic string, key string) (int64, error) {
	panic("implement me")
}

func (p *plugin) GetOffsetN(topic string, key string) (int64, error) {
	panic("implement me")
}

func (p *plugin) GetItemFormPositive(topic string) (sortedset.Item, error) {
	panic("implement me")
}

func (p *plugin) GetItemFromNegative(topic string) (sortedset.Item, error) {
	panic("implement me")
}

func (p *plugin) ExistsItem(topic string, key string) (bool, error) {
	panic("implement me")
}

func (p *plugin) Init() error {
	return nil
}
