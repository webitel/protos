// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: logger_service.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ACTION int32

const (
	ACTION_DEFAULT_NO_ACTION ACTION = 0
	ACTION_CREATE            ACTION = 1
	ACTION_UPDATE            ACTION = 2
	ACTION_READ              ACTION = 3
	ACTION_DELETE            ACTION = 4
)

// Enum value maps for ACTION.
var (
	ACTION_name = map[int32]string{
		0: "DEFAULT_NO_ACTION",
		1: "CREATE",
		2: "UPDATE",
		3: "READ",
		4: "DELETE",
	}
	ACTION_value = map[string]int32{
		"DEFAULT_NO_ACTION": 0,
		"CREATE":            1,
		"UPDATE":            2,
		"READ":              3,
		"DELETE":            4,
	}
)

func (x ACTION) Enum() *ACTION {
	p := new(ACTION)
	*p = x
	return p
}

func (x ACTION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACTION) Descriptor() protoreflect.EnumDescriptor {
	return file_logger_service_proto_enumTypes[0].Descriptor()
}

func (ACTION) Type() protoreflect.EnumType {
	return &file_logger_service_proto_enumTypes[0]
}

func (x ACTION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACTION.Descriptor instead.
func (ACTION) EnumDescriptor() ([]byte, []int) {
	return file_logger_service_proto_rawDescGZIP(), []int{0}
}

type SearchLogByConfigIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q      string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort   string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	// REQUIRED filter
	ConfigId int32 `protobuf:"varint,7,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	// SPECIFIC filter
	User *Lookup `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty"`
	// GENERAL filters
	Action   ACTION `protobuf:"varint,8,opt,name=action,proto3,enum=logger.ACTION" json:"action,omitempty"`
	UserIp   string `protobuf:"bytes,10,opt,name=userIp,proto3" json:"userIp,omitempty"`
	DateFrom int64  `protobuf:"varint,11,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo   int64  `protobuf:"varint,12,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
}

func (x *SearchLogByConfigIdRequest) Reset() {
	*x = SearchLogByConfigIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLogByConfigIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLogByConfigIdRequest) ProtoMessage() {}

func (x *SearchLogByConfigIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logger_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLogByConfigIdRequest.ProtoReflect.Descriptor instead.
func (*SearchLogByConfigIdRequest) Descriptor() ([]byte, []int) {
	return file_logger_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchLogByConfigIdRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchLogByConfigIdRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchLogByConfigIdRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchLogByConfigIdRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchLogByConfigIdRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchLogByConfigIdRequest) GetConfigId() int32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *SearchLogByConfigIdRequest) GetUser() *Lookup {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SearchLogByConfigIdRequest) GetAction() ACTION {
	if x != nil {
		return x.Action
	}
	return ACTION_DEFAULT_NO_ACTION
}

func (x *SearchLogByConfigIdRequest) GetUserIp() string {
	if x != nil {
		return x.UserIp
	}
	return ""
}

func (x *SearchLogByConfigIdRequest) GetDateFrom() int64 {
	if x != nil {
		return x.DateFrom
	}
	return 0
}

func (x *SearchLogByConfigIdRequest) GetDateTo() int64 {
	if x != nil {
		return x.DateTo
	}
	return 0
}

type SearchLogByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q      string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort   string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	// REQUIRED filter
	UserId int32 `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// SPECIFIC filter
	Object *Lookup `protobuf:"bytes,10,opt,name=object,proto3" json:"object,omitempty"`
	// GENERAL filters
	Action   ACTION `protobuf:"varint,8,opt,name=action,proto3,enum=logger.ACTION" json:"action,omitempty"`
	UserIp   string `protobuf:"bytes,9,opt,name=userIp,proto3" json:"userIp,omitempty"`
	DateFrom int64  `protobuf:"varint,11,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo   int64  `protobuf:"varint,12,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
}

func (x *SearchLogByUserIdRequest) Reset() {
	*x = SearchLogByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLogByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLogByUserIdRequest) ProtoMessage() {}

func (x *SearchLogByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logger_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLogByUserIdRequest.ProtoReflect.Descriptor instead.
func (*SearchLogByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_logger_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchLogByUserIdRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchLogByUserIdRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchLogByUserIdRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchLogByUserIdRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchLogByUserIdRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchLogByUserIdRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SearchLogByUserIdRequest) GetObject() *Lookup {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *SearchLogByUserIdRequest) GetAction() ACTION {
	if x != nil {
		return x.Action
	}
	return ACTION_DEFAULT_NO_ACTION
}

func (x *SearchLogByUserIdRequest) GetUserIp() string {
	if x != nil {
		return x.UserIp
	}
	return ""
}

func (x *SearchLogByUserIdRequest) GetDateFrom() int64 {
	if x != nil {
		return x.DateFrom
	}
	return 0
}

func (x *SearchLogByUserIdRequest) GetDateTo() int64 {
	if x != nil {
		return x.DateTo
	}
	return 0
}

type Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Next  bool   `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
	Items []*Log `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Logs) Reset() {
	*x = Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs) ProtoMessage() {}

func (x *Logs) ProtoReflect() protoreflect.Message {
	mi := &file_logger_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs.ProtoReflect.Descriptor instead.
func (*Logs) Descriptor() ([]byte, []int) {
	return file_logger_service_proto_rawDescGZIP(), []int{2}
}

func (x *Logs) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Logs) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *Logs) GetItems() []*Log {
	if x != nil {
		return x.Items
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Action   string  `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Date     int64   `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	User     *Lookup `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	UserIp   string  `protobuf:"bytes,5,opt,name=user_ip,json=userIp,proto3" json:"user_ip,omitempty"`
	NewState string  `protobuf:"bytes,6,opt,name=new_state,json=newState,proto3" json:"new_state,omitempty"`
	ConfigId int32   `protobuf:"varint,7,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Record   *Lookup `protobuf:"bytes,8,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_logger_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_logger_service_proto_rawDescGZIP(), []int{3}
}

func (x *Log) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Log) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Log) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Log) GetUser() *Lookup {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Log) GetUserIp() string {
	if x != nil {
		return x.UserIp
	}
	return ""
}

func (x *Log) GetNewState() string {
	if x != nil {
		return x.NewState
	}
	return ""
}

func (x *Log) GetConfigId() int32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *Log) GetRecord() *Lookup {
	if x != nil {
		return x.Record
	}
	return nil
}

type Lookup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Lookup) Reset() {
	*x = Lookup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lookup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lookup) ProtoMessage() {}

func (x *Lookup) ProtoReflect() protoreflect.Message {
	mi := &file_logger_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lookup.ProtoReflect.Descriptor instead.
func (*Lookup) Descriptor() ([]byte, []int) {
	return file_logger_service_proto_rawDescGZIP(), []int{4}
}

func (x *Lookup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lookup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_logger_service_proto protoreflect.FileDescriptor

var file_logger_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a,
	0x1a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0x51, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x03, 0x4c, 0x6f,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x06,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x4d, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f,
	0x4e, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x32, 0xeb, 0x01, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4c, 0x6f, 0x67, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x70, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c,
	0x6f, 0x67, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x42,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x77, 0x65, 0x62, 0x69, 0x74,
	0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logger_service_proto_rawDescOnce sync.Once
	file_logger_service_proto_rawDescData = file_logger_service_proto_rawDesc
)

func file_logger_service_proto_rawDescGZIP() []byte {
	file_logger_service_proto_rawDescOnce.Do(func() {
		file_logger_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_logger_service_proto_rawDescData)
	})
	return file_logger_service_proto_rawDescData
}

var file_logger_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_logger_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_logger_service_proto_goTypes = []interface{}{
	(ACTION)(0),                        // 0: logger.ACTION
	(*SearchLogByConfigIdRequest)(nil), // 1: logger.SearchLogByConfigIdRequest
	(*SearchLogByUserIdRequest)(nil),   // 2: logger.SearchLogByUserIdRequest
	(*Logs)(nil),                       // 3: logger.Logs
	(*Log)(nil),                        // 4: logger.Log
	(*Lookup)(nil),                     // 5: logger.Lookup
}
var file_logger_service_proto_depIdxs = []int32{
	5, // 0: logger.SearchLogByConfigIdRequest.user:type_name -> logger.Lookup
	0, // 1: logger.SearchLogByConfigIdRequest.action:type_name -> logger.ACTION
	5, // 2: logger.SearchLogByUserIdRequest.object:type_name -> logger.Lookup
	0, // 3: logger.SearchLogByUserIdRequest.action:type_name -> logger.ACTION
	4, // 4: logger.Logs.items:type_name -> logger.Log
	5, // 5: logger.Log.user:type_name -> logger.Lookup
	5, // 6: logger.Log.record:type_name -> logger.Lookup
	2, // 7: logger.LoggerService.SearchLogByUserId:input_type -> logger.SearchLogByUserIdRequest
	1, // 8: logger.LoggerService.SearchLogByConfigId:input_type -> logger.SearchLogByConfigIdRequest
	3, // 9: logger.LoggerService.SearchLogByUserId:output_type -> logger.Logs
	3, // 10: logger.LoggerService.SearchLogByConfigId:output_type -> logger.Logs
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_logger_service_proto_init() }
func file_logger_service_proto_init() {
	if File_logger_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logger_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLogByConfigIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logger_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLogByUserIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logger_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logger_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logger_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lookup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logger_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logger_service_proto_goTypes,
		DependencyIndexes: file_logger_service_proto_depIdxs,
		EnumInfos:         file_logger_service_proto_enumTypes,
		MessageInfos:      file_logger_service_proto_msgTypes,
	}.Build()
	File_logger_service_proto = out.File
	file_logger_service_proto_rawDesc = nil
	file_logger_service_proto_goTypes = nil
	file_logger_service_proto_depIdxs = nil
}
