// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// GetGreeterClient is the client API for GetGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetGreeterClient interface {
	GetCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseReplyList, error)
	GetSearchCourse(ctx context.Context, in *SearchACRequest, opts ...grpc.CallOption) (*SearchACReplyList, error)
	GetSearchECourse(ctx context.Context, in *SearchERequest, opts ...grpc.CallOption) (*SearchEReplyList, error)
	DeleteSCourse(ctx context.Context, in *SdelCRequest, opts ...grpc.CallOption) (*SdelCReply, error)
	PutSCourse(ctx context.Context, in *SputCRequest, opts ...grpc.CallOption) (*SputCReply, error)
	PutTCourse(ctx context.Context, in *TputCRequest, opts ...grpc.CallOption) (*TputCReply, error)
	UpdateTCourse(ctx context.Context, in *TupdateCRequest, opts ...grpc.CallOption) (*TupdateCReply, error)
	DeleteTCourse(ctx context.Context, in *TdeleteCRequest, opts ...grpc.CallOption) (*TdeleteCReply, error)
	GetTCourse(ctx context.Context, in *TgetCRequest, opts ...grpc.CallOption) (*TgetCReplyList, error)
}

type getGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGetGreeterClient(cc grpc.ClientConnInterface) GetGreeterClient {
	return &getGreeterClient{cc}
}

func (c *getGreeterClient) GetCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseReplyList, error) {
	out := new(CourseReplyList)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/getCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) GetSearchCourse(ctx context.Context, in *SearchACRequest, opts ...grpc.CallOption) (*SearchACReplyList, error) {
	out := new(SearchACReplyList)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/getSearchCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) GetSearchECourse(ctx context.Context, in *SearchERequest, opts ...grpc.CallOption) (*SearchEReplyList, error) {
	out := new(SearchEReplyList)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/getSearchECourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) DeleteSCourse(ctx context.Context, in *SdelCRequest, opts ...grpc.CallOption) (*SdelCReply, error) {
	out := new(SdelCReply)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/deleteSCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) PutSCourse(ctx context.Context, in *SputCRequest, opts ...grpc.CallOption) (*SputCReply, error) {
	out := new(SputCReply)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/putSCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) PutTCourse(ctx context.Context, in *TputCRequest, opts ...grpc.CallOption) (*TputCReply, error) {
	out := new(TputCReply)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/putTCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) UpdateTCourse(ctx context.Context, in *TupdateCRequest, opts ...grpc.CallOption) (*TupdateCReply, error) {
	out := new(TupdateCReply)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/updateTCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) DeleteTCourse(ctx context.Context, in *TdeleteCRequest, opts ...grpc.CallOption) (*TdeleteCReply, error) {
	out := new(TdeleteCReply)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/deleteTCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getGreeterClient) GetTCourse(ctx context.Context, in *TgetCRequest, opts ...grpc.CallOption) (*TgetCReplyList, error) {
	out := new(TgetCReplyList)
	err := c.cc.Invoke(ctx, "/CC.getGreeter/getTCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetGreeterServer is the server API for GetGreeter service.
// All implementations must embed UnimplementedGetGreeterServer
// for forward compatibility
type GetGreeterServer interface {
	GetCourse(context.Context, *CourseRequest) (*CourseReplyList, error)
	GetSearchCourse(context.Context, *SearchACRequest) (*SearchACReplyList, error)
	GetSearchECourse(context.Context, *SearchERequest) (*SearchEReplyList, error)
	DeleteSCourse(context.Context, *SdelCRequest) (*SdelCReply, error)
	PutSCourse(context.Context, *SputCRequest) (*SputCReply, error)
	PutTCourse(context.Context, *TputCRequest) (*TputCReply, error)
	UpdateTCourse(context.Context, *TupdateCRequest) (*TupdateCReply, error)
	DeleteTCourse(context.Context, *TdeleteCRequest) (*TdeleteCReply, error)
	GetTCourse(context.Context, *TgetCRequest) (*TgetCReplyList, error)
	mustEmbedUnimplementedGetGreeterServer()
}

// UnimplementedGetGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGetGreeterServer struct {
}

func (UnimplementedGetGreeterServer) GetCourse(context.Context, *CourseRequest) (*CourseReplyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedGetGreeterServer) GetSearchCourse(context.Context, *SearchACRequest) (*SearchACReplyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchCourse not implemented")
}
func (UnimplementedGetGreeterServer) GetSearchECourse(context.Context, *SearchERequest) (*SearchEReplyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchECourse not implemented")
}
func (UnimplementedGetGreeterServer) DeleteSCourse(context.Context, *SdelCRequest) (*SdelCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSCourse not implemented")
}
func (UnimplementedGetGreeterServer) PutSCourse(context.Context, *SputCRequest) (*SputCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSCourse not implemented")
}
func (UnimplementedGetGreeterServer) PutTCourse(context.Context, *TputCRequest) (*TputCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutTCourse not implemented")
}
func (UnimplementedGetGreeterServer) UpdateTCourse(context.Context, *TupdateCRequest) (*TupdateCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTCourse not implemented")
}
func (UnimplementedGetGreeterServer) DeleteTCourse(context.Context, *TdeleteCRequest) (*TdeleteCReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTCourse not implemented")
}
func (UnimplementedGetGreeterServer) GetTCourse(context.Context, *TgetCRequest) (*TgetCReplyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTCourse not implemented")
}
func (UnimplementedGetGreeterServer) mustEmbedUnimplementedGetGreeterServer() {}

// UnsafeGetGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetGreeterServer will
// result in compilation errors.
type UnsafeGetGreeterServer interface {
	mustEmbedUnimplementedGetGreeterServer()
}

func RegisterGetGreeterServer(s grpc.ServiceRegistrar, srv GetGreeterServer) {
	s.RegisterService(&GetGreeter_ServiceDesc, srv)
}

func _GetGreeter_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/getCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).GetCourse(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_GetSearchCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchACRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).GetSearchCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/getSearchCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).GetSearchCourse(ctx, req.(*SearchACRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_GetSearchECourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).GetSearchECourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/getSearchECourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).GetSearchECourse(ctx, req.(*SearchERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_DeleteSCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SdelCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).DeleteSCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/deleteSCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).DeleteSCourse(ctx, req.(*SdelCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_PutSCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SputCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).PutSCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/putSCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).PutSCourse(ctx, req.(*SputCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_PutTCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TputCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).PutTCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/putTCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).PutTCourse(ctx, req.(*TputCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_UpdateTCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TupdateCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).UpdateTCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/updateTCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).UpdateTCourse(ctx, req.(*TupdateCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_DeleteTCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TdeleteCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).DeleteTCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/deleteTCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).DeleteTCourse(ctx, req.(*TdeleteCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetGreeter_GetTCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TgetCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetGreeterServer).GetTCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CC.getGreeter/getTCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetGreeterServer).GetTCourse(ctx, req.(*TgetCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetGreeter_ServiceDesc is the grpc.ServiceDesc for GetGreeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CC.getGreeter",
	HandlerType: (*GetGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCourse",
			Handler:    _GetGreeter_GetCourse_Handler,
		},
		{
			MethodName: "getSearchCourse",
			Handler:    _GetGreeter_GetSearchCourse_Handler,
		},
		{
			MethodName: "getSearchECourse",
			Handler:    _GetGreeter_GetSearchECourse_Handler,
		},
		{
			MethodName: "deleteSCourse",
			Handler:    _GetGreeter_DeleteSCourse_Handler,
		},
		{
			MethodName: "putSCourse",
			Handler:    _GetGreeter_PutSCourse_Handler,
		},
		{
			MethodName: "putTCourse",
			Handler:    _GetGreeter_PutTCourse_Handler,
		},
		{
			MethodName: "updateTCourse",
			Handler:    _GetGreeter_UpdateTCourse_Handler,
		},
		{
			MethodName: "deleteTCourse",
			Handler:    _GetGreeter_DeleteTCourse_Handler,
		},
		{
			MethodName: "getTCourse",
			Handler:    _GetGreeter_GetTCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ChoiCourse.proto",
}