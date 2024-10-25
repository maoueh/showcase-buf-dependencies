// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: acme.proto

package pbacme

import (
	v2 "github.com/streamingfast/firehose-ethereum/types/pb/sf/ethereum/type/v2"
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

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BigInt *v2.BigInt `protobuf:"bytes,1,opt,name=big_int,json=bigInt,proto3" json:"big_int,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_acme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_acme_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetBigInt() *v2.BigInt {
	if x != nil {
		return x.BigInt
	}
	return nil
}

var File_acme_proto protoreflect.FileDescriptor

var file_acme_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x63,
	0x6d, 0x65, 0x1a, 0x1e, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x62, 0x69,
	0x67, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x66,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x06, 0x62, 0x69, 0x67, 0x49, 0x6e, 0x74,
	0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x63, 0x6d, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x70, 0x62, 0x61, 0x63, 0x6d, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acme_proto_rawDescOnce sync.Once
	file_acme_proto_rawDescData = file_acme_proto_rawDesc
)

func file_acme_proto_rawDescGZIP() []byte {
	file_acme_proto_rawDescOnce.Do(func() {
		file_acme_proto_rawDescData = protoimpl.X.CompressGZIP(file_acme_proto_rawDescData)
	})
	return file_acme_proto_rawDescData
}

var file_acme_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_acme_proto_goTypes = []interface{}{
	(*Test)(nil),      // 0: acme.Test
	(*v2.BigInt)(nil), // 1: sf.ethereum.type.v2.BigInt
}
var file_acme_proto_depIdxs = []int32{
	1, // 0: acme.Test.big_int:type_name -> sf.ethereum.type.v2.BigInt
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acme_proto_init() }
func file_acme_proto_init() {
	if File_acme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_acme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acme_proto_goTypes,
		DependencyIndexes: file_acme_proto_depIdxs,
		MessageInfos:      file_acme_proto_msgTypes,
	}.Build()
	File_acme_proto = out.File
	file_acme_proto_rawDesc = nil
	file_acme_proto_goTypes = nil
	file_acme_proto_depIdxs = nil
}
