// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapper.proto

package dapperlib

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DataQueryResults struct {
	Results              []*DataQueryResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataQueryResults) Reset()         { *m = DataQueryResults{} }
func (m *DataQueryResults) String() string { return proto.CompactTextString(m) }
func (*DataQueryResults) ProtoMessage()    {}
func (*DataQueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{0}
}

func (m *DataQueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataQueryResults.Unmarshal(m, b)
}
func (m *DataQueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataQueryResults.Marshal(b, m, deterministic)
}
func (m *DataQueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataQueryResults.Merge(m, src)
}
func (m *DataQueryResults) XXX_Size() int {
	return xxx_messageInfo_DataQueryResults.Size(m)
}
func (m *DataQueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_DataQueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_DataQueryResults proto.InternalMessageInfo

func (m *DataQueryResults) GetResults() []*DataQueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type DataQueryResult struct {
	Domain               string             `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Key                  string             `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Field                string             `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	Records              []*DataQueryRecord `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataQueryResult) Reset()         { *m = DataQueryResult{} }
func (m *DataQueryResult) String() string { return proto.CompactTextString(m) }
func (*DataQueryResult) ProtoMessage()    {}
func (*DataQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{1}
}

func (m *DataQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataQueryResult.Unmarshal(m, b)
}
func (m *DataQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataQueryResult.Marshal(b, m, deterministic)
}
func (m *DataQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataQueryResult.Merge(m, src)
}
func (m *DataQueryResult) XXX_Size() int {
	return xxx_messageInfo_DataQueryResult.Size(m)
}
func (m *DataQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DataQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_DataQueryResult proto.InternalMessageInfo

func (m *DataQueryResult) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DataQueryResult) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataQueryResult) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *DataQueryResult) GetRecords() []*DataQueryRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type DataQueryRecord struct {
	// A unix timestamp representing when the value was recorded
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The value of the record (encoded as a string)
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataQueryRecord) Reset()         { *m = DataQueryRecord{} }
func (m *DataQueryRecord) String() string { return proto.CompactTextString(m) }
func (*DataQueryRecord) ProtoMessage()    {}
func (*DataQueryRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{2}
}

func (m *DataQueryRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataQueryRecord.Unmarshal(m, b)
}
func (m *DataQueryRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataQueryRecord.Marshal(b, m, deterministic)
}
func (m *DataQueryRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataQueryRecord.Merge(m, src)
}
func (m *DataQueryRecord) XXX_Size() int {
	return xxx_messageInfo_DataQueryRecord.Size(m)
}
func (m *DataQueryRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DataQueryRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DataQueryRecord proto.InternalMessageInfo

func (m *DataQueryRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DataQueryRecord) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyMetadataList struct {
	Metadata             []*KeyMetadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KeyMetadataList) Reset()         { *m = KeyMetadataList{} }
func (m *KeyMetadataList) String() string { return proto.CompactTextString(m) }
func (*KeyMetadataList) ProtoMessage()    {}
func (*KeyMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{3}
}

func (m *KeyMetadataList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyMetadataList.Unmarshal(m, b)
}
func (m *KeyMetadataList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyMetadataList.Marshal(b, m, deterministic)
}
func (m *KeyMetadataList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyMetadataList.Merge(m, src)
}
func (m *KeyMetadataList) XXX_Size() int {
	return xxx_messageInfo_KeyMetadataList.Size(m)
}
func (m *KeyMetadataList) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyMetadataList.DiscardUnknown(m)
}

var xxx_messageInfo_KeyMetadataList proto.InternalMessageInfo

func (m *KeyMetadataList) GetMetadata() []*KeyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type KeyMetadata struct {
	// The domain the metadata is associated with
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// The key the metadata is associated with
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The timespan(s) for which this key exists
	Span []*Timespan `protobuf:"bytes,3,rep,name=span,proto3" json:"span,omitempty"`
	// Name/Value pairs of metadata (e.g. 'model: MikroTik')
	Metadata map[string]*Metadata `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// String tags of metadata (e.g. 'LINZ')
	Tags     map[string]*Tag `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Location []*PointSpan    `protobuf:"bytes,6,rep,name=location,proto3" json:"location,omitempty"`
	// The relations to other keys from this metadata, the key of map is to_key
	Relations            map[string]*RelationSpans `protobuf:"bytes,7,rep,name=relations,proto3" json:"relations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KeyMetadata) Reset()         { *m = KeyMetadata{} }
func (m *KeyMetadata) String() string { return proto.CompactTextString(m) }
func (*KeyMetadata) ProtoMessage()    {}
func (*KeyMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{4}
}

func (m *KeyMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyMetadata.Unmarshal(m, b)
}
func (m *KeyMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyMetadata.Marshal(b, m, deterministic)
}
func (m *KeyMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyMetadata.Merge(m, src)
}
func (m *KeyMetadata) XXX_Size() int {
	return xxx_messageInfo_KeyMetadata.Size(m)
}
func (m *KeyMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_KeyMetadata proto.InternalMessageInfo

func (m *KeyMetadata) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *KeyMetadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyMetadata) GetSpan() []*Timespan {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *KeyMetadata) GetMetadata() map[string]*Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *KeyMetadata) GetTags() map[string]*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *KeyMetadata) GetLocation() []*PointSpan {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *KeyMetadata) GetRelations() map[string]*RelationSpans {
	if m != nil {
		return m.Relations
	}
	return nil
}

type Metadata struct {
	// The name of the metadata (e.g. 'model')
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value (and potentially past values)
	Values               []*MetadataValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{5}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetValues() []*MetadataValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type MetadataValue struct {
	// The value of the metadata (e.g. 'MikroTik')
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The timespan this value is valid for
	Span                 *Timespan `protobuf:"bytes,2,opt,name=span,proto3" json:"span,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MetadataValue) Reset()         { *m = MetadataValue{} }
func (m *MetadataValue) String() string { return proto.CompactTextString(m) }
func (*MetadataValue) ProtoMessage()    {}
func (*MetadataValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{6}
}

func (m *MetadataValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataValue.Unmarshal(m, b)
}
func (m *MetadataValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataValue.Marshal(b, m, deterministic)
}
func (m *MetadataValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataValue.Merge(m, src)
}
func (m *MetadataValue) XXX_Size() int {
	return xxx_messageInfo_MetadataValue.Size(m)
}
func (m *MetadataValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataValue.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataValue proto.InternalMessageInfo

func (m *MetadataValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *MetadataValue) GetSpan() *Timespan {
	if m != nil {
		return m.Span
	}
	return nil
}

type Tag struct {
	// The name of the tag (e.g. 'LINZ')
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The timespan this value is valid for
	Span                 []*Timespan `protobuf:"bytes,2,rep,name=span,proto3" json:"span,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{7}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetSpan() []*Timespan {
	if m != nil {
		return m.Span
	}
	return nil
}

type RelationSpans struct {
	// The relations to other keys
	Spans                []*RelationSpan `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RelationSpans) Reset()         { *m = RelationSpans{} }
func (m *RelationSpans) String() string { return proto.CompactTextString(m) }
func (*RelationSpans) ProtoMessage()    {}
func (*RelationSpans) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{8}
}

func (m *RelationSpans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationSpans.Unmarshal(m, b)
}
func (m *RelationSpans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationSpans.Marshal(b, m, deterministic)
}
func (m *RelationSpans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationSpans.Merge(m, src)
}
func (m *RelationSpans) XXX_Size() int {
	return xxx_messageInfo_RelationSpans.Size(m)
}
func (m *RelationSpans) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationSpans.DiscardUnknown(m)
}

var xxx_messageInfo_RelationSpans proto.InternalMessageInfo

func (m *RelationSpans) GetSpans() []*RelationSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

type RelationSpan struct {
	// The type of the relation
	RelType string `protobuf:"bytes,1,opt,name=rel_type,json=relType,proto3" json:"rel_type,omitempty"`
	// The timespan this value is valid for
	Span                 *Timespan `protobuf:"bytes,2,opt,name=span,proto3" json:"span,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RelationSpan) Reset()         { *m = RelationSpan{} }
func (m *RelationSpan) String() string { return proto.CompactTextString(m) }
func (*RelationSpan) ProtoMessage()    {}
func (*RelationSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{9}
}

func (m *RelationSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationSpan.Unmarshal(m, b)
}
func (m *RelationSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationSpan.Marshal(b, m, deterministic)
}
func (m *RelationSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationSpan.Merge(m, src)
}
func (m *RelationSpan) XXX_Size() int {
	return xxx_messageInfo_RelationSpan.Size(m)
}
func (m *RelationSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationSpan.DiscardUnknown(m)
}

var xxx_messageInfo_RelationSpan proto.InternalMessageInfo

func (m *RelationSpan) GetRelType() string {
	if m != nil {
		return m.RelType
	}
	return ""
}

func (m *RelationSpan) GetSpan() *Timespan {
	if m != nil {
		return m.Span
	}
	return nil
}

type SnapshotRelation struct {
	FromKey              string   `protobuf:"bytes,1,opt,name=from_key,json=fromKey,proto3" json:"from_key,omitempty"`
	ToKey                string   `protobuf:"bytes,2,opt,name=to_key,json=toKey,proto3" json:"to_key,omitempty"`
	RelType              string   `protobuf:"bytes,3,opt,name=rel_type,json=relType,proto3" json:"rel_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotRelation) Reset()         { *m = SnapshotRelation{} }
func (m *SnapshotRelation) String() string { return proto.CompactTextString(m) }
func (*SnapshotRelation) ProtoMessage()    {}
func (*SnapshotRelation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{10}
}

func (m *SnapshotRelation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRelation.Unmarshal(m, b)
}
func (m *SnapshotRelation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRelation.Marshal(b, m, deterministic)
}
func (m *SnapshotRelation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRelation.Merge(m, src)
}
func (m *SnapshotRelation) XXX_Size() int {
	return xxx_messageInfo_SnapshotRelation.Size(m)
}
func (m *SnapshotRelation) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRelation.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRelation proto.InternalMessageInfo

func (m *SnapshotRelation) GetFromKey() string {
	if m != nil {
		return m.FromKey
	}
	return ""
}

func (m *SnapshotRelation) GetToKey() string {
	if m != nil {
		return m.ToKey
	}
	return ""
}

func (m *SnapshotRelation) GetRelType() string {
	if m != nil {
		return m.RelType
	}
	return ""
}

type PointSpan struct {
	//The Lat/Lon of the point
	Location *Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	//The timespan the value is valid for
	Span                 *Timespan `protobuf:"bytes,2,opt,name=span,proto3" json:"span,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PointSpan) Reset()         { *m = PointSpan{} }
func (m *PointSpan) String() string { return proto.CompactTextString(m) }
func (*PointSpan) ProtoMessage()    {}
func (*PointSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{11}
}

func (m *PointSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointSpan.Unmarshal(m, b)
}
func (m *PointSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointSpan.Marshal(b, m, deterministic)
}
func (m *PointSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointSpan.Merge(m, src)
}
func (m *PointSpan) XXX_Size() int {
	return xxx_messageInfo_PointSpan.Size(m)
}
func (m *PointSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_PointSpan.DiscardUnknown(m)
}

var xxx_messageInfo_PointSpan proto.InternalMessageInfo

func (m *PointSpan) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *PointSpan) GetSpan() *Timespan {
	if m != nil {
		return m.Span
	}
	return nil
}

type KeyMetadataSnapshotList struct {
	Metadata             []*KeyMetadataSnapshot `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *KeyMetadataSnapshotList) Reset()         { *m = KeyMetadataSnapshotList{} }
func (m *KeyMetadataSnapshotList) String() string { return proto.CompactTextString(m) }
func (*KeyMetadataSnapshotList) ProtoMessage()    {}
func (*KeyMetadataSnapshotList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{12}
}

func (m *KeyMetadataSnapshotList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyMetadataSnapshotList.Unmarshal(m, b)
}
func (m *KeyMetadataSnapshotList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyMetadataSnapshotList.Marshal(b, m, deterministic)
}
func (m *KeyMetadataSnapshotList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyMetadataSnapshotList.Merge(m, src)
}
func (m *KeyMetadataSnapshotList) XXX_Size() int {
	return xxx_messageInfo_KeyMetadataSnapshotList.Size(m)
}
func (m *KeyMetadataSnapshotList) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyMetadataSnapshotList.DiscardUnknown(m)
}

var xxx_messageInfo_KeyMetadataSnapshotList proto.InternalMessageInfo

func (m *KeyMetadataSnapshotList) GetMetadata() []*KeyMetadataSnapshot {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// A simplified version of KeyMetadata without Timespan components (snapshot of metadata at one moment in time)
type KeyMetadataSnapshot struct {
	// The domain the metadata is associated with
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// The key the metadata is associated with
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The moment of the metadata snapshot
	Moment int64 `protobuf:"varint,3,opt,name=moment,proto3" json:"moment,omitempty"`
	// Name/Value pairs of metadata (e.g. 'model: MikroTik')
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// String tags of metadata (e.g. 'LINZ')
	Tags     []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Location *Point   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	// The relation type of to_key
	Relations            []*SnapshotRelation `protobuf:"bytes,7,rep,name=relations,proto3" json:"relations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KeyMetadataSnapshot) Reset()         { *m = KeyMetadataSnapshot{} }
func (m *KeyMetadataSnapshot) String() string { return proto.CompactTextString(m) }
func (*KeyMetadataSnapshot) ProtoMessage()    {}
func (*KeyMetadataSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{13}
}

func (m *KeyMetadataSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyMetadataSnapshot.Unmarshal(m, b)
}
func (m *KeyMetadataSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyMetadataSnapshot.Marshal(b, m, deterministic)
}
func (m *KeyMetadataSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyMetadataSnapshot.Merge(m, src)
}
func (m *KeyMetadataSnapshot) XXX_Size() int {
	return xxx_messageInfo_KeyMetadataSnapshot.Size(m)
}
func (m *KeyMetadataSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyMetadataSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_KeyMetadataSnapshot proto.InternalMessageInfo

func (m *KeyMetadataSnapshot) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *KeyMetadataSnapshot) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyMetadataSnapshot) GetMoment() int64 {
	if m != nil {
		return m.Moment
	}
	return 0
}

func (m *KeyMetadataSnapshot) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *KeyMetadataSnapshot) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *KeyMetadataSnapshot) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *KeyMetadataSnapshot) GetRelations() []*SnapshotRelation {
	if m != nil {
		return m.Relations
	}
	return nil
}

type DomainMetadataList struct {
	// The domain being listed
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All keys in the domain
	Keys []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	// All possible metadata fields & values in the domain
	Metadata map[string]*MetadataValuesList `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// All possible tags in the domain
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DomainMetadataList) Reset()         { *m = DomainMetadataList{} }
func (m *DomainMetadataList) String() string { return proto.CompactTextString(m) }
func (*DomainMetadataList) ProtoMessage()    {}
func (*DomainMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{14}
}

func (m *DomainMetadataList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainMetadataList.Unmarshal(m, b)
}
func (m *DomainMetadataList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainMetadataList.Marshal(b, m, deterministic)
}
func (m *DomainMetadataList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainMetadataList.Merge(m, src)
}
func (m *DomainMetadataList) XXX_Size() int {
	return xxx_messageInfo_DomainMetadataList.Size(m)
}
func (m *DomainMetadataList) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainMetadataList.DiscardUnknown(m)
}

var xxx_messageInfo_DomainMetadataList proto.InternalMessageInfo

func (m *DomainMetadataList) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DomainMetadataList) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *DomainMetadataList) GetMetadata() map[string]*MetadataValuesList {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DomainMetadataList) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type MetadataValuesList struct {
	// The metadata field name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// All potential values for that field
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataValuesList) Reset()         { *m = MetadataValuesList{} }
func (m *MetadataValuesList) String() string { return proto.CompactTextString(m) }
func (*MetadataValuesList) ProtoMessage()    {}
func (*MetadataValuesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{15}
}

func (m *MetadataValuesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataValuesList.Unmarshal(m, b)
}
func (m *MetadataValuesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataValuesList.Marshal(b, m, deterministic)
}
func (m *MetadataValuesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataValuesList.Merge(m, src)
}
func (m *MetadataValuesList) XXX_Size() int {
	return xxx_messageInfo_MetadataValuesList.Size(m)
}
func (m *MetadataValuesList) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataValuesList.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataValuesList proto.InternalMessageInfo

func (m *MetadataValuesList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetadataValuesList) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Timespan struct {
	Start                int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timespan) Reset()         { *m = Timespan{} }
func (m *Timespan) String() string { return proto.CompactTextString(m) }
func (*Timespan) ProtoMessage()    {}
func (*Timespan) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{16}
}

func (m *Timespan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timespan.Unmarshal(m, b)
}
func (m *Timespan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timespan.Marshal(b, m, deterministic)
}
func (m *Timespan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timespan.Merge(m, src)
}
func (m *Timespan) XXX_Size() int {
	return xxx_messageInfo_Timespan.Size(m)
}
func (m *Timespan) XXX_DiscardUnknown() {
	xxx_messageInfo_Timespan.DiscardUnknown(m)
}

var xxx_messageInfo_Timespan proto.InternalMessageInfo

func (m *Timespan) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Timespan) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type Point struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd1f6fc9a1999a7, []int{17}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*DataQueryResults)(nil), "dapper.DataQueryResults")
	proto.RegisterType((*DataQueryResult)(nil), "dapper.DataQueryResult")
	proto.RegisterType((*DataQueryRecord)(nil), "dapper.DataQueryRecord")
	proto.RegisterType((*KeyMetadataList)(nil), "dapper.KeyMetadataList")
	proto.RegisterType((*KeyMetadata)(nil), "dapper.KeyMetadata")
	proto.RegisterMapType((map[string]*Metadata)(nil), "dapper.KeyMetadata.MetadataEntry")
	proto.RegisterMapType((map[string]*RelationSpans)(nil), "dapper.KeyMetadata.RelationsEntry")
	proto.RegisterMapType((map[string]*Tag)(nil), "dapper.KeyMetadata.TagsEntry")
	proto.RegisterType((*Metadata)(nil), "dapper.Metadata")
	proto.RegisterType((*MetadataValue)(nil), "dapper.MetadataValue")
	proto.RegisterType((*Tag)(nil), "dapper.Tag")
	proto.RegisterType((*RelationSpans)(nil), "dapper.RelationSpans")
	proto.RegisterType((*RelationSpan)(nil), "dapper.RelationSpan")
	proto.RegisterType((*SnapshotRelation)(nil), "dapper.SnapshotRelation")
	proto.RegisterType((*PointSpan)(nil), "dapper.PointSpan")
	proto.RegisterType((*KeyMetadataSnapshotList)(nil), "dapper.KeyMetadataSnapshotList")
	proto.RegisterType((*KeyMetadataSnapshot)(nil), "dapper.KeyMetadataSnapshot")
	proto.RegisterMapType((map[string]string)(nil), "dapper.KeyMetadataSnapshot.MetadataEntry")
	proto.RegisterType((*DomainMetadataList)(nil), "dapper.DomainMetadataList")
	proto.RegisterMapType((map[string]*MetadataValuesList)(nil), "dapper.DomainMetadataList.MetadataEntry")
	proto.RegisterType((*MetadataValuesList)(nil), "dapper.MetadataValuesList")
	proto.RegisterType((*Timespan)(nil), "dapper.Timespan")
	proto.RegisterType((*Point)(nil), "dapper.Point")
}

func init() { proto.RegisterFile("dapper.proto", fileDescriptor_1dd1f6fc9a1999a7) }

var fileDescriptor_1dd1f6fc9a1999a7 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdb, 0x6e, 0xd4, 0x3c,
	0x10, 0x56, 0x36, 0x7b, 0xca, 0x6c, 0xfb, 0x77, 0x7f, 0xf7, 0x14, 0x16, 0x90, 0xda, 0x08, 0xa1,
	0x16, 0xd4, 0x42, 0x8b, 0x04, 0x88, 0x0a, 0x51, 0xaa, 0xdd, 0xab, 0x52, 0x01, 0xee, 0x0a, 0x24,
	0x84, 0xb4, 0x72, 0xbb, 0xee, 0x12, 0x35, 0x27, 0x25, 0x5e, 0xa4, 0xdc, 0xf3, 0x34, 0x3c, 0x0e,
	0xaf, 0xc2, 0x0b, 0x20, 0xdb, 0x71, 0xe2, 0xec, 0xa6, 0xa7, 0x3b, 0x8f, 0x67, 0x3c, 0xfe, 0xe6,
	0xcb, 0xcc, 0xe7, 0xc0, 0xc2, 0x98, 0x44, 0x11, 0x8d, 0x77, 0xa3, 0x38, 0x64, 0x21, 0x6a, 0x4a,
	0xcb, 0x19, 0x40, 0xb7, 0x4f, 0x18, 0xf9, 0x3c, 0xa5, 0x71, 0x8a, 0x69, 0x32, 0xf5, 0x58, 0x82,
	0xf6, 0xa0, 0x15, 0xcb, 0xa5, 0x6d, 0x6c, 0x98, 0x5b, 0x9d, 0xfd, 0xf5, 0xdd, 0xec, 0xec, 0x4c,
	0x28, 0x56, 0x71, 0xce, 0x2f, 0x03, 0x96, 0x66, 0x9c, 0x68, 0x0d, 0x9a, 0xe3, 0xd0, 0x27, 0x6e,
	0x60, 0x1b, 0x1b, 0xc6, 0x96, 0x85, 0x33, 0x0b, 0x75, 0xc1, 0xbc, 0xa4, 0xa9, 0x5d, 0x13, 0x9b,
	0x7c, 0x89, 0x56, 0xa0, 0x71, 0xe1, 0x52, 0x6f, 0x6c, 0x9b, 0x62, 0x4f, 0x1a, 0x12, 0xc6, 0x79,
	0x18, 0x8f, 0x13, 0xbb, 0x7e, 0x25, 0x0c, 0xee, 0xc7, 0x2a, 0xce, 0x19, 0x94, 0x50, 0xf0, 0x3d,
	0xf4, 0x00, 0x2c, 0xe6, 0xfa, 0x34, 0x61, 0xc4, 0x8f, 0x04, 0x10, 0x13, 0x17, 0x1b, 0xfc, 0xe6,
	0x9f, 0xc4, 0x9b, 0xd2, 0x0c, 0x8d, 0x34, 0x9c, 0x23, 0x58, 0x3a, 0xa6, 0xe9, 0x09, 0x65, 0x64,
	0x4c, 0x18, 0xf9, 0xe0, 0x26, 0x0c, 0x3d, 0x83, 0xb6, 0x9f, 0xd9, 0x19, 0x29, 0xcb, 0x0a, 0x8d,
	0x16, 0x8a, 0xf3, 0x20, 0xe7, 0x77, 0x1d, 0x3a, 0x9a, 0xe7, 0x0e, 0x6c, 0x3c, 0x82, 0x7a, 0x12,
	0x91, 0xc0, 0x36, 0xc5, 0x35, 0x5d, 0x75, 0xcd, 0x90, 0x83, 0x8e, 0x48, 0x80, 0x85, 0x17, 0xbd,
	0xd5, 0x00, 0x49, 0x7a, 0x36, 0x2b, 0x00, 0xed, 0xaa, 0xc5, 0x20, 0x60, 0x71, 0x5a, 0xc0, 0x43,
	0x7b, 0x50, 0x67, 0x64, 0x92, 0xd8, 0x0d, 0x71, 0xf4, 0x61, 0xd5, 0xd1, 0x21, 0x99, 0x24, 0xf2,
	0x98, 0x08, 0x45, 0x3b, 0xd0, 0xf6, 0xc2, 0x73, 0xc2, 0xdc, 0x30, 0xb0, 0x9b, 0xe2, 0xd8, 0xff,
	0xea, 0xd8, 0xa7, 0xd0, 0x0d, 0xd8, 0x29, 0x07, 0x97, 0x87, 0xa0, 0x43, 0xb0, 0x62, 0xea, 0x89,
	0x75, 0x62, 0xb7, 0x44, 0xbc, 0x53, 0x75, 0x0d, 0x56, 0x41, 0xf2, 0xae, 0xe2, 0x50, 0xef, 0x04,
	0x16, 0x4b, 0xf0, 0x15, 0x57, 0x46, 0xc1, 0xd5, 0x63, 0xfd, 0xfb, 0x69, 0x64, 0xe5, 0x1f, 0x44,
	0xba, 0xdf, 0xd4, 0x5e, 0x1b, 0xbd, 0x3e, 0x58, 0x79, 0x49, 0x15, 0xa9, 0x36, 0xcb, 0xa9, 0x3a,
	0x39, 0xef, 0x64, 0xa2, 0x67, 0x39, 0x85, 0xff, 0xca, 0x88, 0x2b, 0x52, 0x3d, 0x2d, 0xa7, 0x5a,
	0x55, 0xa9, 0xd4, 0x41, 0xce, 0x54, 0xa2, 0x25, 0x75, 0x4e, 0xa0, 0x9d, 0x37, 0x0a, 0x82, 0x7a,
	0x40, 0x7c, 0x9a, 0xe5, 0x13, 0x6b, 0xb4, 0x03, 0x4d, 0x11, 0x9c, 0xd8, 0x35, 0x41, 0xe4, 0xea,
	0x6c, 0x9d, 0x5f, 0xb8, 0x17, 0x67, 0x41, 0xce, 0x71, 0x41, 0x9c, 0x70, 0x14, 0x6d, 0x6e, 0x68,
	0x6d, 0x9e, 0x37, 0xda, 0x0c, 0x77, 0xe5, 0x46, 0x73, 0xde, 0x81, 0x39, 0x24, 0x93, 0x4a, 0x58,
	0x45, 0x82, 0x6b, 0x3a, 0xd5, 0x39, 0x80, 0xc5, 0x52, 0xe1, 0xe8, 0x09, 0x34, 0xb8, 0x43, 0xa9,
	0xcb, 0x4a, 0x15, 0x3d, 0x58, 0x86, 0x38, 0x1f, 0x61, 0x41, 0xdf, 0x46, 0xf7, 0xa0, 0x1d, 0x53,
	0x6f, 0xc4, 0xd2, 0x48, 0x41, 0x69, 0xc5, 0xd4, 0x1b, 0xa6, 0xd1, 0x6d, 0xcb, 0x19, 0x41, 0xf7,
	0x34, 0x20, 0x51, 0xf2, 0x23, 0x64, 0x2a, 0x31, 0x4f, 0x7a, 0x11, 0x87, 0xfe, 0xa8, 0xf8, 0x8c,
	0x2d, 0x6e, 0x1f, 0xd3, 0x14, 0xad, 0x42, 0x93, 0x85, 0xa3, 0x62, 0x42, 0x1b, 0x2c, 0xe4, 0xdb,
	0x3a, 0x0c, 0xb3, 0x04, 0xc3, 0xf9, 0x0e, 0x56, 0x3e, 0x0e, 0x68, 0x5b, 0x9b, 0x19, 0x43, 0xe0,
	0x5a, 0x2c, 0xcd, 0x8c, 0x36, 0x2f, 0xb7, 0x83, 0x8f, 0x61, 0x5d, 0x1b, 0x1e, 0x55, 0x89, 0x90,
	0xa8, 0x57, 0x73, 0x12, 0x75, 0xbf, 0x62, 0xde, 0xf2, 0xe2, 0x0b, 0xa9, 0xfa, 0x53, 0x83, 0xe5,
	0x8a, 0x88, 0x3b, 0x48, 0xd6, 0x1a, 0x34, 0xfd, 0xd0, 0xa7, 0x01, 0x13, 0x64, 0x98, 0x38, 0xb3,
	0xd0, 0x60, 0x4e, 0xa4, 0xb6, 0xaf, 0x81, 0x74, 0xa5, 0x58, 0x21, 0x4d, 0xac, 0xac, 0x4c, 0x8d,
	0xb6, 0x4b, 0x6a, 0x74, 0x2d, 0xb3, 0x2f, 0xe7, 0x95, 0xc8, 0x56, 0xb1, 0xb3, 0xbd, 0xa0, 0xeb,
	0xcf, 0xc1, 0xcd, 0xfa, 0x53, 0xf9, 0x7e, 0x88, 0x91, 0xfe, 0x6b, 0x00, 0xea, 0x0b, 0xbe, 0x4a,
	0xef, 0xc8, 0x55, 0x9c, 0x22, 0xa8, 0x5f, 0xd2, 0x54, 0xce, 0xb7, 0x85, 0xc5, 0x1a, 0xf5, 0x35,
	0xf6, 0xe4, 0x63, 0xb0, 0x95, 0xbf, 0x80, 0x73, 0x99, 0x6f, 0x24, 0xaf, 0x5e, 0x90, 0xd7, 0xfb,
	0x7a, 0x73, 0x65, 0xcf, 0xcb, 0x1a, 0xd6, 0xab, 0x54, 0x9c, 0x84, 0xdf, 0xac, 0x57, 0x7d, 0x08,
	0x68, 0x3e, 0xa0, 0x52, 0x3b, 0xd6, 0x4a, 0x92, 0x66, 0xe5, 0xda, 0xb5, 0x0f, 0x6d, 0xd5, 0xf2,
	0x9c, 0xdd, 0x84, 0x91, 0x98, 0x65, 0xef, 0xb6, 0x34, 0x38, 0x56, 0x1a, 0x8c, 0x05, 0x2e, 0x13,
	0xf3, 0xa5, 0xf3, 0x1e, 0x1a, 0xe2, 0x9b, 0xa3, 0x1e, 0xb4, 0xf9, 0xc7, 0x63, 0xd3, 0xb1, 0xbc,
	0xac, 0x86, 0x73, 0x9b, 0xff, 0x08, 0x78, 0x61, 0x30, 0x91, 0xce, 0x9a, 0x70, 0x16, 0x1b, 0x47,
	0x9d, 0x6f, 0x96, 0x2c, 0xd0, 0x73, 0xcf, 0xce, 0x9a, 0xe2, 0x1f, 0xe9, 0xc5, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0x97, 0x79, 0xa5, 0x33, 0x09, 0x00, 0x00,
}