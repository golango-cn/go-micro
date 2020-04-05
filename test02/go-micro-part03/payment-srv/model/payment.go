package model

import (
	"context"
	"fmt"

	//"fmt"
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/proto/payment"
	"time"

	micro "github.com/micro/go-micro/v2"
	invS "go-micro/go-micro-part03/proto/inventory"
	ordS "go-micro/go-micro-part03/proto/order"
	"sync"
)

var (
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher

	paymentService *PaymentService
)

func Init() {
	invClient = invS.NewInventoryService("mu.micro.book.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher("mu.micro.book.topic.payment.done", client.DefaultClient)
	paymentService = &PaymentService{}
}

func GetPaymentService() *PaymentService {
	return paymentService
}

type PaymentService struct {
}

func (s *PaymentService) PayOrder(orderId int64) (err error) {

	//获取支付单
	orderRsp, err := ordSClient.GetOrder(context.TODO(), &ordS.Request{
		OrderId: orderId,
	})
	if err != nil {
		log.Errorf("[PayOrder] 查询 订单信息失败，orderId：%d, %s", orderId, err)
		return
	}

	// 订单不存在
	if orderRsp == nil || !orderRsp.Success || orderRsp.Order == nil {
		err = fmt.Errorf("[PayOrder] 支付单不存在")
		log.Errorf("[PayOrder] 查询 订单信息失败，orderId：%d, %s", orderId, err)
		return
	}

	// 确认出库
	invRsp, err := invClient.Confirm(context.TODO(), &invS.Request{
		HistoryId: orderRsp.Order.Id,
	})
	if err != nil || invRsp == nil || !invRsp.Success {
		err = fmt.Errorf("[PayOrder] 确认出库失败，%s", err)
		log.Errorf("%s", err)
		return
	}

	// 广播支付成功
	s.sendPayDoneEvt(orderId, 1)
	return

}

// sendPayDoneEvt 发送支付事件
func (s *PaymentService) sendPayDoneEvt(orderId int64, state int32) {

	// 构建事件
	ev := &payment.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Infof("[sendPayDoneEvt] 发送支付事件，%+v\n", ev)

	// 广播
	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Errorf("[sendPayDoneEvt] 异常: %v", err)
	}
}
