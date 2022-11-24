// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beacon/beacon/whois.proto

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

type Whois struct {
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	Latest    string `protobuf:"bytes,2,opt,name=latest,proto3" json:"latest,omitempty"`
	Jitter    string `protobuf:"bytes,3,opt,name=jitter,proto3" json:"jitter,omitempty"`
}

func (m *Whois) Reset()         { *m = Whois{} }
func (m *Whois) String() string { return proto.CompactTextString(m) }
func (*Whois) ProtoMessage()    {}
func (*Whois) Descriptor() ([]byte, []int) {
	return fileDescriptor_22fad40849351577, []int{0}
}
func (m *Whois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Whois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Whois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Whois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Whois.Merge(m, src)
}
func (m *Whois) XXX_Size() int {
	return m.Size()
}
func (m *Whois) XXX_DiscardUnknown() {
	xxx_messageInfo_Whois.DiscardUnknown(m)
}

var xxx_messageInfo_Whois proto.InternalMessageInfo

func (m *Whois) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *Whois) GetLatest() string {
	if m != nil {
		return m.Latest
	}
	return ""
}

func (m *Whois) GetJitter() string {
	if m != nil {
		return m.Jitter
	}
	return ""
}

func init() {
	proto.RegisterType((*Whois)(nil), "beacon.beacon.Whois")
}

func init() { proto.RegisterFile("beacon/beacon/whois.proto", fileDescriptor_22fad40849351577) }

var fileDescriptor_22fad40849351577 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4a, 0x4d, 0x4c,
	0xce, 0xcf, 0xd3, 0x87, 0x52, 0xe5, 0x19, 0xf9, 0x99, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xbc, 0x10, 0x31, 0x3d, 0x08, 0xa5, 0x14, 0xca, 0xc5, 0x1a, 0x0e, 0x92, 0x15, 0x92, 0xe1,
	0xe2, 0x2c, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x42, 0x08, 0x08, 0x89, 0x71, 0xb1, 0xe5, 0x24, 0x96, 0xa4, 0x16, 0x97, 0x48, 0x30, 0x81,
	0xa5, 0xa0, 0x3c, 0x90, 0x78, 0x56, 0x66, 0x49, 0x49, 0x6a, 0x91, 0x04, 0x33, 0x44, 0x1c, 0xc2,
	0x73, 0xd2, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x51, 0xa8, 0x9b,
	0x2a, 0x60, 0x8e, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xce, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0x0c, 0x67, 0x2b, 0xba, 0x00, 0x00, 0x00,
}

func (m *Whois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Whois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Jitter) > 0 {
		i -= len(m.Jitter)
		copy(dAtA[i:], m.Jitter)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Jitter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Latest) > 0 {
		i -= len(m.Latest)
		copy(dAtA[i:], m.Latest)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Latest)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhois(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhois(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Whois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Latest)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Jitter)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	return n
}

func sovWhois(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhois(x uint64) (n int) {
	return sovWhois(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Whois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhois
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
			return fmt.Errorf("proto: Whois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Whois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Latest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jitter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jitter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhois(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhois
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
func skipWhois(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
				return 0, ErrInvalidLengthWhois
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhois
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhois
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhois        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhois          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhois = fmt.Errorf("proto: unexpected end of group")
)