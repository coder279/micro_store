// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package go_micro_service_user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRegisterRequest
	UserRegisterResponse
	UserLoginRequest
	UserLoginResponse
	UserInfoRequest
	UserInfoResponse
*/
package go_micro_service_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRegisterRequest struct {
	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *UserRegisterRequest) Reset()                    { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()               {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserRegisterResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserRegisterResponse) Reset()                    { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()               {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserLoginRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *UserLoginRequest) Reset()                    { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()               {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserLoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserLoginResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserLoginResponse) Reset()                    { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()               {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserLoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserInfoRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *UserInfoRequest) Reset()                    { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()               {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserInfoResponse struct {
	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
}

func (m *UserInfoResponse) Reset()                    { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()               {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserInfoResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRegisterRequest)(nil), "go.micro.service.user.UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "go.micro.service.user.UserRegisterResponse")
	proto.RegisterType((*UserLoginRequest)(nil), "go.micro.service.user.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "go.micro.service.user.UserLoginResponse")
	proto.RegisterType((*UserInfoRequest)(nil), "go.micro.service.user.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "go.micro.service.user.UserInfoResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0x76, 0xeb, 0x5f, 0x3b, 0x1e, 0xd4, 0xa8, 0x50, 0x2a, 0x82, 0xe4, 0xe0, 0x2e, 0x8a, 0x41,
	0xf4, 0x21, 0x44, 0x94, 0x3d, 0x2c, 0x78, 0x13, 0xb4, 0xae, 0xb3, 0x25, 0x87, 0x34, 0x35, 0x93,
	0xd5, 0xa7, 0xf2, 0x1d, 0x25, 0x89, 0x91, 0x2a, 0xcb, 0xb6, 0xb0, 0xa7, 0x92, 0xf9, 0xfe, 0x86,
	0x6f, 0x0a, 0x30, 0x27, 0x34, 0xa2, 0x31, 0xda, 0x6a, 0x76, 0x54, 0x69, 0xa1, 0xe4, 0xd4, 0x68,
	0x41, 0x68, 0x3e, 0xe4, 0x14, 0x85, 0x03, 0xb9, 0x82, 0x83, 0x47, 0x42, 0x33, 0xc1, 0x4a, 0x92,
	0x75, 0xdf, 0xf7, 0x39, 0x92, 0x65, 0xc7, 0x90, 0x39, 0xf8, 0xb9, 0x2e, 0x15, 0xe6, 0x83, 0xd3,
	0xc1, 0x28, 0x9b, 0xa4, 0x6e, 0x30, 0x2e, 0x15, 0xb2, 0x13, 0x80, 0x99, 0x34, 0x64, 0x03, 0x9a,
	0x78, 0x34, 0xf3, 0x13, 0x0f, 0x17, 0x90, 0x36, 0x25, 0xd1, 0xa7, 0x36, 0x6f, 0xf9, 0x7a, 0x90,
	0xc6, 0x37, 0xbf, 0x82, 0xc3, 0xbf, 0x71, 0xd4, 0xe8, 0x9a, 0x90, 0xe5, 0xb0, 0xad, 0x90, 0xa8,
	0xac, 0x62, 0x5a, 0x7c, 0xf2, 0x7b, 0xd8, 0x73, 0x8a, 0x07, 0x5d, 0xc9, 0xba, 0xd7, 0x76, 0xed,
	0xf8, 0xe4, 0x5f, 0xfc, 0x25, 0xec, 0xb7, 0xcc, 0x3a, 0xb3, 0x05, 0xec, 0x3a, 0xfa, 0x5d, 0x3d,
	0xd3, 0x7d, 0xa2, 0xf9, 0x38, 0xec, 0x1a, 0xf8, 0x3f, 0xee, 0x2b, 0x34, 0x79, 0xfd, 0x95, 0xc0,
	0x86, 0x33, 0x64, 0x08, 0x69, 0xac, 0x8c, 0x9d, 0x8b, 0x85, 0x97, 0x14, 0x0b, 0xce, 0x58, 0x5c,
	0xf4, 0xe2, 0x86, 0x4d, 0xf9, 0x1a, 0x7b, 0x82, 0x4d, 0x5f, 0x0d, 0x1b, 0x2e, 0xd1, 0xb5, 0x2f,
	0x51, 0x8c, 0xba, 0x89, 0xbf, 0xee, 0x2f, 0xb0, 0x73, 0x8b, 0x36, 0x16, 0xc4, 0xce, 0x96, 0x48,
	0x5b, 0x8d, 0x17, 0xc3, 0x4e, 0x5e, 0x4c, 0x78, 0xdd, 0xf2, 0xbf, 0xfa, 0xcd, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x67, 0xa0, 0xa4, 0xd5, 0xf8, 0x02, 0x00, 0x00,
}
