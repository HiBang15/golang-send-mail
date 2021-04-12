// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mail_message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateMailTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateMailTemplateRequest) Reset() {
	*x = CreateMailTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMailTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMailTemplateRequest) ProtoMessage() {}

func (x *CreateMailTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMailTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateMailTemplateRequest) Descriptor() ([]byte, []int) {
	return file_mail_message_proto_rawDescGZIP(), []int{0}
}

type MailTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MailTemplate) Reset() {
	*x = MailTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailTemplate) ProtoMessage() {}

func (x *MailTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_mail_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailTemplate.ProtoReflect.Descriptor instead.
func (*MailTemplate) Descriptor() ([]byte, []int) {
	return file_mail_message_proto_rawDescGZIP(), []int{1}
}

type UpdateMailTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMailTemplateRequest) Reset() {
	*x = UpdateMailTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMailTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMailTemplateRequest) ProtoMessage() {}

func (x *UpdateMailTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMailTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateMailTemplateRequest) Descriptor() ([]byte, []int) {
	return file_mail_message_proto_rawDescGZIP(), []int{2}
}

type DeleteMailTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMailTemplateRequest) Reset() {
	*x = DeleteMailTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMailTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMailTemplateRequest) ProtoMessage() {}

func (x *DeleteMailTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMailTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteMailTemplateRequest) Descriptor() ([]byte, []int) {
	return file_mail_message_proto_rawDescGZIP(), []int{3}
}

type DeleteMailTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMailTemplateResponse) Reset() {
	*x = DeleteMailTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMailTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMailTemplateResponse) ProtoMessage() {}

func (x *DeleteMailTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMailTemplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteMailTemplateResponse) Descriptor() ([]byte, []int) {
	return file_mail_message_proto_rawDescGZIP(), []int{4}
}

var File_mail_message_proto protoreflect.FileDescriptor

var file_mail_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1b,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x4d,
	0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mail_message_proto_rawDescOnce sync.Once
	file_mail_message_proto_rawDescData = file_mail_message_proto_rawDesc
)

func file_mail_message_proto_rawDescGZIP() []byte {
	file_mail_message_proto_rawDescOnce.Do(func() {
		file_mail_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_message_proto_rawDescData)
	})
	return file_mail_message_proto_rawDescData
}

var file_mail_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mail_message_proto_goTypes = []interface{}{
	(*CreateMailTemplateRequest)(nil),  // 0: sendmail.CreateMailTemplateRequest
	(*MailTemplate)(nil),               // 1: sendmail.MailTemplate
	(*UpdateMailTemplateRequest)(nil),  // 2: sendmail.UpdateMailTemplateRequest
	(*DeleteMailTemplateRequest)(nil),  // 3: sendmail.DeleteMailTemplateRequest
	(*DeleteMailTemplateResponse)(nil), // 4: sendmail.DeleteMailTemplateResponse
}
var file_mail_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mail_message_proto_init() }
func file_mail_message_proto_init() {
	if File_mail_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMailTemplateRequest); i {
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
		file_mail_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailTemplate); i {
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
		file_mail_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMailTemplateRequest); i {
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
		file_mail_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMailTemplateRequest); i {
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
		file_mail_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMailTemplateResponse); i {
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
			RawDescriptor: file_mail_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mail_message_proto_goTypes,
		DependencyIndexes: file_mail_message_proto_depIdxs,
		MessageInfos:      file_mail_message_proto_msgTypes,
	}.Build()
	File_mail_message_proto = out.File
	file_mail_message_proto_rawDesc = nil
	file_mail_message_proto_goTypes = nil
	file_mail_message_proto_depIdxs = nil
}
