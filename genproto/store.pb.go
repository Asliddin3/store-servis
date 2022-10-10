// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store.proto

package store

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StoreResponse struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Addresses            []*AddressResp `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0}
}
func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return m.Size()
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StoreResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoreResponse) GetAddresses() []*AddressResp {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type GetStoreInfoById struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoreInfoById) Reset()         { *m = GetStoreInfoById{} }
func (m *GetStoreInfoById) String() string { return proto.CompactTextString(m) }
func (*GetStoreInfoById) ProtoMessage()    {}
func (*GetStoreInfoById) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{1}
}
func (m *GetStoreInfoById) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStoreInfoById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStoreInfoById.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStoreInfoById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreInfoById.Merge(m, src)
}
func (m *GetStoreInfoById) XXX_Size() int {
	return m.Size()
}
func (m *GetStoreInfoById) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreInfoById.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreInfoById proto.InternalMessageInfo

func (m *GetStoreInfoById) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type StoreRequest struct {
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Addresses            []*Address `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{2}
}
func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return m.Size()
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoreRequest) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type AddressResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	District             string   `protobuf:"bytes,2,opt,name=district,proto3" json:"district"`
	Street               string   `protobuf:"bytes,3,opt,name=street,proto3" json:"street"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressResp) Reset()         { *m = AddressResp{} }
func (m *AddressResp) String() string { return proto.CompactTextString(m) }
func (*AddressResp) ProtoMessage()    {}
func (*AddressResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{3}
}
func (m *AddressResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressResp.Merge(m, src)
}
func (m *AddressResp) XXX_Size() int {
	return m.Size()
}
func (m *AddressResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddressResp proto.InternalMessageInfo

func (m *AddressResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AddressResp) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *AddressResp) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

type Address struct {
	District             string   `protobuf:"bytes,1,opt,name=district,proto3" json:"district"`
	Street               string   `protobuf:"bytes,2,opt,name=street,proto3" json:"street"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{4}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Address.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return m.Size()
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreResponse)(nil), "store.StoreResponse")
	proto.RegisterType((*GetStoreInfoById)(nil), "store.GetStoreInfoById")
	proto.RegisterType((*StoreRequest)(nil), "store.StoreRequest")
	proto.RegisterType((*AddressResp)(nil), "store.AddressResp")
	proto.RegisterType((*Address)(nil), "store.Address")
}

func init() { proto.RegisterFile("store.proto", fileDescriptor_98bbca36ef968dfc) }

var fileDescriptor_98bbca36ef968dfc = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0xb9, 0x04, 0x4a, 0xeb, 0x40, 0x55, 0x19, 0x04, 0x51, 0x87, 0xa8, 0xca, 0x94, 0x01,
	0x55, 0xa8, 0x2c, 0x30, 0x30, 0x50, 0x06, 0xd4, 0x0d, 0xc2, 0x13, 0x84, 0x9e, 0x91, 0x6e, 0x20,
	0x57, 0xce, 0x07, 0x12, 0x33, 0x2f, 0xc1, 0x23, 0x31, 0xf2, 0x08, 0x28, 0xbc, 0x08, 0xea, 0xf5,
	0x80, 0x34, 0x2d, 0x9b, 0xed, 0xf8, 0xf7, 0x67, 0xff, 0x39, 0x88, 0xd8, 0x6a, 0x43, 0xc3, 0x99,
	0xd1, 0x56, 0xe3, 0x96, 0x4b, 0x52, 0x82, 0xdd, 0xdb, 0x79, 0x90, 0x13, 0xcf, 0x74, 0xc9, 0x84,
	0x5d, 0x08, 0x94, 0x8c, 0xc5, 0x40, 0x64, 0x61, 0x1e, 0x28, 0x89, 0x08, 0x9b, 0x65, 0xf1, 0x40,
	0x71, 0x30, 0x10, 0x59, 0x27, 0x77, 0x31, 0x1e, 0x43, 0xa7, 0x90, 0xd2, 0x10, 0x33, 0x71, 0x1c,
	0x0e, 0xc2, 0x2c, 0x1a, 0xe1, 0x70, 0x31, 0xfc, 0x62, 0x51, 0x9f, 0x8f, 0xcb, 0xff, 0x9a, 0xd2,
	0x14, 0x7a, 0x57, 0x64, 0x1d, 0x69, 0x52, 0xde, 0xeb, 0xf1, 0xcb, 0x44, 0x36, 0x49, 0xe9, 0x35,
	0xec, 0xf8, 0x55, 0x1e, 0x9f, 0x88, 0xed, 0x5a, 0xf2, 0xd1, 0x2a, 0xb9, 0xdb, 0x20, 0xd7, 0xa8,
	0x37, 0x10, 0xd5, 0xf6, 0x59, 0x39, 0xad, 0x0f, 0x6d, 0xa9, 0xd8, 0x1a, 0x35, 0xb5, 0x1e, 0xf2,
	0x9b, 0xe3, 0x01, 0xb4, 0xd8, 0x1a, 0x22, 0x1b, 0x87, 0xee, 0x8b, 0xcf, 0xd2, 0x73, 0xd8, 0xf6,
	0x23, 0x97, 0xe4, 0xe2, 0x5f, 0x79, 0x50, 0x97, 0x8f, 0x5e, 0x85, 0x3f, 0x92, 0xc9, 0x3c, 0xab,
	0x29, 0xe1, 0x29, 0x44, 0x97, 0x86, 0x0a, 0x4b, 0xae, 0x8a, 0x7b, 0xfe, 0x98, 0xba, 0x11, 0xfd,
	0xfd, 0xe5, 0xa2, 0xff, 0x51, 0x67, 0xd0, 0xfe, 0xb1, 0x14, 0x0f, 0x7d, 0x47, 0xd3, 0xe3, 0xf5,
	0xd2, 0x71, 0xef, 0xbd, 0x4a, 0xc4, 0x47, 0x95, 0x88, 0xcf, 0x2a, 0x11, 0x6f, 0x5f, 0xc9, 0xc6,
	0x5d, 0xcb, 0x3d, 0x8a, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xfa, 0x1c, 0xb5, 0x23,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreserviceClient is the client API for Storeservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreserviceClient interface {
	CreateStore(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	GetStore(ctx context.Context, in *GetStoreInfoById, opts ...grpc.CallOption) (*StoreResponse, error)
}

type storeserviceClient struct {
	cc *grpc.ClientConn
}

func NewStoreserviceClient(cc *grpc.ClientConn) StoreserviceClient {
	return &storeserviceClient{cc}
}

func (c *storeserviceClient) CreateStore(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/store.Storeservice/CreateStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeserviceClient) GetStore(ctx context.Context, in *GetStoreInfoById, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/store.Storeservice/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreserviceServer is the server API for Storeservice service.
type StoreserviceServer interface {
	CreateStore(context.Context, *StoreRequest) (*StoreResponse, error)
	GetStore(context.Context, *GetStoreInfoById) (*StoreResponse, error)
}

// UnimplementedStoreserviceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreserviceServer struct {
}

func (*UnimplementedStoreserviceServer) CreateStore(ctx context.Context, req *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (*UnimplementedStoreserviceServer) GetStore(ctx context.Context, req *GetStoreInfoById) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStore not implemented")
}

func RegisterStoreserviceServer(s *grpc.Server, srv StoreserviceServer) {
	s.RegisterService(&_Storeservice_serviceDesc, srv)
}

func _Storeservice_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreserviceServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Storeservice/CreateStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreserviceServer).CreateStore(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storeservice_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreInfoById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreserviceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Storeservice/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreserviceServer).GetStore(ctx, req.(*GetStoreInfoById))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storeservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store.Storeservice",
	HandlerType: (*StoreserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStore",
			Handler:    _Storeservice_CreateStore_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _Storeservice_GetStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}

func (m *StoreResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Addresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetStoreInfoById) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStoreInfoById) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStoreInfoById) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StoreRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Addresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *AddressResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Street) > 0 {
		i -= len(m.Street)
		copy(dAtA[i:], m.Street)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Street)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.District) > 0 {
		i -= len(m.District)
		copy(dAtA[i:], m.District)
		i = encodeVarintStore(dAtA, i, uint64(len(m.District)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Address) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Address) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Street) > 0 {
		i -= len(m.Street)
		copy(dAtA[i:], m.Street)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Street)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.District) > 0 {
		i -= len(m.District)
		copy(dAtA[i:], m.District)
		i = encodeVarintStore(dAtA, i, uint64(len(m.District)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStore(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, e := range m.Addresses {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetStoreInfoById) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStore(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StoreRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, e := range m.Addresses {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AddressResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStore(uint64(m.Id))
	}
	l = len(m.District)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.Street)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.District)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.Street)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: StoreResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, &AddressResp{})
			if err := m.Addresses[len(m.Addresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetStoreInfoById) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: GetStoreInfoById: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStoreInfoById: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: StoreRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, &Address{})
			if err := m.Addresses[len(m.Addresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddressResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: AddressResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field District", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.District = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Street", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Street = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field District", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.District = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Street", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Street = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)