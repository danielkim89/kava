// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/greet/v1beta1/greet.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Greet struct {
	Owner   string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Greet) Reset()         { *m = Greet{} }
func (m *Greet) String() string { return proto.CompactTextString(m) }
func (*Greet) ProtoMessage()    {}
func (*Greet) Descriptor() ([]byte, []int) {
	return fileDescriptor_729b1289029a1b8a, []int{0}
}
func (m *Greet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Greet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Greet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Greet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greet.Merge(m, src)
}
func (m *Greet) XXX_Size() int {
	return m.Size()
}
func (m *Greet) XXX_DiscardUnknown() {
	xxx_messageInfo_Greet.DiscardUnknown(m)
}

var xxx_messageInfo_Greet proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Greet)(nil), "kava.greet.v1beta1.Greet")
}

func init() { proto.RegisterFile("kava/greet/v1beta1/greet.proto", fileDescriptor_729b1289029a1b8a) }

var fileDescriptor_729b1289029a1b8a = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0xa0, 0x82, 0xf0, 0xc0, 0x60, 0x75, 0x30, 0x1d, 0x2e, 0x88, 0x05, 0x16, 0x62,
	0x55, 0xbc, 0x01, 0x12, 0xea, 0xce, 0xc8, 0x82, 0xec, 0xf6, 0x62, 0x2c, 0x6a, 0x6e, 0x14, 0xbb,
	0x05, 0xde, 0x82, 0xc7, 0xea, 0xd8, 0x91, 0x11, 0x92, 0x17, 0x41, 0xb6, 0xd3, 0xed, 0x7c, 0xdf,
	0x39, 0x92, 0x7f, 0x38, 0xbc, 0xe9, 0xad, 0x56, 0xb6, 0x43, 0x8c, 0x6a, 0x3b, 0x37, 0x18, 0xf5,
	0xbc, 0x50, 0xd3, 0x76, 0x14, 0x49, 0x88, 0xd4, 0x37, 0xc5, 0x8c, 0xfd, 0xec, 0x7c, 0x49, 0xc1,
	0x53, 0x78, 0xce, 0x0b, 0x55, 0xa0, 0xcc, 0x67, 0x53, 0x4b, 0x96, 0x8a, 0x4f, 0x69, 0xb4, 0x17,
	0x96, 0xc8, 0xae, 0x51, 0x65, 0x32, 0x9b, 0x17, 0x15, 0x9d, 0xc7, 0x10, 0xb5, 0x6f, 0xcb, 0xe0,
	0x6a, 0xc1, 0x27, 0x8b, 0x74, 0x84, 0x98, 0xf2, 0x09, 0x7d, 0xbc, 0x63, 0x27, 0xd9, 0x25, 0xbb,
	0x39, 0x7d, 0x2c, 0x20, 0xce, 0x78, 0xed, 0x56, 0xb2, 0xce, 0xaa, 0x76, 0x2b, 0x21, 0xf9, 0x89,
	0xc7, 0x10, 0xb4, 0x45, 0x79, 0x94, 0xe5, 0x01, 0xef, 0x1f, 0x76, 0x7f, 0x50, 0xed, 0x7a, 0x60,
	0xfb, 0x1e, 0xd8, 0x6f, 0x0f, 0xec, 0x7b, 0x80, 0x6a, 0x3f, 0x40, 0xf5, 0x33, 0x40, 0xf5, 0x74,
	0x6d, 0x5d, 0x7c, 0xdd, 0x98, 0x66, 0x49, 0x5e, 0xa5, 0x77, 0xdd, 0xae, 0xb5, 0x09, 0x39, 0xa9,
	0xcf, 0xf1, 0x0f, 0xe2, 0x57, 0x8b, 0xc1, 0x1c, 0xe7, 0x6b, 0xdd, 0xfd, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x0a, 0xac, 0x36, 0x1e, 0x01, 0x00, 0x00,
}

func (m *Greet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Greet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Greet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintGreet(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGreet(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGreet(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGreet(dAtA []byte, offset int, v uint64) int {
	offset -= sovGreet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Greet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGreet(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGreet(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovGreet(uint64(l))
	}
	return n
}

func sovGreet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGreet(x uint64) (n int) {
	return sovGreet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Greet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGreet
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
			return fmt.Errorf("proto: Greet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Greet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGreet
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
				return ErrInvalidLengthGreet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGreet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGreet
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
				return ErrInvalidLengthGreet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGreet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGreet
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
				return ErrInvalidLengthGreet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGreet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGreet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGreet
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
func skipGreet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGreet
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
					return 0, ErrIntOverflowGreet
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
					return 0, ErrIntOverflowGreet
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
				return 0, ErrInvalidLengthGreet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGreet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGreet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGreet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGreet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGreet = fmt.Errorf("proto: unexpected end of group")
)