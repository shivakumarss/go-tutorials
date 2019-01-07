// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SearchRequest.proto

// Simple protobuf file

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SearchRequestMessage struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequestMessage) Reset()         { *m = SearchRequestMessage{} }
func (m *SearchRequestMessage) String() string { return proto.CompactTextString(m) }
func (*SearchRequestMessage) ProtoMessage()    {}
func (*SearchRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7c0caf246c6cc37, []int{0}
}

func (m *SearchRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequestMessage.Unmarshal(m, b)
}
func (m *SearchRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequestMessage.Marshal(b, m, deterministic)
}
func (m *SearchRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequestMessage.Merge(m, src)
}
func (m *SearchRequestMessage) XXX_Size() int {
	return xxx_messageInfo_SearchRequestMessage.Size(m)
}
func (m *SearchRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequestMessage proto.InternalMessageInfo

func (m *SearchRequestMessage) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequestMessage) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequestMessage) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func init() {
	proto.RegisterType((*SearchRequestMessage)(nil), "example.SearchRequestMessage")
}

func init() { proto.RegisterFile("SearchRequest.proto", fileDescriptor_b7c0caf246c6cc37) }

var fileDescriptor_b7c0caf246c6cc37 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0e, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55, 0x2a, 0xe5, 0x12, 0x41, 0x91, 0xf7, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x2c, 0x4d, 0x2d, 0xaa, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xe4, 0xb9, 0xb8, 0x0b, 0x12, 0xd3, 0x53, 0xe3,
	0xf3, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0xb8, 0x40, 0x42,
	0x7e, 0x60, 0x11, 0x21, 0x35, 0x2e, 0xfe, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x92, 0xf8, 0x82, 0xd4,
	0xa2, 0x78, 0x90, 0x84, 0x04, 0x33, 0x58, 0x11, 0x2f, 0x44, 0x38, 0x20, 0xb5, 0x28, 0x20, 0x31,
	0x3d, 0x35, 0x89, 0x0d, 0xec, 0x0c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xa3, 0xe6,
	0x76, 0x9d, 0x00, 0x00, 0x00,
}
