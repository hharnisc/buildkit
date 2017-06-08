// Code generated by protoc-gen-gogo.
// source: control.proto
// DO NOT EDIT!

/*
	Package control is a generated protocol buffer package.

	It is generated from these files:
		control.proto

	It has these top-level messages:
		DiskUsageRequest
		DiskUsageResponse
		UsageRecord
		SolveRequest
		SolveResponse
		VertexStatus
*/
package control

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DiskUsageRequest struct {
}

func (m *DiskUsageRequest) Reset()                    { *m = DiskUsageRequest{} }
func (*DiskUsageRequest) ProtoMessage()               {}
func (*DiskUsageRequest) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{0} }

type DiskUsageResponse struct {
	Record []*UsageRecord `protobuf:"bytes,1,rep,name=record" json:"record,omitempty"`
}

func (m *DiskUsageResponse) Reset()                    { *m = DiskUsageResponse{} }
func (*DiskUsageResponse) ProtoMessage()               {}
func (*DiskUsageResponse) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{1} }

func (m *DiskUsageResponse) GetRecord() []*UsageRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type UsageRecord struct {
	ID      string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Mutable bool   `protobuf:"varint,2,opt,name=Mutable,proto3" json:"Mutable,omitempty"`
	InUse   bool   `protobuf:"varint,3,opt,name=InUse,proto3" json:"InUse,omitempty"`
	Size_   int64  `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (m *UsageRecord) Reset()                    { *m = UsageRecord{} }
func (*UsageRecord) ProtoMessage()               {}
func (*UsageRecord) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{2} }

func (m *UsageRecord) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UsageRecord) GetMutable() bool {
	if m != nil {
		return m.Mutable
	}
	return false
}

func (m *UsageRecord) GetInUse() bool {
	if m != nil {
		return m.InUse
	}
	return false
}

func (m *UsageRecord) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

type SolveRequest struct {
	Ref        string   `protobuf:"bytes,1,opt,name=Ref,proto3" json:"Ref,omitempty"`
	Definition [][]byte `protobuf:"bytes,2,rep,name=Definition" json:"Definition,omitempty"`
}

func (m *SolveRequest) Reset()                    { *m = SolveRequest{} }
func (*SolveRequest) ProtoMessage()               {}
func (*SolveRequest) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{3} }

func (m *SolveRequest) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *SolveRequest) GetDefinition() [][]byte {
	if m != nil {
		return m.Definition
	}
	return nil
}

type SolveResponse struct {
	Vertex []*VertexStatus `protobuf:"bytes,1,rep,name=vertex" json:"vertex,omitempty"`
}

func (m *SolveResponse) Reset()                    { *m = SolveResponse{} }
func (*SolveResponse) ProtoMessage()               {}
func (*SolveResponse) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{4} }

func (m *SolveResponse) GetVertex() []*VertexStatus {
	if m != nil {
		return m.Vertex
	}
	return nil
}

type VertexStatus struct {
}

func (m *VertexStatus) Reset()                    { *m = VertexStatus{} }
func (*VertexStatus) ProtoMessage()               {}
func (*VertexStatus) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{5} }

func init() {
	proto.RegisterType((*DiskUsageRequest)(nil), "control.DiskUsageRequest")
	proto.RegisterType((*DiskUsageResponse)(nil), "control.DiskUsageResponse")
	proto.RegisterType((*UsageRecord)(nil), "control.UsageRecord")
	proto.RegisterType((*SolveRequest)(nil), "control.SolveRequest")
	proto.RegisterType((*SolveResponse)(nil), "control.SolveResponse")
	proto.RegisterType((*VertexStatus)(nil), "control.VertexStatus")
}
func (this *DiskUsageRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DiskUsageRequest)
	if !ok {
		that2, ok := that.(DiskUsageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *DiskUsageResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DiskUsageResponse)
	if !ok {
		that2, ok := that.(DiskUsageResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Record) != len(that1.Record) {
		return false
	}
	for i := range this.Record {
		if !this.Record[i].Equal(that1.Record[i]) {
			return false
		}
	}
	return true
}
func (this *UsageRecord) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UsageRecord)
	if !ok {
		that2, ok := that.(UsageRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.Mutable != that1.Mutable {
		return false
	}
	if this.InUse != that1.InUse {
		return false
	}
	if this.Size_ != that1.Size_ {
		return false
	}
	return true
}
func (this *SolveRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SolveRequest)
	if !ok {
		that2, ok := that.(SolveRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Ref != that1.Ref {
		return false
	}
	if len(this.Definition) != len(that1.Definition) {
		return false
	}
	for i := range this.Definition {
		if !bytes.Equal(this.Definition[i], that1.Definition[i]) {
			return false
		}
	}
	return true
}
func (this *SolveResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SolveResponse)
	if !ok {
		that2, ok := that.(SolveResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Vertex) != len(that1.Vertex) {
		return false
	}
	for i := range this.Vertex {
		if !this.Vertex[i].Equal(that1.Vertex[i]) {
			return false
		}
	}
	return true
}
func (this *VertexStatus) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*VertexStatus)
	if !ok {
		that2, ok := that.(VertexStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *DiskUsageRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&control.DiskUsageRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DiskUsageResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&control.DiskUsageResponse{")
	if this.Record != nil {
		s = append(s, "Record: "+fmt.Sprintf("%#v", this.Record)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UsageRecord) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&control.UsageRecord{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Mutable: "+fmt.Sprintf("%#v", this.Mutable)+",\n")
	s = append(s, "InUse: "+fmt.Sprintf("%#v", this.InUse)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SolveRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&control.SolveRequest{")
	s = append(s, "Ref: "+fmt.Sprintf("%#v", this.Ref)+",\n")
	s = append(s, "Definition: "+fmt.Sprintf("%#v", this.Definition)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SolveResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&control.SolveResponse{")
	if this.Vertex != nil {
		s = append(s, "Vertex: "+fmt.Sprintf("%#v", this.Vertex)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VertexStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&control.VertexStatus{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringControl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Control service

type ControlClient interface {
	DiskUsage(ctx context.Context, in *DiskUsageRequest, opts ...grpc.CallOption) (*DiskUsageResponse, error)
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
}

type controlClient struct {
	cc *grpc.ClientConn
}

func NewControlClient(cc *grpc.ClientConn) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) DiskUsage(ctx context.Context, in *DiskUsageRequest, opts ...grpc.CallOption) (*DiskUsageResponse, error) {
	out := new(DiskUsageResponse)
	err := grpc.Invoke(ctx, "/control.Control/DiskUsage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := grpc.Invoke(ctx, "/control.Control/Solve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Control service

type ControlServer interface {
	DiskUsage(context.Context, *DiskUsageRequest) (*DiskUsageResponse, error)
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
}

func RegisterControlServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

func _Control_DiskUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).DiskUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.Control/DiskUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).DiskUsage(ctx, req.(*DiskUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.Control/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiskUsage",
			Handler:    _Control_DiskUsage_Handler,
		},
		{
			MethodName: "Solve",
			Handler:    _Control_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

func (m *DiskUsageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiskUsageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DiskUsageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiskUsageResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Record) > 0 {
		for _, msg := range m.Record {
			dAtA[i] = 0xa
			i++
			i = encodeVarintControl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *UsageRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UsageRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Mutable {
		dAtA[i] = 0x10
		i++
		if m.Mutable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.InUse {
		dAtA[i] = 0x18
		i++
		if m.InUse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Size_ != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.Size_))
	}
	return i, nil
}

func (m *SolveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SolveRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ref) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Ref)))
		i += copy(dAtA[i:], m.Ref)
	}
	if len(m.Definition) > 0 {
		for _, b := range m.Definition {
			dAtA[i] = 0x12
			i++
			i = encodeVarintControl(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *SolveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SolveResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Vertex) > 0 {
		for _, msg := range m.Vertex {
			dAtA[i] = 0xa
			i++
			i = encodeVarintControl(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *VertexStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VertexStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Control(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Control(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintControl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DiskUsageRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *DiskUsageResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Record) > 0 {
		for _, e := range m.Record {
			l = e.Size()
			n += 1 + l + sovControl(uint64(l))
		}
	}
	return n
}

func (m *UsageRecord) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	if m.Mutable {
		n += 2
	}
	if m.InUse {
		n += 2
	}
	if m.Size_ != 0 {
		n += 1 + sovControl(uint64(m.Size_))
	}
	return n
}

func (m *SolveRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	if len(m.Definition) > 0 {
		for _, b := range m.Definition {
			l = len(b)
			n += 1 + l + sovControl(uint64(l))
		}
	}
	return n
}

func (m *SolveResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Vertex) > 0 {
		for _, e := range m.Vertex {
			l = e.Size()
			n += 1 + l + sovControl(uint64(l))
		}
	}
	return n
}

func (m *VertexStatus) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovControl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozControl(x uint64) (n int) {
	return sovControl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DiskUsageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DiskUsageRequest{`,
		`}`,
	}, "")
	return s
}
func (this *DiskUsageResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DiskUsageResponse{`,
		`Record:` + strings.Replace(fmt.Sprintf("%v", this.Record), "UsageRecord", "UsageRecord", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UsageRecord) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UsageRecord{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Mutable:` + fmt.Sprintf("%v", this.Mutable) + `,`,
		`InUse:` + fmt.Sprintf("%v", this.InUse) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SolveRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SolveRequest{`,
		`Ref:` + fmt.Sprintf("%v", this.Ref) + `,`,
		`Definition:` + fmt.Sprintf("%v", this.Definition) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SolveResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SolveResponse{`,
		`Vertex:` + strings.Replace(fmt.Sprintf("%v", this.Vertex), "VertexStatus", "VertexStatus", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VertexStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VertexStatus{`,
		`}`,
	}, "")
	return s
}
func valueToStringControl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DiskUsageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DiskUsageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiskUsageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DiskUsageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DiskUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiskUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Record = append(m.Record, &UsageRecord{})
			if err := m.Record[len(m.Record)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UsageRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UsageRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UsageRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mutable = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InUse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InUse = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SolveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SolveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SolveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Definition", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Definition = append(m.Definition, make([]byte, postIndex-iNdEx))
			copy(m.Definition[len(m.Definition)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SolveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SolveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SolveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vertex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vertex = append(m.Vertex, &VertexStatus{})
			if err := m.Vertex[len(m.Vertex)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VertexStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VertexStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VertexStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipControl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthControl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowControl
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
				next, err := skipControl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthControl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControl   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("control.proto", fileDescriptorControl) }

var fileDescriptorControl = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xbb, 0x4e, 0x02, 0x41,
	0x14, 0x86, 0x77, 0x76, 0xb9, 0xc8, 0xe1, 0x12, 0x9c, 0xa0, 0x19, 0x29, 0x26, 0x9b, 0xad, 0xb6,
	0x40, 0x0a, 0x4c, 0x2c, 0x8d, 0xe2, 0x36, 0x14, 0x36, 0x43, 0xb0, 0x5f, 0x70, 0x30, 0x1b, 0xc9,
	0x0e, 0xee, 0x0c, 0xc4, 0x58, 0xd9, 0xd8, 0xfb, 0x18, 0x3e, 0x8a, 0x25, 0xa5, 0xa5, 0x8c, 0x8d,
	0x25, 0x8f, 0x60, 0x58, 0x06, 0xdc, 0x10, 0xbb, 0x73, 0xbe, 0x73, 0x99, 0xff, 0x3f, 0x19, 0xa8,
	0x8e, 0x44, 0xac, 0x12, 0x31, 0x69, 0x4f, 0x13, 0xa1, 0x04, 0x2e, 0x9a, 0xd4, 0xc3, 0x50, 0x0f,
	0x22, 0xf9, 0x30, 0x90, 0xe1, 0x3d, 0x67, 0xfc, 0x71, 0xc6, 0xa5, 0xf2, 0xae, 0xe0, 0x30, 0xc3,
	0xe4, 0x54, 0xc4, 0x92, 0xe3, 0x16, 0x14, 0x12, 0x3e, 0x12, 0xc9, 0x1d, 0x41, 0xae, 0xe3, 0x97,
	0x3b, 0x8d, 0xf6, 0x76, 0xa3, 0xe9, 0x5b, 0xd7, 0x98, 0xe9, 0xf1, 0x42, 0x28, 0x67, 0x30, 0xae,
	0x81, 0xdd, 0x0b, 0x08, 0x72, 0x91, 0x5f, 0x62, 0x76, 0x2f, 0xc0, 0x04, 0x8a, 0x37, 0x33, 0x15,
	0x0e, 0x27, 0x9c, 0xd8, 0x2e, 0xf2, 0x0f, 0xd8, 0x36, 0xc5, 0x0d, 0xc8, 0xf7, 0xe2, 0x81, 0xe4,
	0xc4, 0x49, 0xf9, 0x26, 0xc1, 0x18, 0x72, 0xfd, 0xe8, 0x99, 0x93, 0x9c, 0x8b, 0x7c, 0x87, 0xa5,
	0xb1, 0x77, 0x09, 0x95, 0xbe, 0x98, 0xcc, 0xb7, 0xaa, 0x71, 0x1d, 0x1c, 0xc6, 0xc7, 0xe6, 0x91,
	0x75, 0x88, 0x29, 0x40, 0xc0, 0xc7, 0x51, 0x1c, 0xa9, 0x48, 0xc4, 0xc4, 0x76, 0x1d, 0xbf, 0xc2,
	0x32, 0xc4, 0xbb, 0x80, 0xaa, 0xd9, 0x60, 0x3c, 0x9e, 0x42, 0x61, 0xce, 0x13, 0xc5, 0x9f, 0x8c,
	0xc7, 0xa3, 0x9d, 0xc7, 0xdb, 0x14, 0xf7, 0x55, 0xa8, 0x66, 0x92, 0x99, 0x26, 0xaf, 0x06, 0x95,
	0x2c, 0xef, 0xbc, 0x22, 0x28, 0x5e, 0x6f, 0x06, 0x70, 0x17, 0x4a, 0xbb, 0x1b, 0xe2, 0x93, 0xdd,
	0x9e, 0xfd, 0x5b, 0x37, 0x9b, 0xff, 0x95, 0x8c, 0x9c, 0x73, 0xc8, 0xa7, 0xfa, 0xf0, 0x9f, 0x8e,
	0xac, 0xe3, 0xe6, 0xf1, 0x3e, 0xde, 0xcc, 0x75, 0x5b, 0x8b, 0x25, 0xb5, 0x3e, 0x97, 0xd4, 0x5a,
	0x2d, 0x29, 0x7a, 0xd1, 0x14, 0xbd, 0x6b, 0x8a, 0x3e, 0x34, 0x45, 0x0b, 0x4d, 0xd1, 0x97, 0xa6,
	0xe8, 0x47, 0x53, 0x6b, 0xa5, 0x29, 0x7a, 0xfb, 0xa6, 0xd6, 0xb0, 0x90, 0xfe, 0x88, 0xb3, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0xe3, 0xfa, 0x6f, 0x22, 0x02, 0x00, 0x00,
}
