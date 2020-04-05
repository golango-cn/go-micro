package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/payment-srv/model"
	"go-micro/go-micro-part03/proto/payment"
)

type PaymentHander struct{}

func (p PaymentHander) PayOrder(ctx context.Context, requ *payment.Request, resp *payment.Response) error {

	paymentService := model.GetPaymentService()

	log.Infof("[PayOrder] 收到支付请求")
	err := paymentService.PayOrder(requ.OrderId)
	if err != nil {
		resp.Success = false
		resp.Error = &payment.Error{
			Detail: err.Error(),
		}
		return nil
	}

	resp.Success = true
	return nil
}
