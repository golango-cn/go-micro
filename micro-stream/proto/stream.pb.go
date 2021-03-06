// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package proto

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

type StreamRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StreamRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type StreamResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StreamResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "proto.StreamResponse")
}

func init() {
	proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54)
}

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xb6, 0x5c, 0xbc,
	0xc1, 0x60, 0xe1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x19, 0x2e, 0xce, 0xb4, 0xcc,
	0xa2, 0xe2, 0x12, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x80,
	0x92, 0x15, 0x17, 0x1f, 0x4c, 0x7b, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x86, 0x7e, 0x09, 0x2e,
	0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x6e, 0x18, 0xd7, 0xe8, 0x2c, 0x23, 0x17, 0x4f,
	0x70, 0x49, 0x6a, 0x62, 0x6e, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3d, 0x17, 0x97,
	0x7b, 0x6a, 0x09, 0xc4, 0x3c, 0x43, 0x21, 0x11, 0x88, 0x43, 0xf5, 0x50, 0x9c, 0x27, 0x25, 0x8a,
	0x26, 0x0a, 0xb1, 0x55, 0x89, 0x41, 0x83, 0xd1, 0x80, 0x51, 0xc8, 0x16, 0xc9, 0x00, 0x23, 0x12,
	0x0d, 0x40, 0xd3, 0x6e, 0x4c, 0xb2, 0xfd, 0x49, 0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x62, 0x68, 0xaa, 0x68, 0x01, 0x00, 0x00,
}
