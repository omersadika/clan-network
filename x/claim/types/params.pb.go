// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: clan/claim/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the claim module's parameters.
type Params struct {
	AirdropEnabled     bool          `protobuf:"varint,1,opt,name=airdrop_enabled,json=airdropEnabled,proto3" json:"airdrop_enabled,omitempty"`
	AirdropStartTime   time.Time     `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	DurationUntilDecay time.Duration `protobuf:"bytes,3,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay,omitempty" yaml:"duration_until_decay"`
	DurationOfDecay    time.Duration `protobuf:"bytes,4,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay,omitempty" yaml:"duration_of_decay"`
	// denom of claimable asset
	ClaimDenom string `protobuf:"bytes,5,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae47f011891d1d2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAirdropEnabled() bool {
	if m != nil {
		return m.AirdropEnabled
	}
	return false
}

func (m *Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *Params) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "ClanNetwork.clannetwork.claim.v1beta1.Params")
}

func init() { proto.RegisterFile("clan/claim/v1beta1/params.proto", fileDescriptor_6ae47f011891d1d2) }

var fileDescriptor_6ae47f011891d1d2 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x6f, 0xd4, 0x30,
	0x00, 0x8d, 0x69, 0xa9, 0xc0, 0x95, 0x28, 0x44, 0x1d, 0xd2, 0xab, 0x64, 0x9f, 0x22, 0x55, 0xdc,
	0x00, 0xb1, 0x0e, 0xb6, 0x8e, 0xd7, 0x63, 0x40, 0x48, 0x80, 0x0e, 0x58, 0x58, 0x22, 0x27, 0xf1,
	0x05, 0x8b, 0xd8, 0x8e, 0x12, 0x07, 0xc8, 0x5f, 0x60, 0xea, 0x84, 0x3a, 0xf2, 0x73, 0x3a, 0x76,
	0x64, 0x0a, 0xe8, 0x6e, 0x63, 0x41, 0xba, 0x5f, 0x80, 0x9c, 0x38, 0x7c, 0xdc, 0x9d, 0xd4, 0xcd,
	0x7e, 0x1f, 0x7e, 0xcf, 0xd2, 0x83, 0x38, 0xce, 0xa8, 0x24, 0x71, 0x46, 0xb9, 0x20, 0x1f, 0xc6,
	0x11, 0xd3, 0x74, 0x4c, 0x72, 0x5a, 0x50, 0x51, 0x06, 0x79, 0xa1, 0xb4, 0x72, 0x4f, 0xce, 0x32,
	0x2a, 0x9f, 0x33, 0xfd, 0x51, 0x15, 0xef, 0x03, 0x23, 0x96, 0x7f, 0xcf, 0x5c, 0x04, 0xd6, 0x33,
	0x38, 0x4c, 0x55, 0xaa, 0x5a, 0x07, 0x31, 0xa7, 0xce, 0x3c, 0x40, 0xa9, 0x52, 0x69, 0xc6, 0x48,
	0x7b, 0x8b, 0xaa, 0x39, 0x49, 0xaa, 0x82, 0x6a, 0xae, 0xa4, 0xe5, 0xf1, 0x3a, 0xaf, 0xb9, 0x60,
	0xa5, 0xa6, 0x22, 0xef, 0x04, 0xfe, 0xaf, 0x1d, 0xb8, 0xf7, 0xb2, 0xad, 0xe3, 0xde, 0x87, 0x07,
	0x94, 0x17, 0x49, 0xa1, 0xf2, 0x90, 0x49, 0x1a, 0x65, 0x2c, 0xf1, 0xc0, 0x10, 0x8c, 0x6e, 0xcd,
	0xee, 0x58, 0xf8, 0x49, 0x87, 0xba, 0x0a, 0xba, 0xbd, 0xb0, 0xd4, 0xb4, 0xd0, 0xa1, 0x79, 0xd4,
	0xbb, 0x31, 0x04, 0xa3, 0xfd, 0x47, 0x83, 0xa0, 0x4b, 0x0c, 0xfa, 0xc4, 0xe0, 0x75, 0x9f, 0x38,
	0x39, 0xb9, 0x6c, 0xb0, 0xb3, 0x6a, 0xf0, 0x51, 0x4d, 0x45, 0x76, 0xea, 0x6f, 0xbe, 0xe1, 0x9f,
	0x7f, 0xc7, 0x60, 0x76, 0xd7, 0x12, 0xaf, 0x0c, 0x6e, 0xdc, 0xee, 0x17, 0x00, 0x0f, 0xfb, 0x8f,
	0x85, 0x95, 0xd4, 0x3c, 0x0b, 0x13, 0x16, 0xd3, 0xda, 0xdb, 0x69, 0x33, 0x8f, 0x36, 0x32, 0xa7,
	0x56, 0x3c, 0x79, 0x6a, 0x22, 0x7f, 0x36, 0x18, 0x6d, 0xb3, 0x3f, 0x50, 0x82, 0x6b, 0x26, 0x72,
	0x5d, 0xaf, 0x1a, 0x7c, 0xdc, 0x95, 0xda, 0xa6, 0xf3, 0x2f, 0x4c, 0x2d, 0xb7, 0xa7, 0xde, 0x18,
	0x66, 0x6a, 0x08, 0xf7, 0x33, 0x80, 0xf7, 0xfe, 0x38, 0xd4, 0xdc, 0xb6, 0xda, 0xbd, 0xae, 0xd5,
	0x99, 0x6d, 0x75, 0xbc, 0xe1, 0xfd, 0xaf, 0x92, 0xb7, 0x56, 0xa9, 0x17, 0x75, 0x7d, 0x0e, 0x7a,
	0xfc, 0xc5, 0xbc, 0x2b, 0x83, 0xe1, 0x7e, 0x3b, 0x99, 0x30, 0x61, 0x52, 0x09, 0xef, 0xe6, 0x10,
	0x8c, 0x6e, 0xcf, 0x60, 0x0b, 0x4d, 0x0d, 0x72, 0xba, 0x7b, 0xf1, 0x15, 0x3b, 0x93, 0x67, 0x97,
	0x0b, 0x04, 0xae, 0x16, 0x08, 0xfc, 0x58, 0x20, 0x70, 0xbe, 0x44, 0xce, 0xd5, 0x12, 0x39, 0xdf,
	0x96, 0xc8, 0x79, 0x3b, 0x4e, 0xb9, 0x7e, 0x57, 0x45, 0x41, 0xac, 0x04, 0xf9, 0x67, 0x94, 0x66,
	0xbc, 0xf2, 0xa1, 0x5d, 0x25, 0xf9, 0x64, 0xb7, 0xac, 0xeb, 0x9c, 0x95, 0xd1, 0x5e, 0xfb, 0xb9,
	0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x00, 0xfa, 0xfb, 0xe4, 0xe6, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x2a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if m.AirdropEnabled {
		i--
		if m.AirdropEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropEnabled {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AirdropEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
