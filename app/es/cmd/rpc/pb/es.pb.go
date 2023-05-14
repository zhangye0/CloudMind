// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: es.proto

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

type SearchForReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchForReq) Reset() {
	*x = SearchForReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchForReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchForReq) ProtoMessage() {}

func (x *SearchForReq) ProtoReflect() protoreflect.Message {
	mi := &file_es_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchForReq.ProtoReflect.Descriptor instead.
func (*SearchForReq) Descriptor() ([]byte, []int) {
	return file_es_proto_rawDescGZIP(), []int{0}
}

type SearchForResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchForResp) Reset() {
	*x = SearchForResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchForResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchForResp) ProtoMessage() {}

func (x *SearchForResp) ProtoReflect() protoreflect.Message {
	mi := &file_es_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchForResp.ProtoReflect.Descriptor instead.
func (*SearchForResp) Descriptor() ([]byte, []int) {
	return file_es_proto_rawDescGZIP(), []int{1}
}

type AddTextReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index string `protobuf:"bytes,1,opt,name=Index,proto3" json:"Index"`
	Text  string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text"`
}

func (x *AddTextReq) Reset() {
	*x = AddTextReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTextReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTextReq) ProtoMessage() {}

func (x *AddTextReq) ProtoReflect() protoreflect.Message {
	mi := &file_es_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTextReq.ProtoReflect.Descriptor instead.
func (*AddTextReq) Descriptor() ([]byte, []int) {
	return file_es_proto_rawDescGZIP(), []int{2}
}

func (x *AddTextReq) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *AddTextReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AddTextResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error"`
}

func (x *AddTextResp) Reset() {
	*x = AddTextResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTextResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTextResp) ProtoMessage() {}

func (x *AddTextResp) ProtoReflect() protoreflect.Message {
	mi := &file_es_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTextResp.ProtoReflect.Descriptor instead.
func (*AddTextResp) Descriptor() ([]byte, []int) {
	return file_es_proto_rawDescGZIP(), []int{3}
}

func (x *AddTextResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_es_proto protoreflect.FileDescriptor

var file_es_proto_rawDesc = []byte{
	0x0a, 0x08, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x0e,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x22, 0x0f,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x36, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x62, 0x0a, 0x02,
	0x65, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x12,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_es_proto_rawDescOnce sync.Once
	file_es_proto_rawDescData = file_es_proto_rawDesc
)

func file_es_proto_rawDescGZIP() []byte {
	file_es_proto_rawDescOnce.Do(func() {
		file_es_proto_rawDescData = protoimpl.X.CompressGZIP(file_es_proto_rawDescData)
	})
	return file_es_proto_rawDescData
}

var file_es_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_es_proto_goTypes = []interface{}{
	(*SearchForReq)(nil),  // 0: pb.SearchForReq
	(*SearchForResp)(nil), // 1: pb.SearchForResp
	(*AddTextReq)(nil),    // 2: pb.AddTextReq
	(*AddTextResp)(nil),   // 3: pb.AddTextResp
}
var file_es_proto_depIdxs = []int32{
	0, // 0: pb.es.SearchFor:input_type -> pb.SearchForReq
	2, // 1: pb.es.AddText:input_type -> pb.AddTextReq
	1, // 2: pb.es.SearchFor:output_type -> pb.SearchForResp
	3, // 3: pb.es.AddText:output_type -> pb.AddTextResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_es_proto_init() }
func file_es_proto_init() {
	if File_es_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_es_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchForReq); i {
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
		file_es_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchForResp); i {
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
		file_es_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTextReq); i {
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
		file_es_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTextResp); i {
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
			RawDescriptor: file_es_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_es_proto_goTypes,
		DependencyIndexes: file_es_proto_depIdxs,
		MessageInfos:      file_es_proto_msgTypes,
	}.Build()
	File_es_proto = out.File
	file_es_proto_rawDesc = nil
	file_es_proto_goTypes = nil
	file_es_proto_depIdxs = nil
}