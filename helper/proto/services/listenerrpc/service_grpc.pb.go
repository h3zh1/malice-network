// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: services/listenerrpc/service.proto

package listenerrpc

import (
	context "context"
	clientpb "github.com/chainreactors/malice-network/helper/proto/client/clientpb"
	implantpb "github.com/chainreactors/malice-network/helper/proto/implant/implantpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ListenerRPC_Register_FullMethodName         = "/listenerrpc.ListenerRPC/Register"
	ListenerRPC_SysInfo_FullMethodName          = "/listenerrpc.ListenerRPC/SysInfo"
	ListenerRPC_Checkin_FullMethodName          = "/listenerrpc.ListenerRPC/Checkin"
	ListenerRPC_InitBindSession_FullMethodName  = "/listenerrpc.ListenerRPC/InitBindSession"
	ListenerRPC_RegisterListener_FullMethodName = "/listenerrpc.ListenerRPC/RegisterListener"
	ListenerRPC_RegisterPipeline_FullMethodName = "/listenerrpc.ListenerRPC/RegisterPipeline"
	ListenerRPC_RegisterWebsite_FullMethodName  = "/listenerrpc.ListenerRPC/RegisterWebsite"
	ListenerRPC_StartPipeline_FullMethodName    = "/listenerrpc.ListenerRPC/StartPipeline"
	ListenerRPC_StopPipeline_FullMethodName     = "/listenerrpc.ListenerRPC/StopPipeline"
	ListenerRPC_DeletePipeline_FullMethodName   = "/listenerrpc.ListenerRPC/DeletePipeline"
	ListenerRPC_ListPipelines_FullMethodName    = "/listenerrpc.ListenerRPC/ListPipelines"
	ListenerRPC_StartWebsite_FullMethodName     = "/listenerrpc.ListenerRPC/StartWebsite"
	ListenerRPC_StopWebsite_FullMethodName      = "/listenerrpc.ListenerRPC/StopWebsite"
	ListenerRPC_UploadWebsite_FullMethodName    = "/listenerrpc.ListenerRPC/UploadWebsite"
	ListenerRPC_ListWebsites_FullMethodName     = "/listenerrpc.ListenerRPC/ListWebsites"
	ListenerRPC_GetArtifact_FullMethodName      = "/listenerrpc.ListenerRPC/GetArtifact"
	ListenerRPC_SpiteStream_FullMethodName      = "/listenerrpc.ListenerRPC/SpiteStream"
	ListenerRPC_JobStream_FullMethodName        = "/listenerrpc.ListenerRPC/JobStream"
)

// ListenerRPCClient is the client API for ListenerRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListenerRPCClient interface {
	// implant
	Register(ctx context.Context, in *clientpb.RegisterSession, opts ...grpc.CallOption) (*clientpb.Empty, error)
	SysInfo(ctx context.Context, in *implantpb.SysInfo, opts ...grpc.CallOption) (*clientpb.Empty, error)
	Checkin(ctx context.Context, in *implantpb.Ping, opts ...grpc.CallOption) (*clientpb.Empty, error)
	InitBindSession(ctx context.Context, in *implantpb.Request, opts ...grpc.CallOption) (*clientpb.Empty, error)
	// pipeline
	RegisterListener(ctx context.Context, in *clientpb.RegisterListener, opts ...grpc.CallOption) (*clientpb.Empty, error)
	RegisterPipeline(ctx context.Context, in *clientpb.Pipeline, opts ...grpc.CallOption) (*clientpb.Empty, error)
	RegisterWebsite(ctx context.Context, in *clientpb.Pipeline, opts ...grpc.CallOption) (*clientpb.WebsiteResponse, error)
	StartPipeline(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error)
	StopPipeline(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error)
	DeletePipeline(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error)
	ListPipelines(ctx context.Context, in *clientpb.Listener, opts ...grpc.CallOption) (*clientpb.Pipelines, error)
	StartWebsite(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error)
	StopWebsite(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error)
	UploadWebsite(ctx context.Context, in *clientpb.Website, opts ...grpc.CallOption) (*clientpb.Empty, error)
	ListWebsites(ctx context.Context, in *clientpb.Listener, opts ...grpc.CallOption) (*clientpb.Pipelines, error)
	GetArtifact(ctx context.Context, in *clientpb.Builder, opts ...grpc.CallOption) (*clientpb.Builder, error)
	// rpc ListenerCtrl(clientpb.CtrlStatus) returns (stream clientpb.CtrlPipeline);
	SpiteStream(ctx context.Context, opts ...grpc.CallOption) (ListenerRPC_SpiteStreamClient, error)
	JobStream(ctx context.Context, opts ...grpc.CallOption) (ListenerRPC_JobStreamClient, error)
}

type listenerRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewListenerRPCClient(cc grpc.ClientConnInterface) ListenerRPCClient {
	return &listenerRPCClient{cc}
}

func (c *listenerRPCClient) Register(ctx context.Context, in *clientpb.RegisterSession, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) SysInfo(ctx context.Context, in *implantpb.SysInfo, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_SysInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) Checkin(ctx context.Context, in *implantpb.Ping, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_Checkin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) InitBindSession(ctx context.Context, in *implantpb.Request, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_InitBindSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) RegisterListener(ctx context.Context, in *clientpb.RegisterListener, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_RegisterListener_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) RegisterPipeline(ctx context.Context, in *clientpb.Pipeline, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_RegisterPipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) RegisterWebsite(ctx context.Context, in *clientpb.Pipeline, opts ...grpc.CallOption) (*clientpb.WebsiteResponse, error) {
	out := new(clientpb.WebsiteResponse)
	err := c.cc.Invoke(ctx, ListenerRPC_RegisterWebsite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) StartPipeline(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_StartPipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) StopPipeline(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_StopPipeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) DeletePipeline(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_DeletePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) ListPipelines(ctx context.Context, in *clientpb.Listener, opts ...grpc.CallOption) (*clientpb.Pipelines, error) {
	out := new(clientpb.Pipelines)
	err := c.cc.Invoke(ctx, ListenerRPC_ListPipelines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) StartWebsite(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_StartWebsite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) StopWebsite(ctx context.Context, in *clientpb.CtrlPipeline, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_StopWebsite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) UploadWebsite(ctx context.Context, in *clientpb.Website, opts ...grpc.CallOption) (*clientpb.Empty, error) {
	out := new(clientpb.Empty)
	err := c.cc.Invoke(ctx, ListenerRPC_UploadWebsite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) ListWebsites(ctx context.Context, in *clientpb.Listener, opts ...grpc.CallOption) (*clientpb.Pipelines, error) {
	out := new(clientpb.Pipelines)
	err := c.cc.Invoke(ctx, ListenerRPC_ListWebsites_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) GetArtifact(ctx context.Context, in *clientpb.Builder, opts ...grpc.CallOption) (*clientpb.Builder, error) {
	out := new(clientpb.Builder)
	err := c.cc.Invoke(ctx, ListenerRPC_GetArtifact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerRPCClient) SpiteStream(ctx context.Context, opts ...grpc.CallOption) (ListenerRPC_SpiteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ListenerRPC_ServiceDesc.Streams[0], ListenerRPC_SpiteStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerRPCSpiteStreamClient{stream}
	return x, nil
}

type ListenerRPC_SpiteStreamClient interface {
	Send(*clientpb.SpiteResponse) error
	Recv() (*clientpb.SpiteRequest, error)
	grpc.ClientStream
}

type listenerRPCSpiteStreamClient struct {
	grpc.ClientStream
}

func (x *listenerRPCSpiteStreamClient) Send(m *clientpb.SpiteResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerRPCSpiteStreamClient) Recv() (*clientpb.SpiteRequest, error) {
	m := new(clientpb.SpiteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerRPCClient) JobStream(ctx context.Context, opts ...grpc.CallOption) (ListenerRPC_JobStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ListenerRPC_ServiceDesc.Streams[1], ListenerRPC_JobStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerRPCJobStreamClient{stream}
	return x, nil
}

type ListenerRPC_JobStreamClient interface {
	Send(*clientpb.JobStatus) error
	Recv() (*clientpb.JobCtrl, error)
	grpc.ClientStream
}

type listenerRPCJobStreamClient struct {
	grpc.ClientStream
}

func (x *listenerRPCJobStreamClient) Send(m *clientpb.JobStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerRPCJobStreamClient) Recv() (*clientpb.JobCtrl, error) {
	m := new(clientpb.JobCtrl)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ListenerRPCServer is the server API for ListenerRPC service.
// All implementations must embed UnimplementedListenerRPCServer
// for forward compatibility
type ListenerRPCServer interface {
	// implant
	Register(context.Context, *clientpb.RegisterSession) (*clientpb.Empty, error)
	SysInfo(context.Context, *implantpb.SysInfo) (*clientpb.Empty, error)
	Checkin(context.Context, *implantpb.Ping) (*clientpb.Empty, error)
	InitBindSession(context.Context, *implantpb.Request) (*clientpb.Empty, error)
	// pipeline
	RegisterListener(context.Context, *clientpb.RegisterListener) (*clientpb.Empty, error)
	RegisterPipeline(context.Context, *clientpb.Pipeline) (*clientpb.Empty, error)
	RegisterWebsite(context.Context, *clientpb.Pipeline) (*clientpb.WebsiteResponse, error)
	StartPipeline(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error)
	StopPipeline(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error)
	DeletePipeline(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error)
	ListPipelines(context.Context, *clientpb.Listener) (*clientpb.Pipelines, error)
	StartWebsite(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error)
	StopWebsite(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error)
	UploadWebsite(context.Context, *clientpb.Website) (*clientpb.Empty, error)
	ListWebsites(context.Context, *clientpb.Listener) (*clientpb.Pipelines, error)
	GetArtifact(context.Context, *clientpb.Builder) (*clientpb.Builder, error)
	// rpc ListenerCtrl(clientpb.CtrlStatus) returns (stream clientpb.CtrlPipeline);
	SpiteStream(ListenerRPC_SpiteStreamServer) error
	JobStream(ListenerRPC_JobStreamServer) error
	mustEmbedUnimplementedListenerRPCServer()
}

// UnimplementedListenerRPCServer must be embedded to have forward compatible implementations.
type UnimplementedListenerRPCServer struct {
}

func (UnimplementedListenerRPCServer) Register(context.Context, *clientpb.RegisterSession) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedListenerRPCServer) SysInfo(context.Context, *implantpb.SysInfo) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysInfo not implemented")
}
func (UnimplementedListenerRPCServer) Checkin(context.Context, *implantpb.Ping) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkin not implemented")
}
func (UnimplementedListenerRPCServer) InitBindSession(context.Context, *implantpb.Request) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitBindSession not implemented")
}
func (UnimplementedListenerRPCServer) RegisterListener(context.Context, *clientpb.RegisterListener) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterListener not implemented")
}
func (UnimplementedListenerRPCServer) RegisterPipeline(context.Context, *clientpb.Pipeline) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPipeline not implemented")
}
func (UnimplementedListenerRPCServer) RegisterWebsite(context.Context, *clientpb.Pipeline) (*clientpb.WebsiteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWebsite not implemented")
}
func (UnimplementedListenerRPCServer) StartPipeline(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPipeline not implemented")
}
func (UnimplementedListenerRPCServer) StopPipeline(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPipeline not implemented")
}
func (UnimplementedListenerRPCServer) DeletePipeline(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedListenerRPCServer) ListPipelines(context.Context, *clientpb.Listener) (*clientpb.Pipelines, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelines not implemented")
}
func (UnimplementedListenerRPCServer) StartWebsite(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartWebsite not implemented")
}
func (UnimplementedListenerRPCServer) StopWebsite(context.Context, *clientpb.CtrlPipeline) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopWebsite not implemented")
}
func (UnimplementedListenerRPCServer) UploadWebsite(context.Context, *clientpb.Website) (*clientpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadWebsite not implemented")
}
func (UnimplementedListenerRPCServer) ListWebsites(context.Context, *clientpb.Listener) (*clientpb.Pipelines, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWebsites not implemented")
}
func (UnimplementedListenerRPCServer) GetArtifact(context.Context, *clientpb.Builder) (*clientpb.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtifact not implemented")
}
func (UnimplementedListenerRPCServer) SpiteStream(ListenerRPC_SpiteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SpiteStream not implemented")
}
func (UnimplementedListenerRPCServer) JobStream(ListenerRPC_JobStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method JobStream not implemented")
}
func (UnimplementedListenerRPCServer) mustEmbedUnimplementedListenerRPCServer() {}

// UnsafeListenerRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListenerRPCServer will
// result in compilation errors.
type UnsafeListenerRPCServer interface {
	mustEmbedUnimplementedListenerRPCServer()
}

func RegisterListenerRPCServer(s grpc.ServiceRegistrar, srv ListenerRPCServer) {
	s.RegisterService(&ListenerRPC_ServiceDesc, srv)
}

func _ListenerRPC_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.RegisterSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).Register(ctx, req.(*clientpb.RegisterSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_SysInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(implantpb.SysInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).SysInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_SysInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).SysInfo(ctx, req.(*implantpb.SysInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_Checkin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(implantpb.Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).Checkin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_Checkin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).Checkin(ctx, req.(*implantpb.Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_InitBindSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(implantpb.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).InitBindSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_InitBindSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).InitBindSession(ctx, req.(*implantpb.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_RegisterListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.RegisterListener)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).RegisterListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_RegisterListener_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).RegisterListener(ctx, req.(*clientpb.RegisterListener))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_RegisterPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).RegisterPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_RegisterPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).RegisterPipeline(ctx, req.(*clientpb.Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_RegisterWebsite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).RegisterWebsite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_RegisterWebsite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).RegisterWebsite(ctx, req.(*clientpb.Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_StartPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.CtrlPipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).StartPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_StartPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).StartPipeline(ctx, req.(*clientpb.CtrlPipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_StopPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.CtrlPipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).StopPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_StopPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).StopPipeline(ctx, req.(*clientpb.CtrlPipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.CtrlPipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_DeletePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).DeletePipeline(ctx, req.(*clientpb.CtrlPipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.Listener)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_ListPipelines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).ListPipelines(ctx, req.(*clientpb.Listener))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_StartWebsite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.CtrlPipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).StartWebsite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_StartWebsite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).StartWebsite(ctx, req.(*clientpb.CtrlPipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_StopWebsite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.CtrlPipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).StopWebsite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_StopWebsite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).StopWebsite(ctx, req.(*clientpb.CtrlPipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_UploadWebsite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.Website)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).UploadWebsite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_UploadWebsite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).UploadWebsite(ctx, req.(*clientpb.Website))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_ListWebsites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.Listener)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).ListWebsites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_ListWebsites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).ListWebsites(ctx, req.(*clientpb.Listener))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_GetArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clientpb.Builder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerRPCServer).GetArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerRPC_GetArtifact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerRPCServer).GetArtifact(ctx, req.(*clientpb.Builder))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerRPC_SpiteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerRPCServer).SpiteStream(&listenerRPCSpiteStreamServer{stream})
}

type ListenerRPC_SpiteStreamServer interface {
	Send(*clientpb.SpiteRequest) error
	Recv() (*clientpb.SpiteResponse, error)
	grpc.ServerStream
}

type listenerRPCSpiteStreamServer struct {
	grpc.ServerStream
}

func (x *listenerRPCSpiteStreamServer) Send(m *clientpb.SpiteRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerRPCSpiteStreamServer) Recv() (*clientpb.SpiteResponse, error) {
	m := new(clientpb.SpiteResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerRPC_JobStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerRPCServer).JobStream(&listenerRPCJobStreamServer{stream})
}

type ListenerRPC_JobStreamServer interface {
	Send(*clientpb.JobCtrl) error
	Recv() (*clientpb.JobStatus, error)
	grpc.ServerStream
}

type listenerRPCJobStreamServer struct {
	grpc.ServerStream
}

func (x *listenerRPCJobStreamServer) Send(m *clientpb.JobCtrl) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerRPCJobStreamServer) Recv() (*clientpb.JobStatus, error) {
	m := new(clientpb.JobStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ListenerRPC_ServiceDesc is the grpc.ServiceDesc for ListenerRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListenerRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listenerrpc.ListenerRPC",
	HandlerType: (*ListenerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ListenerRPC_Register_Handler,
		},
		{
			MethodName: "SysInfo",
			Handler:    _ListenerRPC_SysInfo_Handler,
		},
		{
			MethodName: "Checkin",
			Handler:    _ListenerRPC_Checkin_Handler,
		},
		{
			MethodName: "InitBindSession",
			Handler:    _ListenerRPC_InitBindSession_Handler,
		},
		{
			MethodName: "RegisterListener",
			Handler:    _ListenerRPC_RegisterListener_Handler,
		},
		{
			MethodName: "RegisterPipeline",
			Handler:    _ListenerRPC_RegisterPipeline_Handler,
		},
		{
			MethodName: "RegisterWebsite",
			Handler:    _ListenerRPC_RegisterWebsite_Handler,
		},
		{
			MethodName: "StartPipeline",
			Handler:    _ListenerRPC_StartPipeline_Handler,
		},
		{
			MethodName: "StopPipeline",
			Handler:    _ListenerRPC_StopPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _ListenerRPC_DeletePipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _ListenerRPC_ListPipelines_Handler,
		},
		{
			MethodName: "StartWebsite",
			Handler:    _ListenerRPC_StartWebsite_Handler,
		},
		{
			MethodName: "StopWebsite",
			Handler:    _ListenerRPC_StopWebsite_Handler,
		},
		{
			MethodName: "UploadWebsite",
			Handler:    _ListenerRPC_UploadWebsite_Handler,
		},
		{
			MethodName: "ListWebsites",
			Handler:    _ListenerRPC_ListWebsites_Handler,
		},
		{
			MethodName: "GetArtifact",
			Handler:    _ListenerRPC_GetArtifact_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SpiteStream",
			Handler:       _ListenerRPC_SpiteStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "JobStream",
			Handler:       _ListenerRPC_JobStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services/listenerrpc/service.proto",
}
