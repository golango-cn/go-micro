package handler

import (
	"context"
	"go-micro/go-micro-part03/proto/common"
	"go-micro/go-micro-part03/svr/model"
	go_svr_proto_proto "go-micro/go-micro-part03/proto"
)

type EmployeeHander struct {
}

func (e *EmployeeHander) Login(ctx context.Context, requ *go_svr_proto_proto.LoginRequest, resp *go_svr_proto_proto.LoginResponse) error {
	resp.BaseResponse = &common.BaseResponse{
		IsSuccess:    true,
		ErrorMessage: "测试成功",
	}
	employee, err := model.Login(requ.FirstName)

	if err != nil {
		return err
	}
	resp.Employee = employee
	return nil
}

func (e *EmployeeHander) Logout(context.Context, *go_svr_proto_proto.LogoutRequest, *go_svr_proto_proto.LogoutResponse) error {
	return nil
}

func (e *EmployeeHander) GetEmployees(ctx context.Context, requ *go_svr_proto_proto.GetEmployeeRequest, resp *go_svr_proto_proto.GetEmployeeResponse) error {

	resp.BaseResponse = &common.BaseResponse{
		IsSuccess:    true,
		ErrorMessage: "测试成功",
	}

	employees, err := model.GetEmployees()
	if err != nil {
		return err
	}
	resp.Employees = employees
	return nil
}
