// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: protogen/book.proto

package bookpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Year   int32  `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	Isbn   string `protobuf:"bytes,5,opt,name=isbn,proto3" json:"isbn,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_protogen_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Book) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

type CreateBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookReq) Reset() {
	*x = CreateBookReq{}
	mi := &file_protogen_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookReq) ProtoMessage() {}

func (x *CreateBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookReq.ProtoReflect.Descriptor instead.
func (*CreateBookReq) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookReq) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type CreateBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBookResp) Reset() {
	*x = CreateBookResp{}
	mi := &file_protogen_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResp) ProtoMessage() {}

func (x *CreateBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResp.ProtoReflect.Descriptor instead.
func (*CreateBookResp) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookReq) Reset() {
	*x = GetBookReq{}
	mi := &file_protogen_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReq) ProtoMessage() {}

func (x *GetBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReq.ProtoReflect.Descriptor instead.
func (*GetBookReq) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResp) Reset() {
	*x = GetBookResp{}
	mi := &file_protogen_book_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResp) ProtoMessage() {}

func (x *GetBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResp.ProtoReflect.Descriptor instead.
func (*GetBookResp) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookResp) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type UpdateBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *UpdateBookReq) Reset() {
	*x = UpdateBookReq{}
	mi := &file_protogen_book_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookReq) ProtoMessage() {}

func (x *UpdateBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookReq.ProtoReflect.Descriptor instead.
func (*UpdateBookReq) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookReq) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type UpdateBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateBookResp) Reset() {
	*x = UpdateBookResp{}
	mi := &file_protogen_book_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookResp) ProtoMessage() {}

func (x *UpdateBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookResp.ProtoReflect.Descriptor instead.
func (*UpdateBookResp) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBookResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookReq) Reset() {
	*x = DeleteBookReq{}
	mi := &file_protogen_book_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookReq) ProtoMessage() {}

func (x *DeleteBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookReq.ProtoReflect.Descriptor instead.
func (*DeleteBookReq) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteBookResp) Reset() {
	*x = DeleteBookResp{}
	mi := &file_protogen_book_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResp) ProtoMessage() {}

func (x *DeleteBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResp.ProtoReflect.Descriptor instead.
func (*DeleteBookResp) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBookResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListBooksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBooksReq) Reset() {
	*x = ListBooksReq{}
	mi := &file_protogen_book_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBooksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksReq) ProtoMessage() {}

func (x *ListBooksReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksReq.ProtoReflect.Descriptor instead.
func (*ListBooksReq) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{9}
}

type ListBooksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *ListBooksResp) Reset() {
	*x = ListBooksResp{}
	mi := &file_protogen_book_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBooksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResp) ProtoMessage() {}

func (x *ListBooksResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResp.ProtoReflect.Descriptor instead.
func (*ListBooksResp) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{10}
}

func (x *ListBooksResp) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_protogen_book_proto protoreflect.FileDescriptor

var file_protogen_book_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x22, 0x6c, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x22, 0x31, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x20,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22,
	0x31, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1f,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x22, 0x33, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x32, 0xb2, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x15,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x15, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protogen_book_proto_rawDescOnce sync.Once
	file_protogen_book_proto_rawDescData = file_protogen_book_proto_rawDesc
)

func file_protogen_book_proto_rawDescGZIP() []byte {
	file_protogen_book_proto_rawDescOnce.Do(func() {
		file_protogen_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_protogen_book_proto_rawDescData)
	})
	return file_protogen_book_proto_rawDescData
}

var file_protogen_book_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protogen_book_proto_goTypes = []any{
	(*Book)(nil),           // 0: bookpb.Book
	(*CreateBookReq)(nil),  // 1: bookpb.CreateBookReq
	(*CreateBookResp)(nil), // 2: bookpb.CreateBookResp
	(*GetBookReq)(nil),     // 3: bookpb.GetBookReq
	(*GetBookResp)(nil),    // 4: bookpb.GetBookResp
	(*UpdateBookReq)(nil),  // 5: bookpb.UpdateBookReq
	(*UpdateBookResp)(nil), // 6: bookpb.UpdateBookResp
	(*DeleteBookReq)(nil),  // 7: bookpb.DeleteBookReq
	(*DeleteBookResp)(nil), // 8: bookpb.DeleteBookResp
	(*ListBooksReq)(nil),   // 9: bookpb.ListBooksReq
	(*ListBooksResp)(nil),  // 10: bookpb.ListBooksResp
}
var file_protogen_book_proto_depIdxs = []int32{
	0,  // 0: bookpb.CreateBookReq.book:type_name -> bookpb.Book
	0,  // 1: bookpb.GetBookResp.book:type_name -> bookpb.Book
	0,  // 2: bookpb.UpdateBookReq.book:type_name -> bookpb.Book
	0,  // 3: bookpb.ListBooksResp.books:type_name -> bookpb.Book
	1,  // 4: bookpb.BookService.CreateBook:input_type -> bookpb.CreateBookReq
	3,  // 5: bookpb.BookService.GetBook:input_type -> bookpb.GetBookReq
	5,  // 6: bookpb.BookService.UpdateBook:input_type -> bookpb.UpdateBookReq
	7,  // 7: bookpb.BookService.DeleteBook:input_type -> bookpb.DeleteBookReq
	9,  // 8: bookpb.BookService.ListBooks:input_type -> bookpb.ListBooksReq
	2,  // 9: bookpb.BookService.CreateBook:output_type -> bookpb.CreateBookResp
	4,  // 10: bookpb.BookService.GetBook:output_type -> bookpb.GetBookResp
	6,  // 11: bookpb.BookService.UpdateBook:output_type -> bookpb.UpdateBookResp
	8,  // 12: bookpb.BookService.DeleteBook:output_type -> bookpb.DeleteBookResp
	10, // 13: bookpb.BookService.ListBooks:output_type -> bookpb.ListBooksResp
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protogen_book_proto_init() }
func file_protogen_book_proto_init() {
	if File_protogen_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protogen_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protogen_book_proto_goTypes,
		DependencyIndexes: file_protogen_book_proto_depIdxs,
		MessageInfos:      file_protogen_book_proto_msgTypes,
	}.Build()
	File_protogen_book_proto = out.File
	file_protogen_book_proto_rawDesc = nil
	file_protogen_book_proto_goTypes = nil
	file_protogen_book_proto_depIdxs = nil
}
