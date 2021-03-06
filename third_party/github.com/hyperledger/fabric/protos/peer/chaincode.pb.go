/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode.proto

package peer

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}

var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}

func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{2, 0}
}

//ChaincodeID contains the path as specified by the deploy transaction
//that created it as well as the hashCode that is generated by the
//system for the path. From the user level (ie, CLI, REST API and so on)
//deploy transaction is expected to provide the path and other requests
//are expected to provide the hashCode. The other value will be ignored.
//Internally, the structure could contain both values. For instance, the
//hashCode will be set when first generated using the path
type ChaincodeID struct {
	//deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	//all other requests will use the name (really a hashcode) generated by
	//the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//user friendly version name for the chaincode
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeID) Reset()         { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()    {}
func (*ChaincodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{0}
}

func (m *ChaincodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeID.Unmarshal(m, b)
}
func (m *ChaincodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeID.Marshal(b, m, deterministic)
}
func (m *ChaincodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeID.Merge(m, src)
}
func (m *ChaincodeID) XXX_Size() int {
	return xxx_messageInfo_ChaincodeID.Size(m)
}
func (m *ChaincodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeID.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeID proto.InternalMessageInfo

func (m *ChaincodeID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args        [][]byte          `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	Decorations map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// is_init is used for the application to signal that an invocation is to be routed
	// to the legacy 'Init' function for compatibility with chaincodes which handled
	// Init in the old way.  New applications should manage their initialized state
	// themselves.
	IsInit               bool     `protobuf:"varint,3,opt,name=is_init,json=isInit,proto3" json:"is_init,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeInput) Reset()         { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()    {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{1}
}

func (m *ChaincodeInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInput.Unmarshal(m, b)
}
func (m *ChaincodeInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInput.Marshal(b, m, deterministic)
}
func (m *ChaincodeInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInput.Merge(m, src)
}
func (m *ChaincodeInput) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInput.Size(m)
}
func (m *ChaincodeInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInput.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInput proto.InternalMessageInfo

func (m *ChaincodeInput) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ChaincodeInput) GetDecorations() map[string][]byte {
	if m != nil {
		return m.Decorations
	}
	return nil
}

func (m *ChaincodeInput) GetIsInit() bool {
	if m != nil {
		return m.IsInit
	}
	return false
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type                 ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,proto3,enum=sdk.protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId          *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	Input                *ChaincodeInput    `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Timeout              int32              `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChaincodeSpec) Reset()         { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()    {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{2}
}

func (m *ChaincodeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeSpec.Unmarshal(m, b)
}
func (m *ChaincodeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeSpec.Merge(m, src)
}
func (m *ChaincodeSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeSpec.Size(m)
}
func (m *ChaincodeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeSpec proto.InternalMessageInfo

func (m *ChaincodeSpec) GetType() ChaincodeSpec_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeSpec_UNDEFINED
}

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeSpec) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec        *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	CodePackage          []byte         `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChaincodeDeploymentSpec) Reset()         { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()    {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{3}
}

func (m *ChaincodeDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Unmarshal(m, b)
}
func (m *ChaincodeDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeDeploymentSpec.Merge(m, src)
}
func (m *ChaincodeDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Size(m)
}
func (m *ChaincodeDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeDeploymentSpec proto.InternalMessageInfo

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetCodePackage() []byte {
	if m != nil {
		return m.CodePackage
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec        *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChaincodeInvocationSpec) Reset()         { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()    {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{4}
}

func (m *ChaincodeInvocationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInvocationSpec.Unmarshal(m, b)
}
func (m *ChaincodeInvocationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInvocationSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeInvocationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInvocationSpec.Merge(m, src)
}
func (m *ChaincodeInvocationSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInvocationSpec.Size(m)
}
func (m *ChaincodeInvocationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInvocationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInvocationSpec proto.InternalMessageInfo

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// LifecycleEvent is used as the payload of the chaincode event emitted by LSCC
type LifecycleEvent struct {
	ChaincodeName        string   `protobuf:"bytes,1,opt,name=chaincode_name,json=chaincodeName,proto3" json:"chaincode_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleEvent) Reset()         { *m = LifecycleEvent{} }
func (m *LifecycleEvent) String() string { return proto.CompactTextString(m) }
func (*LifecycleEvent) ProtoMessage()    {}
func (*LifecycleEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{5}
}

func (m *LifecycleEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleEvent.Unmarshal(m, b)
}
func (m *LifecycleEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleEvent.Marshal(b, m, deterministic)
}
func (m *LifecycleEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleEvent.Merge(m, src)
}
func (m *LifecycleEvent) XXX_Size() int {
	return xxx_messageInfo_LifecycleEvent.Size(m)
}
func (m *LifecycleEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleEvent proto.InternalMessageInfo

func (m *LifecycleEvent) GetChaincodeName() string {
	if m != nil {
		return m.ChaincodeName
	}
	return ""
}

func init() {
	proto.RegisterEnum("sdk.protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterType((*ChaincodeID)(nil), "sdk.protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "sdk.protos.ChaincodeInput")
	proto.RegisterMapType((map[string][]byte)(nil), "sdk.protos.ChaincodeInput.DecorationsEntry")
	proto.RegisterType((*ChaincodeSpec)(nil), "sdk.protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "sdk.protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "sdk.protos.ChaincodeInvocationSpec")
	proto.RegisterType((*LifecycleEvent)(nil), "sdk.protos.LifecycleEvent")
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor_202814c635ff5fee) }

var fileDescriptor_202814c635ff5fee = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xb1, 0xd3, 0xa6, 0xe3, 0x34, 0x32, 0x4b, 0xa1, 0x56, 0x4f, 0xc1, 0x12, 0x22, 0x48,
	0xc8, 0x91, 0x82, 0x04, 0x08, 0xa1, 0x4a, 0xa1, 0x0e, 0x55, 0xaa, 0x2a, 0x41, 0xcb, 0xc7, 0x81,
	0x4b, 0xe4, 0xac, 0x27, 0xce, 0xaa, 0xc9, 0xda, 0xb2, 0x37, 0x11, 0xfe, 0x37, 0xfc, 0x1c, 0xfe,
	0x15, 0x68, 0x77, 0xc9, 0x47, 0x69, 0x6f, 0x9c, 0x3c, 0x3b, 0x7e, 0x33, 0xf3, 0xde, 0xdb, 0x1d,
	0x38, 0xc9, 0x11, 0x8b, 0x2e, 0x9b, 0xc7, 0x5c, 0xb0, 0x2c, 0xc1, 0x30, 0x2f, 0x32, 0x99, 0x91,
	0x03, 0xfd, 0x29, 0x83, 0x31, 0xb8, 0x17, 0x9b, 0x5f, 0xc3, 0x88, 0x10, 0x70, 0xf2, 0x58, 0xce,
	0x7d, 0xab, 0x6d, 0x75, 0x8e, 0xa8, 0x8e, 0x55, 0x4e, 0xc4, 0x4b, 0xf4, 0x6b, 0x26, 0xa7, 0x62,
	0xe2, 0xc3, 0xe1, 0x1a, 0x8b, 0x92, 0x67, 0xc2, 0xb7, 0x75, 0x7a, 0x73, 0x0c, 0x7e, 0x59, 0xd0,
	0xda, 0x75, 0x14, 0xf9, 0x4a, 0xaa, 0x06, 0x71, 0x91, 0x96, 0xbe, 0xd5, 0xb6, 0x3b, 0x4d, 0xaa,
	0x63, 0x32, 0x04, 0x37, 0x41, 0x96, 0x15, 0xb1, 0xe4, 0x99, 0x28, 0xfd, 0x5a, 0xdb, 0xee, 0xb8,
	0xbd, 0xe7, 0x86, 0x5c, 0x19, 0xde, 0x6e, 0x10, 0x46, 0x3b, 0xe4, 0x40, 0xc8, 0xa2, 0xa2, 0xfb,
	0xb5, 0xe4, 0x14, 0x0e, 0x79, 0x39, 0xe1, 0x82, 0x4b, 0xcd, 0xa5, 0x41, 0x0f, 0x78, 0x39, 0x14,
	0x5c, 0x9e, 0x9d, 0x83, 0xf7, 0x6f, 0x25, 0xf1, 0xc0, 0xbe, 0xc1, 0xea, 0xaf, 0x3e, 0x15, 0x92,
	0x13, 0xa8, 0xaf, 0xe3, 0xc5, 0xca, 0xe8, 0x6b, 0x52, 0x73, 0x78, 0x57, 0x7b, 0x6b, 0x05, 0xbf,
	0x2d, 0x38, 0xde, 0x32, 0xf9, 0x9c, 0x23, 0x23, 0x21, 0x38, 0xb2, 0xca, 0x51, 0x97, 0xb7, 0x7a,
	0x67, 0x77, 0xe8, 0x2a, 0x50, 0xf8, 0xa5, 0xca, 0x91, 0x6a, 0x1c, 0x79, 0x0d, 0xcd, 0xad, 0xf1,
	0x13, 0x9e, 0xe8, 0x11, 0x6e, 0xef, 0xd1, 0x5d, 0x99, 0x11, 0x75, 0xb7, 0xc0, 0x61, 0x42, 0x5e,
	0x42, 0x9d, 0x2b, 0xe5, 0x5a, 0x90, 0xdb, 0x7b, 0x72, 0xbf, 0x2f, 0xd4, 0x80, 0xd4, 0x65, 0x48,
	0xbe, 0xc4, 0x6c, 0x25, 0x7d, 0xa7, 0x6d, 0x75, 0xea, 0x74, 0x73, 0x0c, 0xce, 0xc1, 0x51, 0x6c,
	0xc8, 0x31, 0x1c, 0x7d, 0x1d, 0x45, 0x83, 0x8f, 0xc3, 0xd1, 0x20, 0xf2, 0x1e, 0x10, 0x80, 0x83,
	0xcb, 0xf1, 0x75, 0x7f, 0x74, 0xe9, 0x59, 0xa4, 0x01, 0xce, 0x68, 0x1c, 0x0d, 0xbc, 0x1a, 0x39,
	0x04, 0xfb, 0xa2, 0x4f, 0x3d, 0x5b, 0xa5, 0xae, 0xfa, 0xdf, 0xfa, 0x9e, 0x13, 0xfc, 0xb4, 0xe0,
	0x74, 0x3b, 0x33, 0xc2, 0x7c, 0x91, 0x55, 0x4b, 0x14, 0x52, 0x7b, 0xf1, 0x1e, 0x5a, 0x3b, 0x6d,
	0x65, 0x8e, 0x4c, 0xbb, 0xe2, 0xf6, 0x1e, 0xdf, 0xeb, 0x0a, 0x3d, 0x66, 0xb7, 0x9c, 0x7c, 0x0a,
	0x4d, 0x5d, 0x98, 0xc7, 0xec, 0x26, 0x4e, 0x51, 0x0b, 0x6d, 0x52, 0x57, 0xe5, 0x3e, 0x99, 0xd4,
	0x95, 0xd3, 0xa8, 0x79, 0xf6, 0x95, 0xd3, 0x70, 0xbc, 0x3a, 0x6d, 0xe1, 0x6c, 0x86, 0x4c, 0xf2,
	0x35, 0x4e, 0x92, 0x58, 0x22, 0x6d, 0xe0, 0x0f, 0x64, 0x13, 0x14, 0xeb, 0x20, 0xdf, 0x63, 0x38,
	0x14, 0xeb, 0x8c, 0xe9, 0xdb, 0xfe, 0x7f, 0x86, 0x66, 0x3c, 0x7d, 0xc8, 0x93, 0x49, 0x8a, 0x02,
	0xcd, 0x23, 0x9a, 0xc4, 0x8b, 0x34, 0x78, 0x03, 0xad, 0x6b, 0x3e, 0x43, 0x56, 0xb1, 0x05, 0x0e,
	0xd6, 0x28, 0x24, 0x79, 0xb6, 0x3f, 0x48, 0xef, 0x8a, 0x79, 0x5f, 0xbb, 0x8e, 0xa3, 0x78, 0x89,
	0x1f, 0xc6, 0x10, 0x64, 0x45, 0x1a, 0xce, 0xab, 0x1c, 0x8b, 0x05, 0x26, 0x29, 0x16, 0xe1, 0x2c,
	0x9e, 0x16, 0x9c, 0x6d, 0xf8, 0xa8, 0x4d, 0xfd, 0xfe, 0x22, 0xe5, 0x72, 0xbe, 0x9a, 0x86, 0x2c,
	0x5b, 0x76, 0xf7, 0xa0, 0x5d, 0x03, 0xed, 0x1a, 0x68, 0x57, 0x41, 0xa7, 0x66, 0x89, 0x5f, 0xfd,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xa1, 0x07, 0xb3, 0xe3, 0x03, 0x00, 0x00,
}
