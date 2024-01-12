// Code generated by protoc-gen-go. DO NOT EDIT.
// source: books.proto

package GRPC

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

type BookRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookRequest) Reset()         { *m = BookRequest{} }
func (m *BookRequest) String() string { return proto.CompactTextString(m) }
func (*BookRequest) ProtoMessage()    {}
func (*BookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e0dc127ded4184, []int{0}
}

func (m *BookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookRequest.Unmarshal(m, b)
}
func (m *BookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookRequest.Marshal(b, m, deterministic)
}
func (m *BookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookRequest.Merge(m, src)
}
func (m *BookRequest) XXX_Size() int {
	return xxx_messageInfo_BookRequest.Size(m)
}
func (m *BookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BookRequest proto.InternalMessageInfo

func (m *BookRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Book struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Isbn                 string   `protobuf:"bytes,2,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Author               *Author  `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e0dc127ded4184, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetIsbn() string {
	if m != nil {
		return m.Isbn
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

type Author struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e0dc127ded4184, []int{2}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Author) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type BookListResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookListResponse) Reset()         { *m = BookListResponse{} }
func (m *BookListResponse) String() string { return proto.CompactTextString(m) }
func (*BookListResponse) ProtoMessage()    {}
func (*BookListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e0dc127ded4184, []int{3}
}

func (m *BookListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookListResponse.Unmarshal(m, b)
}
func (m *BookListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookListResponse.Marshal(b, m, deterministic)
}
func (m *BookListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookListResponse.Merge(m, src)
}
func (m *BookListResponse) XXX_Size() int {
	return xxx_messageInfo_BookListResponse.Size(m)
}
func (m *BookListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookListResponse proto.InternalMessageInfo

func (m *BookListResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type BookResponse struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookResponse) Reset()         { *m = BookResponse{} }
func (m *BookResponse) String() string { return proto.CompactTextString(m) }
func (*BookResponse) ProtoMessage()    {}
func (*BookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e0dc127ded4184, []int{4}
}

func (m *BookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookResponse.Unmarshal(m, b)
}
func (m *BookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookResponse.Marshal(b, m, deterministic)
}
func (m *BookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookResponse.Merge(m, src)
}
func (m *BookResponse) XXX_Size() int {
	return xxx_messageInfo_BookResponse.Size(m)
}
func (m *BookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BookResponse proto.InternalMessageInfo

func (m *BookResponse) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type NoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoRequest) Reset()         { *m = NoRequest{} }
func (m *NoRequest) String() string { return proto.CompactTextString(m) }
func (*NoRequest) ProtoMessage()    {}
func (*NoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e0dc127ded4184, []int{5}
}

func (m *NoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoRequest.Unmarshal(m, b)
}
func (m *NoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoRequest.Marshal(b, m, deterministic)
}
func (m *NoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoRequest.Merge(m, src)
}
func (m *NoRequest) XXX_Size() int {
	return xxx_messageInfo_NoRequest.Size(m)
}
func (m *NoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BookRequest)(nil), "BookRequest")
	proto.RegisterType((*Book)(nil), "Book")
	proto.RegisterType((*Author)(nil), "Author")
	proto.RegisterType((*BookListResponse)(nil), "BookListResponse")
	proto.RegisterType((*BookResponse)(nil), "BookResponse")
	proto.RegisterType((*NoRequest)(nil), "NoRequest")
}

func init() {
	proto.RegisterFile("books.proto", fileDescriptor_01e0dc127ded4184)
}

var fileDescriptor_01e0dc127ded4184 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0x68, 0xf9, 0xf3, 0x15, 0x8d, 0x6e, 0x3c, 0x54, 0xd4, 0x48, 0xf6, 0x60, 0xe0,
	0xc0, 0x92, 0xe0, 0x13, 0x88, 0x07, 0x2f, 0xc6, 0x98, 0x7a, 0xf3, 0x56, 0xe8, 0xa8, 0x1b, 0x90,
	0xc5, 0xee, 0xd6, 0x07, 0xf6, 0x49, 0x4c, 0x77, 0xd7, 0x2a, 0xe8, 0xc1, 0x5b, 0xe7, 0x37, 0xbf,
	0x74, 0xbe, 0x99, 0x2c, 0xe2, 0xb9, 0x52, 0x4b, 0x2d, 0x36, 0x85, 0x32, 0x8a, 0x9f, 0x21, 0x9e,
	0x29, 0xb5, 0x4c, 0xe9, 0xad, 0x24, 0x6d, 0xd8, 0x3e, 0x1a, 0x32, 0x4f, 0x82, 0x41, 0x30, 0xec,
	0xa6, 0x0d, 0x99, 0xf3, 0x0c, 0x61, 0xd5, 0xde, 0xe5, 0x8c, 0x21, 0x94, 0x7a, 0xbe, 0x4e, 0x1a,
	0x96, 0xd8, 0x6f, 0x76, 0x84, 0xc8, 0x48, 0xb3, 0xa2, 0xa4, 0x69, 0xa1, 0x2b, 0xd8, 0x39, 0x5a,
	0x59, 0x69, 0x5e, 0x54, 0x91, 0x84, 0x83, 0x60, 0x18, 0x4f, 0xdb, 0xe2, 0xca, 0x96, 0xa9, 0xc7,
	0x7c, 0x86, 0x96, 0x23, 0xec, 0x14, 0xdd, 0x27, 0x59, 0x68, 0xb3, 0xce, 0x5e, 0xc9, 0xcf, 0xfa,
	0x06, 0xac, 0x8f, 0xce, 0x2a, 0xf3, 0x4d, 0x37, 0xb6, 0xae, 0xf9, 0x04, 0x07, 0x55, 0xcc, 0x5b,
	0xa9, 0x4d, 0x4a, 0x7a, 0xa3, 0xd6, 0x9a, 0xd8, 0x09, 0x22, 0xbb, 0x68, 0x12, 0x0c, 0x9a, 0xc3,
	0x78, 0x1a, 0x09, 0xbb, 0xa7, 0x63, 0x7c, 0x84, 0x9e, 0x5b, 0xdb, 0xcb, 0xc7, 0x08, 0xab, 0x86,
	0x9d, 0x5a, 0xbb, 0x16, 0xf1, 0x18, 0xdd, 0x3b, 0xe5, 0xef, 0x33, 0xfd, 0x08, 0xdc, 0xbd, 0x1e,
	0xa8, 0x78, 0x97, 0x0b, 0x62, 0x23, 0x74, 0x9e, 0xc9, 0x54, 0x44, 0x33, 0x88, 0xda, 0xeb, 0x1f,
	0x8a, 0x5f, 0x79, 0x2e, 0xd0, 0xf6, 0x2a, 0xeb, 0x89, 0x1f, 0x37, 0xef, 0xef, 0x89, 0xad, 0x28,
	0x1c, 0x58, 0x14, 0x94, 0x19, 0xb2, 0xaa, 0x8b, 0xb2, 0xeb, 0x8c, 0x81, 0x72, 0x93, 0x7f, 0x39,
	0xdb, 0xbf, 0xfb, 0x63, 0xf4, 0x18, 0xc8, 0x69, 0x45, 0xff, 0xd4, 0x67, 0xed, 0xc7, 0x68, 0x72,
	0x93, 0xde, 0x5f, 0xcf, 0x5b, 0xf6, 0x8d, 0x5c, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe5,
	0x74, 0xef, 0x32, 0x02, 0x00, 0x00,
}
