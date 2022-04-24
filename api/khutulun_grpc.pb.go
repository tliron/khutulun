// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: khutulun.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConductorClient is the client API for Conductor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConductorClient interface {
	GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Version, error)
	ListHosts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Conductor_ListHostsClient, error)
	AddHost(ctx context.Context, in *HostIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListNamespaces(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Conductor_ListNamespacesClient, error)
	ListPackages(ctx context.Context, in *ListPackages, opts ...grpc.CallOption) (Conductor_ListPackagesClient, error)
	ListPackageFiles(ctx context.Context, in *PackageIdentifier, opts ...grpc.CallOption) (Conductor_ListPackageFilesClient, error)
	GetPackageFiles(ctx context.Context, in *GetPackageFiles, opts ...grpc.CallOption) (Conductor_GetPackageFilesClient, error)
	SetPackageFiles(ctx context.Context, opts ...grpc.CallOption) (Conductor_SetPackageFilesClient, error)
	RemovePackage(ctx context.Context, in *PackageIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeployService(ctx context.Context, in *DeployService, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListResources(ctx context.Context, in *ListResources, opts ...grpc.CallOption) (Conductor_ListResourcesClient, error)
	Interact(ctx context.Context, opts ...grpc.CallOption) (Conductor_InteractClient, error)
}

type conductorClient struct {
	cc grpc.ClientConnInterface
}

func NewConductorClient(cc grpc.ClientConnInterface) ConductorClient {
	return &conductorClient{cc}
}

func (c *conductorClient) GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/khutulun.Conductor/getVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conductorClient) ListHosts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Conductor_ListHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[0], "/khutulun.Conductor/listHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorListHostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Conductor_ListHostsClient interface {
	Recv() (*HostIdentifier, error)
	grpc.ClientStream
}

type conductorListHostsClient struct {
	grpc.ClientStream
}

func (x *conductorListHostsClient) Recv() (*HostIdentifier, error) {
	m := new(HostIdentifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) AddHost(ctx context.Context, in *HostIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/khutulun.Conductor/addHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conductorClient) ListNamespaces(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Conductor_ListNamespacesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[1], "/khutulun.Conductor/listNamespaces", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorListNamespacesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Conductor_ListNamespacesClient interface {
	Recv() (*Namespace, error)
	grpc.ClientStream
}

type conductorListNamespacesClient struct {
	grpc.ClientStream
}

func (x *conductorListNamespacesClient) Recv() (*Namespace, error) {
	m := new(Namespace)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) ListPackages(ctx context.Context, in *ListPackages, opts ...grpc.CallOption) (Conductor_ListPackagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[2], "/khutulun.Conductor/listPackages", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorListPackagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Conductor_ListPackagesClient interface {
	Recv() (*PackageIdentifier, error)
	grpc.ClientStream
}

type conductorListPackagesClient struct {
	grpc.ClientStream
}

func (x *conductorListPackagesClient) Recv() (*PackageIdentifier, error) {
	m := new(PackageIdentifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) ListPackageFiles(ctx context.Context, in *PackageIdentifier, opts ...grpc.CallOption) (Conductor_ListPackageFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[3], "/khutulun.Conductor/listPackageFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorListPackageFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Conductor_ListPackageFilesClient interface {
	Recv() (*PackageFile, error)
	grpc.ClientStream
}

type conductorListPackageFilesClient struct {
	grpc.ClientStream
}

func (x *conductorListPackageFilesClient) Recv() (*PackageFile, error) {
	m := new(PackageFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) GetPackageFiles(ctx context.Context, in *GetPackageFiles, opts ...grpc.CallOption) (Conductor_GetPackageFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[4], "/khutulun.Conductor/getPackageFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorGetPackageFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Conductor_GetPackageFilesClient interface {
	Recv() (*PackageContent, error)
	grpc.ClientStream
}

type conductorGetPackageFilesClient struct {
	grpc.ClientStream
}

func (x *conductorGetPackageFilesClient) Recv() (*PackageContent, error) {
	m := new(PackageContent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) SetPackageFiles(ctx context.Context, opts ...grpc.CallOption) (Conductor_SetPackageFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[5], "/khutulun.Conductor/setPackageFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorSetPackageFilesClient{stream}
	return x, nil
}

type Conductor_SetPackageFilesClient interface {
	Send(*PackageContent) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type conductorSetPackageFilesClient struct {
	grpc.ClientStream
}

func (x *conductorSetPackageFilesClient) Send(m *PackageContent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *conductorSetPackageFilesClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) RemovePackage(ctx context.Context, in *PackageIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/khutulun.Conductor/removePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conductorClient) DeployService(ctx context.Context, in *DeployService, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/khutulun.Conductor/deployService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conductorClient) ListResources(ctx context.Context, in *ListResources, opts ...grpc.CallOption) (Conductor_ListResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[6], "/khutulun.Conductor/listResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorListResourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Conductor_ListResourcesClient interface {
	Recv() (*ResourceIdentifier, error)
	grpc.ClientStream
}

type conductorListResourcesClient struct {
	grpc.ClientStream
}

func (x *conductorListResourcesClient) Recv() (*ResourceIdentifier, error) {
	m := new(ResourceIdentifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *conductorClient) Interact(ctx context.Context, opts ...grpc.CallOption) (Conductor_InteractClient, error) {
	stream, err := c.cc.NewStream(ctx, &Conductor_ServiceDesc.Streams[7], "/khutulun.Conductor/interact", opts...)
	if err != nil {
		return nil, err
	}
	x := &conductorInteractClient{stream}
	return x, nil
}

type Conductor_InteractClient interface {
	Send(*Interaction) error
	Recv() (*Interaction, error)
	grpc.ClientStream
}

type conductorInteractClient struct {
	grpc.ClientStream
}

func (x *conductorInteractClient) Send(m *Interaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *conductorInteractClient) Recv() (*Interaction, error) {
	m := new(Interaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConductorServer is the server API for Conductor service.
// All implementations must embed UnimplementedConductorServer
// for forward compatibility
type ConductorServer interface {
	GetVersion(context.Context, *emptypb.Empty) (*Version, error)
	ListHosts(*emptypb.Empty, Conductor_ListHostsServer) error
	AddHost(context.Context, *HostIdentifier) (*emptypb.Empty, error)
	ListNamespaces(*emptypb.Empty, Conductor_ListNamespacesServer) error
	ListPackages(*ListPackages, Conductor_ListPackagesServer) error
	ListPackageFiles(*PackageIdentifier, Conductor_ListPackageFilesServer) error
	GetPackageFiles(*GetPackageFiles, Conductor_GetPackageFilesServer) error
	SetPackageFiles(Conductor_SetPackageFilesServer) error
	RemovePackage(context.Context, *PackageIdentifier) (*emptypb.Empty, error)
	DeployService(context.Context, *DeployService) (*emptypb.Empty, error)
	ListResources(*ListResources, Conductor_ListResourcesServer) error
	Interact(Conductor_InteractServer) error
	mustEmbedUnimplementedConductorServer()
}

// UnimplementedConductorServer must be embedded to have forward compatible implementations.
type UnimplementedConductorServer struct {
}

func (UnimplementedConductorServer) GetVersion(context.Context, *emptypb.Empty) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedConductorServer) ListHosts(*emptypb.Empty, Conductor_ListHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListHosts not implemented")
}
func (UnimplementedConductorServer) AddHost(context.Context, *HostIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHost not implemented")
}
func (UnimplementedConductorServer) ListNamespaces(*emptypb.Empty, Conductor_ListNamespacesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedConductorServer) ListPackages(*ListPackages, Conductor_ListPackagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPackages not implemented")
}
func (UnimplementedConductorServer) ListPackageFiles(*PackageIdentifier, Conductor_ListPackageFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPackageFiles not implemented")
}
func (UnimplementedConductorServer) GetPackageFiles(*GetPackageFiles, Conductor_GetPackageFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPackageFiles not implemented")
}
func (UnimplementedConductorServer) SetPackageFiles(Conductor_SetPackageFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method SetPackageFiles not implemented")
}
func (UnimplementedConductorServer) RemovePackage(context.Context, *PackageIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePackage not implemented")
}
func (UnimplementedConductorServer) DeployService(context.Context, *DeployService) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployService not implemented")
}
func (UnimplementedConductorServer) ListResources(*ListResources, Conductor_ListResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedConductorServer) Interact(Conductor_InteractServer) error {
	return status.Errorf(codes.Unimplemented, "method Interact not implemented")
}
func (UnimplementedConductorServer) mustEmbedUnimplementedConductorServer() {}

// UnsafeConductorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConductorServer will
// result in compilation errors.
type UnsafeConductorServer interface {
	mustEmbedUnimplementedConductorServer()
}

func RegisterConductorServer(s grpc.ServiceRegistrar, srv ConductorServer) {
	s.RegisterService(&Conductor_ServiceDesc, srv)
}

func _Conductor_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConductorServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khutulun.Conductor/getVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConductorServer).GetVersion(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conductor_ListHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConductorServer).ListHosts(m, &conductorListHostsServer{stream})
}

type Conductor_ListHostsServer interface {
	Send(*HostIdentifier) error
	grpc.ServerStream
}

type conductorListHostsServer struct {
	grpc.ServerStream
}

func (x *conductorListHostsServer) Send(m *HostIdentifier) error {
	return x.ServerStream.SendMsg(m)
}

func _Conductor_AddHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConductorServer).AddHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khutulun.Conductor/addHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConductorServer).AddHost(ctx, req.(*HostIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conductor_ListNamespaces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConductorServer).ListNamespaces(m, &conductorListNamespacesServer{stream})
}

type Conductor_ListNamespacesServer interface {
	Send(*Namespace) error
	grpc.ServerStream
}

type conductorListNamespacesServer struct {
	grpc.ServerStream
}

func (x *conductorListNamespacesServer) Send(m *Namespace) error {
	return x.ServerStream.SendMsg(m)
}

func _Conductor_ListPackages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPackages)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConductorServer).ListPackages(m, &conductorListPackagesServer{stream})
}

type Conductor_ListPackagesServer interface {
	Send(*PackageIdentifier) error
	grpc.ServerStream
}

type conductorListPackagesServer struct {
	grpc.ServerStream
}

func (x *conductorListPackagesServer) Send(m *PackageIdentifier) error {
	return x.ServerStream.SendMsg(m)
}

func _Conductor_ListPackageFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PackageIdentifier)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConductorServer).ListPackageFiles(m, &conductorListPackageFilesServer{stream})
}

type Conductor_ListPackageFilesServer interface {
	Send(*PackageFile) error
	grpc.ServerStream
}

type conductorListPackageFilesServer struct {
	grpc.ServerStream
}

func (x *conductorListPackageFilesServer) Send(m *PackageFile) error {
	return x.ServerStream.SendMsg(m)
}

func _Conductor_GetPackageFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPackageFiles)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConductorServer).GetPackageFiles(m, &conductorGetPackageFilesServer{stream})
}

type Conductor_GetPackageFilesServer interface {
	Send(*PackageContent) error
	grpc.ServerStream
}

type conductorGetPackageFilesServer struct {
	grpc.ServerStream
}

func (x *conductorGetPackageFilesServer) Send(m *PackageContent) error {
	return x.ServerStream.SendMsg(m)
}

func _Conductor_SetPackageFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConductorServer).SetPackageFiles(&conductorSetPackageFilesServer{stream})
}

type Conductor_SetPackageFilesServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*PackageContent, error)
	grpc.ServerStream
}

type conductorSetPackageFilesServer struct {
	grpc.ServerStream
}

func (x *conductorSetPackageFilesServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *conductorSetPackageFilesServer) Recv() (*PackageContent, error) {
	m := new(PackageContent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Conductor_RemovePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConductorServer).RemovePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khutulun.Conductor/removePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConductorServer).RemovePackage(ctx, req.(*PackageIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conductor_DeployService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConductorServer).DeployService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khutulun.Conductor/deployService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConductorServer).DeployService(ctx, req.(*DeployService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conductor_ListResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListResources)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConductorServer).ListResources(m, &conductorListResourcesServer{stream})
}

type Conductor_ListResourcesServer interface {
	Send(*ResourceIdentifier) error
	grpc.ServerStream
}

type conductorListResourcesServer struct {
	grpc.ServerStream
}

func (x *conductorListResourcesServer) Send(m *ResourceIdentifier) error {
	return x.ServerStream.SendMsg(m)
}

func _Conductor_Interact_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConductorServer).Interact(&conductorInteractServer{stream})
}

type Conductor_InteractServer interface {
	Send(*Interaction) error
	Recv() (*Interaction, error)
	grpc.ServerStream
}

type conductorInteractServer struct {
	grpc.ServerStream
}

func (x *conductorInteractServer) Send(m *Interaction) error {
	return x.ServerStream.SendMsg(m)
}

func (x *conductorInteractServer) Recv() (*Interaction, error) {
	m := new(Interaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Conductor_ServiceDesc is the grpc.ServiceDesc for Conductor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conductor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "khutulun.Conductor",
	HandlerType: (*ConductorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getVersion",
			Handler:    _Conductor_GetVersion_Handler,
		},
		{
			MethodName: "addHost",
			Handler:    _Conductor_AddHost_Handler,
		},
		{
			MethodName: "removePackage",
			Handler:    _Conductor_RemovePackage_Handler,
		},
		{
			MethodName: "deployService",
			Handler:    _Conductor_DeployService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listHosts",
			Handler:       _Conductor_ListHosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listNamespaces",
			Handler:       _Conductor_ListNamespaces_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listPackages",
			Handler:       _Conductor_ListPackages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listPackageFiles",
			Handler:       _Conductor_ListPackageFiles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "getPackageFiles",
			Handler:       _Conductor_GetPackageFiles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "setPackageFiles",
			Handler:       _Conductor_SetPackageFiles_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "listResources",
			Handler:       _Conductor_ListResources_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "interact",
			Handler:       _Conductor_Interact_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "khutulun.proto",
}

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	Instantiate(ctx context.Context, in *Config, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Interact(ctx context.Context, opts ...grpc.CallOption) (Plugin_InteractClient, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Instantiate(ctx context.Context, in *Config, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/khutulun.Plugin/instantiate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Interact(ctx context.Context, opts ...grpc.CallOption) (Plugin_InteractClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugin_ServiceDesc.Streams[0], "/khutulun.Plugin/interact", opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginInteractClient{stream}
	return x, nil
}

type Plugin_InteractClient interface {
	Send(*Interaction) error
	Recv() (*Interaction, error)
	grpc.ClientStream
}

type pluginInteractClient struct {
	grpc.ClientStream
}

func (x *pluginInteractClient) Send(m *Interaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginInteractClient) Recv() (*Interaction, error) {
	m := new(Interaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	Instantiate(context.Context, *Config) (*emptypb.Empty, error)
	Interact(Plugin_InteractServer) error
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) Instantiate(context.Context, *Config) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Instantiate not implemented")
}
func (UnimplementedPluginServer) Interact(Plugin_InteractServer) error {
	return status.Errorf(codes.Unimplemented, "method Interact not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_Instantiate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Instantiate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/khutulun.Plugin/instantiate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Instantiate(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Interact_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginServer).Interact(&pluginInteractServer{stream})
}

type Plugin_InteractServer interface {
	Send(*Interaction) error
	Recv() (*Interaction, error)
	grpc.ServerStream
}

type pluginInteractServer struct {
	grpc.ServerStream
}

func (x *pluginInteractServer) Send(m *Interaction) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginInteractServer) Recv() (*Interaction, error) {
	m := new(Interaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "khutulun.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "instantiate",
			Handler:    _Plugin_Instantiate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "interact",
			Handler:       _Plugin_Interact_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "khutulun.proto",
}
