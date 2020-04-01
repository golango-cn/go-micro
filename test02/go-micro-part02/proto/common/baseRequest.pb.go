// Code generated by protoc-gen-go. DO NOT EDIT.
// source: baseRequest.proto

package common

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

type BaseRequest struct {
	AppVersion           int32    `protobuf:"varint,1,opt,name=appVersion,proto3" json:"appVersion,omitempty"`
	Imei                 string   `protobuf:"bytes,2,opt,name=imei,proto3" json:"imei,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d0305dc517399c, []int{0}
}

func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (m *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(m, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetAppVersion() int32 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

func (m *BaseRequest) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "common.baseRequest")
}

func init() {
	proto.RegisterFile("baseRequest.proto", fileDescriptor_b3d0305dc517399c)
}

var fileDescriptor_b3d0305dc517399c = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4a, 0x2c, 0x4e,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x53, 0x72, 0xe4, 0xe2, 0x46, 0x92, 0x14, 0x92, 0xe3, 0xe2, 0x4a,
	0x2c, 0x28, 0x08, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d,
	0x42, 0x12, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0xcc, 0x4d, 0xcd, 0x94, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0x26, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xc5,
	0x07, 0xd5, 0x66, 0x00, 0x00, 0x00,
}
