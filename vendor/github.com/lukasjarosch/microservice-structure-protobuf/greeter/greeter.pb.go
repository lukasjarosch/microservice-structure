// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GreetingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingRequest) Reset()         { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()    {}
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_0a068844aff3a0b4, []int{0}
}
func (m *GreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingRequest.Unmarshal(m, b)
}
func (m *GreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingRequest.Marshal(b, m, deterministic)
}
func (dst *GreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingRequest.Merge(dst, src)
}
func (m *GreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingRequest.Size(m)
}
func (m *GreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingRequest proto.InternalMessageInfo

func (m *GreetingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetingResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingResponse) Reset()         { *m = GreetingResponse{} }
func (m *GreetingResponse) String() string { return proto.CompactTextString(m) }
func (*GreetingResponse) ProtoMessage()    {}
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_0a068844aff3a0b4, []int{1}
}
func (m *GreetingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingResponse.Unmarshal(m, b)
}
func (m *GreetingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingResponse.Marshal(b, m, deterministic)
}
func (dst *GreetingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingResponse.Merge(dst, src)
}
func (m *GreetingResponse) XXX_Size() int {
	return xxx_messageInfo_GreetingResponse.Size(m)
}
func (m *GreetingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingResponse proto.InternalMessageInfo

func (m *GreetingResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type FarewellRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FarewellRequest) Reset()         { *m = FarewellRequest{} }
func (m *FarewellRequest) String() string { return proto.CompactTextString(m) }
func (*FarewellRequest) ProtoMessage()    {}
func (*FarewellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_0a068844aff3a0b4, []int{2}
}
func (m *FarewellRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FarewellRequest.Unmarshal(m, b)
}
func (m *FarewellRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FarewellRequest.Marshal(b, m, deterministic)
}
func (dst *FarewellRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FarewellRequest.Merge(dst, src)
}
func (m *FarewellRequest) XXX_Size() int {
	return xxx_messageInfo_FarewellRequest.Size(m)
}
func (m *FarewellRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FarewellRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FarewellRequest proto.InternalMessageInfo

func (m *FarewellRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FarewellResponse struct {
	Farewell             string   `protobuf:"bytes,1,opt,name=farewell,proto3" json:"farewell,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FarewellResponse) Reset()         { *m = FarewellResponse{} }
func (m *FarewellResponse) String() string { return proto.CompactTextString(m) }
func (*FarewellResponse) ProtoMessage()    {}
func (*FarewellResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_0a068844aff3a0b4, []int{3}
}
func (m *FarewellResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FarewellResponse.Unmarshal(m, b)
}
func (m *FarewellResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FarewellResponse.Marshal(b, m, deterministic)
}
func (dst *FarewellResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FarewellResponse.Merge(dst, src)
}
func (m *FarewellResponse) XXX_Size() int {
	return xxx_messageInfo_FarewellResponse.Size(m)
}
func (m *FarewellResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FarewellResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FarewellResponse proto.InternalMessageInfo

func (m *FarewellResponse) GetFarewell() string {
	if m != nil {
		return m.Farewell
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingRequest)(nil), "greeter.GreetingRequest")
	proto.RegisterType((*GreetingResponse)(nil), "greeter.GreetingResponse")
	proto.RegisterType((*FarewellRequest)(nil), "greeter.FarewellRequest")
	proto.RegisterType((*FarewellResponse)(nil), "greeter.FarewellResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error)
	Farewell(ctx context.Context, in *FarewellRequest, opts ...grpc.CallOption) (*FarewellResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/greeter.Hello/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Farewell(ctx context.Context, in *FarewellRequest, opts ...grpc.CallOption) (*FarewellResponse, error) {
	out := new(FarewellResponse)
	err := c.cc.Invoke(ctx, "/greeter.Hello/Farewell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error)
	Farewell(context.Context, *FarewellRequest) (*FarewellResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.Hello/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Farewell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FarewellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Farewell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.Hello/Farewell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Farewell(ctx, req.(*FarewellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _Hello_Greeting_Handler,
		},
		{
			MethodName: "Farewell",
			Handler:    _Hello_Farewell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter.proto",
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_greeter_0a068844aff3a0b4) }

var fileDescriptor_greeter_0a068844aff3a0b4 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x64, 0xd2,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12,
	0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xca, 0xa4, 0x74, 0xc0, 0x54, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e,
	0x6e, 0x71, 0x79, 0x62, 0x7a, 0x7a, 0x6a, 0x91, 0x7e, 0x7e, 0x01, 0x58, 0x05, 0xa6, 0x6a, 0x25,
	0x55, 0x2e, 0x7e, 0x77, 0x90, 0xb1, 0x99, 0x79, 0xe9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60,
	0xb6, 0x92, 0x1e, 0x97, 0x00, 0x42, 0x59, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17,
	0x47, 0x3a, 0x54, 0x0c, 0xaa, 0x16, 0xce, 0x07, 0x19, 0xeb, 0x96, 0x58, 0x94, 0x5a, 0x9e, 0x9a,
	0x93, 0x43, 0xc0, 0x58, 0x84, 0x32, 0x84, 0xb1, 0x69, 0x50, 0x31, 0x98, 0xb1, 0x30, 0xbe, 0xd1,
	0x6e, 0x46, 0x2e, 0x56, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0xa1, 0x08, 0x2e, 0x0e, 0x98, 0x83, 0x84,
	0x24, 0xf4, 0x60, 0x01, 0x85, 0xe6, 0x15, 0x29, 0x49, 0x2c, 0x32, 0x10, 0x6b, 0x94, 0xc4, 0x9b,
	0x2e, 0x3f, 0x99, 0xcc, 0x24, 0xa8, 0xc4, 0xa3, 0x5f, 0x66, 0xa8, 0x0f, 0x73, 0xb7, 0x15, 0xa3,
	0x16, 0xc8, 0x64, 0x98, 0x9b, 0x90, 0x4c, 0x46, 0xf3, 0x0d, 0x92, 0xc9, 0xe8, 0x1e, 0x40, 0x35,
	0x19, 0xe6, 0x74, 0x2b, 0x46, 0xad, 0x24, 0x36, 0x70, 0x90, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0x02, 0xb8, 0xfa, 0xd8, 0x01, 0x00, 0x00,
}
