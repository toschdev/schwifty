// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schwifty/schwifty/collection.proto

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

type Collection struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner       string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Ticker      string `protobuf:"bytes,5,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Uri         string `protobuf:"bytes,6,opt,name=uri,proto3" json:"uri,omitempty"`
	UriHash     string `protobuf:"bytes,7,opt,name=uriHash,proto3" json:"uriHash,omitempty"`
	Data        string `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Creator     string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_243d85adbd83889b, []int{0}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Collection) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Collection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Collection) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Collection) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Collection) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Collection) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

func (m *Collection) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Collection) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Collection)(nil), "toschdev.schwifty.schwifty.Collection")
}

func init() {
	proto.RegisterFile("schwifty/schwifty/collection.proto", fileDescriptor_243d85adbd83889b)
}

var fileDescriptor_243d85adbd83889b = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0xeb, 0x34, 0x4d, 0xe9, 0x8f, 0x84, 0x90, 0x85, 0xd0, 0x2f, 0x06, 0x2b, 0xea, 0xd4,
	0x85, 0x74, 0xe0, 0x0d, 0x60, 0x80, 0xb9, 0x23, 0x5b, 0x6a, 0x1b, 0x62, 0xd1, 0xc6, 0x91, 0xed,
	0x50, 0xfa, 0x16, 0x3c, 0x16, 0x63, 0x47, 0xc4, 0x84, 0x92, 0x17, 0x41, 0x76, 0x9a, 0x94, 0xed,
	0xbb, 0xfb, 0xcf, 0x3e, 0xe9, 0x60, 0x6e, 0x79, 0xb1, 0x53, 0x2f, 0x6e, 0xbf, 0x1c, 0x80, 0xeb,
	0xcd, 0x46, 0x72, 0xa7, 0x74, 0x99, 0x55, 0x46, 0x3b, 0x4d, 0x6f, 0x9c, 0xb6, 0xbc, 0x10, 0xf2,
	0x3d, 0xeb, 0x33, 0x03, 0xcc, 0x7f, 0x08, 0xc0, 0xc3, 0xf0, 0x80, 0x5e, 0x40, 0xa4, 0x04, 0x92,
	0x94, 0x2c, 0xe2, 0x55, 0xa4, 0x04, 0xbd, 0x82, 0x89, 0xde, 0x95, 0xd2, 0x60, 0x94, 0x92, 0xc5,
	0x6c, 0xd5, 0x09, 0x4a, 0x21, 0x2e, 0xf3, 0xad, 0xc4, 0x71, 0x30, 0x03, 0xd3, 0x14, 0xce, 0x85,
	0xb4, 0xdc, 0xa8, 0xca, 0x7f, 0x84, 0x71, 0x38, 0xfd, 0xb7, 0xe8, 0x35, 0x24, 0x4e, 0xf1, 0x37,
	0x69, 0x70, 0x12, 0x8e, 0x47, 0x45, 0x2f, 0x61, 0x5c, 0x1b, 0x85, 0x49, 0x30, 0x3d, 0x52, 0x84,
	0x69, 0x6d, 0xd4, 0x53, 0x6e, 0x0b, 0x9c, 0x06, 0xb7, 0x97, 0xbe, 0x59, 0xe4, 0x2e, 0xc7, 0xb3,
	0xae, 0xd9, 0xb3, 0x4f, 0x73, 0x23, 0x73, 0xa7, 0x0d, 0xce, 0xba, 0xf4, 0x51, 0xde, 0x3f, 0x7e,
	0x35, 0x8c, 0x1c, 0x1a, 0x46, 0x7e, 0x1b, 0x46, 0x3e, 0x5b, 0x36, 0x3a, 0xb4, 0x6c, 0xf4, 0xdd,
	0xb2, 0xd1, 0xf3, 0xed, 0xab, 0x72, 0x45, 0xbd, 0xce, 0xb8, 0xde, 0x2e, 0xfb, 0x75, 0x4e, 0x0b,
	0x7e, 0x9c, 0xd0, 0xed, 0x2b, 0x69, 0xd7, 0x49, 0x18, 0xf2, 0xee, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xe2, 0x61, 0x84, 0x12, 0x6e, 0x01, 0x00, 0x00,
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCollection(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	return n
}

func sovCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollection(x uint64) (n int) {
	return sovCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
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
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollection
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
func skipCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
				return 0, ErrInvalidLengthCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollection = fmt.Errorf("proto: unexpected end of group")
)
