// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: wg/cosmo/webhooks/common.proto

package webhooks

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

type FederatedGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FederatedGraph) Reset() {
	*x = FederatedGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_webhooks_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedGraph) ProtoMessage() {}

func (x *FederatedGraph) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_webhooks_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedGraph.ProtoReflect.Descriptor instead.
func (*FederatedGraph) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_webhooks_common_proto_rawDescGZIP(), []int{0}
}

func (x *FederatedGraph) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FederatedGraph) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_wg_cosmo_webhooks_common_proto protoreflect.FileDescriptor

var file_wg_cosmo_webhooks_common_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x77, 0x67, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0x22, 0x34, 0x0a, 0x0e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xcb, 0x01, 0x0a, 0x15, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x77, 0x67, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0xa2, 0x02, 0x03, 0x57, 0x43, 0x57, 0xaa, 0x02, 0x11, 0x57, 0x67, 0x2e, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0xca, 0x02, 0x11,
	0x57, 0x67, 0x5c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x5c, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0xe2, 0x02, 0x1d, 0x57, 0x67, 0x5c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x5c, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x13, 0x57, 0x67, 0x3a, 0x3a, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x3a, 0x3a, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wg_cosmo_webhooks_common_proto_rawDescOnce sync.Once
	file_wg_cosmo_webhooks_common_proto_rawDescData = file_wg_cosmo_webhooks_common_proto_rawDesc
)

func file_wg_cosmo_webhooks_common_proto_rawDescGZIP() []byte {
	file_wg_cosmo_webhooks_common_proto_rawDescOnce.Do(func() {
		file_wg_cosmo_webhooks_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_wg_cosmo_webhooks_common_proto_rawDescData)
	})
	return file_wg_cosmo_webhooks_common_proto_rawDescData
}

var file_wg_cosmo_webhooks_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wg_cosmo_webhooks_common_proto_goTypes = []interface{}{
	(*FederatedGraph)(nil), // 0: wg.cosmo.webhooks.FederatedGraph
}
var file_wg_cosmo_webhooks_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wg_cosmo_webhooks_common_proto_init() }
func file_wg_cosmo_webhooks_common_proto_init() {
	if File_wg_cosmo_webhooks_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wg_cosmo_webhooks_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedGraph); i {
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
			RawDescriptor: file_wg_cosmo_webhooks_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wg_cosmo_webhooks_common_proto_goTypes,
		DependencyIndexes: file_wg_cosmo_webhooks_common_proto_depIdxs,
		MessageInfos:      file_wg_cosmo_webhooks_common_proto_msgTypes,
	}.Build()
	File_wg_cosmo_webhooks_common_proto = out.File
	file_wg_cosmo_webhooks_common_proto_rawDesc = nil
	file_wg_cosmo_webhooks_common_proto_goTypes = nil
	file_wg_cosmo_webhooks_common_proto_depIdxs = nil
}
