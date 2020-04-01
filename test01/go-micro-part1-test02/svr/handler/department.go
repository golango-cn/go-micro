package handler

import (
	"context"
	"go-micro/go-micro-part1-test02/svr/model"
	go_svr_proto_proto "go-micro/go-micro-part1-test02/svr/proto"
	"go-micro/go-micro-part1-test02/svr/proto/common"
)

type DepartmentHander struct {

}

func (e *DepartmentHander) GetDeparments(ctx context.Context, requ *go_svr_proto_proto.GetDeparmentRequest, resp *go_svr_proto_proto.GetDeparmentResponse) error {

	resp.BaseResponse = &common.BaseResponse{
		IsSuccess:            true,
		ErrorMessage:         "测试成功",
	}

	departments, err := model.GetDepartments()
	if err != nil {
		resp.BaseResponse.IsSuccess = false
		resp.BaseResponse.ErrorMessage = err.Error()
	} else {
		resp.Departments = departments
	}
	return nil
}