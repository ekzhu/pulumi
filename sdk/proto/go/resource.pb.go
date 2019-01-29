// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource.proto

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _struct "github.com/golang/protobuf/ptypes/struct"

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

// ReadResourceRequest contains enough information to uniquely qualify and read a resource's state.
type ReadResourceRequest struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                 string          `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Parent               string          `protobuf:"bytes,4,opt,name=parent" json:"parent,omitempty"`
	Properties           *_struct.Struct `protobuf:"bytes,5,opt,name=properties" json:"properties,omitempty"`
	Dependencies         []string        `protobuf:"bytes,6,rep,name=dependencies" json:"dependencies,omitempty"`
	Provider             string          `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Aliases              []string        `protobuf:"bytes,8,rep,name=aliases" json:"aliases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadResourceRequest) Reset()         { *m = ReadResourceRequest{} }
func (m *ReadResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ReadResourceRequest) ProtoMessage()    {}
func (*ReadResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_fed3f68c4def9ec7, []int{0}
}
func (m *ReadResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResourceRequest.Unmarshal(m, b)
}
func (m *ReadResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResourceRequest.Marshal(b, m, deterministic)
}
func (dst *ReadResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResourceRequest.Merge(dst, src)
}
func (m *ReadResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ReadResourceRequest.Size(m)
}
func (m *ReadResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResourceRequest proto.InternalMessageInfo

func (m *ReadResourceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadResourceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReadResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadResourceRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReadResourceRequest) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ReadResourceRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *ReadResourceRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ReadResourceRequest) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

// ReadResourceResponse contains the result of reading a resource's state.
type ReadResourceResponse struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Properties           *_struct.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadResourceResponse) Reset()         { *m = ReadResourceResponse{} }
func (m *ReadResourceResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResourceResponse) ProtoMessage()    {}
func (*ReadResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_fed3f68c4def9ec7, []int{1}
}
func (m *ReadResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResourceResponse.Unmarshal(m, b)
}
func (m *ReadResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResourceResponse.Marshal(b, m, deterministic)
}
func (dst *ReadResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResourceResponse.Merge(dst, src)
}
func (m *ReadResourceResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResourceResponse.Size(m)
}
func (m *ReadResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResourceResponse proto.InternalMessageInfo

func (m *ReadResourceResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ReadResourceResponse) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

// RegisterResourceRequest contains information about a resource object that was newly allocated.
type RegisterResourceRequest struct {
	Type                 string                                                   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name                 string                                                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Parent               string                                                   `protobuf:"bytes,3,opt,name=parent" json:"parent,omitempty"`
	Custom               bool                                                     `protobuf:"varint,4,opt,name=custom" json:"custom,omitempty"`
	Object               *_struct.Struct                                          `protobuf:"bytes,5,opt,name=object" json:"object,omitempty"`
	Protect              bool                                                     `protobuf:"varint,6,opt,name=protect" json:"protect,omitempty"`
	Dependencies         []string                                                 `protobuf:"bytes,7,rep,name=dependencies" json:"dependencies,omitempty"`
	Provider             string                                                   `protobuf:"bytes,8,opt,name=provider" json:"provider,omitempty"`
	PropertyDependencies map[string]*RegisterResourceRequest_PropertyDependencies `protobuf:"bytes,9,rep,name=propertyDependencies" json:"propertyDependencies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Aliases              []string                                                 `protobuf:"bytes,10,rep,name=aliases" json:"aliases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *RegisterResourceRequest) Reset()         { *m = RegisterResourceRequest{} }
func (m *RegisterResourceRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterResourceRequest) ProtoMessage()    {}
func (*RegisterResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_fed3f68c4def9ec7, []int{2}
}
func (m *RegisterResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceRequest.Unmarshal(m, b)
}
func (m *RegisterResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceRequest.Merge(dst, src)
}
func (m *RegisterResourceRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceRequest.Size(m)
}
func (m *RegisterResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceRequest proto.InternalMessageInfo

func (m *RegisterResourceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RegisterResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterResourceRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RegisterResourceRequest) GetCustom() bool {
	if m != nil {
		return m.Custom
	}
	return false
}

func (m *RegisterResourceRequest) GetObject() *_struct.Struct {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *RegisterResourceRequest) GetProtect() bool {
	if m != nil {
		return m.Protect
	}
	return false
}

func (m *RegisterResourceRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *RegisterResourceRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *RegisterResourceRequest) GetPropertyDependencies() map[string]*RegisterResourceRequest_PropertyDependencies {
	if m != nil {
		return m.PropertyDependencies
	}
	return nil
}

func (m *RegisterResourceRequest) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

// PropertyDependencies describes the resources that a particular property depends on.
type RegisterResourceRequest_PropertyDependencies struct {
	Urns                 []string `protobuf:"bytes,1,rep,name=urns" json:"urns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResourceRequest_PropertyDependencies) Reset() {
	*m = RegisterResourceRequest_PropertyDependencies{}
}
func (m *RegisterResourceRequest_PropertyDependencies) String() string {
	return proto.CompactTextString(m)
}
func (*RegisterResourceRequest_PropertyDependencies) ProtoMessage() {}
func (*RegisterResourceRequest_PropertyDependencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_fed3f68c4def9ec7, []int{2, 0}
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Unmarshal(m, b)
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceRequest_PropertyDependencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Merge(dst, src)
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.Size(m)
}
func (m *RegisterResourceRequest_PropertyDependencies) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceRequest_PropertyDependencies.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceRequest_PropertyDependencies proto.InternalMessageInfo

func (m *RegisterResourceRequest_PropertyDependencies) GetUrns() []string {
	if m != nil {
		return m.Urns
	}
	return nil
}

// RegisterResourceResponse is returned by the engine after a resource has finished being initialized.  It includes the
// auto-assigned URN, the provider-assigned ID, and any other properties initialized by the engine.
type RegisterResourceResponse struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Object               *_struct.Struct `protobuf:"bytes,3,opt,name=object" json:"object,omitempty"`
	Stable               bool            `protobuf:"varint,4,opt,name=stable" json:"stable,omitempty"`
	Stables              []string        `protobuf:"bytes,5,rep,name=stables" json:"stables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterResourceResponse) Reset()         { *m = RegisterResourceResponse{} }
func (m *RegisterResourceResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResourceResponse) ProtoMessage()    {}
func (*RegisterResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_fed3f68c4def9ec7, []int{3}
}
func (m *RegisterResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceResponse.Unmarshal(m, b)
}
func (m *RegisterResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceResponse.Merge(dst, src)
}
func (m *RegisterResourceResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceResponse.Size(m)
}
func (m *RegisterResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceResponse proto.InternalMessageInfo

func (m *RegisterResourceResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *RegisterResourceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterResourceResponse) GetObject() *_struct.Struct {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *RegisterResourceResponse) GetStable() bool {
	if m != nil {
		return m.Stable
	}
	return false
}

func (m *RegisterResourceResponse) GetStables() []string {
	if m != nil {
		return m.Stables
	}
	return nil
}

// RegisterResourceOutputsRequest adds extra resource outputs created by the program after registration has occurred.
type RegisterResourceOutputsRequest struct {
	Urn                  string          `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Outputs              *_struct.Struct `protobuf:"bytes,2,opt,name=outputs" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterResourceOutputsRequest) Reset()         { *m = RegisterResourceOutputsRequest{} }
func (m *RegisterResourceOutputsRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterResourceOutputsRequest) ProtoMessage()    {}
func (*RegisterResourceOutputsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_fed3f68c4def9ec7, []int{4}
}
func (m *RegisterResourceOutputsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResourceOutputsRequest.Unmarshal(m, b)
}
func (m *RegisterResourceOutputsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResourceOutputsRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterResourceOutputsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResourceOutputsRequest.Merge(dst, src)
}
func (m *RegisterResourceOutputsRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterResourceOutputsRequest.Size(m)
}
func (m *RegisterResourceOutputsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResourceOutputsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResourceOutputsRequest proto.InternalMessageInfo

func (m *RegisterResourceOutputsRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *RegisterResourceOutputsRequest) GetOutputs() *_struct.Struct {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadResourceRequest)(nil), "pulumirpc.ReadResourceRequest")
	proto.RegisterType((*ReadResourceResponse)(nil), "pulumirpc.ReadResourceResponse")
	proto.RegisterType((*RegisterResourceRequest)(nil), "pulumirpc.RegisterResourceRequest")
	proto.RegisterMapType((map[string]*RegisterResourceRequest_PropertyDependencies)(nil), "pulumirpc.RegisterResourceRequest.PropertyDependenciesEntry")
	proto.RegisterType((*RegisterResourceRequest_PropertyDependencies)(nil), "pulumirpc.RegisterResourceRequest.PropertyDependencies")
	proto.RegisterType((*RegisterResourceResponse)(nil), "pulumirpc.RegisterResourceResponse")
	proto.RegisterType((*RegisterResourceOutputsRequest)(nil), "pulumirpc.RegisterResourceOutputsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceMonitor service

type ResourceMonitorClient interface {
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error)
	RegisterResource(ctx context.Context, in *RegisterResourceRequest, opts ...grpc.CallOption) (*RegisterResourceResponse, error)
	RegisterResourceOutputs(ctx context.Context, in *RegisterResourceOutputsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type resourceMonitorClient struct {
	cc *grpc.ClientConn
}

func NewResourceMonitorClient(cc *grpc.ClientConn) ResourceMonitorClient {
	return &resourceMonitorClient{cc}
}

func (c *resourceMonitorClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error) {
	out := new(ReadResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/ReadResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterResource(ctx context.Context, in *RegisterResourceRequest, opts ...grpc.CallOption) (*RegisterResourceResponse, error) {
	out := new(RegisterResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceMonitorClient) RegisterResourceOutputs(ctx context.Context, in *RegisterResourceOutputsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceMonitor/RegisterResourceOutputs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceMonitor service

type ResourceMonitorServer interface {
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error)
	RegisterResource(context.Context, *RegisterResourceRequest) (*RegisterResourceResponse, error)
	RegisterResourceOutputs(context.Context, *RegisterResourceOutputsRequest) (*empty.Empty, error)
}

func RegisterResourceMonitorServer(s *grpc.Server, srv ResourceMonitorServer) {
	s.RegisterService(&_ResourceMonitor_serviceDesc, srv)
}

func _ResourceMonitor_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_ReadResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).ReadResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/ReadResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).ReadResource(ctx, req.(*ReadResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterResource(ctx, req.(*RegisterResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceMonitor_RegisterResourceOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterResourceOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceMonitorServer).RegisterResourceOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceMonitor/RegisterResourceOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceMonitorServer).RegisterResourceOutputs(ctx, req.(*RegisterResourceOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceMonitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.ResourceMonitor",
	HandlerType: (*ResourceMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _ResourceMonitor_Invoke_Handler,
		},
		{
			MethodName: "ReadResource",
			Handler:    _ResourceMonitor_ReadResource_Handler,
		},
		{
			MethodName: "RegisterResource",
			Handler:    _ResourceMonitor_RegisterResource_Handler,
		},
		{
			MethodName: "RegisterResourceOutputs",
			Handler:    _ResourceMonitor_RegisterResourceOutputs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}

func init() { proto.RegisterFile("resource.proto", fileDescriptor_resource_fed3f68c4def9ec7) }

var fileDescriptor_resource_fed3f68c4def9ec7 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x9d, 0xd6, 0x49, 0xa6, 0x55, 0xa8, 0x96, 0xa8, 0x75, 0x0d, 0x2a, 0x95, 0xb9, 0x00,
	0x07, 0x47, 0x84, 0x43, 0x11, 0x42, 0xe2, 0x42, 0x0f, 0x1c, 0x2a, 0xc0, 0x9c, 0x41, 0x72, 0x9c,
	0x21, 0x32, 0x4d, 0xbc, 0x66, 0x77, 0x1d, 0x29, 0x37, 0xde, 0x84, 0xc7, 0xe0, 0x85, 0x78, 0x08,
	0x8e, 0xec, 0xae, 0xd7, 0x21, 0xfe, 0x49, 0x53, 0x71, 0xca, 0xfc, 0x7e, 0x3b, 0xfb, 0xcd, 0xb7,
	0x0e, 0x0c, 0x18, 0x72, 0x9a, 0xb3, 0x18, 0x83, 0x8c, 0x51, 0x41, 0x49, 0x3f, 0xcb, 0xe7, 0xf9,
	0x22, 0x61, 0x59, 0xec, 0x3d, 0x98, 0x51, 0x3a, 0x9b, 0xe3, 0x48, 0x27, 0x26, 0xf9, 0xd7, 0x11,
	0x2e, 0x32, 0xb1, 0x2a, 0xea, 0xbc, 0x87, 0xf5, 0x24, 0x17, 0x2c, 0x8f, 0x85, 0xc9, 0x0e, 0xe4,
	0xcf, 0x32, 0x99, 0x22, 0x2b, 0x7c, 0xff, 0x8f, 0x05, 0xf7, 0x43, 0x8c, 0xa6, 0xa1, 0x39, 0x2c,
	0xc4, 0xef, 0x39, 0x72, 0x41, 0x06, 0x60, 0x27, 0x53, 0xd7, 0xba, 0xb0, 0x9e, 0xf4, 0x43, 0x69,
	0x11, 0x02, 0xfb, 0x62, 0x95, 0xa1, 0x6b, 0xeb, 0x88, 0xb6, 0x55, 0x2c, 0x8d, 0x16, 0xe8, 0x76,
	0x8a, 0x98, 0xb2, 0xc9, 0x09, 0x38, 0x59, 0xc4, 0x30, 0x15, 0xee, 0xbe, 0x8e, 0x1a, 0x8f, 0x5c,
	0x02, 0xc8, 0x03, 0x33, 0x64, 0x22, 0x41, 0xee, 0x1e, 0xc8, 0xdc, 0xe1, 0xf8, 0x34, 0x28, 0x46,
	0x0d, 0xca, 0x51, 0x83, 0x4f, 0x7a, 0xd4, 0x70, 0xa3, 0x94, 0xf8, 0x70, 0x34, 0xc5, 0x0c, 0xd3,
	0x29, 0xa6, 0xb1, 0x6a, 0x75, 0x2e, 0x3a, 0x12, 0xb6, 0x12, 0x23, 0x1e, 0xf4, 0xca, 0x6b, 0xb9,
	0x5d, 0x7d, 0xec, 0xda, 0x27, 0x2e, 0x74, 0xa3, 0x79, 0x12, 0x71, 0xd9, 0xda, 0xd3, 0xad, 0xa5,
	0xeb, 0x47, 0x30, 0xac, 0xde, 0x9c, 0x67, 0x34, 0xe5, 0x48, 0x8e, 0xa1, 0x93, 0xb3, 0xd4, 0xdc,
	0x5d, 0x99, 0xb5, 0xe1, 0xed, 0x3b, 0x0f, 0xef, 0xff, 0xda, 0x87, 0xd3, 0x10, 0x67, 0x09, 0x17,
	0xc8, 0xea, 0x0c, 0x97, 0x8c, 0x5a, 0x2d, 0x8c, 0xda, 0xad, 0x8c, 0x76, 0x2a, 0x8c, 0xca, 0x78,
	0x9c, 0x73, 0x41, 0x17, 0x9a, 0xe9, 0x5e, 0x68, 0x3c, 0x32, 0x02, 0x87, 0x4e, 0xbe, 0x61, 0x2c,
	0x76, 0xb1, 0x6c, 0xca, 0x14, 0x43, 0x2a, 0xa5, 0x3a, 0x1c, 0x8d, 0x54, 0xba, 0x0d, 0xee, 0xbb,
	0x3b, 0xb8, 0xef, 0xd5, 0xb8, 0xcf, 0x60, 0x68, 0xc8, 0x58, 0xbd, 0xdd, 0xc4, 0xe9, 0x4b, 0x9c,
	0xc3, 0xf1, 0xeb, 0x60, 0xad, 0xe8, 0x60, 0x0b, 0x49, 0xc1, 0x87, 0x96, 0xf6, 0xab, 0x54, 0xb0,
	0x55, 0xd8, 0x8a, 0xbc, 0xb9, 0x6d, 0xa8, 0x6c, 0xdb, 0x7b, 0x06, 0xc3, 0x36, 0x30, 0x45, 0xb9,
	0x5c, 0x31, 0x97, 0x6b, 0x50, 0xe5, 0xda, 0xf6, 0x7e, 0x58, 0x70, 0xb6, 0xf5, 0x64, 0xa5, 0x8f,
	0x1b, 0x5c, 0x95, 0xfa, 0x90, 0x26, 0xb9, 0x86, 0x83, 0x65, 0x34, 0xcf, 0xd1, 0x48, 0xe3, 0xf2,
	0x3f, 0x2f, 0x16, 0x16, 0x28, 0xaf, 0xec, 0x97, 0x96, 0xff, 0xd3, 0x02, 0xb7, 0xd9, 0xbb, 0x55,
	0xa1, 0xc5, 0x73, 0xb5, 0xd7, 0xcf, 0xf5, 0x9f, 0x08, 0x3a, 0x77, 0x13, 0x81, 0x54, 0x13, 0x17,
	0xd1, 0x64, 0x8e, 0xa5, 0x9a, 0x0a, 0x4f, 0x11, 0x5a, 0x58, 0xea, 0xd1, 0x6a, 0x42, 0x8d, 0xeb,
	0x23, 0x9c, 0xd7, 0x07, 0x7c, 0x9f, 0x8b, 0x2c, 0x17, 0xbc, 0x54, 0x78, 0x73, 0xcc, 0xe7, 0xd0,
	0xa5, 0x45, 0xcd, 0xae, 0x57, 0x54, 0xd6, 0x8d, 0x7f, 0xdb, 0x70, 0xaf, 0xc4, 0xbf, 0xa6, 0x69,
	0x22, 0x28, 0x23, 0x6f, 0xc0, 0x79, 0x97, 0x2e, 0xe9, 0x8d, 0x1c, 0x6f, 0x83, 0xea, 0x22, 0x64,
	0x0e, 0xf7, 0xce, 0x5a, 0x32, 0x05, 0x7d, 0xfe, 0x1e, 0xf9, 0x08, 0x47, 0x9b, 0x4f, 0x9f, 0x9c,
	0x57, 0x36, 0xd6, 0xf8, 0x1a, 0x7a, 0x8f, 0xb6, 0xe6, 0xd7, 0x90, 0x9f, 0xe1, 0xb8, 0x4e, 0x07,
	0xf1, 0x77, 0x0b, 0xc1, 0x7b, 0x7c, 0x6b, 0xcd, 0x1a, 0xfe, 0x4b, 0xf3, 0x43, 0x62, 0xd8, 0x26,
	0x4f, 0x6f, 0x41, 0xa8, 0x6e, 0xc4, 0x3b, 0x69, 0xd0, 0x7d, 0xa5, 0xfe, 0x39, 0xfc, 0xbd, 0x89,
	0xa3, 0x23, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x41, 0xf6, 0x25, 0x77, 0x76, 0x06, 0x00,
	0x00,
}
