// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: tuple.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Operation
type ExpandTreeNode_Operation int32

const (
	ExpandTreeNode_INVALID      ExpandTreeNode_Operation = 0
	ExpandTreeNode_UNION        ExpandTreeNode_Operation = 1
	ExpandTreeNode_INTERSECTION ExpandTreeNode_Operation = 2
	ExpandTreeNode_ROOT         ExpandTreeNode_Operation = 3
)

// Enum value maps for ExpandTreeNode_Operation.
var (
	ExpandTreeNode_Operation_name = map[int32]string{
		0: "INVALID",
		1: "UNION",
		2: "INTERSECTION",
		3: "ROOT",
	}
	ExpandTreeNode_Operation_value = map[string]int32{
		"INVALID":      0,
		"UNION":        1,
		"INTERSECTION": 2,
		"ROOT":         3,
	}
)

func (x ExpandTreeNode_Operation) Enum() *ExpandTreeNode_Operation {
	p := new(ExpandTreeNode_Operation)
	*p = x
	return p
}

func (x ExpandTreeNode_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpandTreeNode_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_tuple_proto_enumTypes[0].Descriptor()
}

func (ExpandTreeNode_Operation) Type() protoreflect.EnumType {
	return &file_tuple_proto_enumTypes[0]
}

func (x ExpandTreeNode_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpandTreeNode_Operation.Descriptor instead.
func (ExpandTreeNode_Operation) EnumDescriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{7, 0}
}

// Tuple
type Tuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relation string   `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Subject  *Subject `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *Tuple) Reset() {
	*x = Tuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuple) ProtoMessage() {}

func (x *Tuple) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuple.ProtoReflect.Descriptor instead.
func (*Tuple) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{0}
}

func (x *Tuple) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Tuple) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *Tuple) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

// Entity
type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{1}
}

func (x *Entity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Entity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EntityAndRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relation string  `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *EntityAndRelation) Reset() {
	*x = EntityAndRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityAndRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityAndRelation) ProtoMessage() {}

func (x *EntityAndRelation) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityAndRelation.ProtoReflect.Descriptor instead.
func (*EntityAndRelation) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{2}
}

func (x *EntityAndRelation) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *EntityAndRelation) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

// Subject
type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{3}
}

func (x *Subject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Subject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subject) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

type TupleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *EntityFilter  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relation string         `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Subject  *SubjectFilter `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *TupleFilter) Reset() {
	*x = TupleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TupleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TupleFilter) ProtoMessage() {}

func (x *TupleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TupleFilter.ProtoReflect.Descriptor instead.
func (*TupleFilter) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{4}
}

func (x *TupleFilter) GetEntity() *EntityFilter {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TupleFilter) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *TupleFilter) GetSubject() *SubjectFilter {
	if x != nil {
		return x.Subject
	}
	return nil
}

type EntityFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EntityFilter) Reset() {
	*x = EntityFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityFilter) ProtoMessage() {}

func (x *EntityFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityFilter.ProtoReflect.Descriptor instead.
func (*EntityFilter) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{5}
}

func (x *EntityFilter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EntityFilter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Subject
type SubjectFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *SubjectFilter) Reset() {
	*x = SubjectFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectFilter) ProtoMessage() {}

func (x *SubjectFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectFilter.ProtoReflect.Descriptor instead.
func (*SubjectFilter) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{6}
}

func (x *SubjectFilter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SubjectFilter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubjectFilter) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

// ExpandTreeNode
type ExpandTreeNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation ExpandTreeNode_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=ExpandTreeNode_Operation" json:"operation,omitempty"`
	Children  []*Expand                `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ExpandTreeNode) Reset() {
	*x = ExpandTreeNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandTreeNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandTreeNode) ProtoMessage() {}

func (x *ExpandTreeNode) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandTreeNode.ProtoReflect.Descriptor instead.
func (*ExpandTreeNode) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{7}
}

func (x *ExpandTreeNode) GetOperation() ExpandTreeNode_Operation {
	if x != nil {
		return x.Operation
	}
	return ExpandTreeNode_INVALID
}

func (x *ExpandTreeNode) GetChildren() []*Expand {
	if x != nil {
		return x.Children
	}
	return nil
}

// Expand
type Expand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *EntityAndRelation `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Types that are assignable to Node:
	//	*Expand_Expand
	//	*Expand_Leaf
	Node isExpand_Node `protobuf_oneof:"node"`
}

func (x *Expand) Reset() {
	*x = Expand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expand) ProtoMessage() {}

func (x *Expand) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expand.ProtoReflect.Descriptor instead.
func (*Expand) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{8}
}

func (x *Expand) GetTarget() *EntityAndRelation {
	if x != nil {
		return x.Target
	}
	return nil
}

func (m *Expand) GetNode() isExpand_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (x *Expand) GetExpand() *ExpandTreeNode {
	if x, ok := x.GetNode().(*Expand_Expand); ok {
		return x.Expand
	}
	return nil
}

func (x *Expand) GetLeaf() *Subjects {
	if x, ok := x.GetNode().(*Expand_Leaf); ok {
		return x.Leaf
	}
	return nil
}

type isExpand_Node interface {
	isExpand_Node()
}

type Expand_Expand struct {
	Expand *ExpandTreeNode `protobuf:"bytes,2,opt,name=expand,proto3,oneof"`
}

type Expand_Leaf struct {
	Leaf *Subjects `protobuf:"bytes,3,opt,name=leaf,proto3,oneof"`
}

func (*Expand_Expand) isExpand_Node() {}

func (*Expand_Leaf) isExpand_Node() {}

// Subjects
type Subjects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exclusion bool       `protobuf:"varint,1,opt,name=exclusion,proto3" json:"exclusion,omitempty"`
	Subjects  []*Subject `protobuf:"bytes,2,rep,name=subjects,proto3" json:"subjects,omitempty"`
}

func (x *Subjects) Reset() {
	*x = Subjects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuple_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subjects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subjects) ProtoMessage() {}

func (x *Subjects) ProtoReflect() protoreflect.Message {
	mi := &file_tuple_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subjects.ProtoReflect.Descriptor instead.
func (*Subjects) Descriptor() ([]byte, []int) {
	return file_tuple_proto_rawDescGZIP(), []int{9}
}

func (x *Subjects) GetExclusion() bool {
	if x != nil {
		return x.Exclusion
	}
	return false
}

func (x *Subjects) GetSubjects() []*Subject {
	if x != nil {
		return x.Subjects
	}
	return nil
}

var File_tuple_proto protoreflect.FileDescriptor

var file_tuple_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x54, 0x75, 0x70, 0x6c, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa,
	0x42, 0x26, 0x72, 0x24, 0x28, 0x40, 0x32, 0x20, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x8e, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x42, 0x26, 0x72, 0x24,
	0x28, 0x40, 0x32, 0x20, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x29, 0x24, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xfa, 0x42, 0x32, 0x72, 0x30, 0x28, 0x80, 0x01,
	0x32, 0x2b, 0x5e, 0x28, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f,
	0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2f, 0x5f, 0x7c, 0x2d, 0x5d,
	0x7b, 0x30, 0x2c, 0x31, 0x32, 0x37, 0x7d, 0x29, 0x7c, 0x5c, 0x2a, 0x29, 0x24, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x45, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x42, 0x26, 0x72, 0x24, 0x28, 0x40, 0x32, 0x20, 0x5e,
	0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b,
	0x31, 0x2c, 0x36, 0x34, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdf, 0x01, 0x0a, 0x07, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x42, 0x26, 0x72, 0x24, 0x28, 0x40, 0x32, 0x20, 0x5e, 0x28,
	0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31,
	0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x35, 0xfa, 0x42, 0x32, 0x72, 0x30, 0x28, 0x80, 0x01, 0x32, 0x2b, 0x5e, 0x28, 0x28, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41,
	0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2f, 0x5f, 0x7c, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x32, 0x37,
	0x7d, 0x29, 0x7c, 0x5c, 0x2a, 0x29, 0x24, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xfa,
	0x42, 0x2f, 0x72, 0x2d, 0x28, 0x40, 0x32, 0x26, 0x5e, 0x28, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a,
	0x5d, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36,
	0x32, 0x7d, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0xd0, 0x01,
	0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x0b,
	0x54, 0x75, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x08,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32,
	0xfa, 0x42, 0x2f, 0x72, 0x2d, 0x28, 0x40, 0x32, 0x26, 0x5e, 0x28, 0x5b, 0x2e, 0x26, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c,
	0x36, 0x32, 0x7d, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0xd0,
	0x01, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x42, 0x26, 0x72, 0x24, 0x28, 0x40, 0x32, 0x20,
	0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d,
	0x7b, 0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x38, 0xfa, 0x42, 0x35, 0x72, 0x33, 0x28, 0x80, 0x01, 0x32, 0x2b, 0x5e, 0x28,
	0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x5b, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2f, 0x5f, 0x7c, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x31,
	0x32, 0x37, 0x7d, 0x29, 0x7c, 0x5c, 0x2a, 0x29, 0x24, 0xd0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xeb, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2c, 0xfa, 0x42, 0x29, 0x72, 0x27, 0x28, 0x40, 0x32, 0x20, 0x5e, 0x28, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x24, 0xd0, 0x01, 0x01, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x38, 0xfa, 0x42, 0x35, 0x72, 0x33, 0x28, 0x80, 0x01, 0x32, 0x2b, 0x5e, 0x28, 0x28, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41,
	0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2f, 0x5f, 0x7c, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x31, 0x32, 0x37,
	0x7d, 0x29, 0x7c, 0x5c, 0x2a, 0x29, 0x24, 0xd0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x32, 0xfa, 0x42, 0x2f, 0x72, 0x2d, 0x28, 0x40, 0x32, 0x26, 0x5e, 0x28, 0x5b, 0x2e, 0x26,
	0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b,
	0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x2e, 0x26, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29,
	0x24, 0xd0, 0x01, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xaf,
	0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x54, 0x72, 0x65, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x54, 0x72, 0x65,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22,
	0x3f, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x49,
	0x4f, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x53, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x03,
	0x22, 0x88, 0x01, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64,
	0x54, 0x72, 0x65, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x48, 0x00, 0x52, 0x04, 0x6c,
	0x65, 0x61, 0x66, 0x42, 0x06, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x4e, 0x0a, 0x08, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x39, 0x42, 0x0a, 0x54,
	0x75, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x66, 0x79, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tuple_proto_rawDescOnce sync.Once
	file_tuple_proto_rawDescData = file_tuple_proto_rawDesc
)

func file_tuple_proto_rawDescGZIP() []byte {
	file_tuple_proto_rawDescOnce.Do(func() {
		file_tuple_proto_rawDescData = protoimpl.X.CompressGZIP(file_tuple_proto_rawDescData)
	})
	return file_tuple_proto_rawDescData
}

var file_tuple_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tuple_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tuple_proto_goTypes = []interface{}{
	(ExpandTreeNode_Operation)(0), // 0: ExpandTreeNode.Operation
	(*Tuple)(nil),                 // 1: Tuple
	(*Entity)(nil),                // 2: Entity
	(*EntityAndRelation)(nil),     // 3: EntityAndRelation
	(*Subject)(nil),               // 4: Subject
	(*TupleFilter)(nil),           // 5: TupleFilter
	(*EntityFilter)(nil),          // 6: EntityFilter
	(*SubjectFilter)(nil),         // 7: SubjectFilter
	(*ExpandTreeNode)(nil),        // 8: ExpandTreeNode
	(*Expand)(nil),                // 9: Expand
	(*Subjects)(nil),              // 10: Subjects
}
var file_tuple_proto_depIdxs = []int32{
	2,  // 0: Tuple.entity:type_name -> Entity
	4,  // 1: Tuple.subject:type_name -> Subject
	2,  // 2: EntityAndRelation.entity:type_name -> Entity
	6,  // 3: TupleFilter.entity:type_name -> EntityFilter
	7,  // 4: TupleFilter.subject:type_name -> SubjectFilter
	0,  // 5: ExpandTreeNode.operation:type_name -> ExpandTreeNode.Operation
	9,  // 6: ExpandTreeNode.children:type_name -> Expand
	3,  // 7: Expand.target:type_name -> EntityAndRelation
	8,  // 8: Expand.expand:type_name -> ExpandTreeNode
	10, // 9: Expand.leaf:type_name -> Subjects
	4,  // 10: Subjects.subjects:type_name -> Subject
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tuple_proto_init() }
func file_tuple_proto_init() {
	if File_tuple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tuple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tuple); i {
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
		file_tuple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_tuple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityAndRelation); i {
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
		file_tuple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_tuple_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TupleFilter); i {
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
		file_tuple_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityFilter); i {
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
		file_tuple_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectFilter); i {
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
		file_tuple_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandTreeNode); i {
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
		file_tuple_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expand); i {
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
		file_tuple_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subjects); i {
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
	file_tuple_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Expand_Expand)(nil),
		(*Expand_Leaf)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tuple_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tuple_proto_goTypes,
		DependencyIndexes: file_tuple_proto_depIdxs,
		EnumInfos:         file_tuple_proto_enumTypes,
		MessageInfos:      file_tuple_proto_msgTypes,
	}.Build()
	File_tuple_proto = out.File
	file_tuple_proto_rawDesc = nil
	file_tuple_proto_goTypes = nil
	file_tuple_proto_depIdxs = nil
}
