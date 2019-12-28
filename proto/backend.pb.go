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
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x51, 0xd1, 0x62, 0xda, 0x28, 0x46, 0x18, 0x25, 0x15, 0xc4, 0x01, 0x37, 0x2e, 0x26,
	0x58, 0x77, 0xba, 0x1b, 0x15, 0xe9, 0x42, 0x0b, 0x23, 0xce, 0xc2, 0x5d, 0x66, 0xe6, 0x59, 0x47,
	0x6d, 0x12, 0x27, 0x69, 0xa1, 0x94, 0x6e, 0xbc, 0x82, 0xa7, 0xf0, 0x3c, 0x5e, 0xc1, 0x83, 0xc8,
	0x24, 0x51, 0x6a, 0x17, 0x42, 0x57, 0x8f, 0x3f, 0xef, 0xfd, 0xdf, 0xff, 0x43, 0x10, 0xce, 0x78,
	0xfe, 0x0c, 0xa2, 0x88, 0x54, 0x25, 0x8d, 0x24, 0xab, 0x76, 0xd0, 0x76, 0x5f, 0xca, 0xfe, 0x0b,
	0x30, 0xab, 0xb2, 0xe1, 0x03, 0x83, 0x81, 0x32, 0x63, 0x77, 0x43, 0xf7, 0xfc, 0x92, 0xab, 0x92,
	0x71, 0x21, 0xa4, 0xe1, 0xa6, 0x94, 0x42, 0xfb, 0x6d, 0x4b, 0xe7, 0x8f, 0x30, 0xe0, 0x4e, 0x75,
	0x3e, 0x56, 0x50, 0x23, 0x76, 0x09, 0xe4, 0x06, 0x35, 0x52, 0xa8, 0x74, 0x29, 0x05, 0x09, 0x22,
	0xc7, 0x88, 0x7e, 0x02, 0xa2, 0xcb, 0x3a, 0x80, 0x06, 0xee, 0x21, 0xf2, 0x77, 0x09, 0x68, 0x25,
	0x85, 0x86, 0x70, 0xfb, 0xed, 0xf3, 0xeb, 0x7d, 0x19, 0x93, 0x26, 0x1b, 0x1d, 0xb3, 0x91, 0x87,
	0xf4, 0x50, 0xeb, 0xbc, 0x02, 0x6e, 0xe0, 0xd6, 0x26, 0x12, 0xec, 0xcd, 0x4e, 0xd2, 0xbf, 0x32,
	0x3c, 0xb0, 0x88, 0x76, 0x18, 0xd4, 0x08, 0xd7, 0x91, 0x4d, 0x54, 0x25, 0x9f, 0x20, 0x37, 0xdd,
	0x8b, 0xe9, 0xe9, 0xd2, 0x11, 0xb9, 0x46, 0xeb, 0x57, 0x60, 0x3c, 0x6d, 0xd7, 0xdb, 0x13, 0x78,
	0x1d, 0x82, 0x36, 0xbd, 0xac, 0x3e, 0x8c, 0xc7, 0xdd, 0x62, 0x1e, 0xbc, 0x63, 0xc1, 0x5b, 0x64,
	0x73, 0x16, 0x5c, 0x16, 0x53, 0x92, 0x22, 0xec, 0xfa, 0x25, 0x90, 0xcb, 0xaa, 0xd0, 0x64, 0xe3,
	0x17, 0x69, 0x35, 0x9d, 0xd3, 0xe1, 0xa1, 0x25, 0xed, 0x87, 0xb4, 0x26, 0xcd, 0x74, 0x63, 0x13,
	0x87, 0xf5, 0x35, 0x53, 0x84, 0xef, 0x54, 0xb1, 0x38, 0xb7, 0xf3, 0x3f, 0x37, 0xc6, 0xf7, 0x4d,
	0xeb, 0x3b, 0x73, 0x9f, 0xb2, 0x66, 0xc7, 0xc9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xaf,
	0x59, 0x88, 0x22, 0x02, 0x00, 0x00,
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
	// Create Schema
	CreateSchema(ctx context.Context, in *Schema, opts ...grpc.CallOption) (*Schema, error)
	// Get Schema
	GetSchema(ctx context.Context, in *RequestObjectById, opts ...grpc.CallOption) (*Schema, error)
	// CreateRecords
	CreateRecords(ctx context.Context, in *Records, opts ...grpc.CallOption) (*Records, error)
	// UpdateRecords
	UpdateRecords(ctx context.Context, in *Records, opts ...grpc.CallOption) (*Records, error)
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

func (c *backendClient) GetSchema(ctx context.Context, in *RequestObjectById, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, "/proto.Backend/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CreateRecords(ctx context.Context, in *Records, opts ...grpc.CallOption) (*Records, error) {
	out := new(Records)
	err := c.cc.Invoke(ctx, "/proto.Backend/CreateRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UpdateRecords(ctx context.Context, in *Records, opts ...grpc.CallOption) (*Records, error) {
	out := new(Records)
	err := c.cc.Invoke(ctx, "/proto.Backend/UpdateRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
type BackendServer interface {
	// Version
	Version(context.Context, *empty.Empty) (*VersionResponse, error)
	// Create Schema
	CreateSchema(context.Context, *Schema) (*Schema, error)
	// Get Schema
	GetSchema(context.Context, *RequestObjectById) (*Schema, error)
	// CreateRecords
	CreateRecords(context.Context, *Records) (*Records, error)
	// UpdateRecords
	UpdateRecords(context.Context, *Records) (*Records, error)
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
func (*UnimplementedBackendServer) GetSchema(ctx context.Context, req *RequestObjectById) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (*UnimplementedBackendServer) CreateRecords(ctx context.Context, req *Records) (*Records, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecords not implemented")
}
func (*UnimplementedBackendServer) UpdateRecords(ctx context.Context, req *Records) (*Records, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecords not implemented")
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

func _Backend_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestObjectById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Backend/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetSchema(ctx, req.(*RequestObjectById))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CreateRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Records)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreateRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Backend/CreateRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreateRecords(ctx, req.(*Records))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UpdateRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Records)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UpdateRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Backend/UpdateRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UpdateRecords(ctx, req.(*Records))
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
		{
			MethodName: "GetSchema",
			Handler:    _Backend_GetSchema_Handler,
		},
		{
			MethodName: "CreateRecords",
			Handler:    _Backend_CreateRecords_Handler,
		},
		{
			MethodName: "UpdateRecords",
			Handler:    _Backend_UpdateRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}
