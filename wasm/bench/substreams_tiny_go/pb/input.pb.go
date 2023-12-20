// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: input.proto

package pb

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

type MapBlockInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params     string `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Block      *Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	ReadStore  uint32 `protobuf:"varint,3,opt,name=read_store,json=readStore,proto3" json:"read_store,omitempty"`
	ReadStore2 uint32 `protobuf:"varint,4,opt,name=read_store2,json=readStore2,proto3" json:"read_store2,omitempty"`
	WriteStore uint32 `protobuf:"varint,5,opt,name=write_store,json=writeStore,proto3" json:"write_store,omitempty"`
}

func (x *MapBlockInput) Reset() {
	*x = MapBlockInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapBlockInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapBlockInput) ProtoMessage() {}

func (x *MapBlockInput) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapBlockInput.ProtoReflect.Descriptor instead.
func (*MapBlockInput) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{0}
}

func (x *MapBlockInput) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *MapBlockInput) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *MapBlockInput) GetReadStore() uint32 {
	if x != nil {
		return x.ReadStore
	}
	return 0
}

func (x *MapBlockInput) GetReadStore2() uint32 {
	if x != nil {
		return x.ReadStore2
	}
	return 0
}

func (x *MapBlockInput) GetWriteStore() uint32 {
	if x != nil {
		return x.WriteStore
	}
	return 0
}

var File_input_proto protoreflect.FileDescriptor

var file_input_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0d, 0x4d, 0x61,
	0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x32, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_input_proto_rawDescOnce sync.Once
	file_input_proto_rawDescData = file_input_proto_rawDesc
)

func file_input_proto_rawDescGZIP() []byte {
	file_input_proto_rawDescOnce.Do(func() {
		file_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_input_proto_rawDescData)
	})
	return file_input_proto_rawDescData
}

var file_input_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_input_proto_goTypes = []interface{}{
	(*MapBlockInput)(nil), // 0: MapBlockInput
	(*Block)(nil),         // 1: Block
}
var file_input_proto_depIdxs = []int32{
	1, // 0: MapBlockInput.block:type_name -> Block
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_input_proto_init() }
func file_input_proto_init() {
	if File_input_proto != nil {
		return
	}
	file_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapBlockInput); i {
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
			RawDescriptor: file_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_input_proto_goTypes,
		DependencyIndexes: file_input_proto_depIdxs,
		MessageInfos:      file_input_proto_msgTypes,
	}.Build()
	File_input_proto = out.File
	file_input_proto_rawDesc = nil
	file_input_proto_goTypes = nil
	file_input_proto_depIdxs = nil
}