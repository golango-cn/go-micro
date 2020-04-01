// Code generated by protoc-gen-go. DO NOT EDIT.
// source: department.proto

package go_svr_proto_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "go-micro/go-micro-part1-test02/web/proto/common"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Department struct {
	DeptNo               string   `protobuf:"bytes,1,opt,name=deptNo,proto3" json:"deptNo,omitempty"`
	DeptName             string   `protobuf:"bytes,2,opt,name=deptName,proto3" json:"deptName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_63863e61582d2703, []int{0}
}

func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (m *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(m, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetDeptNo() string {
	if m != nil {
		return m.DeptNo
	}
	return ""
}

func (m *Department) GetDeptName() string {
	if m != nil {
		return m.DeptName
	}
	return ""
}

type GetDeparmentRequest struct {
	BaseRequest          *common.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetDeparmentRequest) Reset()         { *m = GetDeparmentRequest{} }
func (m *GetDeparmentRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeparmentRequest) ProtoMessage()    {}
func (*GetDeparmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_63863e61582d2703, []int{1}
}

func (m *GetDeparmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeparmentRequest.Unmarshal(m, b)
}
func (m *GetDeparmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeparmentRequest.Marshal(b, m, deterministic)
}
func (m *GetDeparmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeparmentRequest.Merge(m, src)
}
func (m *GetDeparmentRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeparmentRequest.Size(m)
}
func (m *GetDeparmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeparmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeparmentRequest proto.InternalMessageInfo

func (m *GetDeparmentRequest) GetBaseRequest() *common.BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

type GetDeparmentResponse struct {
	BaseResponse         *common.BaseResponse `protobuf:"bytes,1,opt,name=baseResponse,proto3" json:"baseResponse,omitempty"`
	Departments          []*Department        `protobuf:"bytes,2,rep,name=departments,proto3" json:"departments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDeparmentResponse) Reset()         { *m = GetDeparmentResponse{} }
func (m *GetDeparmentResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeparmentResponse) ProtoMessage()    {}
func (*GetDeparmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63863e61582d2703, []int{2}
}

func (m *GetDeparmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeparmentResponse.Unmarshal(m, b)
}
func (m *GetDeparmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeparmentResponse.Marshal(b, m, deterministic)
}
func (m *GetDeparmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeparmentResponse.Merge(m, src)
}
func (m *GetDeparmentResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeparmentResponse.Size(m)
}
func (m *GetDeparmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeparmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeparmentResponse proto.InternalMessageInfo

func (m *GetDeparmentResponse) GetBaseResponse() *common.BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *GetDeparmentResponse) GetDepartments() []*Department {
	if m != nil {
		return m.Departments
	}
	return nil
}

func init() {
	proto.RegisterType((*Department)(nil), "go.svr.proto.proto.department")
	proto.RegisterType((*GetDeparmentRequest)(nil), "go.svr.proto.proto.GetDeparmentRequest")
	proto.RegisterType((*GetDeparmentResponse)(nil), "go.svr.proto.proto.GetDeparmentResponse")
}

func init() {
	proto.RegisterFile("department.proto", fileDescriptor_63863e61582d2703)
}

var fileDescriptor_63863e61582d2703 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xdd, 0x0a, 0x45, 0x67, 0x15, 0x74, 0x5a, 0x64, 0xdd, 0x83, 0x94, 0x5c, 0xdc, 0x53,
	0x84, 0x15, 0xc1, 0x63, 0x0f, 0x05, 0x2f, 0xe2, 0x21, 0xfe, 0x82, 0x6d, 0x77, 0x10, 0x0f, 0xd9,
	0xc4, 0x24, 0x16, 0xfc, 0x1b, 0xfe, 0x62, 0x69, 0x12, 0xba, 0x09, 0x16, 0x7a, 0x09, 0x79, 0xbc,
	0x97, 0x2f, 0x6f, 0x06, 0xae, 0x7a, 0xd2, 0x9d, 0x71, 0x92, 0x06, 0xc7, 0xb5, 0x51, 0x4e, 0x21,
	0x7e, 0x28, 0x6e, 0xb7, 0x26, 0xa8, 0x70, 0xd6, 0xd5, 0x46, 0x49, 0xa9, 0x86, 0x87, 0x75, 0x67,
	0x49, 0xd0, 0xd7, 0x37, 0xd9, 0x98, 0xae, 0x6f, 0x33, 0xc7, 0x6a, 0x35, 0x58, 0x0a, 0x16, 0x5b,
	0x02, 0x8c, 0x70, 0xbc, 0x81, 0x69, 0x4f, 0xda, 0xbd, 0xa9, 0xaa, 0x58, 0x14, 0xcd, 0xb9, 0x88,
	0x0a, 0x6b, 0x38, 0xf3, 0xb7, 0x4e, 0x52, 0x35, 0xf1, 0xce, 0x5e, 0xb3, 0x57, 0x98, 0xbd, 0x90,
	0x5b, 0xed, 0x20, 0x3b, 0x46, 0xfc, 0x19, 0x9f, 0xa0, 0x4c, 0x8a, 0x78, 0x5e, 0xd9, 0xce, 0x78,
	0x68, 0xc2, 0x13, 0x4b, 0xa4, 0x39, 0xf6, 0x5b, 0xc0, 0x3c, 0xc7, 0x85, 0xba, 0xf8, 0x0c, 0x17,
	0x69, 0xfd, 0x08, 0x9c, 0xe7, 0xc0, 0xe0, 0x89, 0x2c, 0x89, 0x4b, 0x28, 0xc7, 0x11, 0x6d, 0x35,
	0x59, 0x9c, 0x36, 0x65, 0x7b, 0xc7, 0xff, 0x6f, 0x90, 0x8f, 0x31, 0x91, 0x3e, 0x69, 0x7f, 0xe0,
	0x7a, 0xb5, 0x97, 0xef, 0x64, 0xb6, 0x9f, 0x1b, 0xc2, 0x1e, 0x2e, 0xd3, 0xa2, 0x16, 0xef, 0x0f,
	0x21, 0x0f, 0xac, 0xa6, 0x6e, 0x8e, 0x07, 0x43, 0x75, 0x76, 0xb2, 0x9e, 0x7a, 0xf7, 0xf1, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x95, 0x84, 0x2a, 0xdc, 0x03, 0x02, 0x00, 0x00,
}
