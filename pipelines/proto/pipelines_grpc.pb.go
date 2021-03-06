// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: pipelines.proto

package proto

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
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (PipelineStatus_GetStatusClient, error)
}

type pipelineStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineStatusClient(cc grpc.ClientConnInterface) PipelineStatusClient {
	return &pipelineStatusClient{cc}
}

func (c *pipelineStatusClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (PipelineStatus_GetStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &PipelineStatus_ServiceDesc.Streams[0], "/pipelines.PipelineStatus/GetStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineStatusGetStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PipelineStatus_GetStatusClient interface {
	Recv() (*GetStatusResponse, error)
	grpc.ClientStream
}

type pipelineStatusGetStatusClient struct {
	grpc.ClientStream
}

func (x *pipelineStatusGetStatusClient) Recv() (*GetStatusResponse, error) {
	m := new(GetStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipelineStatusServer is the server API for PipelineStatus service.
// All implementations must embed UnimplementedPipelineStatusServer
// for forward compatibility
type PipelineStatusServer interface {
	GetStatus(*GetStatusRequest, PipelineStatus_GetStatusServer) error
	mustEmbedUnimplementedPipelineStatusServer()
}

// UnimplementedPipelineStatusServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineStatusServer struct {
}

func (UnimplementedPipelineStatusServer) GetStatus(*GetStatusRequest, PipelineStatus_GetStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
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

func _PipelineStatus_GetStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PipelineStatusServer).GetStatus(m, &pipelineStatusGetStatusServer{stream})
}

type PipelineStatus_GetStatusServer interface {
	Send(*GetStatusResponse) error
	grpc.ServerStream
}

type pipelineStatusGetStatusServer struct {
	grpc.ServerStream
}

func (x *pipelineStatusGetStatusServer) Send(m *GetStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PipelineStatus_ServiceDesc is the grpc.ServiceDesc for PipelineStatus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineStatus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipelines.PipelineStatus",
	HandlerType: (*PipelineStatusServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatus",
			Handler:       _PipelineStatus_GetStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pipelines.proto",
}
