// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/web/v1/interaction_service.proto

package v1 // import "github.com/census-instrumentation/opencensus-proto/gen-go/agent/web/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "github.com/census-instrumentation/opencensus-proto/gen-go/web/v1"

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

type ExportInteractionsRequest struct {
	Interactions         []*v1.Interaction `protobuf:"bytes,1,rep,name=interactions,proto3" json:"interactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExportInteractionsRequest) Reset()         { *m = ExportInteractionsRequest{} }
func (m *ExportInteractionsRequest) String() string { return proto.CompactTextString(m) }
func (*ExportInteractionsRequest) ProtoMessage()    {}
func (*ExportInteractionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_interaction_service_4ef3c17027e79b4b, []int{0}
}
func (m *ExportInteractionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportInteractionsRequest.Unmarshal(m, b)
}
func (m *ExportInteractionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportInteractionsRequest.Marshal(b, m, deterministic)
}
func (dst *ExportInteractionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportInteractionsRequest.Merge(dst, src)
}
func (m *ExportInteractionsRequest) XXX_Size() int {
	return xxx_messageInfo_ExportInteractionsRequest.Size(m)
}
func (m *ExportInteractionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportInteractionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportInteractionsRequest proto.InternalMessageInfo

func (m *ExportInteractionsRequest) GetInteractions() []*v1.Interaction {
	if m != nil {
		return m.Interactions
	}
	return nil
}

type ExportInteractionsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportInteractionsResponse) Reset()         { *m = ExportInteractionsResponse{} }
func (m *ExportInteractionsResponse) String() string { return proto.CompactTextString(m) }
func (*ExportInteractionsResponse) ProtoMessage()    {}
func (*ExportInteractionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_interaction_service_4ef3c17027e79b4b, []int{1}
}
func (m *ExportInteractionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportInteractionsResponse.Unmarshal(m, b)
}
func (m *ExportInteractionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportInteractionsResponse.Marshal(b, m, deterministic)
}
func (dst *ExportInteractionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportInteractionsResponse.Merge(dst, src)
}
func (m *ExportInteractionsResponse) XXX_Size() int {
	return xxx_messageInfo_ExportInteractionsResponse.Size(m)
}
func (m *ExportInteractionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportInteractionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportInteractionsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportInteractionsRequest)(nil), "opencensus.proto.agent.web.v1.ExportInteractionsRequest")
	proto.RegisterType((*ExportInteractionsResponse)(nil), "opencensus.proto.agent.web.v1.ExportInteractionsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InteractionServiceClient interface {
	Export(ctx context.Context, in *ExportInteractionsRequest, opts ...grpc.CallOption) (*ExportInteractionsResponse, error)
}

type interactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewInteractionServiceClient(cc *grpc.ClientConn) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) Export(ctx context.Context, in *ExportInteractionsRequest, opts ...grpc.CallOption) (*ExportInteractionsResponse, error) {
	out := new(ExportInteractionsResponse)
	err := c.cc.Invoke(ctx, "/opencensus.proto.agent.web.v1.InteractionService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
type InteractionServiceServer interface {
	Export(context.Context, *ExportInteractionsRequest) (*ExportInteractionsResponse, error)
}

func RegisterInteractionServiceServer(s *grpc.Server, srv InteractionServiceServer) {
	s.RegisterService(&_InteractionService_serviceDesc, srv)
}

func _InteractionService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportInteractionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opencensus.proto.agent.web.v1.InteractionService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).Export(ctx, req.(*ExportInteractionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InteractionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opencensus.proto.agent.web.v1.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _InteractionService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/web/v1/interaction_service.proto",
}

func init() {
	proto.RegisterFile("github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/web/v1/interaction_service.proto", fileDescriptor_interaction_service_4ef3c17027e79b4b)
}

var fileDescriptor_interaction_service_4ef3c17027e79b4b = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4f, 0xfb, 0x40,
	0x0c, 0xc5, 0xff, 0xd1, 0x5f, 0xea, 0x70, 0x30, 0xdd, 0x02, 0x44, 0x20, 0x95, 0x88, 0x81, 0xa5,
	0x3e, 0xb5, 0x2c, 0xb0, 0x56, 0x02, 0xc1, 0x56, 0x95, 0x8d, 0x05, 0x25, 0x27, 0x2b, 0xdc, 0x50,
	0x5f, 0x38, 0x3b, 0x29, 0x1b, 0x13, 0x1f, 0x81, 0xef, 0x8b, 0x92, 0xab, 0xd4, 0x54, 0x81, 0x0e,
	0x5d, 0xdf, 0xf3, 0xef, 0x3d, 0x5b, 0xb2, 0xa2, 0xd2, 0xc9, 0x5b, 0x5d, 0x80, 0xf5, 0x2b, 0x63,
	0x91, 0xb8, 0xe6, 0x89, 0x23, 0x96, 0x50, 0xaf, 0x90, 0x24, 0x17, 0xe7, 0xc9, 0xf8, 0x0a, 0x69,
	0x63, 0x55, 0xc1, 0x8b, 0x37, 0x1c, 0x6c, 0x4f, 0x34, 0x51, 0xcc, 0x4b, 0x24, 0x31, 0x6b, 0x2c,
	0x4c, 0x33, 0x35, 0x8e, 0x04, 0x43, 0x6e, 0x5b, 0xfc, 0x95, 0x31, 0x34, 0xce, 0x22, 0x74, 0x63,
	0xfa, 0x62, 0x0b, 0x46, 0x05, 0x3a, 0x10, 0xd6, 0x58, 0x40, 0x33, 0x4d, 0x2f, 0x07, 0xb9, 0x9b,
	0xc4, 0xd6, 0xef, 0x84, 0x0c, 0xd5, 0xd9, 0xfd, 0x47, 0xe5, 0x83, 0x3c, 0x6d, 0x4b, 0x78, 0x89,
	0xef, 0x35, 0xb2, 0xe8, 0x47, 0x75, 0xdc, 0xeb, 0xe6, 0xd3, 0x64, 0xfc, 0xff, 0xfa, 0x68, 0x76,
	0x05, 0x83, 0xd6, 0xd8, 0x07, 0xbd, 0x8c, 0xe5, 0x0e, 0x99, 0x9d, 0xab, 0xf4, 0xb7, 0x1a, 0xae,
	0x3c, 0x31, 0xce, 0xbe, 0x13, 0xa5, 0x7b, 0xc6, 0x73, 0xbc, 0x51, 0x7f, 0xaa, 0x51, 0x84, 0xf4,
	0x2d, 0xec, 0x3d, 0x14, 0xfe, 0x3c, 0x21, 0xbd, 0x3b, 0x80, 0x8c, 0x5b, 0x65, 0xff, 0xe6, 0x5f,
	0x89, 0x1a, 0x3b, 0xbf, 0x3f, 0x61, 0x7e, 0x32, 0xdc, 0x7c, 0xd1, 0x4e, 0x2d, 0x92, 0x97, 0x87,
	0x03, 0xde, 0xa1, 0x44, 0x9a, 0x94, 0xbb, 0x0f, 0x50, 0x8c, 0x3a, 0xeb, 0xe6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xfb, 0x97, 0x51, 0x5f, 0x02, 0x00, 0x00,
}
