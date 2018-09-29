// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GeneralRequest struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	ByteInput            []byte   `protobuf:"bytes,2,opt,name=byteInput,proto3" json:"byteInput,omitempty"`
	Inputs               []string `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralRequest) Reset()         { *m = GeneralRequest{} }
func (m *GeneralRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralRequest) ProtoMessage()    {}
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *GeneralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralRequest.Unmarshal(m, b)
}
func (m *GeneralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralRequest.Marshal(b, m, deterministic)
}
func (m *GeneralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralRequest.Merge(m, src)
}
func (m *GeneralRequest) XXX_Size() int {
	return xxx_messageInfo_GeneralRequest.Size(m)
}
func (m *GeneralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralRequest proto.InternalMessageInfo

func (m *GeneralRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *GeneralRequest) GetByteInput() []byte {
	if m != nil {
		return m.ByteInput
	}
	return nil
}

func (m *GeneralRequest) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type GeneralResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (m *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(m, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GeneralRequest)(nil), "common.GeneralRequest")
	proto.RegisterType((*GeneralResponse)(nil), "common.GeneralResponse")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x4b, 0x51, 0x8d, 0x3c, 0xa2, 0xad, 0xba, 0xa2, 0x80, 0x2a, 0x0e, 0x88, 0x13, 0x52,
	0x25, 0x0e, 0x54, 0xed, 0xa5, 0x52, 0x25, 0x70, 0x0b, 0x25, 0xff, 0xe4, 0x18, 0x92, 0x5c, 0x72,
	0x59, 0xec, 0x09, 0xb6, 0x62, 0xef, 0x3a, 0xbb, 0x63, 0x14, 0xbe, 0x5a, 0x3e, 0x5d, 0x64, 0x1b,
	0x02, 0x49, 0x6e, 0x9b, 0x8b, 0xa5, 0x37, 0xde, 0xfd, 0xcd, 0x7b, 0xa3, 0xdd, 0x85, 0xba, 0x2f,
	0x93, 0x44, 0x8a, 0x41, 0xaa, 0x24, 0x49, 0x66, 0x95, 0xaa, 0x77, 0x0d, 0x9f, 0xa6, 0x28, 0x50,
	0xf1, 0xd8, 0xc3, 0xbb, 0x0c, 0x35, 0xb1, 0x06, 0x7c, 0x88, 0x44, 0x9a, 0x51, 0xbb, 0xd2, 0xad,
	0xf4, 0x6d, 0xaf, 0x14, 0xac, 0x03, 0xf6, 0x72, 0x43, 0x38, 0x2b, 0xfe, 0xbc, 0xef, 0x56, 0xfa,
	0x75, 0x6f, 0x5f, 0x60, 0x4d, 0xb0, 0x8a, 0x65, 0xba, 0x5d, 0xed, 0x56, 0xfb, 0xb6, 0xb7, 0x55,
	0xbd, 0xef, 0xf0, 0xf9, 0x89, 0xae, 0x53, 0x29, 0x34, 0xb2, 0x36, 0xd4, 0x12, 0xd4, 0x9a, 0xaf,
	0x70, 0xdb, 0x60, 0x27, 0x87, 0x0f, 0x35, 0xb0, 0x9c, 0xc2, 0x15, 0x3b, 0x82, 0xaf, 0x2e, 0x57,
	0x1a, 0xe7, 0xa4, 0x22, 0xb1, 0x3a, 0x45, 0x0a, 0x65, 0xe0, 0xf0, 0x38, 0x66, 0xcd, 0xc1, 0x36,
	0xc5, 0x73, 0xd3, 0xdf, 0x5a, 0xaf, 0xea, 0x65, 0xbb, 0xde, 0x3b, 0x36, 0x81, 0x2f, 0x07, 0x2c,
	0x97, 0x2b, 0x9e, 0x68, 0x13, 0xce, 0x39, 0x74, 0x1c, 0x29, 0xd6, 0xa8, 0xa8, 0x24, 0x2d, 0xa4,
	0x87, 0x37, 0x31, 0xfa, 0x74, 0xc9, 0xe3, 0x0c, 0x8d, 0x90, 0x27, 0xd0, 0x2a, 0x59, 0xf9, 0x37,
	0xf5, 0xd0, 0xc7, 0x68, 0x8d, 0xca, 0x34, 0xe8, 0x31, 0x34, 0x0f, 0x68, 0x2e, 0x57, 0x28, 0x28,
	0x44, 0xfd, 0x46, 0x6b, 0x13, 0x24, 0x3f, 0xcc, 0x2d, 0xed, 0xec, 0x99, 0xd0, 0x46, 0x50, 0x77,
	0x42, 0xf4, 0x6f, 0x47, 0x41, 0xa0, 0x50, 0x1b, 0x19, 0x9a, 0x41, 0x63, 0x8a, 0xf4, 0xef, 0x9e,
	0x66, 0x6e, 0x4e, 0xb9, 0x8a, 0x28, 0xbc, 0x48, 0xcf, 0x5c, 0xc3, 0x6c, 0x2f, 0x51, 0x32, 0x23,
	0x53, 0x9a, 0x93, 0xdf, 0x20, 0x72, 0x32, 0x95, 0xcf, 0x7b, 0x11, 0x25, 0x68, 0x02, 0x19, 0xc3,
	0xc7, 0x3d, 0xe4, 0x6f, 0x64, 0x34, 0xe4, 0xdf, 0x60, 0xcd, 0xff, 0x8f, 0x86, 0x3f, 0x7f, 0x99,
	0x6c, 0xfe, 0x03, 0xf6, 0x1c, 0x45, 0x30, 0xde, 0x90, 0xd1, 0x79, 0x59, 0x5a, 0xc5, 0xb3, 0xf2,
	0xe3, 0x31, 0x00, 0x00, 0xff, 0xff, 0xff, 0x80, 0xdd, 0x17, 0x66, 0x04, 0x00, 0x00,
}