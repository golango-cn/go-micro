package main

import (
	"fmt"
	"go-micro/broker/broadcast"
	"go-micro/broker/broadcast/queue"
	"time"
)

func main() {

	broker, err := queue.New()
	fmt.Println(broker, err)

	broadcast.Init(broker)

	watcher := broadcast.GetBroadCast().Watch()

	index := 0

	for {
		value, err := watcher.Next()

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("获取数据", value)
		if index == 5 {
			watcher.Stop()
		}

		index++
	}

	time.Sleep(3 * time.Second)

}
