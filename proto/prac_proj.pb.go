// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/prac_proj.proto

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

type IVStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity float32 `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *IVStruct) Reset() {
	*x = IVStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prac_proj_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IVStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IVStruct) ProtoMessage() {}

func (x *IVStruct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prac_proj_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IVStruct.ProtoReflect.Descriptor instead.
func (*IVStruct) Descriptor() ([]byte, []int) {
	return file_proto_prac_proj_proto_rawDescGZIP(), []int{0}
}

func (x *IVStruct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IVStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IVStruct) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Response_IV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctd string `protobuf:"bytes,1,opt,name=ctd,proto3" json:"ctd,omitempty"`
}

func (x *Response_IV) Reset() {
	*x = Response_IV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prac_proj_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_IV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_IV) ProtoMessage() {}

func (x *Response_IV) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prac_proj_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_IV.ProtoReflect.Descriptor instead.
func (*Response_IV) Descriptor() ([]byte, []int) {
	return file_proto_prac_proj_proto_rawDescGZIP(), []int{1}
}

func (x *Response_IV) GetCtd() string {
	if x != nil {
		return x.Ctd
	}
	return ""
}

var File_proto_prac_proj_proto protoreflect.FileDescriptor

var file_proto_prac_proj_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x61, 0x63, 0x5f, 0x70, 0x72, 0x6f,
	0x6a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x61, 0x63, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x22, 0x4b, 0x0a, 0x09, 0x49, 0x56, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x49, 0x56, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x74, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x74, 0x64,
	0x32, 0x47, 0x0a, 0x0a, 0x49, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x49, 0x56, 0x12, 0x14, 0x2e, 0x70, 0x72,
	0x61, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x49, 0x56, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x49, 0x56, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x64, 0x68, 0x61, 0x72, 0x64,
	0x68, 0x6b, 0x32, 0x34, 0x2f, 0x70, 0x72, 0x61, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_prac_proj_proto_rawDescOnce sync.Once
	file_proto_prac_proj_proto_rawDescData = file_proto_prac_proj_proto_rawDesc
)

func file_proto_prac_proj_proto_rawDescGZIP() []byte {
	file_proto_prac_proj_proto_rawDescOnce.Do(func() {
		file_proto_prac_proj_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prac_proj_proto_rawDescData)
	})
	return file_proto_prac_proj_proto_rawDescData
}

var file_proto_prac_proj_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_prac_proj_proto_goTypes = []interface{}{
	(*IVStruct)(nil),    // 0: prac_proj.IV_struct
	(*Response_IV)(nil), // 1: prac_proj.response_IV
}
var file_proto_prac_proj_proto_depIdxs = []int32{
	0, // 0: prac_proj.Iv_service.create_IV:input_type -> prac_proj.IV_struct
	1, // 1: prac_proj.Iv_service.create_IV:output_type -> prac_proj.response_IV
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_prac_proj_proto_init() }
func file_proto_prac_proj_proto_init() {
	if File_proto_prac_proj_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prac_proj_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IVStruct); i {
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
		file_proto_prac_proj_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_IV); i {
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
			RawDescriptor: file_proto_prac_proj_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_prac_proj_proto_goTypes,
		DependencyIndexes: file_proto_prac_proj_proto_depIdxs,
		MessageInfos:      file_proto_prac_proj_proto_msgTypes,
	}.Build()
	File_proto_prac_proj_proto = out.File
	file_proto_prac_proj_proto_rawDesc = nil
	file_proto_prac_proj_proto_goTypes = nil
	file_proto_prac_proj_proto_depIdxs = nil
}
