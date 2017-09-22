// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prompb.proto

/*
Package prompb is a generated protocol buffer package.

It is generated from these files:
	prompb.proto

It has these top-level messages:
	Label
	Sample
	TimeSeries
	LabelMatcher
	Query
	QueryResult
	ReadRequest
	ReadResponse
	WriteRequest
*/
package prompb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LabelMatcher_Type int32

const (
	LabelMatcher_EQ  LabelMatcher_Type = 0
	LabelMatcher_NEQ LabelMatcher_Type = 1
	LabelMatcher_RE  LabelMatcher_Type = 2
	LabelMatcher_NRE LabelMatcher_Type = 3
)

var LabelMatcher_Type_name = map[int32]string{
	0: "EQ",
	1: "NEQ",
	2: "RE",
	3: "NRE",
}
var LabelMatcher_Type_value = map[string]int32{
	"EQ":  0,
	"NEQ": 1,
	"RE":  2,
	"NRE": 3,
}

func (x LabelMatcher_Type) String() string {
	return proto.EnumName(LabelMatcher_Type_name, int32(x))
}
func (LabelMatcher_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Label struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Sample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Sample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sample) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

type TimeSeries struct {
	Labels  []*Label  `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type LabelMatcher struct {
	Type  LabelMatcher_Type `protobuf:"varint,1,opt,name=type,enum=prompb.LabelMatcher_Type" json:"type,omitempty"`
	Name  string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value string            `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()                    { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string            { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()               {}
func (*LabelMatcher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LabelMatcher) GetType() LabelMatcher_Type {
	if m != nil {
		return m.Type
	}
	return LabelMatcher_EQ
}

func (m *LabelMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Query struct {
	StartTimestampMs int64           `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64           `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs" json:"end_timestamp_ms,omitempty"`
	Matchers         []*LabelMatcher `protobuf:"bytes,3,rep,name=matchers" json:"matchers,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Query) GetStartTimestampMs() int64 {
	if m != nil {
		return m.StartTimestampMs
	}
	return 0
}

func (m *Query) GetEndTimestampMs() int64 {
	if m != nil {
		return m.EndTimestampMs
	}
	return 0
}

func (m *Query) GetMatchers() []*LabelMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

type QueryResult struct {
	TimeSeries []*TimeSeries `protobuf:"bytes,1,rep,name=time_series,json=timeSeries" json:"time_series,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryResult) GetTimeSeries() []*TimeSeries {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

type ReadRequest struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadRequest) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ReadResponse struct {
	// In same order as the request's queries.
	Results []*QueryResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type WriteRequest struct {
	TimeSeries []*TimeSeries `protobuf:"bytes,1,rep,name=time_series,json=timeSeries" json:"time_series,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *WriteRequest) GetTimeSeries() []*TimeSeries {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "prompb.Label")
	proto.RegisterType((*Sample)(nil), "prompb.Sample")
	proto.RegisterType((*TimeSeries)(nil), "prompb.TimeSeries")
	proto.RegisterType((*LabelMatcher)(nil), "prompb.LabelMatcher")
	proto.RegisterType((*Query)(nil), "prompb.Query")
	proto.RegisterType((*QueryResult)(nil), "prompb.QueryResult")
	proto.RegisterType((*ReadRequest)(nil), "prompb.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "prompb.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "prompb.WriteRequest")
	proto.RegisterEnum("prompb.LabelMatcher_Type", LabelMatcher_Type_name, LabelMatcher_Type_value)
}

func init() { proto.RegisterFile("prompb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0x4d, 0x9b, 0xc2, 0x34, 0x54, 0xd1, 0xb0, 0x87, 0x72, 0x5b, 0x2c, 0x21, 0x72, 0x60,
	0x2b, 0xd8, 0x95, 0xb8, 0x71, 0x00, 0xd4, 0x1b, 0x8b, 0x54, 0x6f, 0x25, 0x4e, 0x28, 0x72, 0xe9,
	0x48, 0x44, 0x8a, 0x13, 0xaf, 0xed, 0x20, 0xf5, 0x67, 0xf0, 0x8f, 0x51, 0x26, 0x1f, 0x4d, 0xa5,
	0x9e, 0xf6, 0x66, 0xcf, 0x7b, 0x6f, 0xe6, 0xf9, 0x8d, 0x0c, 0xb1, 0xb1, 0x95, 0x36, 0xfb, 0xb5,
	0xb1, 0x95, 0xaf, 0x30, 0x6a, 0x6f, 0xe2, 0x23, 0xcc, 0xbe, 0xab, 0x3d, 0x15, 0x88, 0x30, 0x2d,
	0x95, 0xa6, 0x55, 0x70, 0x1d, 0xa4, 0x2f, 0x24, 0x9f, 0xf1, 0x0a, 0x66, 0x7f, 0x55, 0x51, 0xd3,
	0x6a, 0xc2, 0xc5, 0xf6, 0x22, 0xbe, 0x40, 0xf4, 0xa0, 0xb4, 0x29, 0x46, 0x78, 0x23, 0x0a, 0x3a,
	0x1c, 0xdf, 0x40, 0xec, 0x73, 0x4d, 0xce, 0x2b, 0x6d, 0x32, 0xed, 0x58, 0x1c, 0xca, 0xc5, 0x50,
	0xbb, 0x77, 0xe2, 0x17, 0xc0, 0x2e, 0xd7, 0xf4, 0x40, 0x36, 0x27, 0x87, 0x6f, 0x21, 0x2a, 0x1a,
	0x0f, 0x6e, 0x15, 0x5c, 0x87, 0xe9, 0xe2, 0xf6, 0xe5, 0xba, 0xb3, 0xca, 0xce, 0x64, 0x07, 0x62,
	0x0a, 0x73, 0xc7, 0x73, 0x9b, 0x96, 0x0d, 0x6f, 0xd9, 0xf3, 0x5a, 0x3b, 0xb2, 0x87, 0xc5, 0xbf,
	0x00, 0x62, 0xd6, 0xde, 0x2b, 0xff, 0xfb, 0x0f, 0x59, 0xbc, 0x81, 0xa9, 0x3f, 0x9a, 0xd6, 0xe7,
	0xf2, 0xf6, 0xf5, 0x59, 0xff, 0x8e, 0xb3, 0xde, 0x1d, 0x0d, 0x49, 0xa6, 0x0d, 0x59, 0x4c, 0x2e,
	0x65, 0x11, 0x8e, 0xb3, 0x48, 0x61, 0xda, 0xe8, 0x30, 0x82, 0xc9, 0x66, 0x9b, 0x3c, 0xc3, 0x39,
	0x84, 0x3f, 0x36, 0xdb, 0x24, 0x68, 0x0a, 0x72, 0x93, 0x4c, 0xb8, 0x20, 0x37, 0x49, 0xd8, 0x78,
	0x9a, 0x6d, 0x6b, 0xb2, 0x47, 0x7c, 0x0f, 0xe8, 0xbc, 0xb2, 0x3e, 0x3b, 0x4b, 0x29, 0xe0, 0x94,
	0x12, 0x46, 0x76, 0xa7, 0xa8, 0x30, 0x85, 0x84, 0xca, 0x43, 0x76, 0x21, 0xd1, 0x25, 0x95, 0x87,
	0x31, 0xf3, 0x03, 0x3c, 0xd7, 0xed, 0x5b, 0xdc, 0x2a, 0xe4, 0x80, 0xae, 0x2e, 0x3d, 0x54, 0x0e,
	0x2c, 0xf1, 0x15, 0x16, 0x6c, 0x49, 0x92, 0xab, 0x0b, 0x8f, 0x77, 0xc0, 0x4b, 0xca, 0x1c, 0xaf,
	0xa5, 0x5b, 0x06, 0xf6, 0x3d, 0x4e, 0x0b, 0x93, 0xe0, 0x87, 0xb3, 0xf8, 0x04, 0x0b, 0x49, 0xea,
	0x20, 0xe9, 0xb1, 0x26, 0xe7, 0xf1, 0x1d, 0xcc, 0x1f, 0xeb, 0xb1, 0x7e, 0x58, 0x66, 0x3b, 0xa9,
	0x47, 0xc5, 0x67, 0x88, 0x5b, 0x9d, 0x33, 0x55, 0xe9, 0x08, 0x6f, 0x60, 0x6e, 0xd9, 0x46, 0x2f,
	0x7c, 0x75, 0x2e, 0x64, 0x4c, 0xf6, 0x1c, 0xf1, 0x0d, 0xe2, 0x9f, 0x36, 0xf7, 0xd4, 0xcf, 0x7d,
	0x8a, 0xf7, 0x7d, 0xc4, 0x7f, 0xe1, 0xee, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x31, 0x7c,
	0xb6, 0x1b, 0x03, 0x00, 0x00,
}
