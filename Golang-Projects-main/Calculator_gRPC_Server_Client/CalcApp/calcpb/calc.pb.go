// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: calcpb/calc.proto

package calcpb

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

type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"zigzag32,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"zigzag32,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *CalcRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type CalcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"zigzag32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalcResponse) Reset() {
	*x = CalcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcResponse) ProtoMessage() {}

func (x *CalcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcResponse.ProtoReflect.Descriptor instead.
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return file_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calcpb_calc_proto protoreflect.FileDescriptor

var file_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x35, 0x0a, 0x0b, 0x43, 0x61, 0x6c,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32,
	0x22, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xe0, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6c,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x6d, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e,
	0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x11,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calcpb_calc_proto_rawDescOnce sync.Once
	file_calcpb_calc_proto_rawDescData = file_calcpb_calc_proto_rawDesc
)

func file_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calcpb_calc_proto_rawDescData)
	})
	return file_calcpb_calc_proto_rawDescData
}

var file_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calcpb_calc_proto_goTypes = []interface{}{
	(*CalcRequest)(nil),  // 0: calc.CalcRequest
	(*CalcResponse)(nil), // 1: calc.CalcResponse
}
var file_calcpb_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalcService.GetSum:input_type -> calc.CalcRequest
	0, // 1: calc.CalcService.GetMultiply:input_type -> calc.CalcRequest
	0, // 2: calc.CalcService.GetDivision:input_type -> calc.CalcRequest
	0, // 3: calc.CalcService.GetSubtract:input_type -> calc.CalcRequest
	1, // 4: calc.CalcService.GetSum:output_type -> calc.CalcResponse
	1, // 5: calc.CalcService.GetMultiply:output_type -> calc.CalcResponse
	1, // 6: calc.CalcService.GetDivision:output_type -> calc.CalcResponse
	1, // 7: calc.CalcService.GetSubtract:output_type -> calc.CalcResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calcpb_calc_proto_init() }
func file_calcpb_calc_proto_init() {
	if File_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
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
		file_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcResponse); i {
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
			RawDescriptor: file_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calcpb_calc_proto = out.File
	file_calcpb_calc_proto_rawDesc = nil
	file_calcpb_calc_proto_goTypes = nil
	file_calcpb_calc_proto_depIdxs = nil
}
