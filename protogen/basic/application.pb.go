// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.2
// source: proto/basic/application.proto

package basic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	job "my-protobuf/protogen/job"
	software "my-protobuf/protogen/software"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *job.Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	JobId       uint32           `protobuf:"varint,2,opt,name=job_id,proto3" json:"job_id,omitempty"`
}

func (x *JobApplication) Reset() {
	*x = JobApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobApplication) ProtoMessage() {}

func (x *JobApplication) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobApplication.ProtoReflect.Descriptor instead.
func (*JobApplication) Descriptor() ([]byte, []int) {
	return file_proto_basic_application_proto_rawDescGZIP(), []int{0}
}

func (x *JobApplication) GetApplication() *job.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

func (x *JobApplication) GetJobId() uint32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

type SoftwareApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *software.Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	SoftwareId  uint32                `protobuf:"varint,2,opt,name=software_id,proto3" json:"software_id,omitempty"`
}

func (x *SoftwareApplication) Reset() {
	*x = SoftwareApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoftwareApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoftwareApplication) ProtoMessage() {}

func (x *SoftwareApplication) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoftwareApplication.ProtoReflect.Descriptor instead.
func (*SoftwareApplication) Descriptor() ([]byte, []int) {
	return file_proto_basic_application_proto_rawDescGZIP(), []int{1}
}

func (x *SoftwareApplication) GetApplication() *software.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

func (x *SoftwareApplication) GetSoftwareId() uint32 {
	if x != nil {
		return x.SoftwareId
	}
	return 0
}

var File_proto_basic_application_proto protoreflect.FileDescriptor

var file_proto_basic_application_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x22, 0x7a, 0x0a, 0x13, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61,
	0x72, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f,
	0x69, 0x64, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_application_proto_rawDescOnce sync.Once
	file_proto_basic_application_proto_rawDescData = file_proto_basic_application_proto_rawDesc
)

func file_proto_basic_application_proto_rawDescGZIP() []byte {
	file_proto_basic_application_proto_rawDescOnce.Do(func() {
		file_proto_basic_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_application_proto_rawDescData)
	})
	return file_proto_basic_application_proto_rawDescData
}

var file_proto_basic_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_basic_application_proto_goTypes = []interface{}{
	(*JobApplication)(nil),       // 0: JobApplication
	(*SoftwareApplication)(nil),  // 1: SoftwareApplication
	(*job.Application)(nil),      // 2: job.proto.pkg.Application
	(*software.Application)(nil), // 3: software.proto.pkg.Application
}
var file_proto_basic_application_proto_depIdxs = []int32{
	2, // 0: JobApplication.application:type_name -> job.proto.pkg.Application
	3, // 1: SoftwareApplication.application:type_name -> software.proto.pkg.Application
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_basic_application_proto_init() }
func file_proto_basic_application_proto_init() {
	if File_proto_basic_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_basic_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobApplication); i {
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
		file_proto_basic_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoftwareApplication); i {
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
			RawDescriptor: file_proto_basic_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_application_proto_goTypes,
		DependencyIndexes: file_proto_basic_application_proto_depIdxs,
		MessageInfos:      file_proto_basic_application_proto_msgTypes,
	}.Build()
	File_proto_basic_application_proto = out.File
	file_proto_basic_application_proto_rawDesc = nil
	file_proto_basic_application_proto_goTypes = nil
	file_proto_basic_application_proto_depIdxs = nil
}