// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/collectsub/collectsub/collectsub.proto

package collectsub

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ColectSubscriberServiceClient is the client API for ColectSubscriberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ColectSubscriberServiceClient interface {
	AddCollectEntries(ctx context.Context, in *AddCollectEntriesRequest, opts ...grpc.CallOption) (*AddCollectEntriesResponse, error)
	GetCollectEntries(ctx context.Context, in *GetCollectEntriesRequest, opts ...grpc.CallOption) (*GetCollectEntriesResponse, error)
}

type colectSubscriberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewColectSubscriberServiceClient(cc grpc.ClientConnInterface) ColectSubscriberServiceClient {
	return &colectSubscriberServiceClient{cc}
}

func (c *colectSubscriberServiceClient) AddCollectEntries(ctx context.Context, in *AddCollectEntriesRequest, opts ...grpc.CallOption) (*AddCollectEntriesResponse, error) {
	out := new(AddCollectEntriesResponse)
	err := c.cc.Invoke(ctx, "/gaucsec.guac.collect_subscriber.schema.ColectSubscriberService/AddCollectEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colectSubscriberServiceClient) GetCollectEntries(ctx context.Context, in *GetCollectEntriesRequest, opts ...grpc.CallOption) (*GetCollectEntriesResponse, error) {
	out := new(GetCollectEntriesResponse)
	err := c.cc.Invoke(ctx, "/gaucsec.guac.collect_subscriber.schema.ColectSubscriberService/GetCollectEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ColectSubscriberServiceServer is the server API for ColectSubscriberService service.
// All implementations must embed UnimplementedColectSubscriberServiceServer
// for forward compatibility
type ColectSubscriberServiceServer interface {
	AddCollectEntries(context.Context, *AddCollectEntriesRequest) (*AddCollectEntriesResponse, error)
	GetCollectEntries(context.Context, *GetCollectEntriesRequest) (*GetCollectEntriesResponse, error)
	mustEmbedUnimplementedColectSubscriberServiceServer()
}

// UnimplementedColectSubscriberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedColectSubscriberServiceServer struct {
}

func (UnimplementedColectSubscriberServiceServer) AddCollectEntries(context.Context, *AddCollectEntriesRequest) (*AddCollectEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollectEntries not implemented")
}
func (UnimplementedColectSubscriberServiceServer) GetCollectEntries(context.Context, *GetCollectEntriesRequest) (*GetCollectEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectEntries not implemented")
}
func (UnimplementedColectSubscriberServiceServer) mustEmbedUnimplementedColectSubscriberServiceServer() {
}

// UnsafeColectSubscriberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColectSubscriberServiceServer will
// result in compilation errors.
type UnsafeColectSubscriberServiceServer interface {
	mustEmbedUnimplementedColectSubscriberServiceServer()
}

func RegisterColectSubscriberServiceServer(s grpc.ServiceRegistrar, srv ColectSubscriberServiceServer) {
	s.RegisterService(&ColectSubscriberService_ServiceDesc, srv)
}

func _ColectSubscriberService_AddCollectEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollectEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColectSubscriberServiceServer).AddCollectEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaucsec.guac.collect_subscriber.schema.ColectSubscriberService/AddCollectEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColectSubscriberServiceServer).AddCollectEntries(ctx, req.(*AddCollectEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColectSubscriberService_GetCollectEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColectSubscriberServiceServer).GetCollectEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaucsec.guac.collect_subscriber.schema.ColectSubscriberService/GetCollectEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColectSubscriberServiceServer).GetCollectEntries(ctx, req.(*GetCollectEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ColectSubscriberService_ServiceDesc is the grpc.ServiceDesc for ColectSubscriberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColectSubscriberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gaucsec.guac.collect_subscriber.schema.ColectSubscriberService",
	HandlerType: (*ColectSubscriberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCollectEntries",
			Handler:    _ColectSubscriberService_AddCollectEntries_Handler,
		},
		{
			MethodName: "GetCollectEntries",
			Handler:    _ColectSubscriberService_GetCollectEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/collectsub/collectsub/collectsub.proto",
}