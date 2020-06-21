// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Port
	Graph
	Error
	Message
	Payload
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "sync"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Port struct {
	mapLocker sync.Mutex

	Seq       int32             `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Url       string            `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	GraphName string            `protobuf:"bytes,3,opt,name=graph_name,json=graphName" json:"graph_name,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Port) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Port) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Port) GetGraphName() string {
	if m != nil {
		return m.GraphName
	}
	return ""
}

func (m *Port) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Graph struct {
	Name  string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Seq   int32   `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	Ports []*Port `protobuf:"bytes,3,rep,name=ports" json:"ports,omitempty"`
}

func (m *Graph) Reset()                    { *m = Graph{} }
func (m *Graph) String() string            { return proto.CompactTextString(m) }
func (*Graph) ProtoMessage()               {}
func (*Graph) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Graph) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Graph) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Graph) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type Error struct {
	Namespace   string            `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Code        int64             `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Stack       string            `protobuf:"bytes,4,opt,name=stack" json:"stack,omitempty"`
	Context     map[string]string `protobuf:"bytes,5,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Error) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Error) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Error) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

func (m *Error) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

type Message struct {
	headerLocker sync.Mutex

	Id     string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Header map[string]string `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Err    *Error            `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Message) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type Payload struct {
	Id           string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Timestamp    int64             `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	CurrentGraph string            `protobuf:"bytes,3,opt,name=current_graph,json=currentGraph" json:"current_graph,omitempty"`
	Graphs       map[string]*Graph `protobuf:"bytes,4,rep,name=graphs" json:"graphs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Context      map[string]string `protobuf:"bytes,5,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Message      *Message          `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Payload) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Payload) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Payload) GetCurrentGraph() string {
	if m != nil {
		return m.CurrentGraph
	}
	return ""
}

func (m *Payload) GetGraphs() map[string]*Graph {
	if m != nil {
		return m.Graphs
	}
	return nil
}

func (m *Payload) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Payload) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Port)(nil), "protocol.Port")
	proto.RegisterType((*Graph)(nil), "protocol.Graph")
	proto.RegisterType((*Error)(nil), "protocol.Error")
	proto.RegisterType((*Message)(nil), "protocol.Message")
	proto.RegisterType((*Payload)(nil), "protocol.Payload")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0x92, 0xa6, 0x5d, 0x2e, 0x5d, 0x01, 0x8b, 0x07, 0xab, 0xea, 0x50, 0x29, 0x20, 0x55,
	0x42, 0xca, 0xc3, 0xd0, 0x50, 0x19, 0x8f, 0x68, 0x02, 0x21, 0x0d, 0x4d, 0xe6, 0x07, 0x4c, 0x5e,
	0x62, 0x6d, 0xd1, 0x9a, 0x38, 0xd8, 0x2e, 0xa2, 0xaf, 0xfc, 0x32, 0xc4, 0x7f, 0xe1, 0x7f, 0x20,
	0x9f, 0x9d, 0x35, 0x6d, 0xf7, 0x32, 0xed, 0xed, 0x7c, 0xdf, 0x9d, 0xef, 0xfb, 0xbe, 0x3b, 0x18,
	0x35, 0x4a, 0x1a, 0x99, 0xcb, 0x65, 0x86, 0x01, 0x39, 0x68, 0xdf, 0xb3, 0x3f, 0x01, 0xf4, 0x2e,
	0xa4, 0x32, 0xe4, 0x29, 0x44, 0x5a, 0xfc, 0xa0, 0xc1, 0x34, 0x98, 0xc7, 0xcc, 0x86, 0x36, 0xb3,
	0x52, 0x4b, 0x1a, 0x4e, 0x83, 0x79, 0xc2, 0x6c, 0x48, 0x8e, 0x00, 0xae, 0x15, 0x6f, 0x6e, 0x2e,
	0x6b, 0x5e, 0x09, 0x1a, 0x21, 0x90, 0x60, 0xe6, 0x1b, 0xaf, 0x04, 0x59, 0xc0, 0x41, 0x25, 0x0c,
	0x2f, 0xb8, 0xe1, 0xb4, 0x37, 0x8d, 0xe6, 0xe9, 0xf1, 0x24, 0xbb, 0x1b, 0x6c, 0x87, 0x64, 0xe7,
	0x1e, 0x3e, 0xab, 0x8d, 0x5a, 0xb3, 0xbb, 0xea, 0xf1, 0x47, 0x38, 0xdc, 0x82, 0xec, 0xec, 0x5b,
	0xb1, 0x46, 0x36, 0x09, 0xb3, 0x21, 0x79, 0x0e, 0xf1, 0x4f, 0xbe, 0x5c, 0x09, 0xcf, 0xc7, 0x3d,
	0x4e, 0xc3, 0x45, 0x30, 0xfb, 0x0e, 0xf1, 0x67, 0xcb, 0x81, 0x10, 0xe8, 0x21, 0x31, 0xd7, 0x85,
	0x71, 0x2b, 0x2b, 0xdc, 0xc8, 0x7a, 0x0d, 0x71, 0x23, 0x95, 0xd1, 0x34, 0x42, 0x8a, 0xa3, 0x6d,
	0x8a, 0xcc, 0x81, 0xb3, 0x7f, 0x01, 0xc4, 0x67, 0x4a, 0x49, 0x45, 0x26, 0x90, 0xd8, 0x9f, 0x74,
	0xc3, 0xf3, 0xf6, 0xeb, 0x4d, 0xc2, 0xce, 0xcc, 0x65, 0xe1, 0x58, 0x45, 0x0c, 0x63, 0x32, 0x85,
	0xb4, 0x10, 0x3a, 0x57, 0x65, 0x63, 0x4a, 0x59, 0x7b, 0x9f, 0xba, 0x29, 0x2b, 0x46, 0x1b, 0x9e,
	0xdf, 0xd2, 0x9e, 0x13, 0x83, 0x0f, 0xf2, 0x1e, 0x06, 0xb9, 0xac, 0x8d, 0xf8, 0x65, 0x68, 0xbc,
	0x6b, 0x1f, 0x72, 0xc9, 0x3e, 0x39, 0xd8, 0xd9, 0xd7, 0x16, 0x8f, 0x4f, 0x61, 0xd8, 0x05, 0x1e,
	0x64, 0xde, 0xdf, 0x00, 0x06, 0xe7, 0x42, 0x6b, 0x7e, 0x2d, 0xc8, 0x08, 0xc2, 0xb2, 0xf0, 0x6d,
	0x61, 0x59, 0x90, 0x13, 0xe8, 0xdf, 0x08, 0x5e, 0x08, 0x45, 0x43, 0xa4, 0x73, 0xb4, 0xa1, 0xe3,
	0x5b, 0xb2, 0x2f, 0x88, 0x3b, 0x3e, 0xbe, 0xd8, 0x5a, 0x72, 0x25, 0x8b, 0x35, 0xea, 0x1e, 0x32,
	0x8c, 0xc9, 0x4b, 0x88, 0x84, 0x52, 0x28, 0x37, 0x3d, 0x7e, 0xb2, 0x23, 0x8b, 0x59, 0x6c, 0xfc,
	0x01, 0xd2, 0xce, 0x6f, 0x0f, 0x12, 0xf1, 0x3b, 0x82, 0xc1, 0x05, 0x5f, 0x2f, 0x25, 0x2f, 0xf6,
	0x44, 0x4c, 0x20, 0x31, 0x65, 0x25, 0xb4, 0xe1, 0x55, 0xe3, 0xb7, 0xb4, 0x49, 0x90, 0x57, 0x70,
	0x98, 0xaf, 0x94, 0x12, 0xb5, 0xb9, 0xc4, 0x3b, 0xf6, 0xcb, 0x1a, 0xfa, 0xa4, 0xbb, 0xab, 0x13,
	0xe8, 0x23, 0xa8, 0xfd, 0x55, 0x77, 0x7c, 0xf0, 0x53, 0x33, 0x2c, 0xd4, 0xde, 0x07, 0x57, 0x4c,
	0x16, 0xbb, 0xeb, 0x7c, 0xb1, 0xdf, 0x77, 0xef, 0x42, 0xc9, 0x5b, 0x18, 0x54, 0xce, 0x60, 0xda,
	0x47, 0xc7, 0x9e, 0xed, 0x39, 0xcf, 0xda, 0x8a, 0xf1, 0x57, 0x48, 0x3b, 0xd3, 0xef, 0xf1, 0xed,
	0x4d, 0xd7, 0xb7, 0x2d, 0xf7, 0xb1, 0xaf, 0x63, 0xe4, 0x63, 0x2e, 0xe9, 0xaa, 0x8f, 0xdf, 0xbe,
	0xfb, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x42, 0x29, 0x00, 0x02, 0x6c, 0x04, 0x00, 0x00,
}
