// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/marker/v1/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// MarkerTransferAuthorization gives the grantee permissions to execute
// a marker transfer on behalf of the granter's account.
type MarkerTransferAuthorization struct {
	// transfer_limit is the total amount the grantee can transfer
	TransferLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=transfer_limit,json=transferLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"transfer_limit"`
}

func (m *MarkerTransferAuthorization) Reset()         { *m = MarkerTransferAuthorization{} }
func (m *MarkerTransferAuthorization) String() string { return proto.CompactTextString(m) }
func (*MarkerTransferAuthorization) ProtoMessage()    {}
func (*MarkerTransferAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86b03f937f368fb, []int{0}
}
func (m *MarkerTransferAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarkerTransferAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarkerTransferAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarkerTransferAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkerTransferAuthorization.Merge(m, src)
}
func (m *MarkerTransferAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MarkerTransferAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkerTransferAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MarkerTransferAuthorization proto.InternalMessageInfo

func (m *MarkerTransferAuthorization) GetTransferLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TransferLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*MarkerTransferAuthorization)(nil), "provenance.marker.v1.MarkerTransferAuthorization")
}

func init() { proto.RegisterFile("provenance/marker/v1/authz.proto", fileDescriptor_e86b03f937f368fb) }

var fileDescriptor_e86b03f937f368fb = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0x4b, 0xcd, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd2, 0x2f, 0x33,
	0xd4, 0x4f, 0x2c, 0x2d, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41, 0xa8,
	0xd0, 0x83, 0xa8, 0xd0, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd0, 0x07,
	0xb1, 0x20, 0x6a, 0xa5, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21, 0x12, 0x10, 0x0e,
	0x54, 0x4a, 0x0e, 0xc2, 0xd3, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x4f, 0xce, 0xcf, 0xcc, 0x83, 0xc8, 0x2b, 0x2d, 0x61, 0xe4, 0x92, 0xf6, 0x05, 0x1b,
	0x1f, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96, 0x5a, 0xe4, 0x58, 0x5a, 0x92, 0x91, 0x5f, 0x94, 0x59,
	0x95, 0x58, 0x92, 0x99, 0x9f, 0x27, 0x54, 0xc4, 0xc5, 0x57, 0x02, 0x95, 0x88, 0xcf, 0xc9, 0xcc,
	0xcd, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd4, 0x83, 0x5a, 0x03, 0x32, 0x58,
	0x0f, 0x6a, 0xb0, 0x9e, 0x73, 0x7e, 0x66, 0x9e, 0x93, 0xc1, 0x89, 0x7b, 0xf2, 0x0c, 0xab, 0xee,
	0xcb, 0x6b, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x42, 0xdd, 0x04, 0xa5,
	0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xc1, 0x1a, 0x8a, 0x83, 0x78, 0x61,
	0x56, 0xf8, 0x80, 0x6c, 0xb0, 0x12, 0x3c, 0xb5, 0x45, 0x97, 0x17, 0xc5, 0x19, 0x4e, 0xe9, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0, 0x25, 0x9e, 0x09, 0x0e, 0x05, 0x8c, 0xa0,
	0x0a, 0x60, 0x8c, 0x32, 0x42, 0x72, 0x00, 0x42, 0x89, 0x6e, 0x66, 0x3e, 0x12, 0x4f, 0xbf, 0x02,
	0x16, 0xfe, 0x60, 0x07, 0x25, 0xb1, 0x81, 0x83, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x3f, 0x79, 0x12, 0xa1, 0x01, 0x00, 0x00,
}

func (m *MarkerTransferAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarkerTransferAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarkerTransferAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TransferLimit) > 0 {
		for iNdEx := len(m.TransferLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TransferLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarkerTransferAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TransferLimit) > 0 {
		for _, e := range m.TransferLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarkerTransferAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: MarkerTransferAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarkerTransferAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferLimit = append(m.TransferLimit, types.Coin{})
			if err := m.TransferLimit[len(m.TransferLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
