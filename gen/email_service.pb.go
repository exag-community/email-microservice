// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: email_service.proto

package exag

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_email_service_proto protoreflect.FileDescriptor

var file_email_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x45, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1e,
	0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6f,
	0x64, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x6c, 0x6c, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_email_service_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil), // 0: email.SendMessageRequest
	(*Empty)(nil),              // 1: common.Empty
}
var file_email_service_proto_depIdxs = []int32{
	0, // 0: email.EmailService.send_mail:input_type -> email.SendMessageRequest
	1, // 1: email.EmailService.send_mail:output_type -> common.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_email_service_proto_init() }
func file_email_service_proto_init() {
	if File_email_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_email_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_email_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_service_proto_goTypes,
		DependencyIndexes: file_email_service_proto_depIdxs,
	}.Build()
	File_email_service_proto = out.File
	file_email_service_proto_rawDesc = nil
	file_email_service_proto_goTypes = nil
	file_email_service_proto_depIdxs = nil
}
