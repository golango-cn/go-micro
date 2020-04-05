package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/v2/client"
	auth "go-micro/go-micro-part03/proto"
	payment "go-micro/go-micro-part03/proto/payment"
)

var (
	payMentService payment.PaymentService
	authClient    auth.TokenService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	payMentService = payment.NewPaymentService("mu.micro.book.srv.payment", client.DefaultClient)
	authClient = auth.NewTokenService("mu.micro.book.srv.auth", client.DefaultClient)
}

// PayOrder 支付订单
func PayOrder(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Infof("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	orderId, _ := strconv.ParseInt(r.Form.Get("orderId"), 10, 64)

	log.Info("调用PayOrder接口", orderId)

	// 调用后台服务
	_, err := payMentService.PayOrder(context.TODO(), &payment.Request{
		OrderId: orderId,
	})

	// 返回结果
	response := map[string]interface{}{}

	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
