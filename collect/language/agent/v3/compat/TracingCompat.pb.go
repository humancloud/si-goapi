//
// Copyright (c) 2019 Stackinsights to present
// All rights reserved
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// language-agent/TracingCompat.proto is a deprecated file.

package compat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v31 "demo.stackinsights.ai/repo/goapi/collect/common/v3"
	v3 "demo.stackinsights.ai/repo/goapi/collect/language/agent/v3"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_language_agent_TracingCompat_proto protoreflect.FileDescriptor

var file_language_agent_TracingCompat_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x01, 0x0a, 0x19, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x12, 0x1c, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x20, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x42, 0xab, 0x01, 0x0a, 0x3a, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x50, 0x01, 0x5a, 0x41, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0xb8, 0x01,
	0x01, 0xaa, 0x02, 0x24, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56,
	0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_language_agent_TracingCompat_proto_goTypes = []interface{}{
	(*v3.SegmentObject)(nil),     // 0: stackinsights.v3.SegmentObject
	(*v3.SegmentCollection)(nil), // 1: stackinsights.v3.SegmentCollection
	(*v31.Commands)(nil),         // 2: stackinsights.v3.Commands
}
var file_language_agent_TracingCompat_proto_depIdxs = []int32{
	0, // 0: TraceSegmentReportService.collect:input_type -> stackinsights.v3.SegmentObject
	1, // 1: TraceSegmentReportService.collectInSync:input_type -> stackinsights.v3.SegmentCollection
	2, // 2: TraceSegmentReportService.collect:output_type -> stackinsights.v3.Commands
	2, // 3: TraceSegmentReportService.collectInSync:output_type -> stackinsights.v3.Commands
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_language_agent_TracingCompat_proto_init() }
func file_language_agent_TracingCompat_proto_init() {
	if File_language_agent_TracingCompat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_language_agent_TracingCompat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_language_agent_TracingCompat_proto_goTypes,
		DependencyIndexes: file_language_agent_TracingCompat_proto_depIdxs,
	}.Build()
	File_language_agent_TracingCompat_proto = out.File
	file_language_agent_TracingCompat_proto_rawDesc = nil
	file_language_agent_TracingCompat_proto_goTypes = nil
	file_language_agent_TracingCompat_proto_depIdxs = nil
}
