// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prcibc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf6b2ed3f2682f8, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf6b2ed3f2682f8, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetIbcMsgRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetIbcMsgRequest) Reset()         { *m = QueryGetIbcMsgRequest{} }
func (m *QueryGetIbcMsgRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetIbcMsgRequest) ProtoMessage()    {}
func (*QueryGetIbcMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf6b2ed3f2682f8, []int{2}
}
func (m *QueryGetIbcMsgRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetIbcMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetIbcMsgRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetIbcMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetIbcMsgRequest.Merge(m, src)
}
func (m *QueryGetIbcMsgRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetIbcMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetIbcMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetIbcMsgRequest proto.InternalMessageInfo

func (m *QueryGetIbcMsgRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetIbcMsgResponse struct {
	IbcMsg IbcMsg `protobuf:"bytes,1,opt,name=IbcMsg,proto3" json:"IbcMsg"`
}

func (m *QueryGetIbcMsgResponse) Reset()         { *m = QueryGetIbcMsgResponse{} }
func (m *QueryGetIbcMsgResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetIbcMsgResponse) ProtoMessage()    {}
func (*QueryGetIbcMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf6b2ed3f2682f8, []int{3}
}
func (m *QueryGetIbcMsgResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetIbcMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetIbcMsgResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetIbcMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetIbcMsgResponse.Merge(m, src)
}
func (m *QueryGetIbcMsgResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetIbcMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetIbcMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetIbcMsgResponse proto.InternalMessageInfo

func (m *QueryGetIbcMsgResponse) GetIbcMsg() IbcMsg {
	if m != nil {
		return m.IbcMsg
	}
	return IbcMsg{}
}

type QueryAllIbcMsgRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllIbcMsgRequest) Reset()         { *m = QueryAllIbcMsgRequest{} }
func (m *QueryAllIbcMsgRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllIbcMsgRequest) ProtoMessage()    {}
func (*QueryAllIbcMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf6b2ed3f2682f8, []int{4}
}
func (m *QueryAllIbcMsgRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllIbcMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllIbcMsgRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllIbcMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllIbcMsgRequest.Merge(m, src)
}
func (m *QueryAllIbcMsgRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllIbcMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllIbcMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllIbcMsgRequest proto.InternalMessageInfo

func (m *QueryAllIbcMsgRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllIbcMsgResponse struct {
	IbcMsg     []IbcMsg            `protobuf:"bytes,1,rep,name=IbcMsg,proto3" json:"IbcMsg"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllIbcMsgResponse) Reset()         { *m = QueryAllIbcMsgResponse{} }
func (m *QueryAllIbcMsgResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllIbcMsgResponse) ProtoMessage()    {}
func (*QueryAllIbcMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf6b2ed3f2682f8, []int{5}
}
func (m *QueryAllIbcMsgResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllIbcMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllIbcMsgResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllIbcMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllIbcMsgResponse.Merge(m, src)
}
func (m *QueryAllIbcMsgResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllIbcMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllIbcMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllIbcMsgResponse proto.InternalMessageInfo

func (m *QueryAllIbcMsgResponse) GetIbcMsg() []IbcMsg {
	if m != nil {
		return m.IbcMsg
	}
	return nil
}

func (m *QueryAllIbcMsgResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "pricechain.cprc.prcibc.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "pricechain.cprc.prcibc.QueryParamsResponse")
	proto.RegisterType((*QueryGetIbcMsgRequest)(nil), "pricechain.cprc.prcibc.QueryGetIbcMsgRequest")
	proto.RegisterType((*QueryGetIbcMsgResponse)(nil), "pricechain.cprc.prcibc.QueryGetIbcMsgResponse")
	proto.RegisterType((*QueryAllIbcMsgRequest)(nil), "pricechain.cprc.prcibc.QueryAllIbcMsgRequest")
	proto.RegisterType((*QueryAllIbcMsgResponse)(nil), "pricechain.cprc.prcibc.QueryAllIbcMsgResponse")
}

func init() { proto.RegisterFile("prcibc/query.proto", fileDescriptor_7cf6b2ed3f2682f8) }

var fileDescriptor_7cf6b2ed3f2682f8 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6b, 0x14, 0x3f,
	0x18, 0xc6, 0x37, 0xdb, 0xfe, 0x17, 0xfe, 0x11, 0x3c, 0xa4, 0x6b, 0x91, 0x45, 0xa6, 0x35, 0x48,
	0xb7, 0x56, 0x4d, 0x68, 0xbd, 0x7a, 0x69, 0x0b, 0x16, 0x0f, 0xc2, 0xba, 0x82, 0x07, 0x2f, 0x4b,
	0x26, 0x1b, 0xd2, 0xc0, 0xec, 0x24, 0x9d, 0x64, 0xc5, 0x22, 0x5e, 0xbc, 0x7a, 0x11, 0x8a, 0x67,
	0x3f, 0x8c, 0x97, 0x1e, 0x0b, 0x5e, 0x3c, 0x89, 0xec, 0xfa, 0x41, 0x64, 0x92, 0x2c, 0xdd, 0xd9,
	0xad, 0x4e, 0xbd, 0xcd, 0xbc, 0xf3, 0x3c, 0xcf, 0xfb, 0x9b, 0x37, 0xef, 0x0c, 0x44, 0xa6, 0xe0,
	0x2a, 0xe5, 0xf4, 0x64, 0x2c, 0x8a, 0x53, 0x62, 0x0a, 0xed, 0x34, 0x5a, 0x37, 0x85, 0xe2, 0x82,
	0x1f, 0x33, 0x95, 0x13, 0x6e, 0x0a, 0x4e, 0x82, 0xa6, 0xd3, 0x96, 0x5a, 0x6a, 0x2f, 0xa1, 0xe5,
	0x55, 0x50, 0x77, 0xee, 0x48, 0xad, 0x65, 0x26, 0x28, 0x33, 0x8a, 0xb2, 0x3c, 0xd7, 0x8e, 0x39,
	0xa5, 0x73, 0x1b, 0x9f, 0xee, 0x70, 0x6d, 0x47, 0xda, 0xd2, 0x94, 0x59, 0x11, 0x9a, 0xd0, 0x37,
	0xbb, 0xa9, 0x70, 0x6c, 0x97, 0x1a, 0x26, 0x55, 0xee, 0xc5, 0x51, 0xbb, 0x16, 0x59, 0x0c, 0x2b,
	0xd8, 0x68, 0x16, 0xd0, 0x8e, 0x45, 0x95, 0xf2, 0xc1, 0xc8, 0xca, 0x50, 0xc5, 0x6d, 0x88, 0x5e,
	0x94, 0x61, 0x3d, 0x2f, 0xed, 0x8b, 0x93, 0xb1, 0xb0, 0x0e, 0xbf, 0x84, 0x6b, 0x95, 0xaa, 0x35,
	0x3a, 0xb7, 0x02, 0x3d, 0x81, 0xad, 0x10, 0x79, 0x1b, 0x6c, 0x82, 0xed, 0x1b, 0x7b, 0x09, 0xb9,
	0xfa, 0x05, 0x49, 0xf0, 0x1d, 0xac, 0x9e, 0xff, 0xd8, 0x68, 0xf4, 0xa3, 0x07, 0x77, 0xe1, 0x2d,
	0x1f, 0x7a, 0x24, 0xdc, 0xb3, 0x94, 0x3f, 0xb7, 0x32, 0x76, 0x43, 0x37, 0x61, 0x53, 0x0d, 0x7d,
	0xe4, 0x6a, 0xbf, 0xa9, 0x86, 0xf8, 0x15, 0x5c, 0x5f, 0x14, 0x5e, 0x02, 0x84, 0x4a, 0x1d, 0x40,
	0x50, 0xcd, 0x00, 0xc2, 0x1d, 0x1e, 0x44, 0x80, 0xfd, 0x2c, 0xab, 0x02, 0x3c, 0x85, 0xf0, 0x72,
	0x86, 0x31, 0x7a, 0x8b, 0x84, 0x81, 0x93, 0x72, 0xe0, 0x24, 0x9c, 0x6a, 0x1c, 0x38, 0xe9, 0x31,
	0x29, 0xa2, 0xb7, 0x3f, 0xe7, 0xc4, 0x5f, 0x40, 0x24, 0x9f, 0xeb, 0x70, 0x05, 0xf9, 0xca, 0xbf,
	0x92, 0xa3, 0xa3, 0x0a, 0x60, 0xd3, 0x03, 0x76, 0x6b, 0x01, 0x43, 0xeb, 0x79, 0xc2, 0xbd, 0xaf,
	0x2b, 0xf0, 0x3f, 0x4f, 0x88, 0x3e, 0x02, 0xd8, 0x0a, 0xc7, 0x84, 0x76, 0xfe, 0xc4, 0xb2, 0xbc,
	0x19, 0x9d, 0x07, 0xd7, 0xd2, 0x86, 0xce, 0x78, 0xfb, 0xc3, 0xb7, 0x5f, 0x67, 0x4d, 0x8c, 0x36,
	0x69, 0xaf, 0x34, 0x1d, 0x96, 0x26, 0x5a, 0x0c, 0x07, 0xb9, 0x70, 0xb4, 0xb2, 0xa2, 0xe8, 0x33,
	0x98, 0xcd, 0x07, 0x3d, 0xfa, 0x6b, 0x87, 0xc5, 0xe5, 0xe9, 0x90, 0xeb, 0xca, 0x23, 0xd3, 0x43,
	0xcf, 0xb4, 0x85, 0xee, 0xcd, 0x33, 0x95, 0x3e, 0x5a, 0xfd, 0x3e, 0xe8, 0x3b, 0x35, 0x7c, 0x8f,
	0xce, 0x00, 0xfc, 0x3f, 0x04, 0xec, 0x67, 0x59, 0x0d, 0xda, 0xe2, 0x5a, 0xd5, 0xa0, 0x2d, 0xed,
	0x08, 0xee, 0x7a, 0xb4, 0xbb, 0x68, 0xa3, 0x06, 0xed, 0xe0, 0xf0, 0x7c, 0x92, 0x80, 0x8b, 0x49,
	0x02, 0x7e, 0x4e, 0x12, 0xf0, 0x69, 0x9a, 0x34, 0x2e, 0xa6, 0x49, 0xe3, 0xfb, 0x34, 0x69, 0xbc,
	0xbe, 0x2f, 0x95, 0x3b, 0x1e, 0xa7, 0x84, 0xeb, 0xd1, 0x52, 0xc8, 0xdb, 0x59, 0x8c, 0x3b, 0x35,
	0xc2, 0xa6, 0x2d, 0xff, 0x03, 0x78, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x68, 0xec, 0x44, 0x84,
	0xb9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a IbcMsg by id.
	IbcMsg(ctx context.Context, in *QueryGetIbcMsgRequest, opts ...grpc.CallOption) (*QueryGetIbcMsgResponse, error)
	// Queries a list of IbcMsg items.
	IbcMsgAll(ctx context.Context, in *QueryAllIbcMsgRequest, opts ...grpc.CallOption) (*QueryAllIbcMsgResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/pricechain.cprc.prcibc.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IbcMsg(ctx context.Context, in *QueryGetIbcMsgRequest, opts ...grpc.CallOption) (*QueryGetIbcMsgResponse, error) {
	out := new(QueryGetIbcMsgResponse)
	err := c.cc.Invoke(ctx, "/pricechain.cprc.prcibc.Query/IbcMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IbcMsgAll(ctx context.Context, in *QueryAllIbcMsgRequest, opts ...grpc.CallOption) (*QueryAllIbcMsgResponse, error) {
	out := new(QueryAllIbcMsgResponse)
	err := c.cc.Invoke(ctx, "/pricechain.cprc.prcibc.Query/IbcMsgAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a IbcMsg by id.
	IbcMsg(context.Context, *QueryGetIbcMsgRequest) (*QueryGetIbcMsgResponse, error)
	// Queries a list of IbcMsg items.
	IbcMsgAll(context.Context, *QueryAllIbcMsgRequest) (*QueryAllIbcMsgResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) IbcMsg(ctx context.Context, req *QueryGetIbcMsgRequest) (*QueryGetIbcMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcMsg not implemented")
}
func (*UnimplementedQueryServer) IbcMsgAll(ctx context.Context, req *QueryAllIbcMsgRequest) (*QueryAllIbcMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcMsgAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricechain.cprc.prcibc.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IbcMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetIbcMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IbcMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricechain.cprc.prcibc.Query/IbcMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IbcMsg(ctx, req.(*QueryGetIbcMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IbcMsgAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllIbcMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IbcMsgAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricechain.cprc.prcibc.Query/IbcMsgAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IbcMsgAll(ctx, req.(*QueryAllIbcMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pricechain.cprc.prcibc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "IbcMsg",
			Handler:    _Query_IbcMsg_Handler,
		},
		{
			MethodName: "IbcMsgAll",
			Handler:    _Query_IbcMsgAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prcibc/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetIbcMsgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetIbcMsgRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetIbcMsgRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetIbcMsgResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetIbcMsgResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetIbcMsgResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.IbcMsg.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllIbcMsgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllIbcMsgRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllIbcMsgRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllIbcMsgResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllIbcMsgResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllIbcMsgResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.IbcMsg) > 0 {
		for iNdEx := len(m.IbcMsg) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IbcMsg[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetIbcMsgRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetIbcMsgResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.IbcMsg.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllIbcMsgRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllIbcMsgResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IbcMsg) > 0 {
		for _, e := range m.IbcMsg {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetIbcMsgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetIbcMsgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetIbcMsgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetIbcMsgResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetIbcMsgResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetIbcMsgResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IbcMsg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllIbcMsgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllIbcMsgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllIbcMsgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllIbcMsgResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllIbcMsgResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllIbcMsgResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcMsg = append(m.IbcMsg, IbcMsg{})
			if err := m.IbcMsg[len(m.IbcMsg)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
