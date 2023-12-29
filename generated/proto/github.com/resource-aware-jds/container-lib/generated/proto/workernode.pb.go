// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: workernode.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubmitSuccessTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Results [][]byte `protobuf:"bytes,2,rep,name=Results,proto3" json:"Results,omitempty"`
}

func (x *SubmitSuccessTaskRequest) Reset() {
	*x = SubmitSuccessTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workernode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitSuccessTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitSuccessTaskRequest) ProtoMessage() {}

func (x *SubmitSuccessTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workernode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitSuccessTaskRequest.ProtoReflect.Descriptor instead.
func (*SubmitSuccessTaskRequest) Descriptor() ([]byte, []int) {
	return file_workernode_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitSuccessTaskRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SubmitSuccessTaskRequest) GetResults() [][]byte {
	if x != nil {
		return x.Results
	}
	return nil
}

type ReportTaskFailureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ErrorDetail string `protobuf:"bytes,2,opt,name=ErrorDetail,proto3" json:"ErrorDetail,omitempty"`
}

func (x *ReportTaskFailureRequest) Reset() {
	*x = ReportTaskFailureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workernode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportTaskFailureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportTaskFailureRequest) ProtoMessage() {}

func (x *ReportTaskFailureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workernode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportTaskFailureRequest.ProtoReflect.Descriptor instead.
func (*ReportTaskFailureRequest) Descriptor() ([]byte, []int) {
	return file_workernode_proto_rawDescGZIP(), []int{1}
}

func (x *ReportTaskFailureRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ReportTaskFailureRequest) GetErrorDetail() string {
	if x != nil {
		return x.ErrorDetail
	}
	return ""
}

var File_workernode_proto protoreflect.FileDescriptor

var file_workernode_proto_rawDesc = []byte{
	0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x18, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22,
	0x4c, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x32, 0x83, 0x02,
	0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x53, 0x0a,
	0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x24, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x77, 0x61, 0x72, 0x65,
	0x2d, 0x6a, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x6c,
	0x69, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workernode_proto_rawDescOnce sync.Once
	file_workernode_proto_rawDescData = file_workernode_proto_rawDesc
)

func file_workernode_proto_rawDescGZIP() []byte {
	file_workernode_proto_rawDescOnce.Do(func() {
		file_workernode_proto_rawDescData = protoimpl.X.CompressGZIP(file_workernode_proto_rawDescData)
	})
	return file_workernode_proto_rawDescData
}

var file_workernode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_workernode_proto_goTypes = []interface{}{
	(*SubmitSuccessTaskRequest)(nil), // 0: WorkerNode.SubmitSuccessTaskRequest
	(*ReportTaskFailureRequest)(nil), // 1: WorkerNode.ReportTaskFailureRequest
	(*emptypb.Empty)(nil),            // 2: google.protobuf.Empty
	(*Task)(nil),                     // 3: Common.Task
}
var file_workernode_proto_depIdxs = []int32{
	0, // 0: WorkerNode.WorkerNodeContainerReceiver.SubmitSuccessTask:input_type -> WorkerNode.SubmitSuccessTaskRequest
	1, // 1: WorkerNode.WorkerNodeContainerReceiver.ReportTaskFailure:input_type -> WorkerNode.ReportTaskFailureRequest
	2, // 2: WorkerNode.WorkerNodeContainerReceiver.GetTaskFromQueue:input_type -> google.protobuf.Empty
	2, // 3: WorkerNode.WorkerNodeContainerReceiver.SubmitSuccessTask:output_type -> google.protobuf.Empty
	2, // 4: WorkerNode.WorkerNodeContainerReceiver.ReportTaskFailure:output_type -> google.protobuf.Empty
	3, // 5: WorkerNode.WorkerNodeContainerReceiver.GetTaskFromQueue:output_type -> Common.Task
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_workernode_proto_init() }
func file_workernode_proto_init() {
	if File_workernode_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workernode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitSuccessTaskRequest); i {
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
		file_workernode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportTaskFailureRequest); i {
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
			RawDescriptor: file_workernode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workernode_proto_goTypes,
		DependencyIndexes: file_workernode_proto_depIdxs,
		MessageInfos:      file_workernode_proto_msgTypes,
	}.Build()
	File_workernode_proto = out.File
	file_workernode_proto_rawDesc = nil
	file_workernode_proto_goTypes = nil
	file_workernode_proto_depIdxs = nil
}
