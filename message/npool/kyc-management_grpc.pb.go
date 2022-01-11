// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

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

// KycManagementClient is the client API for KycManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KycManagementClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	CreateKycRecord(ctx context.Context, in *CreateKycRequest, opts ...grpc.CallOption) (*CreateKycResponse, error)
	GetKycByUserID(ctx context.Context, in *GetKycByUserIDRequest, opts ...grpc.CallOption) (*GetKycByUserIDResponse, error)
	GetKycByAppID(ctx context.Context, in *GetKycByAppIDRequest, opts ...grpc.CallOption) (*GetKycByAppIDResponse, error)
	GetAllKyc(ctx context.Context, in *GetAllKycRequest, opts ...grpc.CallOption) (*GetAllKycResponse, error)
	UpdateKyc(ctx context.Context, in *UpdateKycRequest, opts ...grpc.CallOption) (*UpdateKycResponse, error)
	UploadKycImage(ctx context.Context, in *UploadKycImageRequest, opts ...grpc.CallOption) (*UploadKycImageResponse, error)
	GetKycImage(ctx context.Context, in *GetKycImageRequest, opts ...grpc.CallOption) (*GetKycImageResponse, error)
}

type kycManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewKycManagementClient(cc grpc.ClientConnInterface) KycManagementClient {
	return &kycManagementClient{cc}
}

func (c *kycManagementClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) CreateKycRecord(ctx context.Context, in *CreateKycRequest, opts ...grpc.CallOption) (*CreateKycResponse, error) {
	out := new(CreateKycResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/CreateKycRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) GetKycByUserID(ctx context.Context, in *GetKycByUserIDRequest, opts ...grpc.CallOption) (*GetKycByUserIDResponse, error) {
	out := new(GetKycByUserIDResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/GetKycByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) GetKycByAppID(ctx context.Context, in *GetKycByAppIDRequest, opts ...grpc.CallOption) (*GetKycByAppIDResponse, error) {
	out := new(GetKycByAppIDResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/GetKycByAppID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) GetAllKyc(ctx context.Context, in *GetAllKycRequest, opts ...grpc.CallOption) (*GetAllKycResponse, error) {
	out := new(GetAllKycResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/GetAllKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) UpdateKyc(ctx context.Context, in *UpdateKycRequest, opts ...grpc.CallOption) (*UpdateKycResponse, error) {
	out := new(UpdateKycResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/UpdateKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) UploadKycImage(ctx context.Context, in *UploadKycImageRequest, opts ...grpc.CallOption) (*UploadKycImageResponse, error) {
	out := new(UploadKycImageResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/UploadKycImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kycManagementClient) GetKycImage(ctx context.Context, in *GetKycImageRequest, opts ...grpc.CallOption) (*GetKycImageResponse, error) {
	out := new(GetKycImageResponse)
	err := c.cc.Invoke(ctx, "/kyc.management.v1.KycManagement/GetKycImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KycManagementServer is the server API for KycManagement service.
// All implementations must embed UnimplementedKycManagementServer
// for forward compatibility
type KycManagementServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	CreateKycRecord(context.Context, *CreateKycRequest) (*CreateKycResponse, error)
	GetKycByUserID(context.Context, *GetKycByUserIDRequest) (*GetKycByUserIDResponse, error)
	GetKycByAppID(context.Context, *GetKycByAppIDRequest) (*GetKycByAppIDResponse, error)
	GetAllKyc(context.Context, *GetAllKycRequest) (*GetAllKycResponse, error)
	UpdateKyc(context.Context, *UpdateKycRequest) (*UpdateKycResponse, error)
	UploadKycImage(context.Context, *UploadKycImageRequest) (*UploadKycImageResponse, error)
	GetKycImage(context.Context, *GetKycImageRequest) (*GetKycImageResponse, error)
	mustEmbedUnimplementedKycManagementServer()
}

// UnimplementedKycManagementServer must be embedded to have forward compatible implementations.
type UnimplementedKycManagementServer struct {
}

func (UnimplementedKycManagementServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedKycManagementServer) CreateKycRecord(context.Context, *CreateKycRequest) (*CreateKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKycRecord not implemented")
}
func (UnimplementedKycManagementServer) GetKycByUserID(context.Context, *GetKycByUserIDRequest) (*GetKycByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycByUserID not implemented")
}
func (UnimplementedKycManagementServer) GetKycByAppID(context.Context, *GetKycByAppIDRequest) (*GetKycByAppIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycByAppID not implemented")
}
func (UnimplementedKycManagementServer) GetAllKyc(context.Context, *GetAllKycRequest) (*GetAllKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllKyc not implemented")
}
func (UnimplementedKycManagementServer) UpdateKyc(context.Context, *UpdateKycRequest) (*UpdateKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKyc not implemented")
}
func (UnimplementedKycManagementServer) UploadKycImage(context.Context, *UploadKycImageRequest) (*UploadKycImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadKycImage not implemented")
}
func (UnimplementedKycManagementServer) GetKycImage(context.Context, *GetKycImageRequest) (*GetKycImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycImage not implemented")
}
func (UnimplementedKycManagementServer) mustEmbedUnimplementedKycManagementServer() {}

// UnsafeKycManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KycManagementServer will
// result in compilation errors.
type UnsafeKycManagementServer interface {
	mustEmbedUnimplementedKycManagementServer()
}

func RegisterKycManagementServer(s grpc.ServiceRegistrar, srv KycManagementServer) {
	s.RegisterService(&KycManagement_ServiceDesc, srv)
}

func _KycManagement_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_CreateKycRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).CreateKycRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/CreateKycRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).CreateKycRecord(ctx, req.(*CreateKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_GetKycByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).GetKycByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/GetKycByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).GetKycByUserID(ctx, req.(*GetKycByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_GetKycByAppID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycByAppIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).GetKycByAppID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/GetKycByAppID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).GetKycByAppID(ctx, req.(*GetKycByAppIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_GetAllKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).GetAllKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/GetAllKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).GetAllKyc(ctx, req.(*GetAllKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_UpdateKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).UpdateKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/UpdateKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).UpdateKyc(ctx, req.(*UpdateKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_UploadKycImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadKycImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).UploadKycImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/UploadKycImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).UploadKycImage(ctx, req.(*UploadKycImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KycManagement_GetKycImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KycManagementServer).GetKycImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyc.management.v1.KycManagement/GetKycImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KycManagementServer).GetKycImage(ctx, req.(*GetKycImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KycManagement_ServiceDesc is the grpc.ServiceDesc for KycManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KycManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kyc.management.v1.KycManagement",
	HandlerType: (*KycManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KycManagement_Version_Handler,
		},
		{
			MethodName: "CreateKycRecord",
			Handler:    _KycManagement_CreateKycRecord_Handler,
		},
		{
			MethodName: "GetKycByUserID",
			Handler:    _KycManagement_GetKycByUserID_Handler,
		},
		{
			MethodName: "GetKycByAppID",
			Handler:    _KycManagement_GetKycByAppID_Handler,
		},
		{
			MethodName: "GetAllKyc",
			Handler:    _KycManagement_GetAllKyc_Handler,
		},
		{
			MethodName: "UpdateKyc",
			Handler:    _KycManagement_UpdateKyc_Handler,
		},
		{
			MethodName: "UploadKycImage",
			Handler:    _KycManagement_UploadKycImage_Handler,
		},
		{
			MethodName: "GetKycImage",
			Handler:    _KycManagement_GetKycImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/kyc-management.proto",
}
