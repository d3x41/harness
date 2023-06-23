// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: repo.proto

package rpc

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

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	CreateRepository(ctx context.Context, opts ...grpc.CallOption) (RepositoryService_CreateRepositoryClient, error)
	GetTreeNode(ctx context.Context, in *GetTreeNodeRequest, opts ...grpc.CallOption) (*GetTreeNodeResponse, error)
	ListTreeNodes(ctx context.Context, in *ListTreeNodesRequest, opts ...grpc.CallOption) (RepositoryService_ListTreeNodesClient, error)
	GetSubmodule(ctx context.Context, in *GetSubmoduleRequest, opts ...grpc.CallOption) (*GetSubmoduleResponse, error)
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (RepositoryService_GetBlobClient, error)
	ListCommits(ctx context.Context, in *ListCommitsRequest, opts ...grpc.CallOption) (RepositoryService_ListCommitsClient, error)
	GetCommit(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*GetCommitResponse, error)
	GetCommitDivergences(ctx context.Context, in *GetCommitDivergencesRequest, opts ...grpc.CallOption) (*GetCommitDivergencesResponse, error)
	DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*DeleteRepositoryResponse, error)
	SyncRepository(ctx context.Context, in *SyncRepositoryRequest, opts ...grpc.CallOption) (*SyncRepositoryResponse, error)
	HashRepository(ctx context.Context, in *HashRepositoryRequest, opts ...grpc.CallOption) (*HashRepositoryResponse, error)
	MergeBase(ctx context.Context, in *MergeBaseRequest, opts ...grpc.CallOption) (*MergeBaseResponse, error)
}

type repositoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryServiceClient(cc grpc.ClientConnInterface) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) CreateRepository(ctx context.Context, opts ...grpc.CallOption) (RepositoryService_CreateRepositoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &RepositoryService_ServiceDesc.Streams[0], "/rpc.RepositoryService/CreateRepository", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryServiceCreateRepositoryClient{stream}
	return x, nil
}

type RepositoryService_CreateRepositoryClient interface {
	Send(*CreateRepositoryRequest) error
	CloseAndRecv() (*CreateRepositoryResponse, error)
	grpc.ClientStream
}

type repositoryServiceCreateRepositoryClient struct {
	grpc.ClientStream
}

func (x *repositoryServiceCreateRepositoryClient) Send(m *CreateRepositoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *repositoryServiceCreateRepositoryClient) CloseAndRecv() (*CreateRepositoryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateRepositoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repositoryServiceClient) GetTreeNode(ctx context.Context, in *GetTreeNodeRequest, opts ...grpc.CallOption) (*GetTreeNodeResponse, error) {
	out := new(GetTreeNodeResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/GetTreeNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) ListTreeNodes(ctx context.Context, in *ListTreeNodesRequest, opts ...grpc.CallOption) (RepositoryService_ListTreeNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RepositoryService_ServiceDesc.Streams[1], "/rpc.RepositoryService/ListTreeNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryServiceListTreeNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RepositoryService_ListTreeNodesClient interface {
	Recv() (*ListTreeNodesResponse, error)
	grpc.ClientStream
}

type repositoryServiceListTreeNodesClient struct {
	grpc.ClientStream
}

func (x *repositoryServiceListTreeNodesClient) Recv() (*ListTreeNodesResponse, error) {
	m := new(ListTreeNodesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repositoryServiceClient) GetSubmodule(ctx context.Context, in *GetSubmoduleRequest, opts ...grpc.CallOption) (*GetSubmoduleResponse, error) {
	out := new(GetSubmoduleResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/GetSubmodule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (RepositoryService_GetBlobClient, error) {
	stream, err := c.cc.NewStream(ctx, &RepositoryService_ServiceDesc.Streams[2], "/rpc.RepositoryService/GetBlob", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryServiceGetBlobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RepositoryService_GetBlobClient interface {
	Recv() (*GetBlobResponse, error)
	grpc.ClientStream
}

type repositoryServiceGetBlobClient struct {
	grpc.ClientStream
}

func (x *repositoryServiceGetBlobClient) Recv() (*GetBlobResponse, error) {
	m := new(GetBlobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repositoryServiceClient) ListCommits(ctx context.Context, in *ListCommitsRequest, opts ...grpc.CallOption) (RepositoryService_ListCommitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RepositoryService_ServiceDesc.Streams[3], "/rpc.RepositoryService/ListCommits", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryServiceListCommitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RepositoryService_ListCommitsClient interface {
	Recv() (*ListCommitsResponse, error)
	grpc.ClientStream
}

type repositoryServiceListCommitsClient struct {
	grpc.ClientStream
}

func (x *repositoryServiceListCommitsClient) Recv() (*ListCommitsResponse, error) {
	m := new(ListCommitsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repositoryServiceClient) GetCommit(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*GetCommitResponse, error) {
	out := new(GetCommitResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/GetCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) GetCommitDivergences(ctx context.Context, in *GetCommitDivergencesRequest, opts ...grpc.CallOption) (*GetCommitDivergencesResponse, error) {
	out := new(GetCommitDivergencesResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/GetCommitDivergences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*DeleteRepositoryResponse, error) {
	out := new(DeleteRepositoryResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/DeleteRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) SyncRepository(ctx context.Context, in *SyncRepositoryRequest, opts ...grpc.CallOption) (*SyncRepositoryResponse, error) {
	out := new(SyncRepositoryResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/SyncRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) HashRepository(ctx context.Context, in *HashRepositoryRequest, opts ...grpc.CallOption) (*HashRepositoryResponse, error) {
	out := new(HashRepositoryResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/HashRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) MergeBase(ctx context.Context, in *MergeBaseRequest, opts ...grpc.CallOption) (*MergeBaseResponse, error) {
	out := new(MergeBaseResponse)
	err := c.cc.Invoke(ctx, "/rpc.RepositoryService/MergeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
// All implementations must embed UnimplementedRepositoryServiceServer
// for forward compatibility
type RepositoryServiceServer interface {
	CreateRepository(RepositoryService_CreateRepositoryServer) error
	GetTreeNode(context.Context, *GetTreeNodeRequest) (*GetTreeNodeResponse, error)
	ListTreeNodes(*ListTreeNodesRequest, RepositoryService_ListTreeNodesServer) error
	GetSubmodule(context.Context, *GetSubmoduleRequest) (*GetSubmoduleResponse, error)
	GetBlob(*GetBlobRequest, RepositoryService_GetBlobServer) error
	ListCommits(*ListCommitsRequest, RepositoryService_ListCommitsServer) error
	GetCommit(context.Context, *GetCommitRequest) (*GetCommitResponse, error)
	GetCommitDivergences(context.Context, *GetCommitDivergencesRequest) (*GetCommitDivergencesResponse, error)
	DeleteRepository(context.Context, *DeleteRepositoryRequest) (*DeleteRepositoryResponse, error)
	SyncRepository(context.Context, *SyncRepositoryRequest) (*SyncRepositoryResponse, error)
	HashRepository(context.Context, *HashRepositoryRequest) (*HashRepositoryResponse, error)
	MergeBase(context.Context, *MergeBaseRequest) (*MergeBaseResponse, error)
	mustEmbedUnimplementedRepositoryServiceServer()
}

// UnimplementedRepositoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRepositoryServiceServer struct {
}

func (UnimplementedRepositoryServiceServer) CreateRepository(RepositoryService_CreateRepositoryServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateRepository not implemented")
}
func (UnimplementedRepositoryServiceServer) GetTreeNode(context.Context, *GetTreeNodeRequest) (*GetTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTreeNode not implemented")
}
func (UnimplementedRepositoryServiceServer) ListTreeNodes(*ListTreeNodesRequest, RepositoryService_ListTreeNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTreeNodes not implemented")
}
func (UnimplementedRepositoryServiceServer) GetSubmodule(context.Context, *GetSubmoduleRequest) (*GetSubmoduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmodule not implemented")
}
func (UnimplementedRepositoryServiceServer) GetBlob(*GetBlobRequest, RepositoryService_GetBlobServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlob not implemented")
}
func (UnimplementedRepositoryServiceServer) ListCommits(*ListCommitsRequest, RepositoryService_ListCommitsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCommits not implemented")
}
func (UnimplementedRepositoryServiceServer) GetCommit(context.Context, *GetCommitRequest) (*GetCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommit not implemented")
}
func (UnimplementedRepositoryServiceServer) GetCommitDivergences(context.Context, *GetCommitDivergencesRequest) (*GetCommitDivergencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitDivergences not implemented")
}
func (UnimplementedRepositoryServiceServer) DeleteRepository(context.Context, *DeleteRepositoryRequest) (*DeleteRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepository not implemented")
}
func (UnimplementedRepositoryServiceServer) SyncRepository(context.Context, *SyncRepositoryRequest) (*SyncRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRepository not implemented")
}
func (UnimplementedRepositoryServiceServer) HashRepository(context.Context, *HashRepositoryRequest) (*HashRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashRepository not implemented")
}
func (UnimplementedRepositoryServiceServer) MergeBase(context.Context, *MergeBaseRequest) (*MergeBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeBase not implemented")
}
func (UnimplementedRepositoryServiceServer) mustEmbedUnimplementedRepositoryServiceServer() {}

// UnsafeRepositoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServiceServer will
// result in compilation errors.
type UnsafeRepositoryServiceServer interface {
	mustEmbedUnimplementedRepositoryServiceServer()
}

func RegisterRepositoryServiceServer(s grpc.ServiceRegistrar, srv RepositoryServiceServer) {
	s.RegisterService(&RepositoryService_ServiceDesc, srv)
}

func _RepositoryService_CreateRepository_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RepositoryServiceServer).CreateRepository(&repositoryServiceCreateRepositoryServer{stream})
}

type RepositoryService_CreateRepositoryServer interface {
	SendAndClose(*CreateRepositoryResponse) error
	Recv() (*CreateRepositoryRequest, error)
	grpc.ServerStream
}

type repositoryServiceCreateRepositoryServer struct {
	grpc.ServerStream
}

func (x *repositoryServiceCreateRepositoryServer) SendAndClose(m *CreateRepositoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *repositoryServiceCreateRepositoryServer) Recv() (*CreateRepositoryRequest, error) {
	m := new(CreateRepositoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RepositoryService_GetTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GetTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/GetTreeNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GetTreeNode(ctx, req.(*GetTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_ListTreeNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTreeNodesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepositoryServiceServer).ListTreeNodes(m, &repositoryServiceListTreeNodesServer{stream})
}

type RepositoryService_ListTreeNodesServer interface {
	Send(*ListTreeNodesResponse) error
	grpc.ServerStream
}

type repositoryServiceListTreeNodesServer struct {
	grpc.ServerStream
}

func (x *repositoryServiceListTreeNodesServer) Send(m *ListTreeNodesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RepositoryService_GetSubmodule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmoduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GetSubmodule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/GetSubmodule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GetSubmodule(ctx, req.(*GetSubmoduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_GetBlob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepositoryServiceServer).GetBlob(m, &repositoryServiceGetBlobServer{stream})
}

type RepositoryService_GetBlobServer interface {
	Send(*GetBlobResponse) error
	grpc.ServerStream
}

type repositoryServiceGetBlobServer struct {
	grpc.ServerStream
}

func (x *repositoryServiceGetBlobServer) Send(m *GetBlobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RepositoryService_ListCommits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCommitsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepositoryServiceServer).ListCommits(m, &repositoryServiceListCommitsServer{stream})
}

type RepositoryService_ListCommitsServer interface {
	Send(*ListCommitsResponse) error
	grpc.ServerStream
}

type repositoryServiceListCommitsServer struct {
	grpc.ServerStream
}

func (x *repositoryServiceListCommitsServer) Send(m *ListCommitsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RepositoryService_GetCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GetCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/GetCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GetCommit(ctx, req.(*GetCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_GetCommitDivergences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitDivergencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GetCommitDivergences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/GetCommitDivergences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GetCommitDivergences(ctx, req.(*GetCommitDivergencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_DeleteRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).DeleteRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/DeleteRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).DeleteRepository(ctx, req.(*DeleteRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_SyncRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).SyncRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/SyncRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).SyncRepository(ctx, req.(*SyncRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_HashRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).HashRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/HashRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).HashRepository(ctx, req.(*HashRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_MergeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).MergeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.RepositoryService/MergeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).MergeBase(ctx, req.(*MergeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepositoryService_ServiceDesc is the grpc.ServiceDesc for RepositoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepositoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTreeNode",
			Handler:    _RepositoryService_GetTreeNode_Handler,
		},
		{
			MethodName: "GetSubmodule",
			Handler:    _RepositoryService_GetSubmodule_Handler,
		},
		{
			MethodName: "GetCommit",
			Handler:    _RepositoryService_GetCommit_Handler,
		},
		{
			MethodName: "GetCommitDivergences",
			Handler:    _RepositoryService_GetCommitDivergences_Handler,
		},
		{
			MethodName: "DeleteRepository",
			Handler:    _RepositoryService_DeleteRepository_Handler,
		},
		{
			MethodName: "SyncRepository",
			Handler:    _RepositoryService_SyncRepository_Handler,
		},
		{
			MethodName: "HashRepository",
			Handler:    _RepositoryService_HashRepository_Handler,
		},
		{
			MethodName: "MergeBase",
			Handler:    _RepositoryService_MergeBase_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateRepository",
			Handler:       _RepositoryService_CreateRepository_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListTreeNodes",
			Handler:       _RepositoryService_ListTreeNodes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBlob",
			Handler:       _RepositoryService_GetBlob_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListCommits",
			Handler:       _RepositoryService_ListCommits_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "repo.proto",
}
