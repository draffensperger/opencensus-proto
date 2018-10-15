// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/interaction/v1/interaction_service.proto

package v1 // import "github.com/census-instrumentation/opencensus-proto/gen-go/agent/interaction/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import interactionproto "github.com/draffensperger/opencensus-web/gen-go/interactionproto"

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
	Interactions         []*interactionproto.Interaction `protobuf:"bytes,1,rep,name=interactions,proto3" json:"interactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ExportInteractionsRequest) Reset()         { *m = ExportInteractionsRequest{} }
func (m *ExportInteractionsRequest) String() string { return proto.CompactTextString(m) }
func (*ExportInteractionsRequest) ProtoMessage()    {}
func (*ExportInteractionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_interaction_service_66ea0aa7cf8c72e1, []int{0}
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

func (m *ExportInteractionsRequest) GetInteractions() []*interactionproto.Interaction {
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
	return fileDescriptor_interaction_service_66ea0aa7cf8c72e1, []int{1}
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
	proto.RegisterType((*ExportInteractionsRequest)(nil), "opencensus.proto.agent.interaction.v1.ExportInteractionsRequest")
	proto.RegisterType((*ExportInteractionsResponse)(nil), "opencensus.proto.agent.interaction.v1.ExportInteractionsResponse")
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
	err := c.cc.Invoke(ctx, "/opencensus.proto.agent.interaction.v1.InteractionService/Export", in, out, opts...)
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
		FullMethod: "/opencensus.proto.agent.interaction.v1.InteractionService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).Export(ctx, req.(*ExportInteractionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InteractionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opencensus.proto.agent.interaction.v1.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _InteractionService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/interaction/v1/interaction_service.proto",
}

func init() {
	proto.RegisterFile("github.com/census-instrumentation/opencensus-proto/src/opencensus/proto/agent/interaction/v1/interaction_service.proto", fileDescriptor_interaction_service_66ea0aa7cf8c72e1)
}

var fileDescriptor_interaction_service_66ea0aa7cf8c72e1 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x2a, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x4e, 0xcd, 0x2b, 0x2e, 0x2d, 0xd6, 0xcd, 0xcc,
	0x2b, 0x2e, 0x29, 0x2a, 0xcd, 0x4d, 0xcd, 0x2b, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0xcf, 0x2f,
	0x48, 0xcd, 0x83, 0x4a, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0xeb, 0x17, 0x17, 0x25, 0x23, 0x09, 0xea,
	0x43, 0x04, 0x13, 0xd3, 0x53, 0xf3, 0x4a, 0xf4, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0x12, 0x93, 0xc1,
	0xfa, 0xca, 0x0c, 0x91, 0xb9, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x60, 0xe5,
	0x42, 0xaa, 0x08, 0x03, 0x20, 0x22, 0x7a, 0x60, 0x03, 0xf4, 0x90, 0x74, 0xe8, 0x95, 0x19, 0x4a,
	0x19, 0x62, 0xd8, 0x83, 0xdb, 0x06, 0x88, 0x39, 0x4a, 0xb9, 0x5c, 0x92, 0xae, 0x15, 0x05, 0xf9,
	0x45, 0x25, 0x9e, 0x08, 0xa9, 0xe2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa1, 0x00, 0x2e,
	0x1e, 0x24, 0x1d, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x3a, 0x7a, 0x18, 0xae, 0x41,
	0x75, 0x87, 0x1e, 0x92, 0x59, 0x41, 0x28, 0x26, 0x28, 0xc9, 0x70, 0x49, 0x61, 0xb3, 0xae, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0xd5, 0x68, 0x19, 0x23, 0x97, 0x10, 0x92, 0x44, 0x30, 0x24, 0x0c, 0x84,
	0xfa, 0x19, 0xb9, 0xd8, 0x20, 0xba, 0x84, 0x1c, 0xf4, 0x88, 0x0a, 0x09, 0x3d, 0x9c, 0x7e, 0x92,
	0x72, 0xa4, 0xc0, 0x04, 0x88, 0x33, 0x95, 0x18, 0x9c, 0xa6, 0x32, 0x72, 0x69, 0x64, 0xe6, 0x13,
	0x67, 0x92, 0x93, 0x38, 0xa6, 0x97, 0x02, 0x40, 0xaa, 0x03, 0x18, 0xa3, 0xfc, 0xc8, 0x48, 0x4f,
	0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xd8, 0x53, 0x50, 0x12, 0x1b, 0x58, 0x89, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xb1, 0x49, 0x9a, 0x8b, 0xa8, 0x02, 0x00, 0x00,
}
