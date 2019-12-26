// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("backend.proto", fileDescriptor_5ab9ba5b8d8b2ba5) }

var fileDescriptor_5ab9ba5b8d8b2ba5 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xd2, 0xe9,
	0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49,
	0x25, 0x44, 0x8d, 0x94, 0x0c, 0x54, 0x32, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24,
	0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x18, 0x2a, 0xcb, 0x53, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0x08, 0xe1,
	0x19, 0xcd, 0x63, 0xe4, 0x62, 0x77, 0x82, 0xd8, 0x20, 0xe4, 0xc7, 0xc5, 0x1e, 0x96, 0x5a, 0x54,
	0x9c, 0x99, 0x9f, 0x27, 0x24, 0xa6, 0x07, 0x31, 0x43, 0x0f, 0x66, 0x81, 0x9e, 0x2b, 0xc8, 0x02,
	0x29, 0x31, 0x88, 0x80, 0x1e, 0x54, 0x5d, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x92,
	0x70, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x78, 0x85, 0xb8, 0xf5, 0xcb, 0x0c, 0xf5, 0xcb, 0xa0, 0x86,
	0x38, 0x73, 0xf1, 0x38, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x06, 0x83, 0x6d, 0x14, 0xe2, 0x85, 0x6a,
	0x86, 0x70, 0xa5, 0x50, 0xb9, 0x4a, 0xa2, 0x60, 0x23, 0xf8, 0xa5, 0xb8, 0x40, 0x46, 0x40, 0xdc,
	0x68, 0xc5, 0xa8, 0xe5, 0xc4, 0x1b, 0xc5, 0x0d, 0x56, 0x66, 0x0d, 0x71, 0x09, 0x1b, 0x98, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0xfa, 0x14, 0xa4, 0x17, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendClient interface {
	// Version
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	//Update User Betting
	CreateSchema(ctx context.Context, in *Schema, opts ...grpc.CallOption) (*Schema, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/proto.Backend/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CreateSchema(ctx context.Context, in *Schema, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, "/proto.Backend/CreateSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
type BackendServer interface {
	// Version
	Version(context.Context, *empty.Empty) (*VersionResponse, error)
	//Update User Betting
	CreateSchema(context.Context, *Schema) (*Schema, error)
}

// UnimplementedBackendServer can be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (*UnimplementedBackendServer) Version(ctx context.Context, req *empty.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedBackendServer) CreateSchema(ctx context.Context, req *Schema) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Backend/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CreateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Schema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Backend/CreateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreateSchema(ctx, req.(*Schema))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Backend_Version_Handler,
		},
		{
			MethodName: "CreateSchema",
			Handler:    _Backend_CreateSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}
