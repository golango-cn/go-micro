package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/plugins/session"
	invS "go-micro/go-micro-part03/proto/inventory"
	"net/http"
	"strconv"
	"time"

	"github.com/micro/go-micro/v2/client"
	auth "go-micro/go-micro-part03/proto"
	orders "go-micro/go-micro-part03/proto/order"
)


var (
	orderService orders.OrdersService
	authClient    auth.TokenService
	invClient     invS.InventoryService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	orderService = orders.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	authClient = auth.NewTokenService("mu.micro.book.srv.auth", client.DefaultClient)
}

// New 新增订单入口
func New(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	bookId, _ := strconv.ParseInt(r.Form.Get("bookId"), 10, 10)

	// 返回结果
	response := map[string]interface{}{}

	userId := int64(session.GetSession(r).Values["userId"].(float64))

	// 调用后台服务
	rsp, err := orderService.New(context.TODO(), &orders.Request{
		BookId: bookId,
		UserId: userId,
		//UserId:10003,
	})

	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
		response["orderId"] = rsp.Order.Id
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
