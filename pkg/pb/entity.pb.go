// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: auth/entity.proto

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

type AuthProvider int32

const (
	AuthProvider_SELF   AuthProvider = 0
	AuthProvider_GOOGLE AuthProvider = 1
)

// Enum value maps for AuthProvider.
var (
	AuthProvider_name = map[int32]string{
		0: "SELF",
		1: "GOOGLE",
	}
	AuthProvider_value = map[string]int32{
		"SELF":   0,
		"GOOGLE": 1,
	}
)

func (x AuthProvider) Enum() *AuthProvider {
	p := new(AuthProvider)
	*p = x
	return p
}

func (x AuthProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_entity_proto_enumTypes[0].Descriptor()
}

func (AuthProvider) Type() protoreflect.EnumType {
	return &file_auth_entity_proto_enumTypes[0]
}

func (x AuthProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthProvider.Descriptor instead.
func (AuthProvider) EnumDescriptor() ([]byte, []int) {
	return file_auth_entity_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	DisplayName string       `protobuf:"bytes,3,opt,name=displayName,proto3" json:"displayName,omitempty"`
	ProfilePath string       `protobuf:"bytes,4,opt,name=profilePath,proto3" json:"profilePath,omitempty"`
	Provider    AuthProvider `protobuf:"varint,5,opt,name=provider,proto3,enum=codern.auth.AuthProvider" json:"provider,omitempty"`
	CreatedAt   int64        `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_entity_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetProfilePath() string {
	if x != nil {
		return x.ProfilePath
	}
	return ""
}

func (x *User) GetProvider() AuthProvider {
	if x != nil {
		return x.Provider
	}
	return AuthProvider_SELF
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_auth_entity_proto protoreflect.FileDescriptor

var file_auth_entity_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x22, 0xc5, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x6e, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x24, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x46,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_entity_proto_rawDescOnce sync.Once
	file_auth_entity_proto_rawDescData = file_auth_entity_proto_rawDesc
)

func file_auth_entity_proto_rawDescGZIP() []byte {
	file_auth_entity_proto_rawDescOnce.Do(func() {
		file_auth_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_entity_proto_rawDescData)
	})
	return file_auth_entity_proto_rawDescData
}

var file_auth_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_auth_entity_proto_goTypes = []interface{}{
	(AuthProvider)(0), // 0: codern.auth.AuthProvider
	(*User)(nil),      // 1: codern.auth.User
}
var file_auth_entity_proto_depIdxs = []int32{
	0, // 0: codern.auth.User.provider:type_name -> codern.auth.AuthProvider
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_entity_proto_init() }
func file_auth_entity_proto_init() {
	if File_auth_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_auth_entity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_entity_proto_goTypes,
		DependencyIndexes: file_auth_entity_proto_depIdxs,
		EnumInfos:         file_auth_entity_proto_enumTypes,
		MessageInfos:      file_auth_entity_proto_msgTypes,
	}.Build()
	File_auth_entity_proto = out.File
	file_auth_entity_proto_rawDesc = nil
	file_auth_entity_proto_goTypes = nil
	file_auth_entity_proto_depIdxs = nil
}
