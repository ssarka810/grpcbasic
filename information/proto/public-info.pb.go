// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: information/proto/public-info.proto

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

type BasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Location  string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *BasicInfo) Reset() {
	*x = BasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_information_proto_public_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicInfo) ProtoMessage() {}

func (x *BasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_information_proto_public_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicInfo.ProtoReflect.Descriptor instead.
func (*BasicInfo) Descriptor() ([]byte, []int) {
	return file_information_proto_public_info_proto_rawDescGZIP(), []int{0}
}

func (x *BasicInfo) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *BasicInfo) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *BasicInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type DetailedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DetailedInfo) Reset() {
	*x = DetailedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_information_proto_public_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailedInfo) ProtoMessage() {}

func (x *DetailedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_information_proto_public_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailedInfo.ProtoReflect.Descriptor instead.
func (*DetailedInfo) Descriptor() ([]byte, []int) {
	return file_information_proto_public_info_proto_rawDescGZIP(), []int{1}
}

func (x *DetailedInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DetailedInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *DetailedInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_information_proto_public_info_proto protoreflect.FileDescriptor

var file_information_proto_public_info_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x61, 0x0a, 0x09,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x56, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa3, 0x02, 0x0a, 0x18, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x16, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x16,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x28, 0x01, 0x12,
	0x4c, 0x0a, 0x1d, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x38, 0x31,
	0x30, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_information_proto_public_info_proto_rawDescOnce sync.Once
	file_information_proto_public_info_proto_rawDescData = file_information_proto_public_info_proto_rawDesc
)

func file_information_proto_public_info_proto_rawDescGZIP() []byte {
	file_information_proto_public_info_proto_rawDescOnce.Do(func() {
		file_information_proto_public_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_information_proto_public_info_proto_rawDescData)
	})
	return file_information_proto_public_info_proto_rawDescData
}

var file_information_proto_public_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_information_proto_public_info_proto_goTypes = []interface{}{
	(*BasicInfo)(nil),    // 0: greet.BasicInfo
	(*DetailedInfo)(nil), // 1: greet.DetailedInfo
}
var file_information_proto_public_info_proto_depIdxs = []int32{
	0, // 0: greet.PublicInformationService.GetDetails:input_type -> greet.BasicInfo
	0, // 1: greet.PublicInformationService.ServerStreamGetDetails:input_type -> greet.BasicInfo
	0, // 2: greet.PublicInformationService.ClientStreamGetDetails:input_type -> greet.BasicInfo
	0, // 3: greet.PublicInformationService.BiDirectionalStreamGetDetails:input_type -> greet.BasicInfo
	1, // 4: greet.PublicInformationService.GetDetails:output_type -> greet.DetailedInfo
	1, // 5: greet.PublicInformationService.ServerStreamGetDetails:output_type -> greet.DetailedInfo
	1, // 6: greet.PublicInformationService.ClientStreamGetDetails:output_type -> greet.DetailedInfo
	1, // 7: greet.PublicInformationService.BiDirectionalStreamGetDetails:output_type -> greet.DetailedInfo
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_information_proto_public_info_proto_init() }
func file_information_proto_public_info_proto_init() {
	if File_information_proto_public_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_information_proto_public_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicInfo); i {
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
		file_information_proto_public_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailedInfo); i {
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
			RawDescriptor: file_information_proto_public_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_information_proto_public_info_proto_goTypes,
		DependencyIndexes: file_information_proto_public_info_proto_depIdxs,
		MessageInfos:      file_information_proto_public_info_proto_msgTypes,
	}.Build()
	File_information_proto_public_info_proto = out.File
	file_information_proto_public_info_proto_rawDesc = nil
	file_information_proto_public_info_proto_goTypes = nil
	file_information_proto_public_info_proto_depIdxs = nil
}
