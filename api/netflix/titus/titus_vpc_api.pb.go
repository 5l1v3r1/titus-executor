// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netflix/titus/titus_vpc_api.proto

package titus

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
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

type Family int32

const (
	// Default should never really be used, but we're required to have one due to protobuf
	Family_FAMILY_DEFAULT Family = 0
	Family_FAMILY_V4      Family = 1
	Family_FAMILY_V6      Family = 2
)

var Family_name = map[int32]string{
	0: "FAMILY_DEFAULT",
	1: "FAMILY_V4",
	2: "FAMILY_V6",
}

var Family_value = map[string]int32{
	"FAMILY_DEFAULT": 0,
	"FAMILY_V4":      1,
	"FAMILY_V6":      2,
}

func (x Family) String() string {
	return proto.EnumName(Family_name, int32(x))
}

func (Family) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fee69f8d9b907b12, []int{0}
}

type AllocateAddressRequest struct {
	AddressAllocation    *AddressAllocation `protobuf:"bytes,1,opt,name=addressAllocation,proto3" json:"addressAllocation,omitempty"`
	Family               Family             `protobuf:"varint,2,opt,name=family,proto3,enum=com.netflix.titus.Family" json:"family,omitempty"`
	AccountId            string             `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AllocateAddressRequest) Reset()         { *m = AllocateAddressRequest{} }
func (m *AllocateAddressRequest) String() string { return proto.CompactTextString(m) }
func (*AllocateAddressRequest) ProtoMessage()    {}
func (*AllocateAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fee69f8d9b907b12, []int{0}
}

func (m *AllocateAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateAddressRequest.Unmarshal(m, b)
}
func (m *AllocateAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateAddressRequest.Marshal(b, m, deterministic)
}
func (m *AllocateAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateAddressRequest.Merge(m, src)
}
func (m *AllocateAddressRequest) XXX_Size() int {
	return xxx_messageInfo_AllocateAddressRequest.Size(m)
}
func (m *AllocateAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateAddressRequest proto.InternalMessageInfo

func (m *AllocateAddressRequest) GetAddressAllocation() *AddressAllocation {
	if m != nil {
		return m.AddressAllocation
	}
	return nil
}

func (m *AllocateAddressRequest) GetFamily() Family {
	if m != nil {
		return m.Family
	}
	return Family_FAMILY_DEFAULT
}

func (m *AllocateAddressRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type AllocateAddressResponse struct {
	SignedAddressAllocation *SignedAddressAllocation `protobuf:"bytes,1,opt,name=signedAddressAllocation,proto3" json:"signedAddressAllocation,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *AllocateAddressResponse) Reset()         { *m = AllocateAddressResponse{} }
func (m *AllocateAddressResponse) String() string { return proto.CompactTextString(m) }
func (*AllocateAddressResponse) ProtoMessage()    {}
func (*AllocateAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fee69f8d9b907b12, []int{1}
}

func (m *AllocateAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateAddressResponse.Unmarshal(m, b)
}
func (m *AllocateAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateAddressResponse.Marshal(b, m, deterministic)
}
func (m *AllocateAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateAddressResponse.Merge(m, src)
}
func (m *AllocateAddressResponse) XXX_Size() int {
	return xxx_messageInfo_AllocateAddressResponse.Size(m)
}
func (m *AllocateAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateAddressResponse proto.InternalMessageInfo

func (m *AllocateAddressResponse) GetSignedAddressAllocation() *SignedAddressAllocation {
	if m != nil {
		return m.SignedAddressAllocation
	}
	return nil
}

type GetAllocationRequest struct {
	// Types that are valid to be assigned to SearchParameter:
	//	*GetAllocationRequest_Address
	//	*GetAllocationRequest_Uuid
	SearchParameter      isGetAllocationRequest_SearchParameter `protobuf_oneof:"searchParameter"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *GetAllocationRequest) Reset()         { *m = GetAllocationRequest{} }
func (m *GetAllocationRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllocationRequest) ProtoMessage()    {}
func (*GetAllocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fee69f8d9b907b12, []int{2}
}

func (m *GetAllocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllocationRequest.Unmarshal(m, b)
}
func (m *GetAllocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllocationRequest.Marshal(b, m, deterministic)
}
func (m *GetAllocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllocationRequest.Merge(m, src)
}
func (m *GetAllocationRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllocationRequest.Size(m)
}
func (m *GetAllocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllocationRequest proto.InternalMessageInfo

type isGetAllocationRequest_SearchParameter interface {
	isGetAllocationRequest_SearchParameter()
}

type GetAllocationRequest_Address struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type GetAllocationRequest_Uuid struct {
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3,oneof"`
}

func (*GetAllocationRequest_Address) isGetAllocationRequest_SearchParameter() {}

func (*GetAllocationRequest_Uuid) isGetAllocationRequest_SearchParameter() {}

func (m *GetAllocationRequest) GetSearchParameter() isGetAllocationRequest_SearchParameter {
	if m != nil {
		return m.SearchParameter
	}
	return nil
}

func (m *GetAllocationRequest) GetAddress() string {
	if x, ok := m.GetSearchParameter().(*GetAllocationRequest_Address); ok {
		return x.Address
	}
	return ""
}

func (m *GetAllocationRequest) GetUuid() string {
	if x, ok := m.GetSearchParameter().(*GetAllocationRequest_Uuid); ok {
		return x.Uuid
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetAllocationRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetAllocationRequest_Address)(nil),
		(*GetAllocationRequest_Uuid)(nil),
	}
}

type GetAllocationResponse struct {
	AddressAllocation    *AddressAllocation `protobuf:"bytes,1,opt,name=addressAllocation,proto3" json:"addressAllocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAllocationResponse) Reset()         { *m = GetAllocationResponse{} }
func (m *GetAllocationResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllocationResponse) ProtoMessage()    {}
func (*GetAllocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fee69f8d9b907b12, []int{3}
}

func (m *GetAllocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllocationResponse.Unmarshal(m, b)
}
func (m *GetAllocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllocationResponse.Marshal(b, m, deterministic)
}
func (m *GetAllocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllocationResponse.Merge(m, src)
}
func (m *GetAllocationResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllocationResponse.Size(m)
}
func (m *GetAllocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllocationResponse proto.InternalMessageInfo

func (m *GetAllocationResponse) GetAddressAllocation() *AddressAllocation {
	if m != nil {
		return m.AddressAllocation
	}
	return nil
}

func init() {
	proto.RegisterEnum("com.netflix.titus.Family", Family_name, Family_value)
	proto.RegisterType((*AllocateAddressRequest)(nil), "com.netflix.titus.AllocateAddressRequest")
	proto.RegisterType((*AllocateAddressResponse)(nil), "com.netflix.titus.AllocateAddressResponse")
	proto.RegisterType((*GetAllocationRequest)(nil), "com.netflix.titus.GetAllocationRequest")
	proto.RegisterType((*GetAllocationResponse)(nil), "com.netflix.titus.GetAllocationResponse")
}

func init() { proto.RegisterFile("netflix/titus/titus_vpc_api.proto", fileDescriptor_fee69f8d9b907b12) }

var fileDescriptor_fee69f8d9b907b12 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0xaf, 0xd2, 0x40,
	0x10, 0xc7, 0x5f, 0x9f, 0xca, 0x4b, 0xc7, 0xf0, 0x1e, 0x6c, 0x9e, 0x82, 0x8d, 0x51, 0x6c, 0x4c,
	0xac, 0x1c, 0x4a, 0x44, 0xe3, 0xc1, 0x5b, 0x89, 0xa2, 0x24, 0x98, 0x90, 0x22, 0x26, 0xea, 0x81,
	0x6c, 0xdb, 0x01, 0x36, 0xb6, 0xdd, 0xda, 0xdd, 0x12, 0x3d, 0xf9, 0x97, 0xf9, 0xa7, 0xf8, 0xbf,
	0x18, 0xb7, 0x6b, 0x44, 0x58, 0xa2, 0x17, 0x2f, 0x4d, 0x66, 0xe6, 0x33, 0xbf, 0xbe, 0xd3, 0x85,
	0x7b, 0x39, 0xca, 0x55, 0xca, 0x3e, 0x0f, 0x24, 0x93, 0x95, 0xa8, 0xbf, 0xcb, 0x6d, 0x11, 0x2f,
	0x69, 0xc1, 0xfc, 0xa2, 0xe4, 0x92, 0x93, 0x76, 0xcc, 0x33, 0x5f, 0x63, 0xbe, 0x02, 0x9c, 0xbb,
	0x6b, 0xce, 0xd7, 0x29, 0x0e, 0x14, 0x10, 0x55, 0xab, 0x81, 0x64, 0x19, 0x0a, 0x49, 0xb3, 0xa2,
	0xce, 0x71, 0xee, 0x98, 0xca, 0x46, 0x54, 0x60, 0x1d, 0x77, 0xbf, 0x59, 0x70, 0x33, 0x48, 0x53,
	0x1e, 0x53, 0x89, 0x41, 0x92, 0x94, 0x28, 0x44, 0x88, 0x9f, 0x2a, 0x14, 0x92, 0x84, 0xd0, 0xa6,
	0xb5, 0x47, 0x03, 0x8c, 0xe7, 0x5d, 0xab, 0x67, 0x79, 0xd7, 0x87, 0xf7, 0xfd, 0x83, 0x51, 0xfc,
	0x60, 0x9f, 0x0d, 0x0f, 0xd3, 0xc9, 0x23, 0x68, 0xac, 0x68, 0xc6, 0xd2, 0x2f, 0xdd, 0xd3, 0x9e,
	0xe5, 0x9d, 0x0f, 0x6f, 0x19, 0x0a, 0x8d, 0x15, 0x10, 0x6a, 0x90, 0xdc, 0x06, 0x9b, 0xc6, 0x31,
	0xaf, 0x72, 0x39, 0x49, 0xba, 0x57, 0x7a, 0x96, 0x67, 0x87, 0xbf, 0x1d, 0xee, 0x57, 0xe8, 0x1c,
	0x8c, 0x2f, 0x0a, 0x9e, 0x0b, 0x24, 0x09, 0x74, 0x04, 0x5b, 0xe7, 0x98, 0x04, 0x47, 0xb6, 0xe8,
	0x1b, 0x9a, 0xcf, 0xcd, 0x19, 0xe1, 0xb1, 0x52, 0xee, 0x07, 0xb8, 0x7c, 0x89, 0x72, 0x87, 0xd4,
	0xea, 0x39, 0x70, 0xa6, 0xd7, 0x57, 0xdd, 0xec, 0x57, 0x27, 0xe1, 0x2f, 0x07, 0xb9, 0x84, 0xab,
	0x55, 0xc5, 0x12, 0xa5, 0xc1, 0xcf, 0x80, 0xb2, 0x46, 0x6d, 0xb8, 0x10, 0x48, 0xcb, 0x78, 0x33,
	0xa3, 0x25, 0xcd, 0x50, 0x62, 0xe9, 0x7e, 0x84, 0x1b, 0x7b, 0xc5, 0xf5, 0x6e, 0xff, 0xe1, 0x36,
	0xfd, 0x67, 0xd0, 0xa8, 0xa5, 0x27, 0x04, 0xce, 0xc7, 0xc1, 0xeb, 0xc9, 0xf4, 0xdd, 0xf2, 0xf9,
	0x8b, 0x71, 0xb0, 0x98, 0xbe, 0x69, 0x9d, 0x90, 0x26, 0xd8, 0xda, 0xf7, 0xf6, 0x49, 0xcb, 0xda,
	0x35, 0x9f, 0xb6, 0x4e, 0x87, 0xdf, 0x2d, 0x68, 0x2e, 0x04, 0x96, 0x93, 0xd9, 0x1c, 0xcb, 0x2d,
	0x8b, 0x91, 0x6c, 0xe0, 0x62, 0xef, 0x30, 0xe4, 0xa1, 0x69, 0x32, 0xe3, 0xbf, 0xe7, 0xf4, 0xff,
	0x05, 0xd5, 0x5a, 0x44, 0xd0, 0xfc, 0x43, 0x24, 0xf2, 0xc0, 0x90, 0x6c, 0xba, 0x91, 0xe3, 0xfd,
	0x1d, 0xac, 0x7b, 0x8c, 0xce, 0xde, 0x5f, 0x53, 0xe1, 0xa8, 0xa1, 0x9e, 0xcd, 0xe3, 0x1f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3c, 0xb5, 0x67, 0xb3, 0xaf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserIPServiceClient is the client API for UserIPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserIPServiceClient interface {
	// Static IP Address flow
	AllocateAddress(ctx context.Context, in *AllocateAddressRequest, opts ...grpc.CallOption) (*AllocateAddressResponse, error)
	GetAllocation(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*GetAllocationResponse, error)
}

type userIPServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserIPServiceClient(cc *grpc.ClientConn) UserIPServiceClient {
	return &userIPServiceClient{cc}
}

func (c *userIPServiceClient) AllocateAddress(ctx context.Context, in *AllocateAddressRequest, opts ...grpc.CallOption) (*AllocateAddressResponse, error) {
	out := new(AllocateAddressResponse)
	err := c.cc.Invoke(ctx, "/com.netflix.titus.UserIPService/AllocateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userIPServiceClient) GetAllocation(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*GetAllocationResponse, error) {
	out := new(GetAllocationResponse)
	err := c.cc.Invoke(ctx, "/com.netflix.titus.UserIPService/GetAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserIPServiceServer is the server API for UserIPService service.
type UserIPServiceServer interface {
	// Static IP Address flow
	AllocateAddress(context.Context, *AllocateAddressRequest) (*AllocateAddressResponse, error)
	GetAllocation(context.Context, *GetAllocationRequest) (*GetAllocationResponse, error)
}

func RegisterUserIPServiceServer(s *grpc.Server, srv UserIPServiceServer) {
	s.RegisterService(&_UserIPService_serviceDesc, srv)
}

func _UserIPService_AllocateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserIPServiceServer).AllocateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.netflix.titus.UserIPService/AllocateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserIPServiceServer).AllocateAddress(ctx, req.(*AllocateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserIPService_GetAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserIPServiceServer).GetAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.netflix.titus.UserIPService/GetAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserIPServiceServer).GetAllocation(ctx, req.(*GetAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserIPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.netflix.titus.UserIPService",
	HandlerType: (*UserIPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocateAddress",
			Handler:    _UserIPService_AllocateAddress_Handler,
		},
		{
			MethodName: "GetAllocation",
			Handler:    _UserIPService_GetAllocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netflix/titus/titus_vpc_api.proto",
}
