package database

import (
	"container/list"
	"fmt"
	"go-micro/broker/broadcast"
	"sync"
	"time"
)

type Broker struct {
	sync.RWMutex
	watchers list.List
}

func New() (broadcast.BroadCast, error) {

	var b Broker
	go b.scan()

	return &b, nil
}

func (b *Broker) scan() {

	for {

		// 生产数据
		unix := time.Now().Unix()
		fmt.Println("模拟扫描数据库，生产数据")

		for el := b.watchers.Front(); el != nil; el = el.Next() {
			select {
			case el.Value.(*Watcher).updated <- unix:
			default:

			}
		}

		time.Sleep(time.Second * 2)
	}

}

func (b Broker) Send(i interface{}) error {
	panic("implement me")
}

func (b *Broker) Watch() broadcast.Watcher {

	w := &Watcher{
		exit:    make(chan interface{}),
		updated: make(chan interface{}),
	}

	b.Lock()
	el := b.watchers.PushBack(w)
	b.Unlock()

	go func() {

		<-w.exit
		b.Lock()
		b.watchers.Remove(el)
		b.Unlock()

		fmt.Println("从队列中移除", el)
	}()

	return w

}
