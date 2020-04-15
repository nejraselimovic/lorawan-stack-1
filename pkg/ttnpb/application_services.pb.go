// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/application_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/application_services.proto", fileDescriptor_f6c42f4fe8e3c902)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/application_services.proto", fileDescriptor_f6c42f4fe8e3c902)
}

var fileDescriptor_f6c42f4fe8e3c902 = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4d, 0x6c, 0xe3, 0x44,
	0x14, 0xc7, 0x3d, 0x5d, 0x14, 0x89, 0xd9, 0x85, 0xd5, 0x0e, 0x12, 0x48, 0xde, 0x65, 0x84, 0x8c,
	0x76, 0x8b, 0x96, 0xcd, 0x18, 0x36, 0x02, 0xb4, 0x50, 0x01, 0xfd, 0x40, 0xa1, 0x2a, 0x88, 0xaa,
	0x15, 0x97, 0x5c, 0x8a, 0x93, 0x4e, 0x1d, 0x2b, 0xa9, 0xc7, 0x78, 0x26, 0x2d, 0x69, 0x54, 0xa9,
	0x70, 0xaa, 0x7a, 0x02, 0xf1, 0x21, 0x84, 0x40, 0x42, 0x08, 0x44, 0x2f, 0x48, 0x3d, 0xf6, 0xd8,
	0x63, 0x8f, 0x95, 0xb8, 0xf4, 0xd8, 0xd8, 0x1c, 0x7a, 0x00, 0xa9, 0xc7, 0x1e, 0x38, 0x20, 0x8f,
	0x6d, 0x6a, 0x27, 0xa9, 0xd3, 0x24, 0x7b, 0xcb, 0xcc, 0xbc, 0x79, 0xef, 0x37, 0xef, 0xf9, 0xff,
	0x57, 0xe0, 0x83, 0x3a, 0x73, 0x8d, 0x75, 0xc3, 0xce, 0x73, 0x61, 0x54, 0x6a, 0xba, 0xe1, 0x58,
	0xba, 0xe1, 0x38, 0x75, 0xab, 0x62, 0x08, 0x8b, 0xd9, 0x4b, 0x9c, 0xba, 0x6b, 0x56, 0x85, 0x72,
	0xe2, 0xb8, 0x4c, 0x30, 0xf4, 0xb4, 0x10, 0x36, 0x89, 0x6e, 0x90, 0xb5, 0x82, 0x9a, 0x37, 0x2d,
	0x51, 0x6d, 0x94, 0x49, 0x85, 0xad, 0xea, 0x26, 0x33, 0x99, 0x2e, 0xc3, 0xca, 0x8d, 0x15, 0xb9,
	0x92, 0x0b, 0xf9, 0x2b, 0xbc, 0xae, 0xde, 0x31, 0x19, 0x33, 0xeb, 0x34, 0xac, 0x62, 0xdb, 0x4c,
	0xc8, 0x22, 0x51, 0x72, 0xf5, 0x76, 0x74, 0xfa, 0x7f, 0x0e, 0xba, 0xea, 0x88, 0x66, 0x74, 0xf8,
	0x62, 0x26, 0xe7, 0xe5, 0x41, 0xd6, 0x32, 0xb5, 0x85, 0xb5, 0x62, 0x51, 0x37, 0x2e, 0x83, 0xbb,
	0x83, 0x5c, 0xcb, 0xac, 0x8a, 0xe8, 0xfc, 0xe1, 0x3f, 0x39, 0xf8, 0xcc, 0xe4, 0x45, 0xea, 0x05,
	0x6a, 0x5a, 0x5c, 0xb8, 0x4d, 0xe4, 0x03, 0x98, 0x9b, 0x76, 0xa9, 0x21, 0x28, 0x7a, 0x89, 0xa4,
	0xfb, 0x40, 0xc2, 0xfd, 0xd4, 0xad, 0x4f, 0x1b, 0x94, 0x0b, 0xf5, 0x76, 0x67, 0x64, 0x22, 0x46,
	0xfb, 0x0a, 0x7c, 0xf1, 0xe7, 0x5f, 0x5f, 0x8f, 0xed, 0x00, 0xad, 0xa0, 0x37, 0x38, 0x75, 0xb9,
	0xde, 0xaa, 0xb0, 0x7a, 0xdd, 0x28, 0x33, 0xd7, 0x10, 0xcc, 0x25, 0xc1, 0xde, 0x92, 0xb5, 0xcc,
	0xe3, 0x1f, 0x9b, 0xc9, 0x27, 0xf3, 0x37, 0xc1, 0xfd, 0xd2, 0xbc, 0x36, 0xa7, 0x33, 0xd7, 0x34,
	0x6c, 0x6b, 0x23, 0xdc, 0xec, 0xc8, 0x90, 0x3c, 0x93, 0x99, 0x3a, 0x36, 0xba, 0x32, 0xa2, 0xcf,
	0x01, 0xbc, 0x56, 0xa4, 0x02, 0xdd, 0xed, 0x04, 0x2f, 0x52, 0x31, 0xe8, 0xfb, 0x5e, 0x97, 0xcf,
	0x7b, 0x05, 0x91, 0x54, 0x15, 0xbd, 0x95, 0xfc, 0xc0, 0x02, 0xa8, 0xf4, 0x7a, 0x13, 0xfd, 0x0d,
	0xe0, 0x13, 0x1f, 0x58, 0x5c, 0xa0, 0xf1, 0xce, 0xec, 0xc1, 0x6e, 0xa2, 0x02, 0x8f, 0x31, 0xee,
	0x64, 0x60, 0x70, 0xed, 0xc7, 0xb0, 0xcf, 0xdf, 0x02, 0xf4, 0x54, 0x8a, 0xa4, 0xf4, 0x1a, 0x1a,
	0xa6, 0xf1, 0xa5, 0x0f, 0xd1, 0xe3, 0xec, 0x3a, 0xda, 0x01, 0x30, 0xf7, 0xb1, 0xb3, 0xdc, 0xf3,
	0xc3, 0x0a, 0xf7, 0x07, 0x6d, 0xfc, 0x23, 0xf9, 0xde, 0x82, 0x9a, 0xd1, 0x78, 0xd2, 0xa3, 0xf1,
	0xc1, 0xfc, 0x1d, 0x98, 0x9b, 0xa1, 0x75, 0x2a, 0x28, 0xba, 0x97, 0x51, 0x61, 0xf6, 0x42, 0x55,
	0xea, 0xb3, 0x24, 0xd4, 0x2d, 0x89, 0x75, 0x4b, 0xde, 0x0b, 0x74, 0xab, 0xdd, 0x93, 0x10, 0x2f,
	0xdc, 0xc7, 0x99, 0xd3, 0xdf, 0x7c, 0xf8, 0xef, 0x75, 0x78, 0x2b, 0x91, 0x7a, 0xb2, 0x52, 0xa1,
	0x9c, 0xa3, 0x16, 0x84, 0xc1, 0xb0, 0x17, 0xa4, 0x32, 0x07, 0x60, 0xe9, 0x88, 0x0b, 0xef, 0x6b,
	0x79, 0xc9, 0x32, 0x8e, 0xee, 0x66, 0xb3, 0x44, 0x46, 0x80, 0x7e, 0x00, 0xf0, 0x46, 0x24, 0xe9,
	0xf9, 0xd9, 0x39, 0xda, 0x44, 0xa4, 0xaf, 0xe0, 0xc3, 0xc0, 0x78, 0x3a, 0x5d, 0x1c, 0xe1, 0xb1,
	0x36, 0x25, 0x39, 0x26, 0xb4, 0x37, 0x06, 0x53, 0x44, 0x60, 0x52, 0xf9, 0x1a, 0x6d, 0x4a, 0x85,
	0x7e, 0x07, 0xe0, 0x75, 0xa9, 0x03, 0x99, 0x92, 0xa3, 0x7c, 0x1f, 0x91, 0x44, 0x71, 0x31, 0xda,
	0x73, 0xbd, 0xd1, 0xb8, 0xf6, 0x8e, 0x64, 0x7b, 0x84, 0x86, 0x65, 0x0b, 0xba, 0xf6, 0x64, 0xe0,
	0x12, 0x61, 0xcb, 0x5e, 0xce, 0x36, 0x90, 0xab, 0xf5, 0xeb, 0x7d, 0xc9, 0x34, 0x85, 0xde, 0x1d,
	0x92, 0x49, 0x6f, 0xd5, 0x68, 0x53, 0x7a, 0xca, 0xef, 0x00, 0xde, 0x88, 0xc4, 0x74, 0xc9, 0x48,
	0xbb, 0xa4, 0x76, 0x35, 0xc4, 0x8f, 0x24, 0xe2, 0xac, 0x3a, 0x33, 0x34, 0xa2, 0xe1, 0x58, 0x4b,
	0x35, 0xda, 0x24, 0x91, 0x02, 0xbf, 0xb9, 0x06, 0x6f, 0x16, 0xa9, 0x98, 0x4e, 0x38, 0x0a, 0x7a,
	0x35, 0xbb, 0x99, 0xc9, 0xd8, 0x98, 0x77, 0xbc, 0xc7, 0x95, 0x74, 0x1c, 0x77, 0x98, 0xcd, 0xa9,
	0xf6, 0xeb, 0x98, 0x7c, 0xc1, 0x4f, 0x63, 0xe8, 0xad, 0x01, 0x9f, 0x90, 0x34, 0xbd, 0x52, 0x19,
	0x7d, 0x32, 0xc2, 0x75, 0x69, 0xc3, 0xfd, 0x5c, 0xb8, 0xb4, 0x81, 0x3e, 0x1b, 0xa5, 0x46, 0xd2,
	0x86, 0x07, 0xb5, 0x6c, 0xf4, 0x1b, 0x80, 0x37, 0x17, 0xfb, 0x8d, 0x65, 0xb1, 0xef, 0x58, 0x2e,
	0x73, 0xcb, 0xa2, 0x1c, 0xc2, 0xa4, 0x3a, 0x31, 0xc2, 0x03, 0xa5, 0x3d, 0xfc, 0x01, 0xe0, 0xad,
	0xc0, 0x01, 0x92, 0xc5, 0x39, 0x2a, 0xf4, 0x31, 0x89, 0x54, 0x74, 0xcc, 0xfa, 0x7c, 0x97, 0xeb,
	0x25, 0xa3, 0xb4, 0x19, 0x89, 0xfc, 0x36, 0x1a, 0x09, 0x79, 0xea, 0x17, 0x70, 0xd8, 0xc6, 0xe0,
	0xa8, 0x8d, 0xc1, 0x71, 0x1b, 0x2b, 0x27, 0x6d, 0xac, 0x9c, 0xb6, 0xb1, 0x72, 0xd6, 0xc6, 0xca,
	0x79, 0x1b, 0x83, 0x2d, 0x0f, 0x83, 0x6d, 0x0f, 0x2b, 0xbb, 0x1e, 0x06, 0x7b, 0x1e, 0x56, 0xf6,
	0x3d, 0xac, 0x1c, 0x78, 0x58, 0x39, 0xf4, 0x30, 0x38, 0xf2, 0x30, 0x38, 0xf6, 0xb0, 0x72, 0xe2,
	0x61, 0x70, 0xea, 0x61, 0xe5, 0xcc, 0xc3, 0xe0, 0xdc, 0xc3, 0xca, 0x96, 0x8f, 0x95, 0x6d, 0x1f,
	0x83, 0x2f, 0x7d, 0xac, 0x7c, 0xef, 0x63, 0xf0, 0xb3, 0x8f, 0x95, 0x5d, 0x1f, 0x2b, 0x7b, 0x3e,
	0x06, 0xfb, 0x3e, 0x06, 0x07, 0x3e, 0x06, 0xa5, 0x07, 0x26, 0x23, 0xa2, 0x4a, 0x45, 0xd5, 0xb2,
	0x4d, 0x4e, 0x6c, 0x2a, 0xd6, 0x99, 0x5b, 0xd3, 0xd3, 0x7f, 0x0d, 0x9d, 0x9a, 0xa9, 0x0b, 0x61,
	0x3b, 0xe5, 0x72, 0x4e, 0x4e, 0xab, 0xf0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x33, 0x89,
	0xa9, 0x2e, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationRegistryClient is the client API for ApplicationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationRegistryClient interface {
	// Create a new application. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// Get the application with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// List applications. See request message for details.
	List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*Applications, error)
	Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	Delete(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type applicationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewApplicationRegistryClient(cc *grpc.ClientConn) ApplicationRegistryClient {
	return &applicationRegistryClient{cc}
}

func (c *applicationRegistryClient) Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) List(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*Applications, error) {
	out := new(Applications)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationRegistryClient) Delete(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationRegistryServer is the server API for ApplicationRegistry service.
type ApplicationRegistryServer interface {
	// Create a new application. This also sets the given organization or user as
	// first collaborator with all possible rights.
	Create(context.Context, *CreateApplicationRequest) (*Application, error)
	// Get the application with the given identifiers, selecting the fields given
	// by the field mask. The method may return more or less fields, depending on
	// the rights of the caller.
	Get(context.Context, *GetApplicationRequest) (*Application, error)
	// List applications. See request message for details.
	List(context.Context, *ListApplicationsRequest) (*Applications, error)
	Update(context.Context, *UpdateApplicationRequest) (*Application, error)
	Delete(context.Context, *ApplicationIdentifiers) (*types.Empty, error)
}

// UnimplementedApplicationRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationRegistryServer struct {
}

func (*UnimplementedApplicationRegistryServer) Create(ctx context.Context, req *CreateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedApplicationRegistryServer) Get(ctx context.Context, req *GetApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedApplicationRegistryServer) List(ctx context.Context, req *ListApplicationsRequest) (*Applications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedApplicationRegistryServer) Update(ctx context.Context, req *UpdateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedApplicationRegistryServer) Delete(ctx context.Context, req *ApplicationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterApplicationRegistryServer(s *grpc.Server, srv ApplicationRegistryServer) {
	s.RegisterService(&_ApplicationRegistry_serviceDesc, srv)
}

func _ApplicationRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Create(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Get(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).List(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Update(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegistryServer).Delete(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationRegistry",
	HandlerType: (*ApplicationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApplicationRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApplicationRegistry_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ApplicationRegistry_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApplicationRegistry_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApplicationRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/application_services.proto",
}

// ApplicationAccessClient is the client API for ApplicationAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationAccessClient interface {
	ListRights(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	CreateAPIKey(ctx context.Context, in *CreateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	ListAPIKeys(ctx context.Context, in *ListApplicationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error)
	GetAPIKey(ctx context.Context, in *GetApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Update the rights of an existing application API key. To generate an API key,
	// the CreateAPIKey should be used. To delete an API key, update it
	// with zero rights. It is required for the caller to have all assigned or/and removed rights.
	UpdateAPIKey(ctx context.Context, in *UpdateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	// Get the rights of a collaborator (member) of the application.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(ctx context.Context, in *GetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the application. It is required for the caller to
	// have all assigned or/and removed rights.
	// Setting a collaborator without rights, removes them.
	SetCollaborator(ctx context.Context, in *SetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ListCollaborators(ctx context.Context, in *ListApplicationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error)
}

type applicationAccessClient struct {
	cc *grpc.ClientConn
}

func NewApplicationAccessClient(cc *grpc.ClientConn) ApplicationAccessClient {
	return &applicationAccessClient{cc}
}

func (c *applicationAccessClient) ListRights(ctx context.Context, in *ApplicationIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) CreateAPIKey(ctx context.Context, in *CreateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) ListAPIKeys(ctx context.Context, in *ListApplicationAPIKeysRequest, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) GetAPIKey(ctx context.Context, in *GetApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/GetAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateApplicationAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) GetCollaborator(ctx context.Context, in *GetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*GetCollaboratorResponse, error) {
	out := new(GetCollaboratorResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/GetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) SetCollaborator(ctx context.Context, in *SetApplicationCollaboratorRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/SetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationAccessClient) ListCollaborators(ctx context.Context, in *ListApplicationCollaboratorsRequest, opts ...grpc.CallOption) (*Collaborators, error) {
	out := new(Collaborators)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationAccess/ListCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationAccessServer is the server API for ApplicationAccess service.
type ApplicationAccessServer interface {
	ListRights(context.Context, *ApplicationIdentifiers) (*Rights, error)
	CreateAPIKey(context.Context, *CreateApplicationAPIKeyRequest) (*APIKey, error)
	ListAPIKeys(context.Context, *ListApplicationAPIKeysRequest) (*APIKeys, error)
	GetAPIKey(context.Context, *GetApplicationAPIKeyRequest) (*APIKey, error)
	// Update the rights of an existing application API key. To generate an API key,
	// the CreateAPIKey should be used. To delete an API key, update it
	// with zero rights. It is required for the caller to have all assigned or/and removed rights.
	UpdateAPIKey(context.Context, *UpdateApplicationAPIKeyRequest) (*APIKey, error)
	// Get the rights of a collaborator (member) of the application.
	// Pseudo-rights in the response (such as the "_ALL" right) are not expanded.
	GetCollaborator(context.Context, *GetApplicationCollaboratorRequest) (*GetCollaboratorResponse, error)
	// Set the rights of a collaborator (member) on the application. It is required for the caller to
	// have all assigned or/and removed rights.
	// Setting a collaborator without rights, removes them.
	SetCollaborator(context.Context, *SetApplicationCollaboratorRequest) (*types.Empty, error)
	ListCollaborators(context.Context, *ListApplicationCollaboratorsRequest) (*Collaborators, error)
}

// UnimplementedApplicationAccessServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationAccessServer struct {
}

func (*UnimplementedApplicationAccessServer) ListRights(ctx context.Context, req *ApplicationIdentifiers) (*Rights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRights not implemented")
}
func (*UnimplementedApplicationAccessServer) CreateAPIKey(ctx context.Context, req *CreateApplicationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (*UnimplementedApplicationAccessServer) ListAPIKeys(ctx context.Context, req *ListApplicationAPIKeysRequest) (*APIKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (*UnimplementedApplicationAccessServer) GetAPIKey(ctx context.Context, req *GetApplicationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (*UnimplementedApplicationAccessServer) UpdateAPIKey(ctx context.Context, req *UpdateApplicationAPIKeyRequest) (*APIKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAPIKey not implemented")
}
func (*UnimplementedApplicationAccessServer) GetCollaborator(ctx context.Context, req *GetApplicationCollaboratorRequest) (*GetCollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborator not implemented")
}
func (*UnimplementedApplicationAccessServer) SetCollaborator(ctx context.Context, req *SetApplicationCollaboratorRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCollaborator not implemented")
}
func (*UnimplementedApplicationAccessServer) ListCollaborators(ctx context.Context, req *ListApplicationCollaboratorsRequest) (*Collaborators, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollaborators not implemented")
}

func RegisterApplicationAccessServer(s *grpc.Server, srv ApplicationAccessServer) {
	s.RegisterService(&_ApplicationAccess_serviceDesc, srv)
}

func _ApplicationAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).ListRights(ctx, req.(*ApplicationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).CreateAPIKey(ctx, req.(*CreateApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).ListAPIKeys(ctx, req.(*ListApplicationAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/GetAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).GetAPIKey(ctx, req.(*GetApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).UpdateAPIKey(ctx, req.(*UpdateApplicationAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_GetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).GetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/GetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).GetCollaborator(ctx, req.(*GetApplicationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_SetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetApplicationCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).SetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/SetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).SetCollaborator(ctx, req.(*SetApplicationCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationAccess_ListCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationAccessServer).ListCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationAccess/ListCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationAccessServer).ListCollaborators(ctx, req.(*ListApplicationCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationAccess",
	HandlerType: (*ApplicationAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _ApplicationAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _ApplicationAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _ApplicationAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "GetAPIKey",
			Handler:    _ApplicationAccess_GetAPIKey_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _ApplicationAccess_UpdateAPIKey_Handler,
		},
		{
			MethodName: "GetCollaborator",
			Handler:    _ApplicationAccess_GetCollaborator_Handler,
		},
		{
			MethodName: "SetCollaborator",
			Handler:    _ApplicationAccess_SetCollaborator_Handler,
		},
		{
			MethodName: "ListCollaborators",
			Handler:    _ApplicationAccess_ListCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/application_services.proto",
}
