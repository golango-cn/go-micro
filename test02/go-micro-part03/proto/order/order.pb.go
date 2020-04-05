// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	BookId               int64    `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	OrderId              int64    `protobuf:"varint,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Request) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type Order struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	BookId               int64    `protobuf:"varint,3,opt,name=bookId,proto3" json:"bookId,omitempty"`
	InvHistoryId         int64    `protobuf:"varint,4,opt,name=invHistoryId,proto3" json:"invHistoryId,omitempty"`
	State                int64    `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
	CreatedTime          int64    `protobuf:"varint,6,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,7,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Order) GetInvHistoryId() int64 {
	if m != nil {
		return m.InvHistoryId
	}
	return 0
}

func (m *Order) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Order) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Order) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Order                *Order   `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.orders.Request")
	proto.RegisterType((*Order)(nil), "mu.micro.book.srv.orders.Order")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.orders.Response")
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.orders.Error")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcf, 0x4e, 0x32, 0x31,
	0x10, 0xff, 0x76, 0x97, 0x2e, 0x7c, 0x83, 0xf1, 0xd0, 0x18, 0xd3, 0x78, 0x11, 0x7b, 0xf2, 0xd4,
	0x03, 0xc4, 0x47, 0x30, 0xba, 0x17, 0x4c, 0x8a, 0x2f, 0x00, 0xdb, 0x39, 0x34, 0x0a, 0xc5, 0xb6,
	0x8b, 0xf1, 0x35, 0x7c, 0x0f, 0x1f, 0xc3, 0xf7, 0x32, 0x9d, 0x82, 0x81, 0x03, 0x9c, 0xbc, 0xf5,
	0xf7, 0x67, 0x7e, 0xd3, 0x99, 0x16, 0x86, 0xce, 0x1b, 0xf4, 0x6a, 0xed, 0x5d, 0x74, 0x5c, 0x2c,
	0x3b, 0xb5, 0xb4, 0xad, 0x77, 0x6a, 0xe1, 0xdc, 0x8b, 0x0a, 0x7e, 0xa3, 0x48, 0x0e, 0x72, 0x06,
	0x7d, 0x8d, 0x6f, 0x1d, 0x86, 0xc8, 0x2f, 0xa1, 0xee, 0x02, 0xfa, 0xc6, 0x88, 0x62, 0x54, 0xdc,
	0x56, 0x7a, 0x8b, 0x12, 0x9f, 0xaa, 0x1a, 0x23, 0xca, 0xcc, 0x67, 0xc4, 0x05, 0xf4, 0x29, 0xa4,
	0x31, 0xa2, 0x22, 0x61, 0x07, 0xe5, 0x77, 0x01, 0xec, 0x29, 0x9d, 0xf9, 0x39, 0x94, 0x76, 0x97,
	0x57, 0x5a, 0xb3, 0xd7, 0xa3, 0x3c, 0xd2, 0xa3, 0x3a, 0xe8, 0x21, 0xe1, 0xcc, 0xae, 0x36, 0x8f,
	0x36, 0x44, 0xe7, 0x3f, 0x1a, 0x23, 0x7a, 0xa4, 0x1e, 0x70, 0xfc, 0x02, 0x58, 0x88, 0xf3, 0x88,
	0x82, 0x91, 0x98, 0x01, 0x1f, 0xc1, 0xb0, 0xf5, 0x38, 0x8f, 0x68, 0x9e, 0xed, 0x12, 0x45, 0x4d,
	0xda, 0x3e, 0x95, 0x1c, 0xdd, 0xda, 0xfc, 0x3a, 0xfa, 0xd9, 0xb1, 0x47, 0xc9, 0xcf, 0x02, 0x06,
	0x1a, 0xc3, 0xda, 0xad, 0x02, 0xa6, 0x71, 0x43, 0xd7, 0xb6, 0x18, 0x02, 0xcd, 0x33, 0xd0, 0x3b,
	0xc8, 0xef, 0x80, 0xa1, 0xf7, 0xce, 0xd3, 0x4c, 0xc3, 0xf1, 0xb5, 0x3a, 0xb6, 0x6d, 0x75, 0x9f,
	0x6c, 0x3a, 0xbb, 0x53, 0x19, 0xd1, 0x34, 0xf2, 0xc9, 0x32, 0xda, 0xa5, 0xce, 0x6e, 0x39, 0x01,
	0x46, 0x31, 0x9c, 0x43, 0xaf, 0x75, 0x06, 0xe9, 0x36, 0x4c, 0xd3, 0x39, 0xed, 0xd1, 0x60, 0x9c,
	0xdb, 0x57, 0xba, 0xcb, 0x7f, 0xbd, 0x45, 0xe3, 0xaf, 0x02, 0x6a, 0x4a, 0x09, 0x7c, 0x0a, 0xd5,
	0x14, 0xdf, 0xf9, 0xcd, 0xf1, 0x76, 0xdb, 0x0f, 0x71, 0x25, 0x4f, 0x59, 0xf2, 0x56, 0xe4, 0x3f,
	0x3e, 0x83, 0xc1, 0x03, 0xc6, 0xfc, 0xdc, 0x7f, 0x15, 0xba, 0xa8, 0xe9, 0xdf, 0x4e, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0xe5, 0xbf, 0xc5, 0xc6, 0x02, 0x00, 0x00,
}
