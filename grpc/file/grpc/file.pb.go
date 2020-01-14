// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package file

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FileByte struct {
	Byte                 []byte   `protobuf:"bytes,1,opt,name=byte,proto3" json:"byte,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileByte) Reset()         { *m = FileByte{} }
func (m *FileByte) String() string { return proto.CompactTextString(m) }
func (*FileByte) ProtoMessage()    {}
func (*FileByte) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *FileByte) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileByte.Unmarshal(m, b)
}
func (m *FileByte) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileByte.Marshal(b, m, deterministic)
}
func (m *FileByte) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileByte.Merge(m, src)
}
func (m *FileByte) XXX_Size() int {
	return xxx_messageInfo_FileByte.Size(m)
}
func (m *FileByte) XXX_DiscardUnknown() {
	xxx_messageInfo_FileByte.DiscardUnknown(m)
}

var xxx_messageInfo_FileByte proto.InternalMessageInfo

func (m *FileByte) GetByte() []byte {
	if m != nil {
		return m.Byte
	}
	return nil
}

type FileName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileName) Reset()         { *m = FileName{} }
func (m *FileName) String() string { return proto.CompactTextString(m) }
func (*FileName) ProtoMessage()    {}
func (*FileName) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *FileName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileName.Unmarshal(m, b)
}
func (m *FileName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileName.Marshal(b, m, deterministic)
}
func (m *FileName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileName.Merge(m, src)
}
func (m *FileName) XXX_Size() int {
	return xxx_messageInfo_FileName.Size(m)
}
func (m *FileName) XXX_DiscardUnknown() {
	xxx_messageInfo_FileName.DiscardUnknown(m)
}

var xxx_messageInfo_FileName proto.InternalMessageInfo

func (m *FileName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FileOk struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileOk) Reset()         { *m = FileOk{} }
func (m *FileOk) String() string { return proto.CompactTextString(m) }
func (*FileOk) ProtoMessage()    {}
func (*FileOk) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}

func (m *FileOk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileOk.Unmarshal(m, b)
}
func (m *FileOk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileOk.Marshal(b, m, deterministic)
}
func (m *FileOk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileOk.Merge(m, src)
}
func (m *FileOk) XXX_Size() int {
	return xxx_messageInfo_FileOk.Size(m)
}
func (m *FileOk) XXX_DiscardUnknown() {
	xxx_messageInfo_FileOk.DiscardUnknown(m)
}

var xxx_messageInfo_FileOk proto.InternalMessageInfo

func (m *FileOk) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*FileByte)(nil), "file.FileByte")
	proto.RegisterType((*FileName)(nil), "file.FileName")
	proto.RegisterType((*FileOk)(nil), "file.fileOk")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb8, 0x38, 0xdc, 0x32,
	0x73, 0x52, 0x9d, 0x2a, 0x4b, 0x52, 0x85, 0x84, 0xb8, 0x58, 0x92, 0x2a, 0x4b, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x98, 0xbc, 0x5f, 0x62, 0x2e, 0x58, 0x3e, 0x2f, 0x31,
	0x17, 0x22, 0xcf, 0x19, 0x04, 0x66, 0x2b, 0x49, 0x70, 0xb1, 0x81, 0xcc, 0xf1, 0xcf, 0x16, 0xe2,
	0xe3, 0x62, 0xca, 0xcf, 0x06, 0xcb, 0x71, 0x04, 0x31, 0xe5, 0x67, 0x1b, 0x4d, 0x60, 0xe4, 0x62,
	0x01, 0x69, 0x15, 0xd2, 0xe0, 0x62, 0xf3, 0x2c, 0x06, 0xb3, 0xf8, 0xf4, 0xc0, 0xf6, 0xc3, 0x0c,
	0x94, 0xe2, 0x81, 0xf0, 0x21, 0x06, 0x28, 0x31, 0x08, 0xe9, 0x71, 0x71, 0x04, 0xa5, 0x26, 0xa6,
	0x60, 0x55, 0x8b, 0xc4, 0x07, 0x39, 0x56, 0x89, 0xc1, 0x80, 0x51, 0x48, 0x8f, 0x8b, 0x2b, 0xbc,
	0x28, 0xb3, 0x24, 0xd5, 0x2d, 0x33, 0x35, 0x27, 0x45, 0x08, 0x4d, 0x05, 0xba, 0xe9, 0x1a, 0x8c,
	0x49, 0x6c, 0x60, 0x9f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xb4, 0x0a, 0xcb, 0x07,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileClient interface {
	IsFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileOk, error)
	ReadFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (File_ReadFileClient, error)
	WriteField(ctx context.Context, opts ...grpc.CallOption) (File_WriteFieldClient, error)
}

type fileClient struct {
	cc *grpc.ClientConn
}

func NewFileClient(cc *grpc.ClientConn) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) IsFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileOk, error) {
	out := new(FileOk)
	err := c.cc.Invoke(ctx, "/file.File/IsFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) ReadFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (File_ReadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_File_serviceDesc.Streams[0], "/file.File/ReadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileReadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type File_ReadFileClient interface {
	Recv() (*FileByte, error)
	grpc.ClientStream
}

type fileReadFileClient struct {
	grpc.ClientStream
}

func (x *fileReadFileClient) Recv() (*FileByte, error) {
	m := new(FileByte)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileClient) WriteField(ctx context.Context, opts ...grpc.CallOption) (File_WriteFieldClient, error) {
	stream, err := c.cc.NewStream(ctx, &_File_serviceDesc.Streams[1], "/file.File/WriteField", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileWriteFieldClient{stream}
	return x, nil
}

type File_WriteFieldClient interface {
	Send(*FileByte) error
	CloseAndRecv() (*FileOk, error)
	grpc.ClientStream
}

type fileWriteFieldClient struct {
	grpc.ClientStream
}

func (x *fileWriteFieldClient) Send(m *FileByte) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileWriteFieldClient) CloseAndRecv() (*FileOk, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileOk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServer is the server API for File service.
type FileServer interface {
	IsFile(context.Context, *FileName) (*FileOk, error)
	ReadFile(*FileName, File_ReadFileServer) error
	WriteField(File_WriteFieldServer) error
}

// UnimplementedFileServer can be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (*UnimplementedFileServer) IsFile(ctx context.Context, req *FileName) (*FileOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFile not implemented")
}
func (*UnimplementedFileServer) ReadFile(req *FileName, srv File_ReadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadFile not implemented")
}
func (*UnimplementedFileServer) WriteField(srv File_WriteFieldServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteField not implemented")
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_IsFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).IsFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.File/IsFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).IsFile(ctx, req.(*FileName))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_ReadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServer).ReadFile(m, &fileReadFileServer{stream})
}

type File_ReadFileServer interface {
	Send(*FileByte) error
	grpc.ServerStream
}

type fileReadFileServer struct {
	grpc.ServerStream
}

func (x *fileReadFileServer) Send(m *FileByte) error {
	return x.ServerStream.SendMsg(m)
}

func _File_WriteField_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServer).WriteField(&fileWriteFieldServer{stream})
}

type File_WriteFieldServer interface {
	SendAndClose(*FileOk) error
	Recv() (*FileByte, error)
	grpc.ServerStream
}

type fileWriteFieldServer struct {
	grpc.ServerStream
}

func (x *fileWriteFieldServer) SendAndClose(m *FileOk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileWriteFieldServer) Recv() (*FileByte, error) {
	m := new(FileByte)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsFile",
			Handler:    _File_IsFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadFile",
			Handler:       _File_ReadFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WriteField",
			Handler:       _File_WriteField_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file.proto",
}
