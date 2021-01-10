package redis_sortdset

import (
	"errors"
	"framework/class/sortedset"
	"github.com/go-redis/redis"
)

var _ sortedset.Sortedset = &plugin{}

type plugin struct {
	opts *Options
}

func (p *plugin) Exists(topic string) (bool, error) {
	num, err := p.opts.redisClient.ZCard(topic).Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	return true, nil
}

func (p *plugin) Len(topic string) int64 {
	num, err := p.opts.redisClient.ZCard(topic).Result()
	if err != redis.Nil {
		panic(err)
	}
	return num
}

func (p *plugin) Find(topic string, page int64, pageSize int64) ([]sortedset.Item, error) {
	members, err := p.opts.redisClient.ZRange(topic, page, pageSize).Result()
	if err != nil {
		return nil, err
	}

	i := 0
	items := make([]sortedset.Item, len(members))
	for v, k := range members {
		evt := &item{
			iKey: k,
			iVal: v,
		}
		items[i] = evt
		i++
	}
	return items, nil
}

func (p *plugin) SetItem(topic string, key string, val float64) error {
	_, err := p.opts.redisClient.ZAdd(topic, redis.Z{
		Score:  val,
		Member: key,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) GetOffset(topic string, key string) (int64, error) {
	all, err := p.opts.redisClient.ZLexCount(topic, "-", "+").Result()
	if err != nil {
		return 0, errors.New("没有查到" + topic)
	}
	if all == 0 {
		return 0, errors.New("没有查到" + topic)
	}
	index, err := p.opts.redisClient.ZRank(topic, key).Result()
	if err != nil {
		return 0, err
	}
	return index, nil
}

func (p *plugin) GetOffsetN(topic string, key string) (int64, error) {

	all, err := p.opts.redisClient.ZLexCount(topic, "-", "+").Result()
	if err != nil {
		return 0, errors.New("没有查到" + topic)
	}
	if all == 0 {
		return 0, errors.New("没有查到" + topic)
	}
	index, err := p.opts.redisClient.ZRank(topic, key).Result()
	if err != nil {
		return 0, err
	}
	return all - index - 1, nil
}

func (p *plugin) GetItemFormPositive(topic string) ([]sortedset.Item, error) {
	res, err := p.opts.redisClient.ZRange(topic, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	i := 0
	items := make([]sortedset.Item, len(res))
	for v, k := range res {
		evt := &item{
			iKey: k,
			iVal: v,
		}
		items[i] = evt
		i++
	}
	return items, nil
}

func (p *plugin) GetItemFromNegative(topic string) ([]sortedset.Item, error) {
	res, err := p.opts.redisClient.ZRevRange(topic, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	i := 0
	items := make([]sortedset.Item, len(res))
	for v, k := range res {
		evt := &item{
			iKey: k,
			iVal: v,
		}
		items[i] = evt
		i++
	}
	return items, nil
}

func (p *plugin) ExistsItem(topic string, key string) (bool, error) {
	num, err := p.opts.redisClient.ZScore(topic, key).Result()
	if err != redis.Nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	return true, nil
}

func (p *plugin) DelItem(topic string) (bool, error) {
	num, err := p.opts.redisClient.ZRemRangeByLex(topic, "-", "+").Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, errors.New("不存在" + topic)
	}
	return true, nil
}
func (p *plugin) DelKey(topic string, key string) (bool, error) {
	num, err := p.opts.redisClient.ZRem(topic, key).Result()
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, errors.New("不存在" + topic)
	}
	return true, nil
}
func (p *plugin) Init() error {
	return nil
}
