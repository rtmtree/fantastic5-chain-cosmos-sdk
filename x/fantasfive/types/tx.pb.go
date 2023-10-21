// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fantasfive/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateTeam struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Player0 string `protobuf:"bytes,2,opt,name=player0,proto3" json:"player0,omitempty"`
	Player1 string `protobuf:"bytes,3,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2 string `protobuf:"bytes,4,opt,name=player2,proto3" json:"player2,omitempty"`
	Player3 string `protobuf:"bytes,5,opt,name=player3,proto3" json:"player3,omitempty"`
	Player4 string `protobuf:"bytes,6,opt,name=player4,proto3" json:"player4,omitempty"`
}

func (m *MsgCreateTeam) Reset()         { *m = MsgCreateTeam{} }
func (m *MsgCreateTeam) String() string { return proto.CompactTextString(m) }
func (*MsgCreateTeam) ProtoMessage()    {}
func (*MsgCreateTeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f5978181a662cc, []int{0}
}
func (m *MsgCreateTeam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateTeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateTeam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateTeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateTeam.Merge(m, src)
}
func (m *MsgCreateTeam) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateTeam) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateTeam.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateTeam proto.InternalMessageInfo

func (m *MsgCreateTeam) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateTeam) GetPlayer0() string {
	if m != nil {
		return m.Player0
	}
	return ""
}

func (m *MsgCreateTeam) GetPlayer1() string {
	if m != nil {
		return m.Player1
	}
	return ""
}

func (m *MsgCreateTeam) GetPlayer2() string {
	if m != nil {
		return m.Player2
	}
	return ""
}

func (m *MsgCreateTeam) GetPlayer3() string {
	if m != nil {
		return m.Player3
	}
	return ""
}

func (m *MsgCreateTeam) GetPlayer4() string {
	if m != nil {
		return m.Player4
	}
	return ""
}

type MsgCreateTeamResponse struct {
	TeamId string `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
}

func (m *MsgCreateTeamResponse) Reset()         { *m = MsgCreateTeamResponse{} }
func (m *MsgCreateTeamResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateTeamResponse) ProtoMessage()    {}
func (*MsgCreateTeamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f5978181a662cc, []int{1}
}
func (m *MsgCreateTeamResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateTeamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateTeamResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateTeamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateTeamResponse.Merge(m, src)
}
func (m *MsgCreateTeamResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateTeamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateTeamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateTeamResponse proto.InternalMessageInfo

func (m *MsgCreateTeamResponse) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type MsgAnnounceAndCreateNextMw struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MwId       string `protobuf:"bytes,2,opt,name=mwId,proto3" json:"mwId,omitempty"`
	PlayerPerf string `protobuf:"bytes,3,opt,name=playerPerf,proto3" json:"playerPerf,omitempty"`
}

func (m *MsgAnnounceAndCreateNextMw) Reset()         { *m = MsgAnnounceAndCreateNextMw{} }
func (m *MsgAnnounceAndCreateNextMw) String() string { return proto.CompactTextString(m) }
func (*MsgAnnounceAndCreateNextMw) ProtoMessage()    {}
func (*MsgAnnounceAndCreateNextMw) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f5978181a662cc, []int{2}
}
func (m *MsgAnnounceAndCreateNextMw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnnounceAndCreateNextMw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnnounceAndCreateNextMw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnnounceAndCreateNextMw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnnounceAndCreateNextMw.Merge(m, src)
}
func (m *MsgAnnounceAndCreateNextMw) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnnounceAndCreateNextMw) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnnounceAndCreateNextMw.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnnounceAndCreateNextMw proto.InternalMessageInfo

func (m *MsgAnnounceAndCreateNextMw) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAnnounceAndCreateNextMw) GetMwId() string {
	if m != nil {
		return m.MwId
	}
	return ""
}

func (m *MsgAnnounceAndCreateNextMw) GetPlayerPerf() string {
	if m != nil {
		return m.PlayerPerf
	}
	return ""
}

type MsgAnnounceAndCreateNextMwResponse struct {
	NextMwId string `protobuf:"bytes,1,opt,name=nextMwId,proto3" json:"nextMwId,omitempty"`
}

func (m *MsgAnnounceAndCreateNextMwResponse) Reset()         { *m = MsgAnnounceAndCreateNextMwResponse{} }
func (m *MsgAnnounceAndCreateNextMwResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAnnounceAndCreateNextMwResponse) ProtoMessage()    {}
func (*MsgAnnounceAndCreateNextMwResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f5978181a662cc, []int{3}
}
func (m *MsgAnnounceAndCreateNextMwResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnnounceAndCreateNextMwResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnnounceAndCreateNextMwResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnnounceAndCreateNextMwResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnnounceAndCreateNextMwResponse.Merge(m, src)
}
func (m *MsgAnnounceAndCreateNextMwResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnnounceAndCreateNextMwResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnnounceAndCreateNextMwResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnnounceAndCreateNextMwResponse proto.InternalMessageInfo

func (m *MsgAnnounceAndCreateNextMwResponse) GetNextMwId() string {
	if m != nil {
		return m.NextMwId
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateTeam)(nil), "fantasfive.fantasfive.MsgCreateTeam")
	proto.RegisterType((*MsgCreateTeamResponse)(nil), "fantasfive.fantasfive.MsgCreateTeamResponse")
	proto.RegisterType((*MsgAnnounceAndCreateNextMw)(nil), "fantasfive.fantasfive.MsgAnnounceAndCreateNextMw")
	proto.RegisterType((*MsgAnnounceAndCreateNextMwResponse)(nil), "fantasfive.fantasfive.MsgAnnounceAndCreateNextMwResponse")
}

func init() { proto.RegisterFile("fantasfive/tx.proto", fileDescriptor_f5f5978181a662cc) }

var fileDescriptor_f5f5978181a662cc = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4b, 0xcc, 0x2b,
	0x49, 0x2c, 0x4e, 0xcb, 0x2c, 0x4b, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x45, 0x08, 0xea, 0x21, 0x98, 0x4a, 0xab, 0x19, 0xb9, 0x78, 0x7d, 0x8b, 0xd3, 0x9d, 0x8b,
	0x52, 0x13, 0x4b, 0x52, 0x43, 0x52, 0x13, 0x73, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x41, 0xbc, 0xfc,
	0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x17, 0x24, 0x53, 0x90, 0x93, 0x58, 0x99,
	0x5a, 0x64, 0x20, 0xc1, 0x04, 0x91, 0x81, 0x72, 0x11, 0x32, 0x86, 0x12, 0xcc, 0xc8, 0x32, 0x86,
	0x08, 0x19, 0x23, 0x09, 0x16, 0x64, 0x19, 0x23, 0x84, 0x8c, 0xb1, 0x04, 0x2b, 0xb2, 0x8c, 0x31,
	0x42, 0xc6, 0x44, 0x82, 0x0d, 0x59, 0xc6, 0x44, 0x49, 0x9f, 0x4b, 0x14, 0xc5, 0xb1, 0x41, 0xa9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x25, 0xa9, 0x89, 0xb9, 0x9e, 0x29,
	0x50, 0x37, 0x43, 0x79, 0x4a, 0x59, 0x5c, 0x52, 0xbe, 0xc5, 0xe9, 0x8e, 0x79, 0x79, 0xf9, 0xa5,
	0x79, 0xc9, 0xa9, 0x8e, 0x79, 0x29, 0x10, 0xbd, 0x7e, 0xa9, 0x15, 0x25, 0xbe, 0xe5, 0x78, 0xbc,
	0x2a, 0xc4, 0xc5, 0x92, 0x5b, 0xee, 0x99, 0x02, 0xf5, 0x27, 0x98, 0x2d, 0x24, 0xc7, 0xc5, 0x05,
	0x71, 0x47, 0x40, 0x6a, 0x51, 0x1a, 0xd4, 0x9f, 0x48, 0x22, 0x4a, 0x0e, 0x5c, 0x4a, 0xb8, 0xed,
	0x82, 0xbb, 0x54, 0x8a, 0x8b, 0x23, 0x0f, 0x2c, 0x02, 0x77, 0x2b, 0x9c, 0x6f, 0xf4, 0x91, 0x91,
	0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0x28, 0x81, 0x8b, 0x0b, 0x29, 0x42, 0x54, 0xf4, 0xb0, 0x46, 0x9d,
	0x1e, 0x4a, 0x48, 0x48, 0xe9, 0x10, 0xa3, 0x0a, 0xee, 0x8a, 0x76, 0x46, 0x2e, 0x71, 0x5c, 0xa1,
	0x62, 0x88, 0xdb, 0x24, 0x1c, 0x5a, 0xa4, 0x2c, 0x49, 0xd6, 0x02, 0x73, 0x89, 0x93, 0xf9, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xc9, 0x22, 0x25, 0xe3, 0x0a, 0x7d, 0xe4,
	0x34, 0x5d, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0xd7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0xb6, 0xcc, 0x55, 0xee, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateTeam(ctx context.Context, in *MsgCreateTeam, opts ...grpc.CallOption) (*MsgCreateTeamResponse, error)
	AnnounceAndCreateNextMw(ctx context.Context, in *MsgAnnounceAndCreateNextMw, opts ...grpc.CallOption) (*MsgAnnounceAndCreateNextMwResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateTeam(ctx context.Context, in *MsgCreateTeam, opts ...grpc.CallOption) (*MsgCreateTeamResponse, error) {
	out := new(MsgCreateTeamResponse)
	err := c.cc.Invoke(ctx, "/fantasfive.fantasfive.Msg/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AnnounceAndCreateNextMw(ctx context.Context, in *MsgAnnounceAndCreateNextMw, opts ...grpc.CallOption) (*MsgAnnounceAndCreateNextMwResponse, error) {
	out := new(MsgAnnounceAndCreateNextMwResponse)
	err := c.cc.Invoke(ctx, "/fantasfive.fantasfive.Msg/AnnounceAndCreateNextMw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateTeam(context.Context, *MsgCreateTeam) (*MsgCreateTeamResponse, error)
	AnnounceAndCreateNextMw(context.Context, *MsgAnnounceAndCreateNextMw) (*MsgAnnounceAndCreateNextMwResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateTeam(ctx context.Context, req *MsgCreateTeam) (*MsgCreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (*UnimplementedMsgServer) AnnounceAndCreateNextMw(ctx context.Context, req *MsgAnnounceAndCreateNextMw) (*MsgAnnounceAndCreateNextMwResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceAndCreateNextMw not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTeam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasfive.fantasfive.Msg/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTeam(ctx, req.(*MsgCreateTeam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AnnounceAndCreateNextMw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAnnounceAndCreateNextMw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AnnounceAndCreateNextMw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasfive.fantasfive.Msg/AnnounceAndCreateNextMw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AnnounceAndCreateNextMw(ctx, req.(*MsgAnnounceAndCreateNextMw))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fantasfive.fantasfive.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _Msg_CreateTeam_Handler,
		},
		{
			MethodName: "AnnounceAndCreateNextMw",
			Handler:    _Msg_AnnounceAndCreateNextMw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fantasfive/tx.proto",
}

func (m *MsgCreateTeam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateTeam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateTeam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Player4) > 0 {
		i -= len(m.Player4)
		copy(dAtA[i:], m.Player4)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Player4)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Player3) > 0 {
		i -= len(m.Player3)
		copy(dAtA[i:], m.Player3)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Player3)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Player2) > 0 {
		i -= len(m.Player2)
		copy(dAtA[i:], m.Player2)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Player2)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Player1) > 0 {
		i -= len(m.Player1)
		copy(dAtA[i:], m.Player1)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Player1)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Player0) > 0 {
		i -= len(m.Player0)
		copy(dAtA[i:], m.Player0)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Player0)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateTeamResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateTeamResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateTeamResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TeamId) > 0 {
		i -= len(m.TeamId)
		copy(dAtA[i:], m.TeamId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TeamId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAnnounceAndCreateNextMw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnnounceAndCreateNextMw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnnounceAndCreateNextMw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlayerPerf) > 0 {
		i -= len(m.PlayerPerf)
		copy(dAtA[i:], m.PlayerPerf)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlayerPerf)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MwId) > 0 {
		i -= len(m.MwId)
		copy(dAtA[i:], m.MwId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MwId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAnnounceAndCreateNextMwResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnnounceAndCreateNextMwResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnnounceAndCreateNextMwResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextMwId) > 0 {
		i -= len(m.NextMwId)
		copy(dAtA[i:], m.NextMwId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NextMwId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateTeam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Player0)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Player1)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Player2)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Player3)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Player4)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateTeamResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TeamId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAnnounceAndCreateNextMw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MwId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PlayerPerf)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAnnounceAndCreateNextMwResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NextMwId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateTeam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateTeam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateTeam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player3", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player3 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player4", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player4 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateTeamResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateTeamResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateTeamResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAnnounceAndCreateNextMw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAnnounceAndCreateNextMw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnnounceAndCreateNextMw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MwId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MwId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerPerf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerPerf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAnnounceAndCreateNextMwResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAnnounceAndCreateNextMwResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnnounceAndCreateNextMwResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextMwId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextMwId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
