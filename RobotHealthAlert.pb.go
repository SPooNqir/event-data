// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: RobotHealthAlert.proto

package event_data

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

type RHReplace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotEmail   string `protobuf:"bytes,1,opt,name=robotEmail,proto3" json:"robotEmail,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Timestamp    string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Status       string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Title        string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	RobotName    string `protobuf:"bytes,6,opt,name=robotName,proto3" json:"robotName,omitempty"`
}

func (x *RHReplace) Reset() {
	*x = RHReplace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RobotHealthAlert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RHReplace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RHReplace) ProtoMessage() {}

func (x *RHReplace) ProtoReflect() protoreflect.Message {
	mi := &file_RobotHealthAlert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RHReplace.ProtoReflect.Descriptor instead.
func (*RHReplace) Descriptor() ([]byte, []int) {
	return file_RobotHealthAlert_proto_rawDescGZIP(), []int{0}
}

func (x *RHReplace) GetRobotEmail() string {
	if x != nil {
		return x.RobotEmail
	}
	return ""
}

func (x *RHReplace) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *RHReplace) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *RHReplace) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RHReplace) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RHReplace) GetRobotName() string {
	if x != nil {
		return x.RobotName
	}
	return ""
}

type RobotHealthAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"timestamp"
	Publishtime string `protobuf:"bytes,1,opt,name=publishtime,proto3" json:"timestamp"`
	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,2,opt,name=DataType,proto3" json:"DataType"`
	// @inject_tag: json:"uid"
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid"`
	// @inject_tag: json:"tos"
	Tos []string `protobuf:"bytes,4,rep,name=tos,proto3" json:"tos"`
	// @inject_tag: json:"replaces"
	Replaces *RHReplace `protobuf:"bytes,5,opt,name=replaces,proto3" json:"replaces"`
	// @inject_tag: json:"from"
	From string `protobuf:"bytes,6,opt,name=from,proto3" json:"from"`
	// @inject_tag: json:"identity"
	Identity string `protobuf:"bytes,7,opt,name=identity,proto3" json:"identity"`
}

func (x *RobotHealthAlert) Reset() {
	*x = RobotHealthAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RobotHealthAlert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotHealthAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotHealthAlert) ProtoMessage() {}

func (x *RobotHealthAlert) ProtoReflect() protoreflect.Message {
	mi := &file_RobotHealthAlert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotHealthAlert.ProtoReflect.Descriptor instead.
func (*RobotHealthAlert) Descriptor() ([]byte, []int) {
	return file_RobotHealthAlert_proto_rawDescGZIP(), []int{1}
}

func (x *RobotHealthAlert) GetPublishtime() string {
	if x != nil {
		return x.Publishtime
	}
	return ""
}

func (x *RobotHealthAlert) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *RobotHealthAlert) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RobotHealthAlert) GetTos() []string {
	if x != nil {
		return x.Tos
	}
	return nil
}

func (x *RobotHealthAlert) GetReplaces() *RHReplace {
	if x != nil {
		return x.Replaces
	}
	return nil
}

func (x *RobotHealthAlert) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *RobotHealthAlert) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

var File_RobotHealthAlert_proto protoreflect.FileDescriptor

var file_RobotHealthAlert_proto_rawDesc = []byte{
	0x0a, 0x16, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x62, 0x71, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x62, 0x71, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x09, 0x52, 0x48, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x87, 0x02, 0x0a, 0x10, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xea, 0x3f, 0x0d, 0x08,
	0x01, 0x12, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x52, 0x0b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xea, 0x3f, 0x02,
	0x08, 0x01, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x6f, 0x73,
	0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x52, 0x48, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x3a, 0x15, 0xea, 0x3f, 0x12, 0x0a, 0x10, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x61, 0x76, 0x61, 0x79, 0x73,
	0x73, 0x69, 0x65, 0x72, 0x65, 0x2d, 0x73, 0x70, 0x6f, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RobotHealthAlert_proto_rawDescOnce sync.Once
	file_RobotHealthAlert_proto_rawDescData = file_RobotHealthAlert_proto_rawDesc
)

func file_RobotHealthAlert_proto_rawDescGZIP() []byte {
	file_RobotHealthAlert_proto_rawDescOnce.Do(func() {
		file_RobotHealthAlert_proto_rawDescData = protoimpl.X.CompressGZIP(file_RobotHealthAlert_proto_rawDescData)
	})
	return file_RobotHealthAlert_proto_rawDescData
}

var file_RobotHealthAlert_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RobotHealthAlert_proto_goTypes = []interface{}{
	(*RHReplace)(nil),        // 0: event_data.RHReplace
	(*RobotHealthAlert)(nil), // 1: event_data.RobotHealthAlert
}
var file_RobotHealthAlert_proto_depIdxs = []int32{
	0, // 0: event_data.RobotHealthAlert.replaces:type_name -> event_data.RHReplace
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RobotHealthAlert_proto_init() }
func file_RobotHealthAlert_proto_init() {
	if File_RobotHealthAlert_proto != nil {
		return
	}
	file_bq_table_proto_init()
	file_bq_field_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RobotHealthAlert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RHReplace); i {
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
		file_RobotHealthAlert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotHealthAlert); i {
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
			RawDescriptor: file_RobotHealthAlert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RobotHealthAlert_proto_goTypes,
		DependencyIndexes: file_RobotHealthAlert_proto_depIdxs,
		MessageInfos:      file_RobotHealthAlert_proto_msgTypes,
	}.Build()
	File_RobotHealthAlert_proto = out.File
	file_RobotHealthAlert_proto_rawDesc = nil
	file_RobotHealthAlert_proto_goTypes = nil
	file_RobotHealthAlert_proto_depIdxs = nil
}
