// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: eigr/functions/protocol/actors/extensions.proto

package actors

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_eigr_functions_protocol_actors_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         9999,
		Name:          "eigr.functions.protocol.actors.actor_id",
		Tag:           "varint,9999,opt,name=actor_id",
		Filename:      "eigr/functions/protocol/actors/extensions.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool actor_id = 9999;
	E_ActorId = &file_eigr_functions_protocol_actors_extensions_proto_extTypes[0]
)

var File_eigr_functions_protocol_actors_extensions_proto protoreflect.FileDescriptor

var file_eigr_functions_protocol_actors_extensions_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x69, 0x67, 0x72, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x65, 0x69, 0x67, 0x72, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3a, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8f,
	0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x42, 0x26,
	0x5a, 0x24, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x2f, 0x65, 0x69, 0x67, 0x72, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eigr_functions_protocol_actors_extensions_proto_goTypes = []any{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_eigr_functions_protocol_actors_extensions_proto_depIdxs = []int32{
	0, // 0: eigr.functions.protocol.actors.actor_id:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eigr_functions_protocol_actors_extensions_proto_init() }
func file_eigr_functions_protocol_actors_extensions_proto_init() {
	if File_eigr_functions_protocol_actors_extensions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eigr_functions_protocol_actors_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_eigr_functions_protocol_actors_extensions_proto_goTypes,
		DependencyIndexes: file_eigr_functions_protocol_actors_extensions_proto_depIdxs,
		ExtensionInfos:    file_eigr_functions_protocol_actors_extensions_proto_extTypes,
	}.Build()
	File_eigr_functions_protocol_actors_extensions_proto = out.File
	file_eigr_functions_protocol_actors_extensions_proto_rawDesc = nil
	file_eigr_functions_protocol_actors_extensions_proto_goTypes = nil
	file_eigr_functions_protocol_actors_extensions_proto_depIdxs = nil
}
