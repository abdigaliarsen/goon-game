// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pkg/proto/wikipedia/wikipedia.proto

package wikipedia

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WikipediaService_SetLanguage_FullMethodName        = "/auth.WikipediaService/SetLanguage"
	WikipediaService_GetLanguageUpdates_FullMethodName = "/auth.WikipediaService/GetLanguageUpdates"
	WikipediaService_GetStats_FullMethodName           = "/auth.WikipediaService/GetStats"
)

// WikipediaServiceClient is the client API for WikipediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WikipediaServiceClient interface {
	SetLanguage(ctx context.Context, in *SetLanguageRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetLanguageUpdates(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetLanguageUpdatesResponse, error)
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
}

type wikipediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWikipediaServiceClient(cc grpc.ClientConnInterface) WikipediaServiceClient {
	return &wikipediaServiceClient{cc}
}

func (c *wikipediaServiceClient) SetLanguage(ctx context.Context, in *SetLanguageRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, WikipediaService_SetLanguage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wikipediaServiceClient) GetLanguageUpdates(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetLanguageUpdatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLanguageUpdatesResponse)
	err := c.cc.Invoke(ctx, WikipediaService_GetLanguageUpdates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wikipediaServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, WikipediaService_GetStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WikipediaServiceServer is the server API for WikipediaService service.
// All implementations must embed UnimplementedWikipediaServiceServer
// for forward compatibility.
type WikipediaServiceServer interface {
	SetLanguage(context.Context, *SetLanguageRequest) (*EmptyResponse, error)
	GetLanguageUpdates(context.Context, *EmptyRequest) (*GetLanguageUpdatesResponse, error)
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	mustEmbedUnimplementedWikipediaServiceServer()
}

// UnimplementedWikipediaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWikipediaServiceServer struct{}

func (UnimplementedWikipediaServiceServer) SetLanguage(context.Context, *SetLanguageRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLanguage not implemented")
}
func (UnimplementedWikipediaServiceServer) GetLanguageUpdates(context.Context, *EmptyRequest) (*GetLanguageUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguageUpdates not implemented")
}
func (UnimplementedWikipediaServiceServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedWikipediaServiceServer) mustEmbedUnimplementedWikipediaServiceServer() {}
func (UnimplementedWikipediaServiceServer) testEmbeddedByValue()                          {}

// UnsafeWikipediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WikipediaServiceServer will
// result in compilation errors.
type UnsafeWikipediaServiceServer interface {
	mustEmbedUnimplementedWikipediaServiceServer()
}

func RegisterWikipediaServiceServer(s grpc.ServiceRegistrar, srv WikipediaServiceServer) {
	// If the following call pancis, it indicates UnimplementedWikipediaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WikipediaService_ServiceDesc, srv)
}

func _WikipediaService_SetLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WikipediaServiceServer).SetLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WikipediaService_SetLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WikipediaServiceServer).SetLanguage(ctx, req.(*SetLanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WikipediaService_GetLanguageUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WikipediaServiceServer).GetLanguageUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WikipediaService_GetLanguageUpdates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WikipediaServiceServer).GetLanguageUpdates(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WikipediaService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WikipediaServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WikipediaService_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WikipediaServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WikipediaService_ServiceDesc is the grpc.ServiceDesc for WikipediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WikipediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.WikipediaService",
	HandlerType: (*WikipediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLanguage",
			Handler:    _WikipediaService_SetLanguage_Handler,
		},
		{
			MethodName: "GetLanguageUpdates",
			Handler:    _WikipediaService_GetLanguageUpdates_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _WikipediaService_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/wikipedia/wikipedia.proto",
}
