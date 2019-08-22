// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/config/app/konvoyctl/v1alpha1/config.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Configuration defines configuration of `konvoyctl`.
type Configuration struct {
	// List of known Control Planes.
	ControlPlanes []*ControlPlane `protobuf:"bytes,1,rep,name=control_planes,json=controlPlanes,proto3" json:"control_planes,omitempty"`
	// List of configured `konvoyctl` contexts.
	Contexts []*Context `protobuf:"bytes,2,rep,name=contexts,proto3" json:"contexts,omitempty"`
	// Name of the context to use by default.
	CurrentContext       string   `protobuf:"bytes,3,opt,name=current_context,json=currentContext,proto3" json:"current_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_5794df17731045dd, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return m.Size()
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetControlPlanes() []*ControlPlane {
	if m != nil {
		return m.ControlPlanes
	}
	return nil
}

func (m *Configuration) GetContexts() []*Context {
	if m != nil {
		return m.Contexts
	}
	return nil
}

func (m *Configuration) GetCurrentContext() string {
	if m != nil {
		return m.CurrentContext
	}
	return ""
}

// ControlPlane defines a Control Plane.
type ControlPlane struct {
	// Name defines a reference name for a given Control Plane.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Coordinates defines coordinates of a given Control Plane.
	Coordinates          *ControlPlaneCoordinates `protobuf:"bytes,2,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ControlPlane) Reset()         { *m = ControlPlane{} }
func (m *ControlPlane) String() string { return proto.CompactTextString(m) }
func (*ControlPlane) ProtoMessage()    {}
func (*ControlPlane) Descriptor() ([]byte, []int) {
	return fileDescriptor_5794df17731045dd, []int{1}
}
func (m *ControlPlane) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlPlane) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlPlane.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlPlane) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlane.Merge(m, src)
}
func (m *ControlPlane) XXX_Size() int {
	return m.Size()
}
func (m *ControlPlane) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlane.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlane proto.InternalMessageInfo

func (m *ControlPlane) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ControlPlane) GetCoordinates() *ControlPlaneCoordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

// ControlPlaneCoordinates defines coordinates of a Control Plane.
type ControlPlaneCoordinates struct {
	ApiServer            *ControlPlaneCoordinates_ApiServer `protobuf:"bytes,1,opt,name=api_server,json=apiServer,proto3" json:"api_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ControlPlaneCoordinates) Reset()         { *m = ControlPlaneCoordinates{} }
func (m *ControlPlaneCoordinates) String() string { return proto.CompactTextString(m) }
func (*ControlPlaneCoordinates) ProtoMessage()    {}
func (*ControlPlaneCoordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_5794df17731045dd, []int{2}
}
func (m *ControlPlaneCoordinates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlPlaneCoordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlPlaneCoordinates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlPlaneCoordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlaneCoordinates.Merge(m, src)
}
func (m *ControlPlaneCoordinates) XXX_Size() int {
	return m.Size()
}
func (m *ControlPlaneCoordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlaneCoordinates.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlaneCoordinates proto.InternalMessageInfo

func (m *ControlPlaneCoordinates) GetApiServer() *ControlPlaneCoordinates_ApiServer {
	if m != nil {
		return m.ApiServer
	}
	return nil
}

type ControlPlaneCoordinates_ApiServer struct {
	// URL defines URL of the Control Plane API Server.
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlPlaneCoordinates_ApiServer) Reset()         { *m = ControlPlaneCoordinates_ApiServer{} }
func (m *ControlPlaneCoordinates_ApiServer) String() string { return proto.CompactTextString(m) }
func (*ControlPlaneCoordinates_ApiServer) ProtoMessage()    {}
func (*ControlPlaneCoordinates_ApiServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5794df17731045dd, []int{2, 0}
}
func (m *ControlPlaneCoordinates_ApiServer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlPlaneCoordinates_ApiServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlPlaneCoordinates_ApiServer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlPlaneCoordinates_ApiServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlaneCoordinates_ApiServer.Merge(m, src)
}
func (m *ControlPlaneCoordinates_ApiServer) XXX_Size() int {
	return m.Size()
}
func (m *ControlPlaneCoordinates_ApiServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlaneCoordinates_ApiServer.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlaneCoordinates_ApiServer proto.InternalMessageInfo

func (m *ControlPlaneCoordinates_ApiServer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Context defines a context in which individual `konvoyctl` commands run.
type Context struct {
	// Name defines a reference name for a given context.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ControlPlane defines a reference to a known Control Plane.
	ControlPlane string `protobuf:"bytes,2,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	// Defaults defines default settings for a given context.
	Defaults             *Context_Defaults `protobuf:"bytes,3,opt,name=defaults,proto3" json:"defaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_5794df17731045dd, []int{3}
}
func (m *Context) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Context.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return m.Size()
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Context) GetControlPlane() string {
	if m != nil {
		return m.ControlPlane
	}
	return ""
}

func (m *Context) GetDefaults() *Context_Defaults {
	if m != nil {
		return m.Defaults
	}
	return nil
}

// Defaults defines default settings for a context.
type Context_Defaults struct {
	// Mesh defines a Mesh to use in requests if one is not provided explicitly.
	Mesh                 string   `protobuf:"bytes,1,opt,name=mesh,proto3" json:"mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context_Defaults) Reset()         { *m = Context_Defaults{} }
func (m *Context_Defaults) String() string { return proto.CompactTextString(m) }
func (*Context_Defaults) ProtoMessage()    {}
func (*Context_Defaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_5794df17731045dd, []int{3, 0}
}
func (m *Context_Defaults) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Context_Defaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Context_Defaults.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Context_Defaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context_Defaults.Merge(m, src)
}
func (m *Context_Defaults) XXX_Size() int {
	return m.Size()
}
func (m *Context_Defaults) XXX_DiscardUnknown() {
	xxx_messageInfo_Context_Defaults.DiscardUnknown(m)
}

var xxx_messageInfo_Context_Defaults proto.InternalMessageInfo

func (m *Context_Defaults) GetMesh() string {
	if m != nil {
		return m.Mesh
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "konvoyctl.config.v1alpha1.Configuration")
	proto.RegisterType((*ControlPlane)(nil), "konvoyctl.config.v1alpha1.ControlPlane")
	proto.RegisterType((*ControlPlaneCoordinates)(nil), "konvoyctl.config.v1alpha1.ControlPlaneCoordinates")
	proto.RegisterType((*ControlPlaneCoordinates_ApiServer)(nil), "konvoyctl.config.v1alpha1.ControlPlaneCoordinates.ApiServer")
	proto.RegisterType((*Context)(nil), "konvoyctl.config.v1alpha1.Context")
	proto.RegisterType((*Context_Defaults)(nil), "konvoyctl.config.v1alpha1.Context.Defaults")
}

func init() {
	proto.RegisterFile("pkg/config/app/konvoyctl/v1alpha1/config.proto", fileDescriptor_5794df17731045dd)
}

var fileDescriptor_5794df17731045dd = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x8e, 0xd3, 0x30,
	0x1c, 0xc6, 0xe5, 0x4b, 0x81, 0xe4, 0x9f, 0xeb, 0x81, 0x3c, 0x70, 0xa1, 0x82, 0xe8, 0x94, 0xe5,
	0x8a, 0x90, 0x1c, 0x5d, 0x58, 0x11, 0x12, 0x2d, 0x12, 0x1b, 0x42, 0x61, 0x63, 0x29, 0x26, 0xf5,
	0xf5, 0xa2, 0xcb, 0xd9, 0x96, 0xe3, 0x44, 0xe5, 0x0d, 0x10, 0x0b, 0x4f, 0xc2, 0x0b, 0x30, 0xc1,
	0xd6, 0x91, 0x47, 0x40, 0xdd, 0x78, 0x0b, 0x14, 0xc7, 0x09, 0xe9, 0x50, 0x15, 0x36, 0xe7, 0xcb,
	0xf7, 0xfb, 0xf2, 0xe5, 0xef, 0x3f, 0x10, 0x79, 0xbd, 0x8a, 0x33, 0xc1, 0x2f, 0xf3, 0x55, 0x4c,
	0xa5, 0x8c, 0xaf, 0x05, 0xaf, 0xc5, 0xc7, 0x4c, 0x17, 0x71, 0x7d, 0x41, 0x0b, 0x79, 0x45, 0x2f,
	0xec, 0x5b, 0x22, 0x95, 0xd0, 0x02, 0x3f, 0xe8, 0x0d, 0xc4, 0xea, 0x9d, 0x6f, 0x72, 0x5a, 0xd3,
	0x22, 0x5f, 0x52, 0xcd, 0xe2, 0xee, 0xd0, 0x32, 0xd1, 0x06, 0xc1, 0x78, 0x6e, 0xcc, 0x95, 0xa2,
	0x3a, 0x17, 0x1c, 0xbf, 0x86, 0x93, 0x4c, 0x70, 0xad, 0x44, 0xb1, 0x90, 0x05, 0xe5, 0xac, 0x0c,
	0xd0, 0x99, 0x33, 0xf5, 0x93, 0x73, 0xb2, 0x37, 0x9e, 0xcc, 0x5b, 0xe0, 0x4d, 0xe3, 0x4f, 0xc7,
	0xd9, 0xe0, 0xa9, 0xc4, 0xcf, 0xc1, 0x6d, 0x04, 0xb6, 0xd6, 0x65, 0x70, 0x64, 0x92, 0xa2, 0x03,
	0x49, 0x6c, 0xad, 0xd3, 0x9e, 0xc1, 0xe7, 0x70, 0x37, 0xab, 0x94, 0x62, 0x5c, 0x2f, 0xac, 0x16,
	0x38, 0x67, 0x68, 0xea, 0xa5, 0x27, 0x56, 0xb6, 0x48, 0xf4, 0x05, 0xc1, 0xf1, 0xb0, 0x08, 0x7e,
	0x04, 0x23, 0x4e, 0x6f, 0x58, 0x80, 0x1a, 0xfb, 0xcc, 0xfb, 0xf6, 0xfb, 0xbb, 0x33, 0x52, 0x47,
	0xf7, 0x50, 0x6a, 0x64, 0xfc, 0x1e, 0xfc, 0x4c, 0x08, 0xb5, 0xcc, 0x39, 0xd5, 0xac, 0xe9, 0x86,
	0xa6, 0x7e, 0x92, 0xfc, 0xe3, 0x5f, 0xce, 0xff, 0x92, 0x33, 0x68, 0x92, 0x6f, 0x7d, 0x46, 0x4d,
	0xf4, 0x30, 0x32, 0xfa, 0x8a, 0xe0, 0x74, 0x0f, 0x84, 0x57, 0x00, 0x54, 0xe6, 0x8b, 0x92, 0xa9,
	0x9a, 0x29, 0x53, 0xd1, 0x4f, 0x9e, 0xfd, 0xff, 0xc7, 0xc9, 0x0b, 0x99, 0xbf, 0x35, 0x19, 0x3b,
	0x35, 0x3c, 0xda, 0xc9, 0x93, 0xc7, 0xe0, 0xf5, 0x1e, 0xfc, 0x10, 0x9c, 0x4a, 0x15, 0x76, 0x22,
	0x2d, 0xa0, 0x9c, 0x4f, 0x08, 0xa5, 0x8d, 0x1c, 0xfd, 0x40, 0x70, 0xc7, 0x4e, 0xf3, 0xd0, 0xf0,
	0x08, 0x8c, 0x77, 0xb6, 0xc4, 0x8c, 0x6f, 0xc7, 0x77, 0x3c, 0x5c, 0x03, 0xfc, 0x0a, 0xdc, 0x25,
	0xbb, 0xa4, 0x55, 0xa1, 0x4b, 0x73, 0x7d, 0x7e, 0xf2, 0xe4, 0xf0, 0x16, 0x90, 0x97, 0x16, 0x49,
	0x7b, 0x78, 0x12, 0x82, 0xdb, 0xa9, 0x18, 0xc3, 0xe8, 0x86, 0x95, 0x57, 0x6d, 0xc7, 0xd4, 0x9c,
	0x67, 0xf7, 0x37, 0xdb, 0x10, 0xfd, 0xdc, 0x86, 0xe8, 0xd7, 0x36, 0x44, 0xef, 0xdc, 0x2e, 0xf2,
	0xc3, 0x6d, 0xb3, 0xef, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x54, 0xa5, 0x57, 0x37, 0x55,
	0x03, 0x00, 0x00,
}

func (m *Configuration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Configuration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ControlPlanes) > 0 {
		for _, msg := range m.ControlPlanes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Contexts) > 0 {
		for _, msg := range m.Contexts {
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.CurrentContext) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.CurrentContext)))
		i += copy(dAtA[i:], m.CurrentContext)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ControlPlane) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlPlane) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Coordinates != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Coordinates.Size()))
		n1, err := m.Coordinates.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ControlPlaneCoordinates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlPlaneCoordinates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApiServer != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ApiServer.Size()))
		n2, err := m.ApiServer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ControlPlaneCoordinates_ApiServer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlPlaneCoordinates_ApiServer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Context) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Context) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.ControlPlane) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ControlPlane)))
		i += copy(dAtA[i:], m.ControlPlane)
	}
	if m.Defaults != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Defaults.Size()))
		n3, err := m.Defaults.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Context_Defaults) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Context_Defaults) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Mesh) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Mesh)))
		i += copy(dAtA[i:], m.Mesh)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Configuration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ControlPlanes) > 0 {
		for _, e := range m.ControlPlanes {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if len(m.Contexts) > 0 {
		for _, e := range m.Contexts {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.CurrentContext)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ControlPlane) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Coordinates != nil {
		l = m.Coordinates.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ControlPlaneCoordinates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApiServer != nil {
		l = m.ApiServer.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ControlPlaneCoordinates_ApiServer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Context) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ControlPlane)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Defaults != nil {
		l = m.Defaults.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Context_Defaults) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Mesh)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Configuration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Configuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Configuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControlPlanes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ControlPlanes = append(m.ControlPlanes, &ControlPlane{})
			if err := m.ControlPlanes[len(m.ControlPlanes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contexts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contexts = append(m.Contexts, &Context{})
			if err := m.Contexts[len(m.Contexts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlPlane) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlPlane: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlPlane: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coordinates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Coordinates == nil {
				m.Coordinates = &ControlPlaneCoordinates{}
			}
			if err := m.Coordinates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlPlaneCoordinates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlPlaneCoordinates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlPlaneCoordinates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiServer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApiServer == nil {
				m.ApiServer = &ControlPlaneCoordinates_ApiServer{}
			}
			if err := m.ApiServer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlPlaneCoordinates_ApiServer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ApiServer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiServer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Context) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Context: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Context: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControlPlane", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ControlPlane = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Defaults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Defaults == nil {
				m.Defaults = &Context_Defaults{}
			}
			if err := m.Defaults.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Context_Defaults) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Defaults: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Defaults: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mesh", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mesh = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)