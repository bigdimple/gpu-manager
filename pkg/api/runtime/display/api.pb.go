// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/runtime/display/api.proto

/*
Package display is a generated protocol buffer package.

It is generated from these files:
	pkg/api/runtime/display/api.proto

It has these top-level messages:
	GraphResponse
	UsageResponse
	ContainerStat
	Devices
	DeviceInfo
	VersionResponse
	Spec
*/
package display

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GraphResponse struct {
	Graph string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
}

func (m *GraphResponse) Reset()                    { *m = GraphResponse{} }
func (m *GraphResponse) String() string            { return proto.CompactTextString(m) }
func (*GraphResponse) ProtoMessage()               {}
func (*GraphResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GraphResponse) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

type UsageResponse struct {
	Usage map[string]*ContainerStat `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UsageResponse) Reset()                    { *m = UsageResponse{} }
func (m *UsageResponse) String() string            { return proto.CompactTextString(m) }
func (*UsageResponse) ProtoMessage()               {}
func (*UsageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UsageResponse) GetUsage() map[string]*ContainerStat {
	if m != nil {
		return m.Usage
	}
	return nil
}

type ContainerStat struct {
	Stat    map[string]*Devices `protobuf:"bytes,1,rep,name=stat" json:"stat,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Project string              `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	User    string              `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Cluster string              `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
	Spec    map[string]*Spec    `protobuf:"bytes,5,rep,name=spec" json:"spec,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ContainerStat) Reset()                    { *m = ContainerStat{} }
func (m *ContainerStat) String() string            { return proto.CompactTextString(m) }
func (*ContainerStat) ProtoMessage()               {}
func (*ContainerStat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContainerStat) GetStat() map[string]*Devices {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *ContainerStat) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ContainerStat) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ContainerStat) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *ContainerStat) GetSpec() map[string]*Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type Devices struct {
	Dev []*DeviceInfo `protobuf:"bytes,1,rep,name=dev" json:"dev,omitempty"`
}

func (m *Devices) Reset()                    { *m = Devices{} }
func (m *Devices) String() string            { return proto.CompactTextString(m) }
func (*Devices) ProtoMessage()               {}
func (*Devices) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Devices) GetDev() []*DeviceInfo {
	if m != nil {
		return m.Dev
	}
	return nil
}

type DeviceInfo struct {
	Id      string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CardIdx string  `protobuf:"bytes,2,opt,name=card_idx,json=cardIdx" json:"card_idx,omitempty"`
	Gpu     float32 `protobuf:"fixed32,10,opt,name=gpu" json:"gpu,omitempty"`
	Mem     float32 `protobuf:"fixed32,11,opt,name=mem" json:"mem,omitempty"`
	Pids    []int32 `protobuf:"varint,12,rep,packed,name=pids" json:"pids,omitempty"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeviceInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceInfo) GetCardIdx() string {
	if m != nil {
		return m.CardIdx
	}
	return ""
}

func (m *DeviceInfo) GetGpu() float32 {
	if m != nil {
		return m.Gpu
	}
	return 0
}

func (m *DeviceInfo) GetMem() float32 {
	if m != nil {
		return m.Mem
	}
	return 0
}

func (m *DeviceInfo) GetPids() []int32 {
	if m != nil {
		return m.Pids
	}
	return nil
}

type VersionResponse struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Spec struct {
	Gpu float32 `protobuf:"fixed32,1,opt,name=gpu" json:"gpu,omitempty"`
	Mem float32 `protobuf:"fixed32,2,opt,name=mem" json:"mem,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Spec) GetGpu() float32 {
	if m != nil {
		return m.Gpu
	}
	return 0
}

func (m *Spec) GetMem() float32 {
	if m != nil {
		return m.Mem
	}
	return 0
}

func init() {
	proto.RegisterType((*GraphResponse)(nil), "display.GraphResponse")
	proto.RegisterType((*UsageResponse)(nil), "display.UsageResponse")
	proto.RegisterType((*ContainerStat)(nil), "display.ContainerStat")
	proto.RegisterType((*Devices)(nil), "display.Devices")
	proto.RegisterType((*DeviceInfo)(nil), "display.DeviceInfo")
	proto.RegisterType((*VersionResponse)(nil), "display.VersionResponse")
	proto.RegisterType((*Spec)(nil), "display.Spec")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GPUDisplay service

type GPUDisplayClient interface {
	// PrintGraph returns the text graph of allocator state
	PrintGraph(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*GraphResponse, error)
	// GPU usages
	PrintUsages(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*UsageResponse, error)
	// Version
	Version(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type gPUDisplayClient struct {
	cc *grpc.ClientConn
}

func NewGPUDisplayClient(cc *grpc.ClientConn) GPUDisplayClient {
	return &gPUDisplayClient{cc}
}

func (c *gPUDisplayClient) PrintGraph(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*GraphResponse, error) {
	out := new(GraphResponse)
	err := grpc.Invoke(ctx, "/display.GPUDisplay/PrintGraph", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUDisplayClient) PrintUsages(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*UsageResponse, error) {
	out := new(UsageResponse)
	err := grpc.Invoke(ctx, "/display.GPUDisplay/PrintUsages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPUDisplayClient) Version(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/display.GPUDisplay/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GPUDisplay service

type GPUDisplayServer interface {
	// PrintGraph returns the text graph of allocator state
	PrintGraph(context.Context, *google_protobuf1.Empty) (*GraphResponse, error)
	// GPU usages
	PrintUsages(context.Context, *google_protobuf1.Empty) (*UsageResponse, error)
	// Version
	Version(context.Context, *google_protobuf1.Empty) (*VersionResponse, error)
}

func RegisterGPUDisplayServer(s *grpc.Server, srv GPUDisplayServer) {
	s.RegisterService(&_GPUDisplay_serviceDesc, srv)
}

func _GPUDisplay_PrintGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDisplayServer).PrintGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/display.GPUDisplay/PrintGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDisplayServer).PrintGraph(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUDisplay_PrintUsages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDisplayServer).PrintUsages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/display.GPUDisplay/PrintUsages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDisplayServer).PrintUsages(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPUDisplay_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPUDisplayServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/display.GPUDisplay/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPUDisplayServer).Version(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GPUDisplay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "display.GPUDisplay",
	HandlerType: (*GPUDisplayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintGraph",
			Handler:    _GPUDisplay_PrintGraph_Handler,
		},
		{
			MethodName: "PrintUsages",
			Handler:    _GPUDisplay_PrintUsages_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _GPUDisplay_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/runtime/display/api.proto",
}

func init() { proto.RegisterFile("pkg/api/runtime/display/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0x94, 0x9d, 0xa4, 0x69, 0xbe, 0x90, 0x52, 0x2d, 0xa8, 0x5a, 0x0c, 0x07, 0xd7, 0xa8, 0x28,
	0x02, 0x64, 0xa3, 0x82, 0x04, 0xe2, 0x4a, 0x4b, 0x55, 0x89, 0x4a, 0x91, 0x51, 0xb9, 0x22, 0xd7,
	0xde, 0x9a, 0xa5, 0xb1, 0x77, 0xe5, 0x5d, 0x47, 0xcd, 0x95, 0x57, 0xe0, 0xc6, 0x6b, 0x21, 0xf1,
	0x04, 0x3c, 0x04, 0x47, 0xb4, 0x3f, 0x76, 0xe2, 0x12, 0x10, 0x97, 0x68, 0xbf, 0xd9, 0xf1, 0xcc,
	0xe7, 0x89, 0x07, 0xf6, 0xf9, 0x55, 0x1e, 0x25, 0x9c, 0x46, 0x55, 0x5d, 0x4a, 0x5a, 0x90, 0x28,
	0xa3, 0x82, 0xcf, 0x93, 0xa5, 0xc2, 0x42, 0x5e, 0x31, 0xc9, 0xd0, 0xd0, 0x42, 0xde, 0x83, 0x9c,
	0xb1, 0x7c, 0x4e, 0x34, 0x3d, 0x29, 0x4b, 0x26, 0x13, 0x49, 0x59, 0x29, 0x0c, 0xcd, 0xbb, 0x6f,
	0x6f, 0xf5, 0x74, 0x51, 0x5f, 0x46, 0xa4, 0xe0, 0x72, 0x69, 0x2e, 0x83, 0x03, 0x98, 0x9c, 0x54,
	0x09, 0xff, 0x14, 0x13, 0xc1, 0x59, 0x29, 0x08, 0xba, 0x0b, 0x83, 0x5c, 0x01, 0xd8, 0xf1, 0x9d,
	0xe9, 0x28, 0x36, 0x43, 0xf0, 0xcd, 0x81, 0xc9, 0xb9, 0x48, 0x72, 0xd2, 0xf2, 0x5e, 0xc2, 0xa0,
	0x56, 0x00, 0x76, 0xfc, 0xde, 0x74, 0x7c, 0xb8, 0x1f, 0xda, 0x65, 0xc2, 0x0e, 0xcd, 0x4c, 0xc7,
	0xa5, 0xac, 0x96, 0xb1, 0xe1, 0x7b, 0x33, 0x80, 0x15, 0x88, 0x76, 0xa1, 0x77, 0x45, 0x96, 0xd6,
	0x4c, 0x1d, 0xd1, 0x53, 0x18, 0x2c, 0x92, 0x79, 0x4d, 0xb0, 0xeb, 0x3b, 0xd3, 0xf1, 0xe1, 0x5e,
	0x2b, 0xfc, 0x86, 0x95, 0x32, 0xa1, 0x25, 0xa9, 0xde, 0xcb, 0x44, 0xc6, 0x86, 0xf4, 0xda, 0x7d,
	0xe5, 0x04, 0x3f, 0x5c, 0x98, 0x74, 0x2e, 0xd1, 0x0b, 0xe8, 0x0b, 0x99, 0x48, 0xbb, 0x9b, 0xbf,
	0x59, 0x22, 0x54, 0x3f, 0x66, 0x35, 0xcd, 0x46, 0x18, 0x86, 0xbc, 0x62, 0x9f, 0x49, 0x2a, 0xb5,
	0xf7, 0x28, 0x6e, 0x46, 0x84, 0xa0, 0x5f, 0x0b, 0x52, 0xe1, 0x9e, 0x86, 0xf5, 0x59, 0xb1, 0xd3,
	0x79, 0x2d, 0x24, 0xa9, 0x70, 0xdf, 0xb0, 0xed, 0xa8, 0xdd, 0x39, 0x49, 0xf1, 0xe0, 0xdf, 0xee,
	0x9c, 0xa4, 0x8d, 0x3b, 0x27, 0xa9, 0x77, 0x0a, 0xa3, 0x76, 0xa1, 0x0d, 0xb1, 0x3c, 0xea, 0xc6,
	0xb2, 0xdb, 0xaa, 0x1e, 0x91, 0x05, 0x4d, 0x89, 0x58, 0x0b, 0xc4, 0x7b, 0x0b, 0xa3, 0x56, 0x7d,
	0x83, 0xd4, 0xc3, 0xae, 0xd4, 0xa4, 0x95, 0x52, 0x0f, 0xad, 0x07, 0xfb, 0x0c, 0x86, 0x56, 0x1d,
	0x1d, 0x40, 0x2f, 0x23, 0x0b, 0x1b, 0xe8, 0x9d, 0x1b, 0xe6, 0xa7, 0xe5, 0x25, 0x8b, 0xd5, 0x7d,
	0xc0, 0x00, 0x56, 0x10, 0xda, 0x01, 0x97, 0x66, 0xd6, 0xd9, 0xa5, 0x19, 0xba, 0x07, 0xdb, 0x69,
	0x52, 0x65, 0x1f, 0x69, 0x76, 0xdd, 0x24, 0xac, 0xe6, 0xd3, 0xec, 0x5a, 0x6d, 0x99, 0xf3, 0x1a,
	0x83, 0xef, 0x4c, 0xdd, 0x58, 0x1d, 0x15, 0x52, 0x90, 0x02, 0x8f, 0x0d, 0x52, 0x90, 0x42, 0xfd,
	0x0b, 0x9c, 0x66, 0x02, 0xdf, 0xf2, 0x7b, 0xd3, 0x41, 0xac, 0xcf, 0xc1, 0x13, 0xb8, 0xfd, 0x81,
	0x54, 0x82, 0xb2, 0xb2, 0xfd, 0x32, 0x31, 0x0c, 0x17, 0x06, 0xb2, 0xd6, 0xcd, 0x18, 0x3c, 0x86,
	0xbe, 0x7a, 0xc5, 0xc6, 0xcc, 0xf9, 0xc3, 0xcc, 0x6d, 0xcd, 0x0e, 0x7f, 0x39, 0x00, 0x27, 0xb3,
	0xf3, 0x23, 0xf3, 0xa2, 0xe8, 0x1d, 0xc0, 0xac, 0xa2, 0xa5, 0xd4, 0x65, 0x41, 0x7b, 0xa1, 0xe9,
	0x54, 0xd8, 0x74, 0x2a, 0x3c, 0x56, 0x9d, 0xf2, 0x56, 0x1f, 0x6b, 0xa7, 0x54, 0xc1, 0xce, 0x97,
	0xef, 0x3f, 0xbf, 0xba, 0xdb, 0x68, 0x2b, 0xd2, 0x75, 0x42, 0x67, 0x30, 0xd6, 0x6a, 0xba, 0x08,
	0xe2, 0x3f, 0xe4, 0x3a, 0xa5, 0x5a, 0x93, 0xd3, 0x95, 0x42, 0x67, 0x30, 0xb4, 0x21, 0xfc, 0x55,
	0x0a, 0xb7, 0x52, 0x37, 0xe2, 0x0a, 0x76, 0xb5, 0x18, 0xa0, 0xed, 0xc8, 0xc6, 0x74, 0xb1, 0xa5,
	0x9f, 0x7d, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x75, 0xcb, 0x20, 0x49, 0x83, 0x04, 0x00, 0x00,
}
