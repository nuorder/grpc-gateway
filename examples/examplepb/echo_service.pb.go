// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/examplepb/echo_service.proto

/*
Package examplepb is a generated protocol buffer package.

Echo Service

Echo Service API consists of a single service which returns
a message.

It is generated from these files:
	examples/examplepb/echo_service.proto
	examples/examplepb/a_bit_of_everything.proto
	examples/examplepb/stream.proto
	examples/examplepb/flow_combination.proto
	examples/examplepb/wrappers.proto

It has these top-level messages:
	SimpleMessage
	ABitOfEverything
	Body
	MessageWithBody
	EmptyProto
	NonEmptyProto
	UnaryProto
	NestedProto
	SingleNestedProto
	Wrappers
*/
package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "context"
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

// SimpleMessage represents a simple message sent to the Echo service.
type SimpleMessage struct {
	// Id represents the message identifier.
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (m *SimpleMessage) Reset()                    { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()               {}
func (*SimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "grpc.gateway.examples.examplepb.SimpleMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/EchoBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/examplepb/echo_service.proto",
}

func init() { proto.RegisterFile("examples/examplepb/echo_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x87, 0x32, 0x0a, 0x92, 0xf4, 0x53, 0x93, 0x33, 0xf2, 0xe3, 0x8b,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0xd3, 0x8b,
	0x0a, 0x92, 0xf5, 0xd2, 0x13, 0x4b, 0x52, 0xcb, 0x13, 0x2b, 0xf5, 0x60, 0x7a, 0xf4, 0xe0, 0x7a,
	0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2,
	0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xda, 0x95, 0x0c, 0xb9, 0x78, 0x83, 0x33,
	0x41, 0x2a, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0x4a, 0x73,
	0x25, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x40, 0x4c, 0xa3, 0xc3, 0x4c, 0x5c, 0xdc, 0xae, 0xc9,
	0x19, 0xf9, 0xc1, 0x10, 0x77, 0x08, 0x2d, 0x61, 0xe4, 0x62, 0x01, 0xf1, 0x85, 0xf4, 0xf4, 0x08,
	0xb8, 0x45, 0x0f, 0xc5, 0x2a, 0x29, 0x12, 0xd5, 0x2b, 0xd9, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0xc9,
	0x4c, 0x49, 0x54, 0xbf, 0xcc, 0x10, 0x16, 0x28, 0xe0, 0x20, 0xd1, 0xaf, 0xce, 0x4c, 0xa9, 0x8d,
	0x92, 0x15, 0x92, 0xc6, 0x2a, 0xa1, 0x5f, 0x9d, 0x57, 0x9a, 0x5b, 0x2b, 0xd4, 0xc3, 0xc8, 0xc5,
	0x01, 0x72, 0xa6, 0x53, 0x7e, 0x4a, 0x25, 0xcd, 0x9d, 0xaa, 0x00, 0x76, 0xaa, 0x14, 0xa6, 0x53,
	0xe3, 0x93, 0xf2, 0x53, 0x2a, 0xad, 0x18, 0xb5, 0x9c, 0xb8, 0xa3, 0x38, 0xe1, 0x9a, 0x93, 0xd8,
	0xc0, 0x91, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x6e, 0x70, 0xce, 0xf4, 0x01, 0x00,
	0x00,
}
