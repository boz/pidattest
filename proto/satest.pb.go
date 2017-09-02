// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/satest.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/satest.proto

It has these top-level messages:
	Request
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto1.RegisterType((*Request)(nil), "proto.Request")
	proto1.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Workload service

type WorkloadClient interface {
	Register(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type workloadClient struct {
	cc *grpc.ClientConn
}

func NewWorkloadClient(cc *grpc.ClientConn) WorkloadClient {
	return &workloadClient{cc}
}

func (c *workloadClient) Register(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Workload/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Workload service

type WorkloadServer interface {
	Register(context.Context, *Request) (*Response, error)
}

func RegisterWorkloadServer(s *grpc.Server, srv WorkloadServer) {
	s.RegisterService(&_Workload_serviceDesc, srv)
}

func _Workload_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Workload/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).Register(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Workload",
	HandlerType: (*WorkloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Workload_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/satest.proto",
}

func init() { proto1.RegisterFile("proto/satest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2c, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x12,
	0x27, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x12, 0x17, 0x17, 0x47, 0x50, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x91, 0x25, 0x17, 0x47, 0x78, 0x7e, 0x51, 0x76, 0x4e, 0x7e,
	0x62, 0x8a, 0x90, 0x2e, 0x48, 0x3c, 0x3d, 0xb3, 0xb8, 0x24, 0xb5, 0x48, 0x88, 0x0f, 0xa2, 0x5b,
	0x0f, 0xaa, 0x47, 0x8a, 0x1f, 0xce, 0x87, 0x68, 0x54, 0x62, 0x48, 0x62, 0x03, 0x8b, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x97, 0x01, 0x48, 0x83, 0x75, 0x00, 0x00, 0x00,
}