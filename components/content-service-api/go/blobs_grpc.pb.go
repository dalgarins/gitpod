// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// BlobServiceClient is the client API for BlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlobServiceClient interface {
	// UploadUrl provides a URL to which clients can upload the content via HTTP PUT.
	UploadUrl(ctx context.Context, in *UploadUrlRequest, opts ...grpc.CallOption) (*UploadUrlResponse, error)
	// DownloadUrl provides a URL from which clients can download the content via HTTP GET.
	DownloadUrl(ctx context.Context, in *DownloadUrlRequest, opts ...grpc.CallOption) (*DownloadUrlResponse, error)
	// Delete deletes the uploaded content.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type blobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlobServiceClient(cc grpc.ClientConnInterface) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) UploadUrl(ctx context.Context, in *UploadUrlRequest, opts ...grpc.CallOption) (*UploadUrlResponse, error) {
	out := new(UploadUrlResponse)
	err := c.cc.Invoke(ctx, "/contentservice.BlobService/UploadUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) DownloadUrl(ctx context.Context, in *DownloadUrlRequest, opts ...grpc.CallOption) (*DownloadUrlResponse, error) {
	out := new(DownloadUrlResponse)
	err := c.cc.Invoke(ctx, "/contentservice.BlobService/DownloadUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/contentservice.BlobService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/contentservice.BlobService/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlobServiceServer is the server API for BlobService service.
// All implementations must embed UnimplementedBlobServiceServer
// for forward compatibility
type BlobServiceServer interface {
	// UploadUrl provides a URL to which clients can upload the content via HTTP PUT.
	UploadUrl(context.Context, *UploadUrlRequest) (*UploadUrlResponse, error)
	// DownloadUrl provides a URL from which clients can download the content via HTTP GET.
	DownloadUrl(context.Context, *DownloadUrlRequest) (*DownloadUrlResponse, error)
	// Delete deletes the uploaded content.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Download(context.Context, *DownloadRequest) (*DownloadResponse, error)
	mustEmbedUnimplementedBlobServiceServer()
}

// UnimplementedBlobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlobServiceServer struct {
}

func (UnimplementedBlobServiceServer) UploadUrl(context.Context, *UploadUrlRequest) (*UploadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadUrl not implemented")
}
func (UnimplementedBlobServiceServer) DownloadUrl(context.Context, *DownloadUrlRequest) (*DownloadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadUrl not implemented")
}
func (UnimplementedBlobServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBlobServiceServer) Download(context.Context, *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedBlobServiceServer) mustEmbedUnimplementedBlobServiceServer() {}

// UnsafeBlobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlobServiceServer will
// result in compilation errors.
type UnsafeBlobServiceServer interface {
	mustEmbedUnimplementedBlobServiceServer()
}

func RegisterBlobServiceServer(s grpc.ServiceRegistrar, srv BlobServiceServer) {
	s.RegisterService(&BlobService_ServiceDesc, srv)
}

func _BlobService_UploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).UploadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.BlobService/UploadUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).UploadUrl(ctx, req.(*UploadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_DownloadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).DownloadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.BlobService/DownloadUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).DownloadUrl(ctx, req.(*DownloadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.BlobService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.BlobService/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlobService_ServiceDesc is the grpc.ServiceDesc for BlobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadUrl",
			Handler:    _BlobService_UploadUrl_Handler,
		},
		{
			MethodName: "DownloadUrl",
			Handler:    _BlobService_DownloadUrl_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BlobService_Delete_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _BlobService_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blobs.proto",
}
