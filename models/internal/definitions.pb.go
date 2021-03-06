// Code generated by protoc-gen-go.
// source: definitions.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	definitions.proto

It has these top-level messages:
	User
	Post
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type User struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,json=userName" json:"UserName,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Post struct {
	Type        string `protobuf:"bytes,1,opt,name=Type,json=type" json:"Type,omitempty"`
	File        string `protobuf:"bytes,2,opt,name=File,json=file" json:"File,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,json=description" json:"Description,omitempty"`
	Key         uint64 `protobuf:"varint,4,opt,name=Key,json=key" json:"Key,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*User)(nil), "internal.User")
	proto.RegisterType((*Post)(nil), "internal.Post")
}

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x49, 0x4d, 0xcb,
	0xcc, 0xcb, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8,
	0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x51, 0x52, 0xe2, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12,
	0x92, 0xe2, 0xe2, 0x00, 0xd1, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x1c, 0xa5, 0x50, 0xbe, 0x52, 0x12, 0x17, 0x4b, 0x40, 0x7e, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b,
	0x48, 0x65, 0x01, 0x4c, 0x9e, 0xa5, 0x04, 0xc8, 0x06, 0x89, 0xb9, 0x65, 0xe6, 0xa4, 0x4a, 0x30,
	0x41, 0xc4, 0xd2, 0x80, 0x6c, 0x21, 0x05, 0x2e, 0x6e, 0x97, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x02,
	0x90, 0x9d, 0x12, 0xcc, 0x60, 0x29, 0xee, 0x14, 0x84, 0x90, 0x90, 0x00, 0x17, 0xb3, 0x77, 0x6a,
	0xa5, 0x04, 0x0b, 0x50, 0x86, 0x25, 0x88, 0x39, 0x3b, 0xb5, 0x32, 0x89, 0x0d, 0xec, 0x30, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xb3, 0xa7, 0x02, 0xad, 0x00, 0x00, 0x00,
}
