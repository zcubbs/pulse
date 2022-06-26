// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: pipelines.proto

package pipelines

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

// PipelineStatusClient is the client API for PipelineStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineStatusClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type pipelineStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineStatusClient(cc grpc.ClientConnInterface) PipelineStatusClient {
	return &pipelineStatusClient{cc}
}

func (c *pipelineStatusClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/pipelines.PipelineStatus/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineStatusServer is the server API for PipelineStatus service.
// All implementations must embed UnimplementedPipelineStatusServer
// for forward compatibility
type PipelineStatusServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	mustEmbedUnimplementedPipelineStatusServer()
}

// UnimplementedPipelineStatusServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineStatusServer struct {
}

func (UnimplementedPipelineStatusServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedPipelineStatusServer) mustEmbedUnimplementedPipelineStatusServer() {}

// UnsafePipelineStatusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineStatusServer will
// result in compilation errors.
type UnsafePipelineStatusServer interface {
	mustEmbedUnimplementedPipelineStatusServer()
}

func RegisterPipelineStatusServer(s grpc.ServiceRegistrar, srv PipelineStatusServer) {
	s.RegisterService(&PipelineStatus_ServiceDesc, srv)
}

func _PipelineStatus_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineStatusServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipelines.PipelineStatus/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineStatusServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineStatus_ServiceDesc is the grpc.ServiceDesc for PipelineStatus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineStatus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipelines.PipelineStatus",
	HandlerType: (*PipelineStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _PipelineStatus_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipelines.proto",
}
