// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/deviceinfo/deviceinfo.proto

package deviceinfo

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/golang/protobuf/proto"
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

type Type int32

const (
	Type_Unknown Type = 0
	Type_Raw     Type = 1
	Type_Json    Type = 2
)

var Type_name = map[int32]string{
	0: "Unknown",
	1: "Raw",
	2: "Json",
}

var Type_value = map[string]int32{
	"Unknown": 0,
	"Raw":     1,
	"Json":    2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_29ed7a3d6003a93e, []int{0}
}

type DeviceInfos struct {
	Infos                []*DeviceInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeviceInfos) Reset()         { *m = DeviceInfos{} }
func (m *DeviceInfos) String() string { return proto.CompactTextString(m) }
func (*DeviceInfos) ProtoMessage()    {}
func (*DeviceInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_29ed7a3d6003a93e, []int{0}
}
func (m *DeviceInfos) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceInfos.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfos.Merge(m, src)
}
func (m *DeviceInfos) XXX_Size() int {
	return m.Size()
}
func (m *DeviceInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfos.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfos proto.InternalMessageInfo

func (m *DeviceInfos) GetInfos() []*DeviceInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

type DeviceInfo struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Link                 string   `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Type                 Type     `protobuf:"varint,5,opt,name=type,proto3,enum=berty.pkg.deviceinfo.Type" json:"type,omitempty"`
	ErrMsg               string   `protobuf:"bytes,6,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Weight               int32    `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_29ed7a3d6003a93e, []int{1}
}
func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return m.Size()
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

func (m *DeviceInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeviceInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DeviceInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *DeviceInfo) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *DeviceInfo) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_Unknown
}

func (m *DeviceInfo) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *DeviceInfo) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterEnum("berty.pkg.deviceinfo.Type", Type_name, Type_value)
	proto.RegisterType((*DeviceInfos)(nil), "berty.pkg.deviceinfo.DeviceInfos")
	proto.RegisterType((*DeviceInfo)(nil), "berty.pkg.deviceinfo.DeviceInfo")
}

func init() { proto.RegisterFile("pkg/deviceinfo/deviceinfo.proto", fileDescriptor_29ed7a3d6003a93e) }

var fileDescriptor_29ed7a3d6003a93e = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xcd, 0xbf, 0x3a, 0x05, 0x09, 0x43, 0xd1, 0xa5, 0x87, 0x18, 0x7a, 0x90, 0xe0,
	0x21, 0x95, 0x0a, 0x3e, 0x80, 0xe8, 0x41, 0xc1, 0x4b, 0xd0, 0x8b, 0x17, 0x69, 0xe3, 0x74, 0x1b,
	0x52, 0x77, 0xc3, 0x66, 0x6d, 0xc9, 0x9b, 0xf8, 0x42, 0x82, 0x47, 0x1f, 0x41, 0xea, 0x8b, 0x48,
	0xb7, 0x62, 0x15, 0x7a, 0xfb, 0xbe, 0xdf, 0xcc, 0x37, 0x0c, 0x7c, 0x70, 0x54, 0x95, 0x62, 0xf8,
	0x44, 0x8b, 0x22, 0xa7, 0x42, 0x4e, 0xd5, 0x1f, 0x99, 0x56, 0x5a, 0x19, 0x85, 0xbd, 0x09, 0x69,
	0xd3, 0xa4, 0x55, 0x29, 0xd2, 0xed, 0x6c, 0x70, 0x05, 0xdd, 0x4b, 0xeb, 0xae, 0xe5, 0x54, 0xd5,
	0x78, 0x0e, 0xde, 0x1a, 0xd7, 0x9c, 0xc5, 0x4e, 0xd2, 0x1d, 0xc5, 0xe9, 0xae, 0x50, 0xba, 0x4d,
	0x64, 0x9b, 0xf5, 0xc1, 0x1b, 0x03, 0xd8, 0x52, 0x0c, 0xc1, 0x29, 0xa9, 0xe1, 0x2c, 0x66, 0xc9,
	0x5e, 0xb6, 0x96, 0xd8, 0x03, 0x6f, 0x31, 0x9e, 0xbf, 0x10, 0x6f, 0x5b, 0xb6, 0x31, 0xd8, 0x87,
	0x4e, 0x3e, 0x36, 0x24, 0x94, 0x6e, 0xb8, 0x63, 0x07, 0xbf, 0x1e, 0x11, 0xdc, 0x79, 0x21, 0x4b,
	0xee, 0x5a, 0x6e, 0x35, 0xa6, 0xe0, 0x9a, 0xa6, 0x22, 0xee, 0xc5, 0x2c, 0xd9, 0x1f, 0xf5, 0x77,
	0x7f, 0x77, 0xd7, 0x54, 0x94, 0xd9, 0x3d, 0x3c, 0x84, 0x80, 0xb4, 0x7e, 0x7c, 0xae, 0x05, 0xf7,
	0xed, 0x19, 0x9f, 0xb4, 0xbe, 0xad, 0x05, 0x1e, 0x80, 0xbf, 0xa4, 0x42, 0xcc, 0x0c, 0x0f, 0x62,
	0x96, 0x78, 0xd9, 0x8f, 0x3b, 0x39, 0x06, 0x77, 0x1d, 0xc7, 0x2e, 0x04, 0xf7, 0xb2, 0x94, 0x6a,
	0x29, 0xc3, 0x16, 0x06, 0xe0, 0x64, 0xe3, 0x65, 0xc8, 0xb0, 0x03, 0xee, 0x4d, 0xad, 0x64, 0xd8,
	0xbe, 0x38, 0x7d, 0x5f, 0x45, 0xec, 0x63, 0x15, 0xb1, 0xcf, 0x55, 0xc4, 0x5e, 0xbf, 0xa2, 0xd6,
	0x43, 0xb4, 0xf9, 0xc5, 0x50, 0x3e, 0x1b, 0xe6, 0x4a, 0xd3, 0xf0, 0x7f, 0x1f, 0x13, 0xdf, 0xb6,
	0x70, 0xf6, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x65, 0x47, 0x0d, 0x5f, 0xa8, 0x01, 0x00, 0x00,
}

func (m *DeviceInfos) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceInfos) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceInfos) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Infos) > 0 {
		for iNdEx := len(m.Infos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Infos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceinfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DeviceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Weight != 0 {
		i = encodeVarintDeviceinfo(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ErrMsg) > 0 {
		i -= len(m.ErrMsg)
		copy(dAtA[i:], m.ErrMsg)
		i = encodeVarintDeviceinfo(dAtA, i, uint64(len(m.ErrMsg)))
		i--
		dAtA[i] = 0x32
	}
	if m.Type != 0 {
		i = encodeVarintDeviceinfo(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Link) > 0 {
		i -= len(m.Link)
		copy(dAtA[i:], m.Link)
		i = encodeVarintDeviceinfo(dAtA, i, uint64(len(m.Link)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = encodeVarintDeviceinfo(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintDeviceinfo(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintDeviceinfo(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeviceinfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceinfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceInfos) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Infos) > 0 {
		for _, e := range m.Infos {
			l = e.Size()
			n += 1 + l + sovDeviceinfo(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeviceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovDeviceinfo(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovDeviceinfo(uint64(l))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovDeviceinfo(uint64(l))
	}
	l = len(m.Link)
	if l > 0 {
		n += 1 + l + sovDeviceinfo(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovDeviceinfo(uint64(m.Type))
	}
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovDeviceinfo(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovDeviceinfo(uint64(m.Weight))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDeviceinfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceinfo(x uint64) (n int) {
	return sovDeviceinfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceInfos) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceinfo
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
			return fmt.Errorf("proto: DeviceInfos: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceInfos: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Infos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
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
				return ErrInvalidLengthDeviceinfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Infos = append(m.Infos, &DeviceInfo{})
			if err := m.Infos[len(m.Infos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceinfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeviceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceinfo
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
			return fmt.Errorf("proto: DeviceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
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
				return ErrInvalidLengthDeviceinfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
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
				return ErrInvalidLengthDeviceinfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
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
				return ErrInvalidLengthDeviceinfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
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
				return ErrInvalidLengthDeviceinfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Link = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
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
				return ErrInvalidLengthDeviceinfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceinfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceinfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceinfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDeviceinfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceinfo
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
					return 0, ErrIntOverflowDeviceinfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceinfo
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
				return 0, ErrInvalidLengthDeviceinfo
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDeviceinfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceinfo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDeviceinfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDeviceinfo
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDeviceinfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceinfo   = fmt.Errorf("proto: integer overflow")
)