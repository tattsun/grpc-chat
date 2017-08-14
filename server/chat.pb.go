// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	JoinRequest
	JoinResponse
	SayRequest
	SayResponse
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type JoinRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JoinRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JoinResponse struct {
}

func (m *JoinResponse) Reset()                    { *m = JoinResponse{} }
func (m *JoinResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()               {}
func (*JoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SayRequest struct {
	Mes string `protobuf:"bytes,1,opt,name=mes" json:"mes,omitempty"`
}

func (m *SayRequest) Reset()                    { *m = SayRequest{} }
func (m *SayRequest) String() string            { return proto.CompactTextString(m) }
func (*SayRequest) ProtoMessage()               {}
func (*SayRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SayRequest) GetMes() string {
	if m != nil {
		return m.Mes
	}
	return ""
}

type SayResponse struct {
}

func (m *SayResponse) Reset()                    { *m = SayResponse{} }
func (m *SayResponse) String() string            { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()               {}
func (*SayResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*JoinRequest)(nil), "JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "JoinResponse")
	proto.RegisterType((*SayRequest)(nil), "SayRequest")
	proto.RegisterType((*SayResponse)(nil), "SayResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := grpc.Invoke(ctx, "/Chat/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := grpc.Invoke(ctx, "/Chat/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatServer interface {
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Say(context.Context, *SayRequest) (*SayResponse, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Chat_Join_Handler,
		},
		{
			MethodName: "Say",
			Handler:    _Chat_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0xe2, 0xf6, 0xca, 0xcf, 0xcc, 0x0b, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xf8, 0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x95, 0xe4, 0xb8, 0xb8, 0x82, 0x13, 0x2b, 0x61, 0x3a, 0x04, 0xb8, 0x98, 0x73, 0x53,
	0x8b, 0xa1, 0x1a, 0x40, 0x4c, 0x25, 0x5e, 0x2e, 0x6e, 0xb0, 0x3c, 0x44, 0xb9, 0x51, 0x20, 0x17,
	0x8b, 0x73, 0x46, 0x62, 0x89, 0x90, 0x2a, 0x17, 0x0b, 0xc8, 0x18, 0x21, 0x1e, 0x3d, 0x24, 0x0b,
	0xa5, 0x78, 0xf5, 0x50, 0xcc, 0x66, 0x10, 0x52, 0xe2, 0x62, 0x0e, 0x4e, 0xac, 0x14, 0xe2, 0xd6,
	0x43, 0xd8, 0x21, 0xc5, 0xa3, 0x87, 0x64, 0xa0, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xed, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x8c, 0x32, 0xba, 0xc9, 0x00, 0x00, 0x00,
}
