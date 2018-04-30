// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adder.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	adder.proto

It has these top-level messages:
	AddRequest
	AddReply
*/
package pb

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

type AddRequest struct {
	Operands []int32 `protobuf:"varint,1,rep,packed,name=operands" json:"operands,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetOperands() []int32 {
	if m != nil {
		return m.Operands
	}
	return nil
}

type AddReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*AddReply)(nil), "pb.AddReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Adder service

type AdderClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
}

type adderClient struct {
	cc *grpc.ClientConn
}

func NewAdderClient(cc *grpc.ClientConn) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/pb.Adder/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Adder service

type AdderServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
}

func RegisterAdderServer(s *grpc.Server, srv AdderServer) {
	s.RegisterService(&_Adder_serviceDesc, srv)
}

func _Adder_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Adder/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Adder_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adder.proto",
}

func init() { proto.RegisterFile("adder.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0x49, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe0, 0xe2, 0x72,
	0x4c, 0x49, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0x2f, 0x48,
	0x2d, 0x4a, 0xcc, 0x4b, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0x60, 0x0d, 0x82, 0xf3, 0x95, 0x94,
	0xb8, 0x38, 0xc0, 0x2a, 0x0b, 0x72, 0x2a, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73,
	0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x23, 0x3d, 0x2e, 0x56, 0x47, 0x90,
	0x05, 0x42, 0xaa, 0x5c, 0xcc, 0x8e, 0x29, 0x29, 0x42, 0x7c, 0x7a, 0x05, 0x49, 0x7a, 0x08, 0xf3,
	0xa5, 0x78, 0xe0, 0xfc, 0x82, 0x9c, 0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x43, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xc6, 0x45, 0x47, 0x97, 0x00, 0x00, 0x00,
}
