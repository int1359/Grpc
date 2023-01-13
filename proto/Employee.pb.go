// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/Employee.proto

package grpc

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_proto_Employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Employee) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ReadEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *ReadEmployeeRequest) Reset() {
	*x = ReadEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEmployeeRequest) ProtoMessage() {}

func (x *ReadEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEmployeeRequest.ProtoReflect.Descriptor instead.
func (*ReadEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_proto_Employee_proto_rawDescGZIP(), []int{1}
}

func (x *ReadEmployeeRequest) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type ReadEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReadEmployeeResponse) Reset() {
	*x = ReadEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEmployeeResponse) ProtoMessage() {}

func (x *ReadEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEmployeeResponse.ProtoReflect.Descriptor instead.
func (*ReadEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_proto_Employee_proto_rawDescGZIP(), []int{2}
}

func (x *ReadEmployeeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_Employee_proto protoreflect.FileDescriptor

var file_proto_Employee_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x42, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x59, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_Employee_proto_rawDescOnce sync.Once
	file_proto_Employee_proto_rawDescData = file_proto_Employee_proto_rawDesc
)

func file_proto_Employee_proto_rawDescGZIP() []byte {
	file_proto_Employee_proto_rawDescOnce.Do(func() {
		file_proto_Employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Employee_proto_rawDescData)
	})
	return file_proto_Employee_proto_rawDescData
}

var file_proto_Employee_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_Employee_proto_goTypes = []interface{}{
	(*Employee)(nil),             // 0: proto.Employee
	(*ReadEmployeeRequest)(nil),  // 1: proto.ReadEmployeeRequest
	(*ReadEmployeeResponse)(nil), // 2: proto.ReadEmployeeResponse
}
var file_proto_Employee_proto_depIdxs = []int32{
	0, // 0: proto.ReadEmployeeRequest.employee:type_name -> proto.Employee
	1, // 1: proto.EmployeeService.GetEmployee:input_type -> proto.ReadEmployeeRequest
	2, // 2: proto.EmployeeService.GetEmployee:output_type -> proto.ReadEmployeeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_Employee_proto_init() }
func file_proto_Employee_proto_init() {
	if File_proto_Employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_proto_Employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEmployeeRequest); i {
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
		file_proto_Employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEmployeeResponse); i {
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
			RawDescriptor: file_proto_Employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Employee_proto_goTypes,
		DependencyIndexes: file_proto_Employee_proto_depIdxs,
		MessageInfos:      file_proto_Employee_proto_msgTypes,
	}.Build()
	File_proto_Employee_proto = out.File
	file_proto_Employee_proto_rawDesc = nil
	file_proto_Employee_proto_goTypes = nil
	file_proto_Employee_proto_depIdxs = nil
}
