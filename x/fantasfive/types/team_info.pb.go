// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fantasfive/team_info.proto

package types

import (
	fmt "fmt"
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

type TeamInfo struct {
	NextId uint64 `protobuf:"varint,1,opt,name=nextId,proto3" json:"nextId,omitempty"`
}

func (m *TeamInfo) Reset()         { *m = TeamInfo{} }
func (m *TeamInfo) String() string { return proto.CompactTextString(m) }
func (*TeamInfo) ProtoMessage()    {}
func (*TeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1590dee628316f41, []int{0}
}
func (m *TeamInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TeamInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamInfo.Merge(m, src)
}
func (m *TeamInfo) XXX_Size() int {
	return m.Size()
}
func (m *TeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeamInfo proto.InternalMessageInfo

func (m *TeamInfo) GetNextId() uint64 {
	if m != nil {
		return m.NextId
	}
	return 0
}

func init() {
	proto.RegisterType((*TeamInfo)(nil), "fantasfive.fantasfive.TeamInfo")
}

func init() { proto.RegisterFile("fantasfive/team_info.proto", fileDescriptor_1590dee628316f41) }

var fileDescriptor_1590dee628316f41 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x4b, 0xcc, 0x2b,
	0x49, 0x2c, 0x4e, 0xcb, 0x2c, 0x4b, 0xd5, 0x2f, 0x49, 0x4d, 0xcc, 0x8d, 0xcf, 0xcc, 0x4b, 0xcb,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x45, 0xc8, 0xe9, 0x21, 0x98, 0x4a, 0x4a, 0x5c,
	0x1c, 0x21, 0xa9, 0x89, 0xb9, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x62, 0x5c, 0x6c, 0x79, 0xa9, 0x15,
	0x25, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x93, 0xf9, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xc9, 0x22, 0x59, 0x58, 0xa1, 0x8f, 0x6c, 0x7b,
	0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68,
	0xb0, 0x84, 0x28, 0x98, 0x00, 0x00, 0x00,
}

func (m *TeamInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TeamInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TeamInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextId != 0 {
		i = encodeVarintTeamInfo(dAtA, i, uint64(m.NextId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTeamInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovTeamInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TeamInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextId != 0 {
		n += 1 + sovTeamInfo(uint64(m.NextId))
	}
	return n
}

func sovTeamInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTeamInfo(x uint64) (n int) {
	return sovTeamInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TeamInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTeamInfo
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
			return fmt.Errorf("proto: TeamInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TeamInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextId", wireType)
			}
			m.NextId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeamInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTeamInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTeamInfo
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
func skipTeamInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTeamInfo
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
					return 0, ErrIntOverflowTeamInfo
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
					return 0, ErrIntOverflowTeamInfo
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
				return 0, ErrInvalidLengthTeamInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTeamInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTeamInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTeamInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTeamInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTeamInfo = fmt.Errorf("proto: unexpected end of group")
)
