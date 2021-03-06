// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: billing/v1/errors.proto

package v1

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

type ServiceErrors int32

const (
	ServiceErrors_ACCOUNT_NOT_FOUND  ServiceErrors = 0
	ServiceErrors_ACCOUNT_NOT_BELONG ServiceErrors = 1
	ServiceErrors_INSUFFICIENT_FUNDS ServiceErrors = 2
)

// Enum value maps for ServiceErrors.
var (
	ServiceErrors_name = map[int32]string{
		0: "ACCOUNT_NOT_FOUND",
		1: "ACCOUNT_NOT_BELONG",
		2: "INSUFFICIENT_FUNDS",
	}
	ServiceErrors_value = map[string]int32{
		"ACCOUNT_NOT_FOUND":  0,
		"ACCOUNT_NOT_BELONG": 1,
		"INSUFFICIENT_FUNDS": 2,
	}
)

func (x ServiceErrors) Enum() *ServiceErrors {
	p := new(ServiceErrors)
	*p = x
	return p
}

func (x ServiceErrors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceErrors) Descriptor() protoreflect.EnumDescriptor {
	return file_billing_v1_errors_proto_enumTypes[0].Descriptor()
}

func (ServiceErrors) Type() protoreflect.EnumType {
	return &file_billing_v1_errors_proto_enumTypes[0]
}

func (x ServiceErrors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceErrors.Descriptor instead.
func (ServiceErrors) EnumDescriptor() ([]byte, []int) {
	return file_billing_v1_errors_proto_rawDescGZIP(), []int{0}
}

var File_billing_v1_errors_proto protoreflect.FileDescriptor

var file_billing_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2a, 0x56, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x45, 0x4c,
	0x4f, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49,
	0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x55, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x42, 0x1b, 0x5a,
	0x19, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_billing_v1_errors_proto_rawDescOnce sync.Once
	file_billing_v1_errors_proto_rawDescData = file_billing_v1_errors_proto_rawDesc
)

func file_billing_v1_errors_proto_rawDescGZIP() []byte {
	file_billing_v1_errors_proto_rawDescOnce.Do(func() {
		file_billing_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_billing_v1_errors_proto_rawDescData)
	})
	return file_billing_v1_errors_proto_rawDescData
}

var file_billing_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_billing_v1_errors_proto_goTypes = []interface{}{
	(ServiceErrors)(0), // 0: billing.v1.ServiceErrors
}
var file_billing_v1_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_billing_v1_errors_proto_init() }
func file_billing_v1_errors_proto_init() {
	if File_billing_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_billing_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_billing_v1_errors_proto_goTypes,
		DependencyIndexes: file_billing_v1_errors_proto_depIdxs,
		EnumInfos:         file_billing_v1_errors_proto_enumTypes,
	}.Build()
	File_billing_v1_errors_proto = out.File
	file_billing_v1_errors_proto_rawDesc = nil
	file_billing_v1_errors_proto_goTypes = nil
	file_billing_v1_errors_proto_depIdxs = nil
}
