// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/token.proto

package protos

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

type GenToekenRequest struct {
	// 用户id
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	// 过期时间戳，单位s
	Exp int64 `protobuf:"varint,2,opt,name=exp" json:"exp,omitempty"`
}

func (m *GenToekenRequest) Reset()                    { *m = GenToekenRequest{} }
func (m *GenToekenRequest) String() string            { return proto.CompactTextString(m) }
func (*GenToekenRequest) ProtoMessage()               {}
func (*GenToekenRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *GenToekenRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GenToekenRequest) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

type GenTokenResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GenTokenResponse) Reset()                    { *m = GenTokenResponse{} }
func (m *GenTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*GenTokenResponse) ProtoMessage()               {}
func (*GenTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *GenTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckToekenRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *CheckToekenRequest) Reset()                    { *m = CheckToekenRequest{} }
func (m *CheckToekenRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckToekenRequest) ProtoMessage()               {}
func (*CheckToekenRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *CheckToekenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckTokenResponse struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CheckTokenResponse) Reset()                    { *m = CheckTokenResponse{} }
func (m *CheckTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenResponse) ProtoMessage()               {}
func (*CheckTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *CheckTokenResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*GenToekenRequest)(nil), "protos.GenToekenRequest")
	proto.RegisterType((*GenTokenResponse)(nil), "protos.GenTokenResponse")
	proto.RegisterType((*CheckToekenRequest)(nil), "protos.CheckToekenRequest")
	proto.RegisterType((*CheckTokenResponse)(nil), "protos.CheckTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TokenService service

type TokenServiceClient interface {
	GenAccessToken(ctx context.Context, in *GenToekenRequest, opts ...grpc.CallOption) (*GenTokenResponse, error)
	CheckToken(ctx context.Context, in *CheckToekenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
}

type tokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewTokenServiceClient(cc *grpc.ClientConn) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) GenAccessToken(ctx context.Context, in *GenToekenRequest, opts ...grpc.CallOption) (*GenTokenResponse, error) {
	out := new(GenTokenResponse)
	err := grpc.Invoke(ctx, "/protos.TokenService/GenAccessToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) CheckToken(ctx context.Context, in *CheckToekenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := grpc.Invoke(ctx, "/protos.TokenService/CheckToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TokenService service

type TokenServiceServer interface {
	GenAccessToken(context.Context, *GenToekenRequest) (*GenTokenResponse, error)
	CheckToken(context.Context, *CheckToekenRequest) (*CheckTokenResponse, error)
}

func RegisterTokenServiceServer(s *grpc.Server, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_GenAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenToekenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GenAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TokenService/GenAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GenAccessToken(ctx, req.(*GenToekenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckToekenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TokenService/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CheckToken(ctx, req.(*CheckToekenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenAccessToken",
			Handler:    _TokenService_GenAccessToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _TokenService_CheckToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/token.proto",
}

func init() { proto.RegisterFile("protos/token.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x62, 0x4a,
	0x66, 0x5c, 0x02, 0xee, 0xa9, 0x79, 0x21, 0xf9, 0xa9, 0xd9, 0xa9, 0x79, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xa5, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x20, 0x26, 0x48, 0x24, 0xb5, 0xa2, 0x40, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc4,
	0x54, 0xd2, 0x80, 0xea, 0x03, 0x6b, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1, 0x62,
	0x05, 0x5b, 0x01, 0xd5, 0x09, 0xe1, 0x28, 0x69, 0x71, 0x09, 0x39, 0x67, 0xa4, 0x26, 0x67, 0xa3,
	0xda, 0x81, 0x5d, 0xad, 0x1a, 0x5c, 0x2d, 0xb2, 0xb9, 0x18, 0xee, 0x31, 0x9a, 0xc7, 0xc8, 0xc5,
	0x03, 0x56, 0x13, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc6, 0xc5, 0xe7, 0x9e, 0x9a,
	0xe7, 0x98, 0x9c, 0x9c, 0x5a, 0x5c, 0x0c, 0x96, 0x10, 0x92, 0x80, 0x78, 0xb4, 0x58, 0x0f, 0xdd,
	0x7b, 0x52, 0xa8, 0x32, 0x48, 0x16, 0x29, 0x31, 0x08, 0xb9, 0x71, 0x71, 0x21, 0x1c, 0x20, 0x24,
	0x05, 0x53, 0x89, 0xe9, 0x01, 0x29, 0x74, 0x39, 0x14, 0x73, 0x9c, 0x14, 0xb9, 0x04, 0x93, 0xf3,
	0x73, 0xf5, 0x12, 0x0b, 0xb2, 0x0b, 0x4a, 0x8b, 0x52, 0x21, 0x4a, 0x9d, 0x50, 0x9c, 0x9c, 0x04,
	0x89, 0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x5a, 0xd3, 0x89, 0x9e, 0x01, 0x00,
	0x00,
}