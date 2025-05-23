// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: etcd.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EtcdKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *EtcdKey) Reset() {
	*x = EtcdKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdKey) ProtoMessage() {}

func (x *EtcdKey) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdKey.ProtoReflect.Descriptor instead.
func (*EtcdKey) Descriptor() ([]byte, []int) {
	return file_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *EtcdKey) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type      string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Delimiter string   `protobuf:"bytes,3,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	Tables    []string `protobuf:"bytes,4,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_etcd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_etcd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *Dataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dataset) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Dataset) GetDelimiter() string {
	if x != nil {
		return x.Delimiter
	}
	return ""
}

func (x *Dataset) GetTables() []string {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_etcd_proto protoreflect.FileDescriptor

var file_etcd_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x45, 0x74, 0x63, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x67, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x32, 0x80, 0x01, 0x0a, 0x04, 0x45,
	0x74, 0x63, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x45, 0x74, 0x63, 0x64, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f,
	0x73, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x4b, 0x65, 0x79, 0x1a, 0x10, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x00, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x72, 0x72,
	0x69, 0x74, 0x30, 0x35, 0x2f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x4f, 0x53, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_etcd_proto_rawDescOnce sync.Once
	file_etcd_proto_rawDescData = file_etcd_proto_rawDesc
)

func file_etcd_proto_rawDescGZIP() []byte {
	file_etcd_proto_rawDescOnce.Do(func() {
		file_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(file_etcd_proto_rawDescData)
	})
	return file_etcd_proto_rawDescData
}

var file_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_etcd_proto_goTypes = []interface{}{
	(*EtcdKey)(nil),       // 0: dynamos.EtcdKey
	(*Dataset)(nil),       // 1: dynamos.Dataset
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_etcd_proto_depIdxs = []int32{
	2, // 0: dynamos.Etcd.InitEtcd:input_type -> google.protobuf.Empty
	0, // 1: dynamos.Etcd.GetDatasetMetadata:input_type -> dynamos.EtcdKey
	2, // 2: dynamos.Etcd.InitEtcd:output_type -> google.protobuf.Empty
	1, // 3: dynamos.Etcd.GetDatasetMetadata:output_type -> dynamos.Dataset
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_etcd_proto_init() }
func file_etcd_proto_init() {
	if File_etcd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_etcd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtcdKey); i {
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
		file_etcd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
			RawDescriptor: file_etcd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_etcd_proto_goTypes,
		DependencyIndexes: file_etcd_proto_depIdxs,
		MessageInfos:      file_etcd_proto_msgTypes,
	}.Build()
	File_etcd_proto = out.File
	file_etcd_proto_rawDesc = nil
	file_etcd_proto_goTypes = nil
	file_etcd_proto_depIdxs = nil
}
