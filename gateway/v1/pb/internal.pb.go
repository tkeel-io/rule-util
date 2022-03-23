// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DispatchRequest struct {
	// Name/ID of the entity which is the destination.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Message body.
	Message              []byte   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DispatchRequest) Reset()         { *m = DispatchRequest{} }
func (m *DispatchRequest) String() string { return proto.CompactTextString(m) }
func (*DispatchRequest) ProtoMessage()    {}
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_e8b92c80bca60b36, []int{0}
}
func (m *DispatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DispatchRequest.Unmarshal(m, b)
}
func (m *DispatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DispatchRequest.Marshal(b, m, deterministic)
}
func (dst *DispatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DispatchRequest.Merge(dst, src)
}
func (m *DispatchRequest) XXX_Size() int {
	return xxx_messageInfo_DispatchRequest.Size(m)
}
func (m *DispatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DispatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DispatchRequest proto.InternalMessageInfo

func (m *DispatchRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *DispatchRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type CommonResponse struct {
	// Success or fail.
	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	// Status code.
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// Status details.
	Details              []byte   `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_e8b92c80bca60b36, []int{1}
}
func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (dst *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(dst, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CommonResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetDetails() []byte {
	if m != nil {
		return m.Details
	}
	return nil
}

type DisconnectRequest struct {
	// Name/ID of the entity which is the destination.
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_e8b92c80bca60b36, []int{2}
}
func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (dst *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(dst, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*DispatchRequest)(nil), "pb.DispatchRequest")
	proto.RegisterType((*CommonResponse)(nil), "pb.CommonResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "pb.DisconnectRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternalSvcClient is the client API for InternalSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalSvcClient interface {
	// Dispatch message to the specific entity.
	Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// Disconnect the specific connection.
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type internalSvcClient struct {
	cc *grpc.ClientConn
}

func NewInternalSvcClient(cc *grpc.ClientConn) InternalSvcClient {
	return &internalSvcClient{cc}
}

func (c *internalSvcClient) Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.InternalSvc/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalSvcClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.InternalSvc/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalSvcServer is the server API for InternalSvc service.
type InternalSvcServer interface {
	// Dispatch message to the specific entity.
	Dispatch(context.Context, *DispatchRequest) (*CommonResponse, error)
	// Disconnect the specific connection.
	Disconnect(context.Context, *DisconnectRequest) (*CommonResponse, error)
}

func RegisterInternalSvcServer(s *grpc.Server, srv InternalSvcServer) {
	s.RegisterService(&_InternalSvc_serviceDesc, srv)
}

func _InternalSvc_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalSvcServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.InternalSvc/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalSvcServer).Dispatch(ctx, req.(*DispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalSvc_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalSvcServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.InternalSvc/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalSvcServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.InternalSvc",
	HandlerType: (*InternalSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _InternalSvc_Dispatch_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _InternalSvc_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal.proto",
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_internal_e8b92c80bca60b36) }

var fileDescriptor_internal_e8b92c80bca60b36 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xdd, 0xaa, 0xb5, 0x8e, 0x52, 0x71, 0x44, 0x09, 0x9e, 0x4a, 0x4e, 0x05, 0x61, 0x0f,
	0x8a, 0x07, 0xcf, 0xed, 0xc5, 0x6b, 0x7a, 0xf1, 0x9a, 0x4d, 0x87, 0x1a, 0xe8, 0x26, 0x71, 0x67,
	0xea, 0xd5, 0xbf, 0x2e, 0xbb, 0x9b, 0x45, 0x56, 0x04, 0x6f, 0xf9, 0xc2, 0x7b, 0x2f, 0xef, 0x05,
	0xe6, 0x3e, 0x08, 0x35, 0xc1, 0xee, 0xcb, 0xd4, 0x44, 0x89, 0x38, 0x49, 0x95, 0x5e, 0xc1, 0xd5,
	0xda, 0x73, 0xb2, 0xe2, 0xde, 0x0d, 0x7d, 0x1c, 0x88, 0x05, 0xef, 0x60, 0x2a, 0xb6, 0xd9, 0x91,
	0xa8, 0x62, 0x51, 0x2c, 0xcf, 0x4d, 0x26, 0x54, 0x70, 0x56, 0x13, 0xb3, 0xdd, 0x91, 0x9a, 0x2c,
	0x8a, 0xe5, 0xa5, 0x19, 0x50, 0xbf, 0xc1, 0x7c, 0x15, 0xeb, 0x3a, 0x06, 0x43, 0x9c, 0x62, 0x60,
	0x6a, 0xb5, 0x9b, 0x83, 0x73, 0xc4, 0xdc, 0x85, 0xcc, 0xcc, 0x80, 0x88, 0x70, 0xe2, 0xe2, 0xb6,
	0x8f, 0x38, 0x35, 0xdd, 0xb9, 0x55, 0x6f, 0x49, 0xac, 0xdf, 0xb3, 0x3a, 0xee, 0x93, 0x33, 0xea,
	0x07, 0xb8, 0x5e, 0x7b, 0x76, 0x31, 0x04, 0x72, 0xf2, 0x4f, 0xc1, 0xc7, 0x2f, 0xb8, 0x78, 0xcd,
	0x0b, 0x37, 0x9f, 0x0e, 0x9f, 0x61, 0x36, 0x4c, 0xc3, 0x9b, 0x32, 0x55, 0xe5, 0xaf, 0xa1, 0xf7,
	0xd8, 0x5e, 0x8e, 0x8b, 0xeb, 0x23, 0x7c, 0x01, 0xf8, 0x79, 0x12, 0x6f, 0xb3, 0x71, 0x5c, 0xe1,
	0x6f, 0x6b, 0x35, 0xed, 0xfe, 0xf5, 0xe9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x45, 0xea, 0x45, 0xe9,
	0x69, 0x01, 0x00, 0x00,
}