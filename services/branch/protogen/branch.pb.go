// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: protogen/branch.proto

package branchpb

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

type MkBranchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Postcode string `protobuf:"bytes,4,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Country  string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *MkBranchReq) Reset() {
	*x = MkBranchReq{}
	mi := &file_protogen_branch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MkBranchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MkBranchReq) ProtoMessage() {}

func (x *MkBranchReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MkBranchReq.ProtoReflect.Descriptor instead.
func (*MkBranchReq) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{0}
}

func (x *MkBranchReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MkBranchReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MkBranchReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MkBranchReq) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

func (x *MkBranchReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	City     string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Postcode string `protobuf:"bytes,5,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Country  string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	mi := &file_protogen_branch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{1}
}

func (x *Branch) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Branch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Branch) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Branch) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Branch) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

func (x *Branch) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type MkBranchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MkBranchResp) Reset() {
	*x = MkBranchResp{}
	mi := &file_protogen_branch_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MkBranchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MkBranchResp) ProtoMessage() {}

func (x *MkBranchResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MkBranchResp.ProtoReflect.Descriptor instead.
func (*MkBranchResp) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{2}
}

func (x *MkBranchResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MkBranchResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetBranchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBranchReq) Reset() {
	*x = GetBranchReq{}
	mi := &file_protogen_branch_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchReq) ProtoMessage() {}

func (x *GetBranchReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchReq.ProtoReflect.Descriptor instead.
func (*GetBranchReq) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{3}
}

func (x *GetBranchReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBranchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branch *Branch `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *GetBranchResp) Reset() {
	*x = GetBranchResp{}
	mi := &file_protogen_branch_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchResp) ProtoMessage() {}

func (x *GetBranchResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchResp.ProtoReflect.Descriptor instead.
func (*GetBranchResp) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{4}
}

func (x *GetBranchResp) GetBranch() *Branch {
	if x != nil {
		return x.Branch
	}
	return nil
}

type UpdateBranchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	City     string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Postcode string `protobuf:"bytes,5,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Country  string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *UpdateBranchReq) Reset() {
	*x = UpdateBranchReq{}
	mi := &file_protogen_branch_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBranchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranchReq) ProtoMessage() {}

func (x *UpdateBranchReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranchReq.ProtoReflect.Descriptor instead.
func (*UpdateBranchReq) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBranchReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBranchReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBranchReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateBranchReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateBranchReq) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

func (x *UpdateBranchReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type UpdateBranchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateBranchResp) Reset() {
	*x = UpdateBranchResp{}
	mi := &file_protogen_branch_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBranchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranchResp) ProtoMessage() {}

func (x *UpdateBranchResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranchResp.ProtoReflect.Descriptor instead.
func (*UpdateBranchResp) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBranchResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateBranchResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteBranchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBranchReq) Reset() {
	*x = DeleteBranchReq{}
	mi := &file_protogen_branch_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBranchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBranchReq) ProtoMessage() {}

func (x *DeleteBranchReq) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBranchReq.ProtoReflect.Descriptor instead.
func (*DeleteBranchReq) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBranchReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBranchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteBranchResp) Reset() {
	*x = DeleteBranchResp{}
	mi := &file_protogen_branch_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBranchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBranchResp) ProtoMessage() {}

func (x *DeleteBranchResp) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_branch_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBranchResp.ProtoReflect.Descriptor instead.
func (*DeleteBranchResp) Descriptor() ([]byte, []int) {
	return file_protogen_branch_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBranchResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteBranchResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protogen_branch_proto protoreflect.FileDescriptor

var file_protogen_branch_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70,
	0x62, 0x22, 0x85, 0x01, 0x0a, 0x0b, 0x4d, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x06, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x0c,
	0x4d, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x28, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70, 0x62, 0x2e, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x99, 0x01, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x46, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e, 0x02, 0x0a, 0x0d, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08,
	0x4d, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x70, 0x62, 0x2e, 0x4d, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70, 0x62, 0x2e, 0x4d, 0x6b, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x12, 0x19, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x3b, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protogen_branch_proto_rawDescOnce sync.Once
	file_protogen_branch_proto_rawDescData = file_protogen_branch_proto_rawDesc
)

func file_protogen_branch_proto_rawDescGZIP() []byte {
	file_protogen_branch_proto_rawDescOnce.Do(func() {
		file_protogen_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_protogen_branch_proto_rawDescData)
	})
	return file_protogen_branch_proto_rawDescData
}

var file_protogen_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protogen_branch_proto_goTypes = []any{
	(*MkBranchReq)(nil),      // 0: branchpb.MkBranchReq
	(*Branch)(nil),           // 1: branchpb.Branch
	(*MkBranchResp)(nil),     // 2: branchpb.MkBranchResp
	(*GetBranchReq)(nil),     // 3: branchpb.GetBranchReq
	(*GetBranchResp)(nil),    // 4: branchpb.GetBranchResp
	(*UpdateBranchReq)(nil),  // 5: branchpb.UpdateBranchReq
	(*UpdateBranchResp)(nil), // 6: branchpb.UpdateBranchResp
	(*DeleteBranchReq)(nil),  // 7: branchpb.DeleteBranchReq
	(*DeleteBranchResp)(nil), // 8: branchpb.DeleteBranchResp
}
var file_protogen_branch_proto_depIdxs = []int32{
	1, // 0: branchpb.GetBranchResp.branch:type_name -> branchpb.Branch
	0, // 1: branchpb.BranchService.MkBranch:input_type -> branchpb.MkBranchReq
	3, // 2: branchpb.BranchService.GetBranch:input_type -> branchpb.GetBranchReq
	5, // 3: branchpb.BranchService.UpdateBranch:input_type -> branchpb.UpdateBranchReq
	7, // 4: branchpb.BranchService.DeleteBranch:input_type -> branchpb.DeleteBranchReq
	2, // 5: branchpb.BranchService.MkBranch:output_type -> branchpb.MkBranchResp
	4, // 6: branchpb.BranchService.GetBranch:output_type -> branchpb.GetBranchResp
	6, // 7: branchpb.BranchService.UpdateBranch:output_type -> branchpb.UpdateBranchResp
	8, // 8: branchpb.BranchService.DeleteBranch:output_type -> branchpb.DeleteBranchResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protogen_branch_proto_init() }
func file_protogen_branch_proto_init() {
	if File_protogen_branch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protogen_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protogen_branch_proto_goTypes,
		DependencyIndexes: file_protogen_branch_proto_depIdxs,
		MessageInfos:      file_protogen_branch_proto_msgTypes,
	}.Build()
	File_protogen_branch_proto = out.File
	file_protogen_branch_proto_rawDesc = nil
	file_protogen_branch_proto_goTypes = nil
	file_protogen_branch_proto_depIdxs = nil
}
