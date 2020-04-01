package handler

import (
	"context"
	"go-micro/go-micro-part1-test01/svr/model"
)
import "go-micro/go-micro-part1-test01/svr/proto"
import "go-micro/go-micro-part1-test01/svr/proto/common"

type EmployeeHander struct {

}


func (e *EmployeeHander)GetEmployees(ctx context.Context, requ *go_svr_proto_proto.GetEmployeeRequest, resp *go_svr_proto_proto.GetEmployeeResponse) error {

	resp.BaseResponse = &common.BaseResponse{
		IsSuccess:            true,
		ErrorMessage:         "测试成功",
	}

	employees, err := model.GetEmployees()
	if err != nil {
		resp.BaseResponse.IsSuccess = false
		resp.BaseResponse.ErrorMessage = err.Error()
	} else {
		resp.Employees = employees
	}
	return nil
}