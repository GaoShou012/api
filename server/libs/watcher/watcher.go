package watcher

import (
	"container/list"
	"sync"
	"time"
)

var mutex sync.Mutex
var items list.List
var newItems []Item

type Item interface {
	IsUpLimit() bool
	UpLimitCallback()
}

func Add(item Item) {
	mutex.Lock()
	defer mutex.Unlock()
	newItems = append(newItems, item)
}

func init() {
	go func() {
		for {
			mutex.Lock()
			tmp := newItems
			newItems = make([]Item, 0)
			mutex.Unlock()
			for _, item := range tmp {
				items.PushBack(item)
			}

			var eleToDel []*list.Element
			for ele := items.Front(); ele != nil; ele = ele.Next() {
				item := ele.Value.(Item)
				if item.IsUpLimit() {
					item.UpLimitCallback()
					eleToDel = append(eleToDel, ele)
				}
			}
			for _, ele := range eleToDel {
				items.Remove(ele)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
}
