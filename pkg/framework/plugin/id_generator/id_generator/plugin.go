package id_generator

import (
	"framework/class/id_generator"
	"sync"
)

var _ id_generator.IdGenerator = &plugin{}

type plugin struct {
	mutex     sync.RWMutex
	cache     map[string]int64
	cacheHash map[string]map[string]int64
	opts      *Options
}

func (p *plugin) Init() error {
	p.cache = make(map[string]int64)
	p.cacheHash = make(map[string]map[string]int64)
	return nil
}

func (p *plugin) Incr(key string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	val := p.cache[key]
	val++
	p.cache[key] = val
	return val, nil
}

func (p *plugin) Get(key string) (int64, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.cache[key], nil
}

func (p *plugin) HIncr(key string, field string) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// 当hash结构不存在的时候，创建hash结构
	hash, ok := p.cacheHash[key]
	if !ok {
		hash = make(map[string]int64)
		p.cacheHash[key] = hash
	}

	val := hash[field]
	val++
	hash[field] = val

	return val, nil
}

func (p *plugin) HGet(key string, field string) (int64, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	hash, ok := p.cacheHash[key]
	if !ok {
		hash = make(map[string]int64)
	}
	return hash[field], nil
}

func (p *plugin) HGetAll(key string) (map[string]int64, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	hash, ok := p.cacheHash[key]
	if !ok {
		hash = make(map[string]int64)
	}
	return hash, nil
}
