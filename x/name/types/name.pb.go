// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/name/v1/name.proto

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

// Params defines the set of params for the name module.
type Params struct {
	// maximum length of name segment to allow
	MaxSegmentLength uint32 `protobuf:"varint,1,opt,name=max_segment_length,json=maxSegmentLength,proto3" json:"max_segment_length,omitempty"`
	// minimum length of name segment to allow
	MinSegmentLength uint32 `protobuf:"varint,2,opt,name=min_segment_length,json=minSegmentLength,proto3" json:"min_segment_length,omitempty"`
	// maximum number of name segments to allow.  Example: `foo.bar.baz` would be 3
	MaxNameLevels uint32 `protobuf:"varint,3,opt,name=max_name_levels,json=maxNameLevels,proto3" json:"max_name_levels,omitempty"`
	// determines if unrestricted name keys are allowed or not
	AllowUnrestrictedNames bool `protobuf:"varint,4,opt,name=allow_unrestricted_names,json=allowUnrestrictedNames,proto3" json:"allow_unrestricted_names,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a314256905bb00ec, []int{0}
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

func (m *Params) GetMaxSegmentLength() uint32 {
	if m != nil {
		return m.MaxSegmentLength
	}
	return 0
}

func (m *Params) GetMinSegmentLength() uint32 {
	if m != nil {
		return m.MinSegmentLength
	}
	return 0
}

func (m *Params) GetMaxNameLevels() uint32 {
	if m != nil {
		return m.MaxNameLevels
	}
	return 0
}

func (m *Params) GetAllowUnrestrictedNames() bool {
	if m != nil {
		return m.AllowUnrestrictedNames
	}
	return false
}

// NameRecord is a structure used to bind ownership of a name hierarchy to a collection of addresses
type NameRecord struct {
	// The bound name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The address the name resolved to.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Whether owner signature is required to add sub-names.
	Restricted bool `protobuf:"varint,3,opt,name=restricted,proto3" json:"restricted,omitempty"`
}

func (m *NameRecord) Reset()      { *m = NameRecord{} }
func (*NameRecord) ProtoMessage() {}
func (*NameRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a314256905bb00ec, []int{1}
}
func (m *NameRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRecord.Merge(m, src)
}
func (m *NameRecord) XXX_Size() int {
	return m.Size()
}
func (m *NameRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRecord.DiscardUnknown(m)
}

var xxx_messageInfo_NameRecord proto.InternalMessageInfo

func (m *NameRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NameRecord) GetRestricted() bool {
	if m != nil {
		return m.Restricted
	}
	return false
}

// CreateRootNameProposal details a proposal to create a new root name
// that is controlled by a given owner and optionally restricted to the owner
// for the sole creation of sub names.
type CreateRootNameProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Owner       string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Restricted  bool   `protobuf:"varint,5,opt,name=restricted,proto3" json:"restricted,omitempty"`
}

func (m *CreateRootNameProposal) Reset()      { *m = CreateRootNameProposal{} }
func (*CreateRootNameProposal) ProtoMessage() {}
func (*CreateRootNameProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a314256905bb00ec, []int{2}
}
func (m *CreateRootNameProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRootNameProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRootNameProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRootNameProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRootNameProposal.Merge(m, src)
}
func (m *CreateRootNameProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateRootNameProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRootNameProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRootNameProposal proto.InternalMessageInfo

// Event emitted when name is bound.
type EventNameBound struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Restricted bool   `protobuf:"varint,3,opt,name=restricted,proto3" json:"restricted,omitempty"`
}

func (m *EventNameBound) Reset()         { *m = EventNameBound{} }
func (m *EventNameBound) String() string { return proto.CompactTextString(m) }
func (*EventNameBound) ProtoMessage()    {}
func (*EventNameBound) Descriptor() ([]byte, []int) {
	return fileDescriptor_a314256905bb00ec, []int{3}
}
func (m *EventNameBound) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNameBound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNameBound.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNameBound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNameBound.Merge(m, src)
}
func (m *EventNameBound) XXX_Size() int {
	return m.Size()
}
func (m *EventNameBound) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNameBound.DiscardUnknown(m)
}

var xxx_messageInfo_EventNameBound proto.InternalMessageInfo

func (m *EventNameBound) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventNameBound) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventNameBound) GetRestricted() bool {
	if m != nil {
		return m.Restricted
	}
	return false
}

// Event emitted when name is unbound.
type EventNameUnbound struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Restricted bool   `protobuf:"varint,3,opt,name=restricted,proto3" json:"restricted,omitempty"`
}

func (m *EventNameUnbound) Reset()         { *m = EventNameUnbound{} }
func (m *EventNameUnbound) String() string { return proto.CompactTextString(m) }
func (*EventNameUnbound) ProtoMessage()    {}
func (*EventNameUnbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_a314256905bb00ec, []int{4}
}
func (m *EventNameUnbound) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNameUnbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNameUnbound.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNameUnbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNameUnbound.Merge(m, src)
}
func (m *EventNameUnbound) XXX_Size() int {
	return m.Size()
}
func (m *EventNameUnbound) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNameUnbound.DiscardUnknown(m)
}

var xxx_messageInfo_EventNameUnbound proto.InternalMessageInfo

func (m *EventNameUnbound) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventNameUnbound) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventNameUnbound) GetRestricted() bool {
	if m != nil {
		return m.Restricted
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "provenance.name.v1.Params")
	proto.RegisterType((*NameRecord)(nil), "provenance.name.v1.NameRecord")
	proto.RegisterType((*CreateRootNameProposal)(nil), "provenance.name.v1.CreateRootNameProposal")
	proto.RegisterType((*EventNameBound)(nil), "provenance.name.v1.EventNameBound")
	proto.RegisterType((*EventNameUnbound)(nil), "provenance.name.v1.EventNameUnbound")
}

func init() { proto.RegisterFile("provenance/name/v1/name.proto", fileDescriptor_a314256905bb00ec) }

var fileDescriptor_a314256905bb00ec = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xb1, 0x6e, 0x13, 0x4d,
	0x10, 0xc7, 0x6f, 0x13, 0x3b, 0x5f, 0x3c, 0x1f, 0x81, 0x68, 0x65, 0xa2, 0x13, 0x12, 0x17, 0xcb,
	0x05, 0x4a, 0x01, 0x3e, 0x22, 0x1a, 0x44, 0x19, 0x44, 0x17, 0x21, 0xeb, 0x50, 0x1a, 0x0a, 0x9c,
	0xf5, 0xdd, 0xe8, 0xb2, 0xd2, 0xdd, 0xee, 0x69, 0x77, 0x7d, 0x31, 0x6f, 0x40, 0x49, 0x49, 0x99,
	0x92, 0x27, 0x41, 0x94, 0x29, 0x29, 0x91, 0xdd, 0xf0, 0x18, 0x68, 0xe7, 0xe2, 0xf8, 0x62, 0x0a,
	0x1a, 0xaa, 0xdb, 0x99, 0xff, 0x7f, 0x66, 0x7e, 0x73, 0xda, 0x85, 0xc7, 0x95, 0xd1, 0x35, 0x2a,
	0xa1, 0x52, 0x8c, 0x95, 0x28, 0x31, 0xae, 0x8f, 0xe9, 0x3b, 0xaa, 0x8c, 0x76, 0x9a, 0xf3, 0xb5,
	0x3c, 0xa2, 0x74, 0x7d, 0xfc, 0xa8, 0x9f, 0xeb, 0x5c, 0x93, 0x1c, 0xfb, 0x53, 0xe3, 0x1c, 0x7e,
	0x63, 0xb0, 0x33, 0x16, 0x46, 0x94, 0x96, 0x3f, 0x05, 0x5e, 0x8a, 0xf9, 0xc4, 0x62, 0x5e, 0xa2,
	0x72, 0x93, 0x02, 0x55, 0xee, 0x2e, 0x42, 0x36, 0x60, 0x47, 0x7b, 0xc9, 0x7e, 0x29, 0xe6, 0xef,
	0x1a, 0xe1, 0x94, 0xf2, 0xe4, 0x96, 0x6a, 0xd3, 0xbd, 0x75, 0xe3, 0x96, 0xea, 0xae, 0xfb, 0x09,
	0x3c, 0xf0, 0xbd, 0x3d, 0xcb, 0xa4, 0xc0, 0x1a, 0x0b, 0x1b, 0x6e, 0x93, 0x75, 0xaf, 0x14, 0xf3,
	0xb7, 0xa2, 0xc4, 0x53, 0x4a, 0xf2, 0x97, 0x10, 0x8a, 0xa2, 0xd0, 0x97, 0x93, 0x99, 0x32, 0x68,
	0x9d, 0x91, 0xa9, 0xc3, 0x8c, 0xca, 0x6c, 0xd8, 0x19, 0xb0, 0xa3, 0xdd, 0xe4, 0x80, 0xf4, 0xb3,
	0x96, 0xec, 0xcb, 0xed, 0xf0, 0x1c, 0xc0, 0x1f, 0x12, 0x4c, 0xb5, 0xc9, 0x38, 0x87, 0x8e, 0x2f,
	0x22, 0xfa, 0x5e, 0x42, 0x67, 0x1e, 0xc2, 0x7f, 0x22, 0xcb, 0x0c, 0x5a, 0x4b, 0x98, 0xbd, 0x64,
	0x15, 0xf2, 0x08, 0x60, 0xdd, 0x8e, 0xc0, 0x76, 0x93, 0x56, 0xe6, 0x55, 0xe7, 0xcb, 0xd5, 0x61,
	0x30, 0xfc, 0xca, 0xe0, 0xe0, 0xb5, 0x41, 0xe1, 0x30, 0xd1, 0xda, 0xf9, 0x61, 0x63, 0xa3, 0x2b,
	0x6d, 0x45, 0xc1, 0xfb, 0xd0, 0x75, 0xd2, 0x15, 0xab, 0x79, 0x4d, 0xc0, 0x07, 0xf0, 0x7f, 0x86,
	0x36, 0x35, 0xb2, 0x72, 0x52, 0xab, 0x9b, 0xa1, 0xed, 0xd4, 0x2d, 0xe6, 0x76, 0x0b, 0xb3, 0x0f,
	0x5d, 0x7d, 0xa9, 0xd0, 0xd0, 0xbe, 0xbd, 0xa4, 0x09, 0x36, 0x10, 0xbb, 0x7f, 0x20, 0xde, 0xfb,
	0x74, 0x75, 0x18, 0x78, 0xcc, 0x5f, 0x1e, 0xf5, 0x03, 0xdc, 0x7f, 0x53, 0xa3, 0x22, 0xc8, 0x13,
	0x3d, 0x53, 0x59, 0x7b, 0x79, 0x76, 0x77, 0xf9, 0x15, 0xc3, 0x56, 0x8b, 0xe1, 0x2f, 0x3f, 0x64,
	0x78, 0x0e, 0xfb, 0xb7, 0xfd, 0xcf, 0xd4, 0xf4, 0xdf, 0x4f, 0x38, 0x49, 0xbf, 0x2f, 0x22, 0x76,
	0xbd, 0x88, 0xd8, 0xcf, 0x45, 0xc4, 0x3e, 0x2f, 0xa3, 0xe0, 0x7a, 0x19, 0x05, 0x3f, 0x96, 0x51,
	0x00, 0x0f, 0x25, 0xdd, 0xdd, 0x8d, 0xeb, 0x3d, 0x66, 0xef, 0x9f, 0xe7, 0xd2, 0x5d, 0xcc, 0xa6,
	0xa3, 0x54, 0x97, 0xf1, 0xda, 0xf0, 0x4c, 0xea, 0x56, 0x14, 0xcf, 0x9b, 0xe7, 0xe2, 0x3e, 0x56,
	0x68, 0xa7, 0x3b, 0xf4, 0x06, 0x5e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x38, 0xb5, 0xff,
	0x4e, 0x03, 0x00, 0x00,
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
	if m.AllowUnrestrictedNames {
		i--
		if m.AllowUnrestrictedNames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.MaxNameLevels != 0 {
		i = encodeVarintName(dAtA, i, uint64(m.MaxNameLevels))
		i--
		dAtA[i] = 0x18
	}
	if m.MinSegmentLength != 0 {
		i = encodeVarintName(dAtA, i, uint64(m.MinSegmentLength))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxSegmentLength != 0 {
		i = encodeVarintName(dAtA, i, uint64(m.MaxSegmentLength))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NameRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Restricted {
		i--
		if m.Restricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintName(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateRootNameProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRootNameProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRootNameProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Restricted {
		i--
		if m.Restricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintName(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintName(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintName(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventNameBound) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNameBound) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNameBound) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Restricted {
		i--
		if m.Restricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintName(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventNameUnbound) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNameUnbound) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNameUnbound) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Restricted {
		i--
		if m.Restricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintName(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintName(dAtA []byte, offset int, v uint64) int {
	offset -= sovName(v)
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
	if m.MaxSegmentLength != 0 {
		n += 1 + sovName(uint64(m.MaxSegmentLength))
	}
	if m.MinSegmentLength != 0 {
		n += 1 + sovName(uint64(m.MinSegmentLength))
	}
	if m.MaxNameLevels != 0 {
		n += 1 + sovName(uint64(m.MaxNameLevels))
	}
	if m.AllowUnrestrictedNames {
		n += 2
	}
	return n
}

func (m *NameRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if m.Restricted {
		n += 2
	}
	return n
}

func (m *CreateRootNameProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if m.Restricted {
		n += 2
	}
	return n
}

func (m *EventNameBound) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if m.Restricted {
		n += 2
	}
	return n
}

func (m *EventNameUnbound) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if m.Restricted {
		n += 2
	}
	return n
}

func sovName(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozName(x uint64) (n int) {
	return sovName(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
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
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSegmentLength", wireType)
			}
			m.MaxSegmentLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSegmentLength |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSegmentLength", wireType)
			}
			m.MinSegmentLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSegmentLength |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNameLevels", wireType)
			}
			m.MaxNameLevels = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNameLevels |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowUnrestrictedNames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
			m.AllowUnrestrictedNames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
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
func (m *NameRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
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
			return fmt.Errorf("proto: NameRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
			m.Restricted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
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
func (m *CreateRootNameProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
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
			return fmt.Errorf("proto: CreateRootNameProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRootNameProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
			m.Restricted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
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
func (m *EventNameBound) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
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
			return fmt.Errorf("proto: EventNameBound: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNameBound: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
			m.Restricted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
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
func (m *EventNameUnbound) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
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
			return fmt.Errorf("proto: EventNameUnbound: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNameUnbound: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
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
			m.Restricted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
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
func skipName(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowName
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
					return 0, ErrIntOverflowName
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
					return 0, ErrIntOverflowName
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
				return 0, ErrInvalidLengthName
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupName
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthName
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthName        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowName          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupName = fmt.Errorf("proto: unexpected end of group")
)
