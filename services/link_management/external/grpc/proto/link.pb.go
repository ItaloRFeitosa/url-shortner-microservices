// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: services/link_management/external/grpc/proto/link.proto

package proto

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

type CreateLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	RedirectTo  string `protobuf:"bytes,2,opt,name=redirect_to,json=redirectTo,proto3" json:"redirect_to,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_management_external_grpc_proto_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_management_external_grpc_proto_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_services_link_management_external_grpc_proto_link_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLinkRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreateLinkRequest) GetRedirectTo() string {
	if x != nil {
		return x.RedirectTo
	}
	return ""
}

func (x *CreateLinkRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateLinkRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type LinkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkspaceId string `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Slug        string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	RedirectTo  string `protobuf:"bytes,4,opt,name=redirect_to,json=redirectTo,proto3" json:"redirect_to,omitempty"`
	Title       string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *LinkData) Reset() {
	*x = LinkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_management_external_grpc_proto_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkData) ProtoMessage() {}

func (x *LinkData) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_management_external_grpc_proto_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkData.ProtoReflect.Descriptor instead.
func (*LinkData) Descriptor() ([]byte, []int) {
	return file_services_link_management_external_grpc_proto_link_proto_rawDescGZIP(), []int{1}
}

func (x *LinkData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LinkData) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *LinkData) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *LinkData) GetRedirectTo() string {
	if x != nil {
		return x.RedirectTo
	}
	return ""
}

func (x *LinkData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LinkData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LinkData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type LinkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string    `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    *LinkData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error   bool      `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LinkReply) Reset() {
	*x = LinkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_management_external_grpc_proto_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkReply) ProtoMessage() {}

func (x *LinkReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_management_external_grpc_proto_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkReply.ProtoReflect.Descriptor instead.
func (*LinkReply) Descriptor() ([]byte, []int) {
	return file_services_link_management_external_grpc_proto_link_proto_rawDescGZIP(), []int{2}
}

func (x *LinkReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LinkReply) GetData() *LinkData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LinkReply) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_services_link_management_external_grpc_proto_link_proto protoreflect.FileDescriptor

var file_services_link_management_external_grpc_proto_link_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60,
	0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x51, 0x0a, 0x15, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_link_management_external_grpc_proto_link_proto_rawDescOnce sync.Once
	file_services_link_management_external_grpc_proto_link_proto_rawDescData = file_services_link_management_external_grpc_proto_link_proto_rawDesc
)

func file_services_link_management_external_grpc_proto_link_proto_rawDescGZIP() []byte {
	file_services_link_management_external_grpc_proto_link_proto_rawDescOnce.Do(func() {
		file_services_link_management_external_grpc_proto_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_link_management_external_grpc_proto_link_proto_rawDescData)
	})
	return file_services_link_management_external_grpc_proto_link_proto_rawDescData
}

var file_services_link_management_external_grpc_proto_link_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_link_management_external_grpc_proto_link_proto_goTypes = []interface{}{
	(*CreateLinkRequest)(nil), // 0: proto.CreateLinkRequest
	(*LinkData)(nil),          // 1: proto.LinkData
	(*LinkReply)(nil),         // 2: proto.LinkReply
}
var file_services_link_management_external_grpc_proto_link_proto_depIdxs = []int32{
	1, // 0: proto.LinkReply.data:type_name -> proto.LinkData
	0, // 1: proto.LinkManagementService.CreateLink:input_type -> proto.CreateLinkRequest
	2, // 2: proto.LinkManagementService.CreateLink:output_type -> proto.LinkReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_link_management_external_grpc_proto_link_proto_init() }
func file_services_link_management_external_grpc_proto_link_proto_init() {
	if File_services_link_management_external_grpc_proto_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_link_management_external_grpc_proto_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkRequest); i {
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
		file_services_link_management_external_grpc_proto_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkData); i {
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
		file_services_link_management_external_grpc_proto_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkReply); i {
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
			RawDescriptor: file_services_link_management_external_grpc_proto_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_link_management_external_grpc_proto_link_proto_goTypes,
		DependencyIndexes: file_services_link_management_external_grpc_proto_link_proto_depIdxs,
		MessageInfos:      file_services_link_management_external_grpc_proto_link_proto_msgTypes,
	}.Build()
	File_services_link_management_external_grpc_proto_link_proto = out.File
	file_services_link_management_external_grpc_proto_link_proto_rawDesc = nil
	file_services_link_management_external_grpc_proto_link_proto_goTypes = nil
	file_services_link_management_external_grpc_proto_link_proto_depIdxs = nil
}
