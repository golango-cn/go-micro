package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/order-srv/model"
	"go-micro/go-micro-part03/proto/payment"
)

var (
	ordersService *model.OrderService
)

// Init 初始化handler
func Init() {
	ordersService = model.GetOrderService()
}
//
// PayOrder 订单支付消息
func PayOrder(ctx context.Context, event *payment.PayEvent) (err error) {
	log.Infof("[PayOrder] 收到支付订单通知，%d，%d", event.OrderId, event.State)

	err = ordersService.Update(event.OrderId)
	if err != nil {
		log.Errorf("[PayOrder] 收到支付订单通知，更新状态异常，%s", err)
		return
	}
	return
}
