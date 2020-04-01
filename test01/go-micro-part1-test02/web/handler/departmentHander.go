package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"

	us "go-micro/go-micro-part1-test02/web/proto"
)

var departmentService us.DepartmentService

func Init() {
	departmentService = us.NewDepartmentService("mu.micro.book.srv.department", client.DefaultClient)
}

// GetDepartments 获取部门信息
func GetDepartments(w http.ResponseWriter, r *http.Request) {

	log.Info("request start")

	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()


	// 调用后台服务
	resp, err := departmentService.GetDeparments(context.TODO(), &us.GetDeparmentRequest{

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
