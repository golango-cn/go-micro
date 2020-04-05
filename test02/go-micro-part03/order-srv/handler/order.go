package handler

import (
	"context"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/order-srv/model"

	proto "go-micro/go-micro-part03/proto/order"
)

type OrderHandler struct{}


// New 新增订单
func (e *OrderHandler) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {



	ordersService := model.GetOrderService()

	orderId, err := ordersService.Create(req.BookId, req.UserId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Order = &proto.Order{
		Id: orderId,
	}
	return
}

// GetOrder 获取订单
func (e *OrderHandler) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {

	ordersService := model.GetOrderService()

	log.Info(fmt.Sprintf("[GetOrder] 收到获取订单请求，%d", req.OrderId))

	rsp.Order, err = ordersService.GetOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}

func Init()  {
	model.Init()
}

