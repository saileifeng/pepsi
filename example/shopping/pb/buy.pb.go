// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buy.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//req->用户编号，[{商品编号，商品数量}]
type BuyGoodsRequest struct {
	UserID               int64        `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	GoodsInfos           []*GoodsInfo `protobuf:"bytes,2,rep,name=goodsInfos,proto3" json:"goodsInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BuyGoodsRequest) Reset()         { *m = BuyGoodsRequest{} }
func (m *BuyGoodsRequest) String() string { return proto.CompactTextString(m) }
func (*BuyGoodsRequest) ProtoMessage()    {}
func (*BuyGoodsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44998f98de420a65, []int{0}
}

func (m *BuyGoodsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyGoodsRequest.Unmarshal(m, b)
}
func (m *BuyGoodsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyGoodsRequest.Marshal(b, m, deterministic)
}
func (m *BuyGoodsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyGoodsRequest.Merge(m, src)
}
func (m *BuyGoodsRequest) XXX_Size() int {
	return xxx_messageInfo_BuyGoodsRequest.Size(m)
}
func (m *BuyGoodsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyGoodsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyGoodsRequest proto.InternalMessageInfo

func (m *BuyGoodsRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *BuyGoodsRequest) GetGoodsInfos() []*GoodsInfo {
	if m != nil {
		return m.GoodsInfos
	}
	return nil
}

//resp->用户编号，订单编号，[{商品编号，商品数量，商品单价}]，商品总价
type BuyGoodsResponse struct {
	OrderID              int64        `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	UserID               int64        `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	GoodsInfos           []*GoodsInfo `protobuf:"bytes,3,rep,name=goodsInfos,proto3" json:"goodsInfos,omitempty"`
	TotalPrice           int64        `protobuf:"varint,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BuyGoodsResponse) Reset()         { *m = BuyGoodsResponse{} }
func (m *BuyGoodsResponse) String() string { return proto.CompactTextString(m) }
func (*BuyGoodsResponse) ProtoMessage()    {}
func (*BuyGoodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44998f98de420a65, []int{1}
}

func (m *BuyGoodsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyGoodsResponse.Unmarshal(m, b)
}
func (m *BuyGoodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyGoodsResponse.Marshal(b, m, deterministic)
}
func (m *BuyGoodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyGoodsResponse.Merge(m, src)
}
func (m *BuyGoodsResponse) XXX_Size() int {
	return xxx_messageInfo_BuyGoodsResponse.Size(m)
}
func (m *BuyGoodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyGoodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuyGoodsResponse proto.InternalMessageInfo

func (m *BuyGoodsResponse) GetOrderID() int64 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

func (m *BuyGoodsResponse) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *BuyGoodsResponse) GetGoodsInfos() []*GoodsInfo {
	if m != nil {
		return m.GoodsInfos
	}
	return nil
}

func (m *BuyGoodsResponse) GetTotalPrice() int64 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func init() {
	proto.RegisterType((*BuyGoodsRequest)(nil), "pb.BuyGoodsRequest")
	proto.RegisterType((*BuyGoodsResponse)(nil), "pb.BuyGoodsResponse")
}

func init() { proto.RegisterFile("buy.proto", fileDescriptor_44998f98de420a65) }

var fileDescriptor_44998f98de420a65 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2a, 0xad, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x4e, 0xcf, 0xcf, 0x4f, 0x29,
	0x86, 0x08, 0x28, 0x45, 0x70, 0xf1, 0x3b, 0x95, 0x56, 0xba, 0x83, 0x44, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0x4a, 0x8b, 0x53, 0x8b, 0x3c, 0x5d, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x98, 0x83, 0xa0, 0x3c, 0x21, 0x5d, 0x2e, 0x2e, 0xb0, 0x4e, 0xcf, 0xbc, 0xb4, 0xfc,
	0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x5e, 0xbd, 0x82, 0x24, 0x3d, 0x77, 0x98, 0x68,
	0x10, 0x92, 0x02, 0xa5, 0xc9, 0x8c, 0x5c, 0x02, 0x08, 0xa3, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x85, 0x24, 0xb8, 0xd8, 0xf3, 0x8b, 0x52, 0x90, 0x0c, 0x87, 0x71, 0x91, 0x6c, 0x65, 0xc2, 0x63,
	0x2b, 0x33, 0x01, 0x5b, 0x85, 0xe4, 0xb8, 0xb8, 0x4a, 0xf2, 0x4b, 0x12, 0x73, 0x02, 0x8a, 0x32,
	0x93, 0x53, 0x25, 0x58, 0xc0, 0x46, 0x21, 0x89, 0x18, 0x39, 0x73, 0x71, 0x39, 0x95, 0x56, 0x06,
	0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x99, 0x72, 0x71, 0xc0, 0x9c, 0x28, 0x24, 0x0c, 0x32,
	0x14, 0x2d, 0x2c, 0xa4, 0x44, 0x50, 0x05, 0x21, 0xbe, 0x48, 0x62, 0x03, 0x87, 0x9d, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xf8, 0xe1, 0xbc, 0x6a, 0x59, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuyServiceClient is the client API for BuyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuyServiceClient interface {
	//生成购买订单
	BuyGoods(ctx context.Context, in *BuyGoodsRequest, opts ...grpc.CallOption) (*BuyGoodsResponse, error)
}

type buyServiceClient struct {
	cc *grpc.ClientConn
}

func NewBuyServiceClient(cc *grpc.ClientConn) BuyServiceClient {
	return &buyServiceClient{cc}
}

func (c *buyServiceClient) BuyGoods(ctx context.Context, in *BuyGoodsRequest, opts ...grpc.CallOption) (*BuyGoodsResponse, error) {
	out := new(BuyGoodsResponse)
	err := c.cc.Invoke(ctx, "/pb.BuyService/BuyGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuyServiceServer is the server API for BuyService service.
type BuyServiceServer interface {
	//生成购买订单
	BuyGoods(context.Context, *BuyGoodsRequest) (*BuyGoodsResponse, error)
}

// UnimplementedBuyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBuyServiceServer struct {
}

func (*UnimplementedBuyServiceServer) BuyGoods(ctx context.Context, req *BuyGoodsRequest) (*BuyGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyGoods not implemented")
}

func RegisterBuyServiceServer(s *grpc.Server, srv BuyServiceServer) {
	s.RegisterService(&_BuyService_serviceDesc, srv)
}

func _BuyService_BuyGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyServiceServer).BuyGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BuyService/BuyGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyServiceServer).BuyGoods(ctx, req.(*BuyGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BuyService",
	HandlerType: (*BuyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyGoods",
			Handler:    _BuyService_BuyGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buy.proto",
}
