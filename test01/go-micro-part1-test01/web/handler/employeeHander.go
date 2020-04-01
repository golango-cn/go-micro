package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"

	us "go-micro/go-micro-part1-test01/web/proto"
)

var employeeService us.EmployeeService

func Init() {
	employeeService = us.NewEmployeeService("mu.micro.book.srv.employee", client.DefaultClient)
}

// GetEmployees 获取用户信息
func GetEmployees(w http.ResponseWriter, r *http.Request) {

	log.Info("request start")

	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()


	// 调用后台服务
	resp, err := employeeService.GetEmployees(context.TODO(), &us.GetEmployeeRequest{

	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
