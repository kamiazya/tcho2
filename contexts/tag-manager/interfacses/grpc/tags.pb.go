// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tags.proto

package tags

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

type Tag struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_e517adfcfbb19b09, []int{0}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_e517adfcfbb19b09, []int{1}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tags_e517adfcfbb19b09, []int{2}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Tag)(nil), "tags.Tag")
	proto.RegisterType((*GetRequest)(nil), "tags.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "tags.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TagsServiceClient is the client API for TagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagsServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type tagsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTagsServiceClient(cc *grpc.ClientConn) TagsServiceClient {
	return &tagsServiceClient{cc}
}

func (c *tagsServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/tags.TagsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServiceServer is the server API for TagsService service.
type TagsServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterTagsServiceServer(s *grpc.Server, srv TagsServiceServer) {
	s.RegisterService(&_TagsService_serviceDesc, srv)
}

func _TagsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tags.TagsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tags.TagsService",
	HandlerType: (*TagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TagsService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tags.proto",
}

func init() { proto.RegisterFile("tags.proto", fileDescriptor_tags_e517adfcfbb19b09) }

var fileDescriptor_tags_e517adfcfbb19b09 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4c, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x34, 0xb9, 0x98, 0x43, 0x12,
	0xd3, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32,
	0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0x22, 0x60, 0xb6, 0x92,
	0x0c, 0x17, 0x97, 0x7b, 0x6a, 0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0xba, 0x0e, 0x25,
	0x1d, 0x2e, 0x6e, 0xb0, 0x6c, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x2c, 0x17, 0xd8, 0x7c,
	0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x4e, 0x3d, 0xb0, 0xc5, 0x21, 0x89, 0xe9, 0x41, 0x60,
	0x61, 0x23, 0x6b, 0x2e, 0xee, 0x90, 0xc4, 0xf4, 0xe2, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0x21, 0x1d, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x01, 0x88, 0x32, 0x84, 0x2d, 0x52, 0x82, 0x48,
	0x22, 0x10, 0x93, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x1e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x57, 0xe6, 0x8e, 0x36, 0xce, 0x00, 0x00, 0x00,
}
