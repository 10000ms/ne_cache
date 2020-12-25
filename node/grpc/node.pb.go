// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: node/grpc/node.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SetValueResponse_Status int32

const (
	SetValueResponse_FAIL SetValueResponse_Status = 0
	SetValueResponse_OK   SetValueResponse_Status = 1
)

// Enum value maps for SetValueResponse_Status.
var (
	SetValueResponse_Status_name = map[int32]string{
		0: "FAIL",
		1: "OK",
	}
	SetValueResponse_Status_value = map[string]int32{
		"FAIL": 0,
		"OK":   1,
	}
)

func (x SetValueResponse_Status) Enum() *SetValueResponse_Status {
	p := new(SetValueResponse_Status)
	*p = x
	return p
}

func (x SetValueResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetValueResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_node_grpc_node_proto_enumTypes[0].Descriptor()
}

func (SetValueResponse_Status) Type() protoreflect.EnumType {
	return &file_node_grpc_node_proto_enumTypes[0]
}

func (x SetValueResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetValueResponse_Status.Descriptor instead.
func (SetValueResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_node_grpc_node_proto_rawDescGZIP(), []int{1, 0}
}

type GetValueResponse_Status int32

const (
	GetValueResponse_FAIL GetValueResponse_Status = 0
	GetValueResponse_OK   GetValueResponse_Status = 1
)

// Enum value maps for GetValueResponse_Status.
var (
	GetValueResponse_Status_name = map[int32]string{
		0: "FAIL",
		1: "OK",
	}
	GetValueResponse_Status_value = map[string]int32{
		"FAIL": 0,
		"OK":   1,
	}
)

func (x GetValueResponse_Status) Enum() *GetValueResponse_Status {
	p := new(GetValueResponse_Status)
	*p = x
	return p
}

func (x GetValueResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetValueResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_node_grpc_node_proto_enumTypes[1].Descriptor()
}

func (GetValueResponse_Status) Type() protoreflect.EnumType {
	return &file_node_grpc_node_proto_enumTypes[1]
}

func (x GetValueResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetValueResponse_Status.Descriptor instead.
func (GetValueResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_node_grpc_node_proto_rawDescGZIP(), []int{3, 0}
}

// set
type SetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *SetValueRequest) Reset() {
	*x = SetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_grpc_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueRequest) ProtoMessage() {}

func (x *SetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_grpc_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueRequest.ProtoReflect.Descriptor instead.
func (*SetValueRequest) Descriptor() ([]byte, []int) {
	return file_node_grpc_node_proto_rawDescGZIP(), []int{0}
}

func (x *SetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetValueRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetValueRequest) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type SetValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SetValueResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ne_cache.node.grpc.SetValueResponse_Status" json:"status,omitempty"`
}

func (x *SetValueResponse) Reset() {
	*x = SetValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_grpc_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueResponse) ProtoMessage() {}

func (x *SetValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_grpc_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueResponse.ProtoReflect.Descriptor instead.
func (*SetValueResponse) Descriptor() ([]byte, []int) {
	return file_node_grpc_node_proto_rawDescGZIP(), []int{1}
}

func (x *SetValueResponse) GetStatus() SetValueResponse_Status {
	if x != nil {
		return x.Status
	}
	return SetValueResponse_FAIL
}

// get
type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_grpc_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_grpc_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_node_grpc_node_proto_rawDescGZIP(), []int{2}
}

func (x *GetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GetValueResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ne_cache.node.grpc.GetValueResponse_Status" json:"status,omitempty"`
	Value  []byte                  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetValueResponse) Reset() {
	*x = GetValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_grpc_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueResponse) ProtoMessage() {}

func (x *GetValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_grpc_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueResponse.ProtoReflect.Descriptor instead.
func (*GetValueResponse) Descriptor() ([]byte, []int) {
	return file_node_grpc_node_proto_rawDescGZIP(), []int{3}
}

func (x *GetValueResponse) GetStatus() GetValueResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetValueResponse_FAIL
}

func (x *GetValueResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_node_grpc_node_proto protoreflect.FileDescriptor

var file_node_grpc_node_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x51, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x73, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x01, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6e,
	0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x01, 0x32, 0xba, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x23, 0x2e, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x2e, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x65,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x0b, 0x5a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_grpc_node_proto_rawDescOnce sync.Once
	file_node_grpc_node_proto_rawDescData = file_node_grpc_node_proto_rawDesc
)

func file_node_grpc_node_proto_rawDescGZIP() []byte {
	file_node_grpc_node_proto_rawDescOnce.Do(func() {
		file_node_grpc_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_grpc_node_proto_rawDescData)
	})
	return file_node_grpc_node_proto_rawDescData
}

var file_node_grpc_node_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_node_grpc_node_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_node_grpc_node_proto_goTypes = []interface{}{
	(SetValueResponse_Status)(0), // 0: ne_cache.node.grpc.SetValueResponse.Status
	(GetValueResponse_Status)(0), // 1: ne_cache.node.grpc.GetValueResponse.Status
	(*SetValueRequest)(nil),      // 2: ne_cache.node.grpc.SetValueRequest
	(*SetValueResponse)(nil),     // 3: ne_cache.node.grpc.SetValueResponse
	(*GetValueRequest)(nil),      // 4: ne_cache.node.grpc.GetValueRequest
	(*GetValueResponse)(nil),     // 5: ne_cache.node.grpc.GetValueResponse
}
var file_node_grpc_node_proto_depIdxs = []int32{
	0, // 0: ne_cache.node.grpc.SetValueResponse.status:type_name -> ne_cache.node.grpc.SetValueResponse.Status
	1, // 1: ne_cache.node.grpc.GetValueResponse.status:type_name -> ne_cache.node.grpc.GetValueResponse.Status
	2, // 2: ne_cache.node.grpc.NodeService.SetValue:input_type -> ne_cache.node.grpc.SetValueRequest
	4, // 3: ne_cache.node.grpc.NodeService.GetValue:input_type -> ne_cache.node.grpc.GetValueRequest
	3, // 4: ne_cache.node.grpc.NodeService.SetValue:output_type -> ne_cache.node.grpc.SetValueResponse
	4, // 5: ne_cache.node.grpc.NodeService.GetValue:output_type -> ne_cache.node.grpc.GetValueRequest
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_grpc_node_proto_init() }
func file_node_grpc_node_proto_init() {
	if File_node_grpc_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_grpc_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueRequest); i {
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
		file_node_grpc_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueResponse); i {
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
		file_node_grpc_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_node_grpc_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueResponse); i {
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
			RawDescriptor: file_node_grpc_node_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_grpc_node_proto_goTypes,
		DependencyIndexes: file_node_grpc_node_proto_depIdxs,
		EnumInfos:         file_node_grpc_node_proto_enumTypes,
		MessageInfos:      file_node_grpc_node_proto_msgTypes,
	}.Build()
	File_node_grpc_node_proto = out.File
	file_node_grpc_node_proto_rawDesc = nil
	file_node_grpc_node_proto_goTypes = nil
	file_node_grpc_node_proto_depIdxs = nil
}