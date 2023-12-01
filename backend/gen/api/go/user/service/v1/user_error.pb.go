// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: user/service/v1/user_error.proto

package v1

import (
	_ "github.com/devexps/go-micro/v2/errors"
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

// UserErrorReason
type UserErrorReason int32

const (
	UserErrorReason_BAD_REQUEST        UserErrorReason = 0
	UserErrorReason_USER_NOT_EXIST     UserErrorReason = 1
	UserErrorReason_INVALID_PASSWORD   UserErrorReason = 2
	UserErrorReason_INVALID_REQUEST    UserErrorReason = 3
	UserErrorReason_INVALID_FIELD_MASK UserErrorReason = 4
	UserErrorReason_QUERY_DATA_FAILED  UserErrorReason = 5
	UserErrorReason_INSERT_DATA_FAILED UserErrorReason = 6
	UserErrorReason_UPDATE_DATA_FAILED UserErrorReason = 7
	UserErrorReason_DELETE_DATA_FAILED UserErrorReason = 8
)

// Enum value maps for UserErrorReason.
var (
	UserErrorReason_name = map[int32]string{
		0: "BAD_REQUEST",
		1: "USER_NOT_EXIST",
		2: "INVALID_PASSWORD",
		3: "INVALID_REQUEST",
		4: "INVALID_FIELD_MASK",
		5: "QUERY_DATA_FAILED",
		6: "INSERT_DATA_FAILED",
		7: "UPDATE_DATA_FAILED",
		8: "DELETE_DATA_FAILED",
	}
	UserErrorReason_value = map[string]int32{
		"BAD_REQUEST":        0,
		"USER_NOT_EXIST":     1,
		"INVALID_PASSWORD":   2,
		"INVALID_REQUEST":    3,
		"INVALID_FIELD_MASK": 4,
		"QUERY_DATA_FAILED":  5,
		"INSERT_DATA_FAILED": 6,
		"UPDATE_DATA_FAILED": 7,
		"DELETE_DATA_FAILED": 8,
	}
)

func (x UserErrorReason) Enum() *UserErrorReason {
	p := new(UserErrorReason)
	*p = x
	return p
}

func (x UserErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_v1_user_error_proto_enumTypes[0].Descriptor()
}

func (UserErrorReason) Type() protoreflect.EnumType {
	return &file_user_service_v1_user_error_proto_enumTypes[0]
}

func (x UserErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserErrorReason.Descriptor instead.
func (UserErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_user_service_v1_user_error_proto_rawDescGZIP(), []int{0}
}

var File_user_service_v1_user_error_proto protoreflect.FileDescriptor

var file_user_service_v1_user_error_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x94, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x0b,
	0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa8,
	0x45, 0x91, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x92, 0x03, 0x12, 0x1a, 0x0a,
	0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x1a, 0x04,
	0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4d, 0x41, 0x53, 0x4b, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45,
	0x95, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x96, 0x03, 0x12,
	0x1c, 0x0a, 0x12, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x97, 0x03, 0x12, 0x1c, 0x0a,
	0x12, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x08, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42,
	0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x76, 0x65, 0x78, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_service_v1_user_error_proto_rawDescOnce sync.Once
	file_user_service_v1_user_error_proto_rawDescData = file_user_service_v1_user_error_proto_rawDesc
)

func file_user_service_v1_user_error_proto_rawDescGZIP() []byte {
	file_user_service_v1_user_error_proto_rawDescOnce.Do(func() {
		file_user_service_v1_user_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_v1_user_error_proto_rawDescData)
	})
	return file_user_service_v1_user_error_proto_rawDescData
}

var file_user_service_v1_user_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_v1_user_error_proto_goTypes = []interface{}{
	(UserErrorReason)(0), // 0: user.service.v1.UserErrorReason
}
var file_user_service_v1_user_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_v1_user_error_proto_init() }
func file_user_service_v1_user_error_proto_init() {
	if File_user_service_v1_user_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_v1_user_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_service_v1_user_error_proto_goTypes,
		DependencyIndexes: file_user_service_v1_user_error_proto_depIdxs,
		EnumInfos:         file_user_service_v1_user_error_proto_enumTypes,
	}.Build()
	File_user_service_v1_user_error_proto = out.File
	file_user_service_v1_user_error_proto_rawDesc = nil
	file_user_service_v1_user_error_proto_goTypes = nil
	file_user_service_v1_user_error_proto_depIdxs = nil
}
