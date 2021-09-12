// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: discount/proto/discount.proto

package discount

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

// productID used to represent a product. Ilustrative only.
type GetDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int32 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *GetDiscountRequest) Reset() {
	*x = GetDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiscountRequest) ProtoMessage() {}

func (x *GetDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiscountRequest.ProtoReflect.Descriptor instead.
func (*GetDiscountRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_discount_proto_rawDescGZIP(), []int{0}
}

func (x *GetDiscountRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

// The discount percentage is a fixed value.
type GetDiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *GetDiscountResponse) Reset() {
	*x = GetDiscountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiscountResponse) ProtoMessage() {}

func (x *GetDiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_discount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiscountResponse.ProtoReflect.Descriptor instead.
func (*GetDiscountResponse) Descriptor() ([]byte, []int) {
	return file_discount_proto_discount_proto_rawDescGZIP(), []int{1}
}

func (x *GetDiscountResponse) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

var File_discount_proto_discount_proto protoreflect.FileDescriptor

var file_discount_proto_discount_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x35, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x32, 0x58, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_proto_discount_proto_rawDescOnce sync.Once
	file_discount_proto_discount_proto_rawDescData = file_discount_proto_discount_proto_rawDesc
)

func file_discount_proto_discount_proto_rawDescGZIP() []byte {
	file_discount_proto_discount_proto_rawDescOnce.Do(func() {
		file_discount_proto_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_proto_discount_proto_rawDescData)
	})
	return file_discount_proto_discount_proto_rawDescData
}

var file_discount_proto_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discount_proto_discount_proto_goTypes = []interface{}{
	(*GetDiscountRequest)(nil),  // 0: discount.GetDiscountRequest
	(*GetDiscountResponse)(nil), // 1: discount.GetDiscountResponse
}
var file_discount_proto_discount_proto_depIdxs = []int32{
	0, // 0: discount.Discount.GetDiscount:input_type -> discount.GetDiscountRequest
	1, // 1: discount.Discount.GetDiscount:output_type -> discount.GetDiscountResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discount_proto_discount_proto_init() }
func file_discount_proto_discount_proto_init() {
	if File_discount_proto_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discount_proto_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiscountRequest); i {
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
		file_discount_proto_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiscountResponse); i {
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
			RawDescriptor: file_discount_proto_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discount_proto_discount_proto_goTypes,
		DependencyIndexes: file_discount_proto_discount_proto_depIdxs,
		MessageInfos:      file_discount_proto_discount_proto_msgTypes,
	}.Build()
	File_discount_proto_discount_proto = out.File
	file_discount_proto_discount_proto_rawDesc = nil
	file_discount_proto_discount_proto_goTypes = nil
	file_discount_proto_discount_proto_depIdxs = nil
}
