// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: board/board.proto

package board

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

// The request message containing the user's name.
type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stone   string `protobuf:"bytes,1,opt,name=stone,proto3" json:"stone,omitempty"`
	X       int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Squares string `protobuf:"bytes,4,opt,name=squares,proto3" json:"squares,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_board_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_board_board_proto_rawDescGZIP(), []int{0}
}

func (x *Board) GetStone() string {
	if x != nil {
		return x.Stone
	}
	return ""
}

func (x *Board) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Board) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Board) GetSquares() string {
	if x != nil {
		return x.Squares
	}
	return ""
}

// The response message containing the greetings
type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Squares string `protobuf:"bytes,1,opt,name=squares,proto3" json:"squares,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_board_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_board_board_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetSquares() string {
	if x != nil {
		return x.Squares
	}
	return ""
}

var File_board_board_proto protoreflect.FileDescriptor

var file_board_board_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x53, 0x0a, 0x05, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73, 0x22,
	0x1f, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73,
	0x32, 0x58, 0x0a, 0x08, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x41, 0x70, 0x69, 0x12, 0x25, 0x0a, 0x07,
	0x50, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x0a, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x0c,
	0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x0a, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_board_board_proto_rawDescOnce sync.Once
	file_board_board_proto_rawDescData = file_board_board_proto_rawDesc
)

func file_board_board_proto_rawDescGZIP() []byte {
	file_board_board_proto_rawDescOnce.Do(func() {
		file_board_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_board_proto_rawDescData)
	})
	return file_board_board_proto_rawDescData
}

var file_board_board_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_board_board_proto_goTypes = []interface{}{
	(*Board)(nil), // 0: board.Board
	(*Res)(nil),   // 1: board.Res
}
var file_board_board_proto_depIdxs = []int32{
	0, // 0: board.BoardApi.Putable:input_type -> board.Board
	0, // 1: board.BoardApi.Reverse:input_type -> board.Board
	1, // 2: board.BoardApi.Putable:output_type -> board.Res
	1, // 3: board.BoardApi.Reverse:output_type -> board.Res
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_board_board_proto_init() }
func file_board_board_proto_init() {
	if File_board_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_board_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_board_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_board_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_board_board_proto_goTypes,
		DependencyIndexes: file_board_board_proto_depIdxs,
		MessageInfos:      file_board_board_proto_msgTypes,
	}.Build()
	File_board_board_proto = out.File
	file_board_board_proto_rawDesc = nil
	file_board_board_proto_goTypes = nil
	file_board_board_proto_depIdxs = nil
}
