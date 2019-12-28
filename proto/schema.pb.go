// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
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

type VersionResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestObjectById struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestObjectById) Reset()         { *m = RequestObjectById{} }
func (m *RequestObjectById) String() string { return proto.CompactTextString(m) }
func (*RequestObjectById) ProtoMessage()    {}
func (*RequestObjectById) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *RequestObjectById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestObjectById.Unmarshal(m, b)
}
func (m *RequestObjectById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestObjectById.Marshal(b, m, deterministic)
}
func (m *RequestObjectById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestObjectById.Merge(m, src)
}
func (m *RequestObjectById) XXX_Size() int {
	return xxx_messageInfo_RequestObjectById.Size(m)
}
func (m *RequestObjectById) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestObjectById.DiscardUnknown(m)
}

var xxx_messageInfo_RequestObjectById proto.InternalMessageInfo

func (m *RequestObjectById) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Column struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	MinLength            int32    `protobuf:"varint,2,opt,name=minLength,proto3" json:"minLength,omitempty"`
	MaxLength            int32    `protobuf:"varint,3,opt,name=maxLength,proto3" json:"maxLength,omitempty"`
	Minimum              int64    `protobuf:"varint,4,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Maximum              int64    `protobuf:"varint,5,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Default              *any.Any `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	Enum                 *any.Any `protobuf:"bytes,7,opt,name=enum,proto3" json:"enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Column) GetMinLength() int32 {
	if m != nil {
		return m.MinLength
	}
	return 0
}

func (m *Column) GetMaxLength() int32 {
	if m != nil {
		return m.MaxLength
	}
	return 0
}

func (m *Column) GetMinimum() int64 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *Column) GetMaximum() int64 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func (m *Column) GetDefault() *any.Any {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *Column) GetEnum() *any.Any {
	if m != nil {
		return m.Enum
	}
	return nil
}

type StringType struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MinLength            int32    `protobuf:"varint,3,opt,name=minLength,proto3" json:"minLength,omitempty"`
	MaxLength            int32    `protobuf:"varint,4,opt,name=maxLength,proto3" json:"maxLength,omitempty"`
	Default              string   `protobuf:"bytes,5,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringType) Reset()         { *m = StringType{} }
func (m *StringType) String() string { return proto.CompactTextString(m) }
func (*StringType) ProtoMessage()    {}
func (*StringType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}

func (m *StringType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringType.Unmarshal(m, b)
}
func (m *StringType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringType.Marshal(b, m, deterministic)
}
func (m *StringType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringType.Merge(m, src)
}
func (m *StringType) XXX_Size() int {
	return xxx_messageInfo_StringType.Size(m)
}
func (m *StringType) XXX_DiscardUnknown() {
	xxx_messageInfo_StringType.DiscardUnknown(m)
}

var xxx_messageInfo_StringType proto.InternalMessageInfo

func (m *StringType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StringType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StringType) GetMinLength() int32 {
	if m != nil {
		return m.MinLength
	}
	return 0
}

func (m *StringType) GetMaxLength() int32 {
	if m != nil {
		return m.MaxLength
	}
	return 0
}

func (m *StringType) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

type IntegerType struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Minimum              int64    `protobuf:"varint,3,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Maximum              int64    `protobuf:"varint,4,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Default              int64    `protobuf:"varint,5,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegerType) Reset()         { *m = IntegerType{} }
func (m *IntegerType) String() string { return proto.CompactTextString(m) }
func (*IntegerType) ProtoMessage()    {}
func (*IntegerType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{4}
}

func (m *IntegerType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerType.Unmarshal(m, b)
}
func (m *IntegerType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerType.Marshal(b, m, deterministic)
}
func (m *IntegerType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerType.Merge(m, src)
}
func (m *IntegerType) XXX_Size() int {
	return xxx_messageInfo_IntegerType.Size(m)
}
func (m *IntegerType) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerType.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerType proto.InternalMessageInfo

func (m *IntegerType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *IntegerType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IntegerType) GetMinimum() int64 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *IntegerType) GetMaximum() int64 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func (m *IntegerType) GetDefault() int64 {
	if m != nil {
		return m.Default
	}
	return 0
}

type FloatType struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Minimum              float64  `protobuf:"fixed64,3,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Maximum              float64  `protobuf:"fixed64,4,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Default              float64  `protobuf:"fixed64,5,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatType) Reset()         { *m = FloatType{} }
func (m *FloatType) String() string { return proto.CompactTextString(m) }
func (*FloatType) ProtoMessage()    {}
func (*FloatType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{5}
}

func (m *FloatType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatType.Unmarshal(m, b)
}
func (m *FloatType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatType.Marshal(b, m, deterministic)
}
func (m *FloatType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatType.Merge(m, src)
}
func (m *FloatType) XXX_Size() int {
	return xxx_messageInfo_FloatType.Size(m)
}
func (m *FloatType) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatType.DiscardUnknown(m)
}

var xxx_messageInfo_FloatType proto.InternalMessageInfo

func (m *FloatType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FloatType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FloatType) GetMinimum() float64 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *FloatType) GetMaximum() float64 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func (m *FloatType) GetDefault() float64 {
	if m != nil {
		return m.Default
	}
	return 0
}

type DateTimeType struct {
	Type                 string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Minimum              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=minimum,proto3" json:"minimum,omitempty"`
	Maximum              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=maximum,proto3" json:"maximum,omitempty"`
	Default              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DateTimeType) Reset()         { *m = DateTimeType{} }
func (m *DateTimeType) String() string { return proto.CompactTextString(m) }
func (*DateTimeType) ProtoMessage()    {}
func (*DateTimeType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{6}
}

func (m *DateTimeType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateTimeType.Unmarshal(m, b)
}
func (m *DateTimeType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateTimeType.Marshal(b, m, deterministic)
}
func (m *DateTimeType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateTimeType.Merge(m, src)
}
func (m *DateTimeType) XXX_Size() int {
	return xxx_messageInfo_DateTimeType.Size(m)
}
func (m *DateTimeType) XXX_DiscardUnknown() {
	xxx_messageInfo_DateTimeType.DiscardUnknown(m)
}

var xxx_messageInfo_DateTimeType proto.InternalMessageInfo

func (m *DateTimeType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DateTimeType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DateTimeType) GetMinimum() *timestamp.Timestamp {
	if m != nil {
		return m.Minimum
	}
	return nil
}

func (m *DateTimeType) GetMaximum() *timestamp.Timestamp {
	if m != nil {
		return m.Maximum
	}
	return nil
}

func (m *DateTimeType) GetDefault() *timestamp.Timestamp {
	if m != nil {
		return m.Default
	}
	return nil
}

type BooleanType struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BooleanType) Reset()         { *m = BooleanType{} }
func (m *BooleanType) String() string { return proto.CompactTextString(m) }
func (*BooleanType) ProtoMessage()    {}
func (*BooleanType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{7}
}

func (m *BooleanType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BooleanType.Unmarshal(m, b)
}
func (m *BooleanType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BooleanType.Marshal(b, m, deterministic)
}
func (m *BooleanType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooleanType.Merge(m, src)
}
func (m *BooleanType) XXX_Size() int {
	return xxx_messageInfo_BooleanType.Size(m)
}
func (m *BooleanType) XXX_DiscardUnknown() {
	xxx_messageInfo_BooleanType.DiscardUnknown(m)
}

var xxx_messageInfo_BooleanType proto.InternalMessageInfo

func (m *BooleanType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BooleanType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Schema struct {
	Id                   string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Properties           map[string]*_struct.Struct `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Required             []string                   `protobuf:"bytes,3,rep,name=required,proto3" json:"required,omitempty"`
	Name                 string                     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ProjectID            string                     `protobuf:"bytes,5,opt,name=projectID,proto3" json:"projectID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{8}
}

func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schema) GetProperties() map[string]*_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Schema) GetRequired() []string {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *Schema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Schema) GetProjectID() string {
	if m != nil {
		return m.ProjectID
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "proto.VersionResponse")
	proto.RegisterType((*RequestObjectById)(nil), "proto.RequestObjectById")
	proto.RegisterType((*Column)(nil), "proto.Column")
	proto.RegisterType((*StringType)(nil), "proto.StringType")
	proto.RegisterType((*IntegerType)(nil), "proto.IntegerType")
	proto.RegisterType((*FloatType)(nil), "proto.FloatType")
	proto.RegisterType((*DateTimeType)(nil), "proto.DateTimeType")
	proto.RegisterType((*BooleanType)(nil), "proto.BooleanType")
	proto.RegisterType((*Schema)(nil), "proto.Schema")
	proto.RegisterMapType((map[string]*_struct.Struct)(nil), "proto.Schema.PropertiesEntry")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x65, 0x92, 0xdd, 0xad, 0xb9, 0xa9, 0x56, 0x87, 0x82, 0x71, 0xa9, 0x18, 0xe2, 0x83, 0x79,
	0x31, 0x85, 0xda, 0x07, 0x51, 0x7c, 0x70, 0x5b, 0x85, 0x05, 0x41, 0xc9, 0x2e, 0x7d, 0xf0, 0x2d,
	0xdd, 0xdc, 0x6e, 0x47, 0x33, 0x33, 0xe9, 0x64, 0x22, 0xcd, 0xaf, 0x10, 0xdf, 0xfc, 0x85, 0xfd,
	0x17, 0x82, 0xec, 0x24, 0xe9, 0x7e, 0xb4, 0xdb, 0x82, 0xf5, 0x29, 0x73, 0xef, 0x39, 0x37, 0x39,
	0xe7, 0x90, 0xb9, 0xb0, 0x59, 0x4c, 0x4e, 0x91, 0x27, 0x51, 0xae, 0xa4, 0x96, 0xb4, 0x6b, 0x1e,
	0xfd, 0x9d, 0xa9, 0x94, 0xd3, 0x0c, 0x77, 0x4d, 0x75, 0x5c, 0x9e, 0xec, 0x16, 0x5a, 0x95, 0x13,
	0x5d, 0x93, 0xfa, 0x4f, 0x56, 0xd1, 0x44, 0x54, 0x0d, 0xf4, 0x6c, 0x15, 0xd2, 0x8c, 0x63, 0xa1,
	0x13, 0x9e, 0xd7, 0x84, 0xe0, 0x05, 0x6c, 0x1d, 0xa1, 0x2a, 0x98, 0x14, 0x31, 0x16, 0xb9, 0x14,
	0x05, 0xd2, 0x6d, 0xe8, 0xfe, 0x48, 0xb2, 0x12, 0x3d, 0xe2, 0x93, 0xd0, 0x89, 0xeb, 0x22, 0x78,
	0x0e, 0x8f, 0x62, 0x3c, 0x2b, 0xb1, 0xd0, 0x9f, 0x8f, 0xbf, 0xe1, 0x44, 0x0f, 0xaa, 0x61, 0x4a,
	0x1f, 0x80, 0xc5, 0xd2, 0x86, 0x67, 0xb1, 0x34, 0xb8, 0x20, 0xd0, 0x3b, 0x90, 0x59, 0xc9, 0x05,
	0xa5, 0xd0, 0xd1, 0x55, 0xde, 0xbe, 0xc4, 0x9c, 0xe9, 0x0e, 0x38, 0x9c, 0x89, 0x4f, 0x28, 0xa6,
	0xfa, 0xd4, 0xb3, 0x7c, 0x12, 0x76, 0xe3, 0x79, 0xc3, 0xa0, 0xc9, 0x79, 0x83, 0xda, 0x0d, 0xda,
	0x36, 0xa8, 0x07, 0x1b, 0x9c, 0x09, 0xc6, 0x4b, 0xee, 0x75, 0x7c, 0x12, 0xda, 0x71, 0x5b, 0x1a,
	0x24, 0x39, 0x37, 0x48, 0xb7, 0x41, 0xea, 0x92, 0x46, 0xb0, 0x91, 0xe2, 0x49, 0x52, 0x66, 0xda,
	0xeb, 0xf9, 0x24, 0x74, 0xf7, 0xb6, 0xa3, 0x3a, 0x8f, 0xa8, 0xcd, 0x23, 0x7a, 0x2f, 0xaa, 0xb8,
	0x25, 0xd1, 0x10, 0x3a, 0x28, 0x4a, 0xee, 0x6d, 0xdc, 0x40, 0x36, 0x8c, 0xe0, 0x37, 0x01, 0x18,
	0x69, 0xc5, 0xc4, 0x74, 0x3c, 0x33, 0x76, 0x9d, 0x59, 0x1f, 0xdc, 0x14, 0x8b, 0x89, 0x62, 0xb9,
	0x66, 0x52, 0x18, 0xbb, 0x4e, 0xbc, 0xd8, 0x5a, 0x8e, 0xc3, 0xbe, 0x31, 0x8e, 0xce, 0x35, 0x71,
	0xb4, 0xd6, 0xba, 0xe6, 0xcd, 0x6d, 0x19, 0xfc, 0x22, 0xe0, 0x0e, 0x85, 0xc6, 0x29, 0xaa, 0x3b,
	0x68, 0x5b, 0x88, 0xdb, 0x5e, 0x1b, 0x77, 0x67, 0x39, 0xee, 0x15, 0x4d, 0xf6, 0x5c, 0xd3, 0x4f,
	0x02, 0xce, 0xc7, 0x4c, 0x26, 0xfa, 0xff, 0x29, 0x22, 0x6b, 0x15, 0x91, 0xb5, 0x8a, 0xc8, 0x5c,
	0xd1, 0x05, 0x81, 0xcd, 0xc3, 0x44, 0xe3, 0x98, 0x71, 0xbc, 0x83, 0xa8, 0xfd, 0x65, 0x51, 0xee,
	0x5e, 0xff, 0xca, 0x4f, 0x33, 0x6e, 0x6f, 0xdc, 0x5c, 0xf0, 0xfe, 0xb2, 0xe0, 0xdb, 0xa6, 0x1a,
	0x33, 0xfb, 0xcb, 0x66, 0x6e, 0x99, 0x6a, 0x8d, 0x1e, 0x80, 0x3b, 0x90, 0x32, 0xc3, 0x44, 0xfc,
	0xbb, 0xcd, 0xe0, 0x0f, 0x81, 0xde, 0xc8, 0xec, 0xa5, 0xd5, 0x2b, 0x4f, 0xdf, 0x01, 0xe4, 0x4a,
	0xe6, 0xa8, 0x34, 0xc3, 0xc2, 0xb3, 0x7c, 0x3b, 0x74, 0xf7, 0x9e, 0xd6, 0x8a, 0xa2, 0x7a, 0x24,
	0xfa, 0x72, 0x89, 0x7f, 0x10, 0x5a, 0x55, 0xf1, 0xc2, 0x00, 0xed, 0xc3, 0x3d, 0x85, 0x67, 0x25,
	0x53, 0x98, 0x7a, 0xb6, 0x6f, 0x87, 0x4e, 0x7c, 0x59, 0xcf, 0xb4, 0x8a, 0x84, 0xa3, 0xc9, 0xc8,
	0x89, 0xcd, 0x79, 0x76, 0x2b, 0x72, 0x25, 0x67, 0x0b, 0x68, 0x78, 0xd8, 0xfc, 0xf9, 0xf3, 0x46,
	0xff, 0x08, 0xb6, 0x56, 0x3e, 0x46, 0x1f, 0x82, 0xfd, 0x1d, 0xab, 0x46, 0xf0, 0xec, 0x48, 0x5f,
	0xb6, 0xfb, 0xcd, 0x32, 0x29, 0x3e, 0xbe, 0x92, 0xe2, 0xc8, 0x2c, 0xd7, 0x66, 0xf1, 0xbd, 0xb1,
	0x5e, 0x93, 0xc1, 0xfd, 0xaf, 0xae, 0x41, 0xdf, 0xd6, 0x9c, 0x9e, 0x79, 0xbc, 0xfa, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x45, 0xff, 0x66, 0xa3, 0xac, 0x05, 0x00, 0x00,
}
