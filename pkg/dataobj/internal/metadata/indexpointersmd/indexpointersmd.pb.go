// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/dataobj/internal/metadata/indexpointersmd/indexpointersmd.proto

package indexpointersmd

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	datasetmd "github.com/grafana/loki/v3/pkg/dataobj/internal/metadata/datasetmd"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

// ColumnType represents the valid types that a indexpointer's column can have.
type ColumnType int32

const (
	// Invalid column type.
	COLUMN_TYPE_UNSPECIFIED ColumnType = 0
	// COLUMN_TYPE_PATH is a column containing the data object path in object storage.
	COLUMN_TYPE_PATH ColumnType = 1
	// COLUMN_TYPE_MIN_TIMESTAMP is a column containing the minimum timestamp of
	// a data object.
	COLUMN_TYPE_MIN_TIMESTAMP ColumnType = 2
	// COLUMN_TYPE_MAX_TIMESTAMP is a column containing the maximum timestamp of
	// a data object.
	COLUMN_TYPE_MAX_TIMESTAMP ColumnType = 3
)

var ColumnType_name = map[int32]string{
	0: "COLUMN_TYPE_UNSPECIFIED",
	1: "COLUMN_TYPE_PATH",
	2: "COLUMN_TYPE_MIN_TIMESTAMP",
	3: "COLUMN_TYPE_MAX_TIMESTAMP",
}

var ColumnType_value = map[string]int32{
	"COLUMN_TYPE_UNSPECIFIED":   0,
	"COLUMN_TYPE_PATH":          1,
	"COLUMN_TYPE_MIN_TIMESTAMP": 2,
	"COLUMN_TYPE_MAX_TIMESTAMP": 3,
}

func (ColumnType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e25fb77c0f3d151, []int{0}
}

// Metadata describes the metadata for the indexpointers section.
type Metadata struct {
	// Columns within the indexpointers section.
	Columns []*ColumnDesc `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (m *Metadata) Reset()      { *m = Metadata{} }
func (*Metadata) ProtoMessage() {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e25fb77c0f3d151, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetColumns() []*ColumnDesc {
	if m != nil {
		return m.Columns
	}
	return nil
}

// ColumnDesc describes an individual column within the indexpointers table.
type ColumnDesc struct {
	// Information about the column.
	Info *datasetmd.ColumnInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	// Column type.
	Type ColumnType `protobuf:"varint,2,opt,name=type,proto3,enum=dataobj.metadata.indexpointers.v1.ColumnType" json:"type,omitempty"`
}

func (m *ColumnDesc) Reset()      { *m = ColumnDesc{} }
func (*ColumnDesc) ProtoMessage() {}
func (*ColumnDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e25fb77c0f3d151, []int{1}
}
func (m *ColumnDesc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ColumnDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ColumnDesc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ColumnDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnDesc.Merge(m, src)
}
func (m *ColumnDesc) XXX_Size() int {
	return m.Size()
}
func (m *ColumnDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnDesc.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnDesc proto.InternalMessageInfo

func (m *ColumnDesc) GetInfo() *datasetmd.ColumnInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ColumnDesc) GetType() ColumnType {
	if m != nil {
		return m.Type
	}
	return COLUMN_TYPE_UNSPECIFIED
}

// ColumnMetadata describes the metadata for a column.
type ColumnMetadata struct {
	// Pages within the column.
	Pages []*PageDesc `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (m *ColumnMetadata) Reset()      { *m = ColumnMetadata{} }
func (*ColumnMetadata) ProtoMessage() {}
func (*ColumnMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e25fb77c0f3d151, []int{2}
}
func (m *ColumnMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ColumnMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ColumnMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ColumnMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnMetadata.Merge(m, src)
}
func (m *ColumnMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ColumnMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnMetadata proto.InternalMessageInfo

func (m *ColumnMetadata) GetPages() []*PageDesc {
	if m != nil {
		return m.Pages
	}
	return nil
}

// PageDesc describes an individual page within a column.
type PageDesc struct {
	// Information about the page.
	Info *datasetmd.PageInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *PageDesc) Reset()      { *m = PageDesc{} }
func (*PageDesc) ProtoMessage() {}
func (*PageDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e25fb77c0f3d151, []int{3}
}
func (m *PageDesc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PageDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PageDesc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PageDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageDesc.Merge(m, src)
}
func (m *PageDesc) XXX_Size() int {
	return m.Size()
}
func (m *PageDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_PageDesc.DiscardUnknown(m)
}

var xxx_messageInfo_PageDesc proto.InternalMessageInfo

func (m *PageDesc) GetInfo() *datasetmd.PageInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterEnum("dataobj.metadata.indexpointers.v1.ColumnType", ColumnType_name, ColumnType_value)
	proto.RegisterType((*Metadata)(nil), "dataobj.metadata.indexpointers.v1.Metadata")
	proto.RegisterType((*ColumnDesc)(nil), "dataobj.metadata.indexpointers.v1.ColumnDesc")
	proto.RegisterType((*ColumnMetadata)(nil), "dataobj.metadata.indexpointers.v1.ColumnMetadata")
	proto.RegisterType((*PageDesc)(nil), "dataobj.metadata.indexpointers.v1.PageDesc")
}

func init() {
	proto.RegisterFile("pkg/dataobj/internal/metadata/indexpointersmd/indexpointersmd.proto", fileDescriptor_5e25fb77c0f3d151)
}

var fileDescriptor_5e25fb77c0f3d151 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2e, 0xc8, 0x4e, 0xd7,
	0x4f, 0x49, 0x2c, 0x49, 0xcc, 0x4f, 0xca, 0xd2, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc,
	0xd1, 0xcf, 0x4d, 0x2d, 0x49, 0x04, 0x09, 0xea, 0x67, 0xe6, 0xa5, 0xa4, 0x56, 0x14, 0xe4, 0x83,
	0x25, 0x8a, 0x73, 0x53, 0xd0, 0xf9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x8a, 0x50, 0x03,
	0xf4, 0x60, 0xfa, 0xf4, 0x50, 0xd4, 0xe9, 0x95, 0x19, 0x4a, 0x99, 0xe3, 0xb7, 0x07, 0x44, 0x14,
	0xa7, 0x96, 0xe4, 0xa6, 0x20, 0x58, 0x10, 0xb3, 0x95, 0x82, 0xb9, 0x38, 0x7c, 0xa1, 0xaa, 0x84,
	0xdc, 0xb9, 0xd8, 0x93, 0xf3, 0x73, 0x4a, 0x73, 0xf3, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8,
	0x8d, 0x74, 0xf5, 0x08, 0xda, 0xac, 0xe7, 0x0c, 0xd6, 0xe1, 0x92, 0x5a, 0x9c, 0x1c, 0x04, 0xd3,
	0xad, 0xd4, 0xc3, 0xc8, 0xc5, 0x85, 0x10, 0x17, 0xb2, 0xe6, 0x62, 0xc9, 0xcc, 0x4b, 0xcb, 0x97,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xc7, 0x34, 0x14, 0xea, 0x28, 0x84, 0x71, 0x9e, 0x79,
	0x69, 0xf9, 0x41, 0x60, 0x4d, 0x42, 0x8e, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x7c, 0x24, 0xb8, 0x28, 0xa4, 0xb2, 0x20, 0x35, 0x08, 0xac, 0x55, 0x29, 0x98, 0x8b,
	0x0f, 0x22, 0x06, 0xf7, 0xa9, 0x23, 0x17, 0x6b, 0x41, 0x62, 0x7a, 0x2a, 0xcc, 0x9f, 0xda, 0x44,
	0x98, 0x1a, 0x90, 0x98, 0x9e, 0x0a, 0xf6, 0x25, 0x44, 0xa7, 0x92, 0x2b, 0x17, 0x07, 0x4c, 0x48,
	0xc8, 0x12, 0xc5, 0x83, 0xaa, 0x78, 0x3d, 0x08, 0xd2, 0x84, 0xf0, 0x9e, 0x56, 0x2d, 0x2c, 0xa4,
	0x40, 0xee, 0x15, 0x92, 0xe6, 0x12, 0x77, 0xf6, 0xf7, 0x09, 0xf5, 0xf5, 0x8b, 0x0f, 0x89, 0x0c,
	0x70, 0x8d, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10,
	0x12, 0xe1, 0x12, 0x40, 0x96, 0x0c, 0x70, 0x0c, 0xf1, 0x10, 0x60, 0x14, 0x92, 0xe5, 0x92, 0x44,
	0x16, 0xf5, 0xf5, 0xf4, 0x8b, 0x0f, 0xf1, 0xf4, 0x75, 0x0d, 0x0e, 0x71, 0xf4, 0x0d, 0x10, 0x60,
	0xc2, 0x90, 0x76, 0x8c, 0x40, 0x92, 0x66, 0x76, 0xaa, 0xbb, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43,
	0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0,
	0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0xe5, 0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x5e, 0x94,
	0x98, 0x96, 0x98, 0x97, 0xa8, 0x9f, 0x93, 0x9f, 0x9d, 0xa9, 0x5f, 0x66, 0xac, 0x4f, 0x52, 0x82,
	0x4f, 0x62, 0x03, 0xa7, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x46, 0x29, 0xcd,
	0x28, 0x03, 0x00, 0x00,
}

func (x ColumnType) String() string {
	s, ok := ColumnType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Metadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metadata)
	if !ok {
		that2, ok := that.(Metadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Columns) != len(that1.Columns) {
		return false
	}
	for i := range this.Columns {
		if !this.Columns[i].Equal(that1.Columns[i]) {
			return false
		}
	}
	return true
}
func (this *ColumnDesc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ColumnDesc)
	if !ok {
		that2, ok := that.(ColumnDesc)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Info.Equal(that1.Info) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *ColumnMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ColumnMetadata)
	if !ok {
		that2, ok := that.(ColumnMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Pages) != len(that1.Pages) {
		return false
	}
	for i := range this.Pages {
		if !this.Pages[i].Equal(that1.Pages[i]) {
			return false
		}
	}
	return true
}
func (this *PageDesc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PageDesc)
	if !ok {
		that2, ok := that.(PageDesc)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Info.Equal(that1.Info) {
		return false
	}
	return true
}
func (this *Metadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&indexpointersmd.Metadata{")
	if this.Columns != nil {
		s = append(s, "Columns: "+fmt.Sprintf("%#v", this.Columns)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ColumnDesc) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&indexpointersmd.ColumnDesc{")
	if this.Info != nil {
		s = append(s, "Info: "+fmt.Sprintf("%#v", this.Info)+",\n")
	}
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ColumnMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&indexpointersmd.ColumnMetadata{")
	if this.Pages != nil {
		s = append(s, "Pages: "+fmt.Sprintf("%#v", this.Pages)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PageDesc) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&indexpointersmd.PageDesc{")
	if this.Info != nil {
		s = append(s, "Info: "+fmt.Sprintf("%#v", this.Info)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIndexpointersmd(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Columns) > 0 {
		for iNdEx := len(m.Columns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Columns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIndexpointersmd(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ColumnDesc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ColumnDesc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ColumnDesc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintIndexpointersmd(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIndexpointersmd(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ColumnMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ColumnMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ColumnMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pages) > 0 {
		for iNdEx := len(m.Pages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIndexpointersmd(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PageDesc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageDesc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PageDesc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIndexpointersmd(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndexpointersmd(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndexpointersmd(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovIndexpointersmd(uint64(l))
		}
	}
	return n
}

func (m *ColumnDesc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovIndexpointersmd(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovIndexpointersmd(uint64(m.Type))
	}
	return n
}

func (m *ColumnMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pages) > 0 {
		for _, e := range m.Pages {
			l = e.Size()
			n += 1 + l + sovIndexpointersmd(uint64(l))
		}
	}
	return n
}

func (m *PageDesc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovIndexpointersmd(uint64(l))
	}
	return n
}

func sovIndexpointersmd(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndexpointersmd(x uint64) (n int) {
	return sovIndexpointersmd(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Metadata) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForColumns := "[]*ColumnDesc{"
	for _, f := range this.Columns {
		repeatedStringForColumns += strings.Replace(f.String(), "ColumnDesc", "ColumnDesc", 1) + ","
	}
	repeatedStringForColumns += "}"
	s := strings.Join([]string{`&Metadata{`,
		`Columns:` + repeatedStringForColumns + `,`,
		`}`,
	}, "")
	return s
}
func (this *ColumnDesc) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ColumnDesc{`,
		`Info:` + strings.Replace(fmt.Sprintf("%v", this.Info), "ColumnInfo", "datasetmd.ColumnInfo", 1) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ColumnMetadata) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPages := "[]*PageDesc{"
	for _, f := range this.Pages {
		repeatedStringForPages += strings.Replace(f.String(), "PageDesc", "PageDesc", 1) + ","
	}
	repeatedStringForPages += "}"
	s := strings.Join([]string{`&ColumnMetadata{`,
		`Pages:` + repeatedStringForPages + `,`,
		`}`,
	}, "")
	return s
}
func (this *PageDesc) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PageDesc{`,
		`Info:` + strings.Replace(fmt.Sprintf("%v", this.Info), "PageInfo", "datasetmd.PageInfo", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIndexpointersmd(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexpointersmd
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexpointersmd
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
				return ErrInvalidLengthIndexpointersmd
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnDesc{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndexpointersmd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndexpointersmd
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
func (m *ColumnDesc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexpointersmd
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
			return fmt.Errorf("proto: ColumnDesc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ColumnDesc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexpointersmd
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
				return ErrInvalidLengthIndexpointersmd
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &datasetmd.ColumnInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexpointersmd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ColumnType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndexpointersmd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndexpointersmd
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
func (m *ColumnMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexpointersmd
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
			return fmt.Errorf("proto: ColumnMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ColumnMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexpointersmd
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
				return ErrInvalidLengthIndexpointersmd
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pages = append(m.Pages, &PageDesc{})
			if err := m.Pages[len(m.Pages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndexpointersmd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndexpointersmd
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
func (m *PageDesc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexpointersmd
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
			return fmt.Errorf("proto: PageDesc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageDesc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexpointersmd
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
				return ErrInvalidLengthIndexpointersmd
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &datasetmd.PageInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndexpointersmd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndexpointersmd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndexpointersmd
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
func skipIndexpointersmd(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexpointersmd
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
					return 0, ErrIntOverflowIndexpointersmd
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
					return 0, ErrIntOverflowIndexpointersmd
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
				return 0, ErrInvalidLengthIndexpointersmd
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIndexpointersmd
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIndexpointersmd
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
				next, err := skipIndexpointersmd(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIndexpointersmd
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
	ErrInvalidLengthIndexpointersmd = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexpointersmd   = fmt.Errorf("proto: integer overflow")
)
