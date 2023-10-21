// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fantasfive/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the fantasfive module's genesis state.
type GenesisState struct {
	Params         Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SystemInfoList []SystemInfo `protobuf:"bytes,2,rep,name=systemInfoList,proto3" json:"systemInfoList"`
	StoredMWList   []StoredMW   `protobuf:"bytes,3,rep,name=storedMWList,proto3" json:"storedMWList"`
	StoredTeamList []StoredTeam `protobuf:"bytes,4,rep,name=storedTeamList,proto3" json:"storedTeamList"`
	MwInfo         MwInfo       `protobuf:"bytes,5,opt,name=mwInfo,proto3" json:"mwInfo"`
	TeamInfo       TeamInfo     `protobuf:"bytes,6,opt,name=teamInfo,proto3" json:"teamInfo"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5158d8bf012844a9, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSystemInfoList() []SystemInfo {
	if m != nil {
		return m.SystemInfoList
	}
	return nil
}

func (m *GenesisState) GetStoredMWList() []StoredMW {
	if m != nil {
		return m.StoredMWList
	}
	return nil
}

func (m *GenesisState) GetStoredTeamList() []StoredTeam {
	if m != nil {
		return m.StoredTeamList
	}
	return nil
}

func (m *GenesisState) GetMwInfo() MwInfo {
	if m != nil {
		return m.MwInfo
	}
	return MwInfo{}
}

func (m *GenesisState) GetTeamInfo() TeamInfo {
	if m != nil {
		return m.TeamInfo
	}
	return TeamInfo{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fantasfive.fantasfive.GenesisState")
}

func init() { proto.RegisterFile("fantasfive/genesis.proto", fileDescriptor_5158d8bf012844a9) }

var fileDescriptor_5158d8bf012844a9 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4b, 0xcc, 0x2b,
	0x49, 0x2c, 0x4e, 0xcb, 0x2c, 0x4b, 0xd5, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x45, 0xc8, 0xe8, 0x21, 0x98, 0x52, 0x22, 0xe9, 0xf9, 0xe9,
	0xf9, 0x60, 0x15, 0xfa, 0x20, 0x16, 0x44, 0xb1, 0x94, 0x38, 0x92, 0x31, 0x05, 0x89, 0x45, 0x89,
	0xb9, 0x50, 0x53, 0xa4, 0x64, 0x90, 0x24, 0x8a, 0x2b, 0x8b, 0x4b, 0x52, 0x73, 0xe3, 0x33, 0xf3,
	0xd2, 0x60, 0xda, 0xa4, 0x90, 0x65, 0x4b, 0xf2, 0x8b, 0x52, 0x53, 0xe2, 0x73, 0xcb, 0xb1, 0xe9,
	0x84, 0xc8, 0x95, 0xa4, 0x26, 0xe6, 0x42, 0x65, 0x91, 0xdd, 0x9d, 0x5b, 0x8e, 0xcb, 0x4c, 0x90,
	0x06, 0x24, 0x39, 0xa5, 0x1d, 0xcc, 0x5c, 0x3c, 0xee, 0x10, 0x5f, 0x06, 0x97, 0x24, 0x96, 0xa4,
	0x0a, 0x59, 0x73, 0xb1, 0x41, 0x9c, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xab, 0x87,
	0xd5, 0xd7, 0x7a, 0x01, 0x60, 0x45, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0x41, 0xb5, 0x08,
	0xf9, 0x73, 0xf1, 0x41, 0xbc, 0xe4, 0x99, 0x97, 0x96, 0xef, 0x93, 0x59, 0x5c, 0x22, 0xc1, 0xa4,
	0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0x88, 0xc3, 0x90, 0x60, 0xb8, 0x62, 0xa8, 0x41, 0x68, 0xda, 0x85,
	0x3c, 0xb9, 0x78, 0x20, 0x3e, 0xf5, 0x0d, 0x07, 0x1b, 0xc7, 0x0c, 0x36, 0x4e, 0x1e, 0x97, 0x71,
	0x50, 0xa5, 0x50, 0xc3, 0x50, 0xb4, 0x82, 0xdd, 0x06, 0xe6, 0x87, 0xa4, 0x26, 0xe6, 0x82, 0x0d,
	0x63, 0xc1, 0xef, 0x36, 0xb8, 0x62, 0xb8, 0xdb, 0x50, 0xb4, 0x83, 0x42, 0x2a, 0xb7, 0x1c, 0xe4,
	0x52, 0x09, 0x56, 0xbc, 0x21, 0xe5, 0x5b, 0x8e, 0xe4, 0x41, 0xa8, 0x16, 0x21, 0x47, 0x2e, 0x0e,
	0x50, 0x54, 0x80, 0xb5, 0xb3, 0x81, 0xb5, 0xe3, 0xf2, 0x54, 0x08, 0x54, 0x19, 0xd4, 0x00, 0xb8,
	0x36, 0x27, 0xf3, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x45, 0x8a,
	0xf0, 0x0a, 0x7d, 0xe4, 0xd8, 0xaf, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x47, 0xbd, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x78, 0x95, 0x60, 0xb1, 0xea, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TeamInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.MwInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.StoredTeamList) > 0 {
		for iNdEx := len(m.StoredTeamList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoredTeamList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StoredMWList) > 0 {
		for iNdEx := len(m.StoredMWList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoredMWList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SystemInfoList) > 0 {
		for iNdEx := len(m.SystemInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SystemInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SystemInfoList) > 0 {
		for _, e := range m.SystemInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StoredMWList) > 0 {
		for _, e := range m.StoredMWList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StoredTeamList) > 0 {
		for _, e := range m.StoredTeamList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.MwInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TeamInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemInfoList = append(m.SystemInfoList, SystemInfo{})
			if err := m.SystemInfoList[len(m.SystemInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredMWList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoredMWList = append(m.StoredMWList, StoredMW{})
			if err := m.StoredMWList[len(m.StoredMWList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredTeamList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoredTeamList = append(m.StoredTeamList, StoredTeam{})
			if err := m.StoredTeamList[len(m.StoredTeamList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MwInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MwInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TeamInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
