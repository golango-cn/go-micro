package model

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	invS "go-micro/go-micro-part03/proto/inventory"
	proto "go-micro/go-micro-part03/proto/order"
	"math/rand"
	"time"
)


type OrderService struct {
	keys map[int64]*proto.Order
}

// 获取订单
func (o *OrderService) GetOrder(orderId int64) (*proto.Order, error) {
	// 模拟查询数据库
	if order, ok := o.keys[orderId]; ok {
		logger.Info("查询订单成功", orderId, order)
		return order, nil
	}
	return nil, fmt.Errorf("没有找到订单信息：%d", orderId)
}

// 创建订单
func (o *OrderService) Create(userId, bookId int64) (int64, error) {



	orderId := rand.Int63n(1000)
	order := &proto.Order{
		Id:           orderId,
		UserId:       userId,
		BookId:       bookId,
		InvHistoryId: 123,
		State:        0,
		CreatedTime:  time.Now().Unix(),
	}
	o.keys[orderId] = order

	// 请求销存
	 invClient.Sell(context.TODO(), &invS.Request{
		BookId: orderId,
	})

	logger.Info("订单创建成功", orderId, order)
	return orderId, nil

}

// 更新订单
func (o *OrderService) Update(orderId int64) error {
	if _, ok := o.keys[orderId]; ok {
		logger.Infof("订单更新成功：%d", orderId)
		return nil
	}

	return fmt.Errorf("没有找到订单信息：%d", orderId)
}

var orderService *OrderService
var invClient invS.InventoryService

func Init() {
	orderService = &OrderService{keys: map[int64]*proto.Order{}}
	invClient = invS.NewInventoryService("mu.micro.book.srv.inventory", client.DefaultClient)
}

func GetOrderService() *OrderService {
	return orderService
}
