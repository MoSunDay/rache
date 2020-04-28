// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master.proto

package kv

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

type JoinArgs struct {
	Groups               map[int32]*Servers `protobuf:"bytes,1,rep,name=Groups,proto3" json:"Groups,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sequence             int64              `protobuf:"zigzag64,2,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	ClientId             int64              `protobuf:"zigzag64,3,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JoinArgs) Reset()         { *m = JoinArgs{} }
func (m *JoinArgs) String() string { return proto.CompactTextString(m) }
func (*JoinArgs) ProtoMessage()    {}
func (*JoinArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{0}
}

func (m *JoinArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinArgs.Unmarshal(m, b)
}
func (m *JoinArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinArgs.Marshal(b, m, deterministic)
}
func (m *JoinArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinArgs.Merge(m, src)
}
func (m *JoinArgs) XXX_Size() int {
	return xxx_messageInfo_JoinArgs.Size(m)
}
func (m *JoinArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinArgs.DiscardUnknown(m)
}

var xxx_messageInfo_JoinArgs proto.InternalMessageInfo

func (m *JoinArgs) GetGroups() map[int32]*Servers {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *JoinArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *JoinArgs) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type JoinReply struct {
	WrongLeader          bool     `protobuf:"varint,1,opt,name=WrongLeader,proto3" json:"WrongLeader,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinReply) Reset()         { *m = JoinReply{} }
func (m *JoinReply) String() string { return proto.CompactTextString(m) }
func (*JoinReply) ProtoMessage()    {}
func (*JoinReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{1}
}

func (m *JoinReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinReply.Unmarshal(m, b)
}
func (m *JoinReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinReply.Marshal(b, m, deterministic)
}
func (m *JoinReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinReply.Merge(m, src)
}
func (m *JoinReply) XXX_Size() int {
	return xxx_messageInfo_JoinReply.Size(m)
}
func (m *JoinReply) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinReply.DiscardUnknown(m)
}

var xxx_messageInfo_JoinReply proto.InternalMessageInfo

func (m *JoinReply) GetWrongLeader() bool {
	if m != nil {
		return m.WrongLeader
	}
	return false
}

func (m *JoinReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// 存放的server的列表，应为IP地址
type Servers struct {
	ServerList           []string `protobuf:"bytes,1,rep,name=ServerList,proto3" json:"ServerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Servers) Reset()         { *m = Servers{} }
func (m *Servers) String() string { return proto.CompactTextString(m) }
func (*Servers) ProtoMessage()    {}
func (*Servers) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{2}
}

func (m *Servers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Servers.Unmarshal(m, b)
}
func (m *Servers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Servers.Marshal(b, m, deterministic)
}
func (m *Servers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Servers.Merge(m, src)
}
func (m *Servers) XXX_Size() int {
	return xxx_messageInfo_Servers.Size(m)
}
func (m *Servers) XXX_DiscardUnknown() {
	xxx_messageInfo_Servers.DiscardUnknown(m)
}

var xxx_messageInfo_Servers proto.InternalMessageInfo

func (m *Servers) GetServerList() []string {
	if m != nil {
		return m.ServerList
	}
	return nil
}

type LeaveArgs struct {
	GIDs                 []int32  `protobuf:"varint,1,rep,packed,name=GIDs,proto3" json:"GIDs,omitempty"`
	Sequence             int64    `protobuf:"zigzag64,2,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	ClientId             int64    `protobuf:"zigzag64,3,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveArgs) Reset()         { *m = LeaveArgs{} }
func (m *LeaveArgs) String() string { return proto.CompactTextString(m) }
func (*LeaveArgs) ProtoMessage()    {}
func (*LeaveArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{3}
}

func (m *LeaveArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveArgs.Unmarshal(m, b)
}
func (m *LeaveArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveArgs.Marshal(b, m, deterministic)
}
func (m *LeaveArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveArgs.Merge(m, src)
}
func (m *LeaveArgs) XXX_Size() int {
	return xxx_messageInfo_LeaveArgs.Size(m)
}
func (m *LeaveArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveArgs.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveArgs proto.InternalMessageInfo

func (m *LeaveArgs) GetGIDs() []int32 {
	if m != nil {
		return m.GIDs
	}
	return nil
}

func (m *LeaveArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *LeaveArgs) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type LeaveReply struct {
	WrongLeader          bool     `protobuf:"varint,1,opt,name=WrongLeader,proto3" json:"WrongLeader,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveReply) Reset()         { *m = LeaveReply{} }
func (m *LeaveReply) String() string { return proto.CompactTextString(m) }
func (*LeaveReply) ProtoMessage()    {}
func (*LeaveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{4}
}

func (m *LeaveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveReply.Unmarshal(m, b)
}
func (m *LeaveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveReply.Marshal(b, m, deterministic)
}
func (m *LeaveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveReply.Merge(m, src)
}
func (m *LeaveReply) XXX_Size() int {
	return xxx_messageInfo_LeaveReply.Size(m)
}
func (m *LeaveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveReply.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveReply proto.InternalMessageInfo

func (m *LeaveReply) GetWrongLeader() bool {
	if m != nil {
		return m.WrongLeader
	}
	return false
}

func (m *LeaveReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type MoveArgs struct {
	Shard                int32    `protobuf:"varint,1,opt,name=Shard,proto3" json:"Shard,omitempty"`
	GID                  int32    `protobuf:"varint,2,opt,name=GID,proto3" json:"GID,omitempty"`
	Sequence             int64    `protobuf:"zigzag64,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	ClientId             int64    `protobuf:"zigzag64,4,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveArgs) Reset()         { *m = MoveArgs{} }
func (m *MoveArgs) String() string { return proto.CompactTextString(m) }
func (*MoveArgs) ProtoMessage()    {}
func (*MoveArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{5}
}

func (m *MoveArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveArgs.Unmarshal(m, b)
}
func (m *MoveArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveArgs.Marshal(b, m, deterministic)
}
func (m *MoveArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveArgs.Merge(m, src)
}
func (m *MoveArgs) XXX_Size() int {
	return xxx_messageInfo_MoveArgs.Size(m)
}
func (m *MoveArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MoveArgs proto.InternalMessageInfo

func (m *MoveArgs) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *MoveArgs) GetGID() int32 {
	if m != nil {
		return m.GID
	}
	return 0
}

func (m *MoveArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *MoveArgs) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type MoveReply struct {
	WrongLeader          bool     `protobuf:"varint,1,opt,name=WrongLeader,proto3" json:"WrongLeader,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveReply) Reset()         { *m = MoveReply{} }
func (m *MoveReply) String() string { return proto.CompactTextString(m) }
func (*MoveReply) ProtoMessage()    {}
func (*MoveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{6}
}

func (m *MoveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveReply.Unmarshal(m, b)
}
func (m *MoveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveReply.Marshal(b, m, deterministic)
}
func (m *MoveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveReply.Merge(m, src)
}
func (m *MoveReply) XXX_Size() int {
	return xxx_messageInfo_MoveReply.Size(m)
}
func (m *MoveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveReply.DiscardUnknown(m)
}

var xxx_messageInfo_MoveReply proto.InternalMessageInfo

func (m *MoveReply) GetWrongLeader() bool {
	if m != nil {
		return m.WrongLeader
	}
	return false
}

func (m *MoveReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type QueryArgs struct {
	Num                  int32    `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
	Sequence             int64    `protobuf:"zigzag64,2,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	ClientId             int64    `protobuf:"zigzag64,3,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryArgs) Reset()         { *m = QueryArgs{} }
func (m *QueryArgs) String() string { return proto.CompactTextString(m) }
func (*QueryArgs) ProtoMessage()    {}
func (*QueryArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{7}
}

func (m *QueryArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryArgs.Unmarshal(m, b)
}
func (m *QueryArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryArgs.Marshal(b, m, deterministic)
}
func (m *QueryArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryArgs.Merge(m, src)
}
func (m *QueryArgs) XXX_Size() int {
	return xxx_messageInfo_QueryArgs.Size(m)
}
func (m *QueryArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryArgs.DiscardUnknown(m)
}

var xxx_messageInfo_QueryArgs proto.InternalMessageInfo

func (m *QueryArgs) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *QueryArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *QueryArgs) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type QueryReply struct {
	WrongLeader          bool     `protobuf:"varint,1,opt,name=WrongLeader,proto3" json:"WrongLeader,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	Config               *Config  `protobuf:"bytes,3,opt,name=Config,proto3" json:"Config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{8}
}

func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (m *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(m, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetWrongLeader() bool {
	if m != nil {
		return m.WrongLeader
	}
	return false
}

func (m *QueryReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *QueryReply) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// master的config信息
type Config struct {
	Num                  int32              `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
	Shards               []int32            `protobuf:"varint,2,rep,packed,name=Shards,proto3" json:"Shards,omitempty"`
	Groups               map[int32]*Servers `protobuf:"bytes,3,rep,name=Groups,proto3" json:"Groups,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{9}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Config) GetShards() []int32 {
	if m != nil {
		return m.Shards
	}
	return nil
}

func (m *Config) GetGroups() map[int32]*Servers {
	if m != nil {
		return m.Groups
	}
	return nil
}

// 写入raft的操作
type MasterOp struct {
	Command  string `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	ClientId int64  `protobuf:"zigzag64,2,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	Sequence int64  `protobuf:"zigzag64,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	// arguments of Join RPC method
	Groups map[int32]*Servers `protobuf:"bytes,4,rep,name=Groups,proto3" json:"Groups,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// arguments of Leave RPC method
	GIDs []int32 `protobuf:"varint,5,rep,packed,name=GIDs,proto3" json:"GIDs,omitempty"`
	// arguments of Move RPC method
	Shard int32 `protobuf:"varint,6,opt,name=Shard,proto3" json:"Shard,omitempty"`
	GID   int32 `protobuf:"varint,7,opt,name=GID,proto3" json:"GID,omitempty"`
	// arguments of Query RPC method
	Num                  int32    `protobuf:"varint,8,opt,name=Num,proto3" json:"Num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MasterOp) Reset()         { *m = MasterOp{} }
func (m *MasterOp) String() string { return proto.CompactTextString(m) }
func (*MasterOp) ProtoMessage()    {}
func (*MasterOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{10}
}

func (m *MasterOp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MasterOp.Unmarshal(m, b)
}
func (m *MasterOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MasterOp.Marshal(b, m, deterministic)
}
func (m *MasterOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterOp.Merge(m, src)
}
func (m *MasterOp) XXX_Size() int {
	return xxx_messageInfo_MasterOp.Size(m)
}
func (m *MasterOp) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterOp.DiscardUnknown(m)
}

var xxx_messageInfo_MasterOp proto.InternalMessageInfo

func (m *MasterOp) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *MasterOp) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *MasterOp) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *MasterOp) GetGroups() map[int32]*Servers {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *MasterOp) GetGIDs() []int32 {
	if m != nil {
		return m.GIDs
	}
	return nil
}

func (m *MasterOp) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *MasterOp) GetGID() int32 {
	if m != nil {
		return m.GID
	}
	return 0
}

func (m *MasterOp) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*JoinArgs)(nil), "kv.JoinArgs")
	proto.RegisterMapType((map[int32]*Servers)(nil), "kv.JoinArgs.GroupsEntry")
	proto.RegisterType((*JoinReply)(nil), "kv.JoinReply")
	proto.RegisterType((*Servers)(nil), "kv.Servers")
	proto.RegisterType((*LeaveArgs)(nil), "kv.LeaveArgs")
	proto.RegisterType((*LeaveReply)(nil), "kv.LeaveReply")
	proto.RegisterType((*MoveArgs)(nil), "kv.MoveArgs")
	proto.RegisterType((*MoveReply)(nil), "kv.MoveReply")
	proto.RegisterType((*QueryArgs)(nil), "kv.QueryArgs")
	proto.RegisterType((*QueryReply)(nil), "kv.QueryReply")
	proto.RegisterType((*Config)(nil), "kv.Config")
	proto.RegisterMapType((map[int32]*Servers)(nil), "kv.Config.GroupsEntry")
	proto.RegisterType((*MasterOp)(nil), "kv.MasterOp")
	proto.RegisterMapType((map[int32]*Servers)(nil), "kv.MasterOp.GroupsEntry")
}

func init() { proto.RegisterFile("master.proto", fileDescriptor_f9c348dec43a6705) }

var fileDescriptor_f9c348dec43a6705 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x3f, 0x6b, 0x8f, 0x0b, 0xaa, 0x56, 0xa8, 0xb2, 0x72, 0x40, 0x61, 0xa5, 0x4a, 0xe1,
	0x62, 0xa1, 0x70, 0x41, 0x5c, 0x0a, 0x4a, 0x4b, 0x14, 0x94, 0x16, 0xb1, 0x11, 0xea, 0xd9, 0x90,
	0x25, 0x84, 0xc4, 0x76, 0x58, 0x7f, 0x48, 0xf9, 0x3f, 0x9c, 0xf8, 0x0d, 0xfc, 0x38, 0xb4, 0xb3,
	0x5e, 0xdb, 0x41, 0x51, 0x0f, 0x56, 0x6f, 0x33, 0x3b, 0x9b, 0xb7, 0x2f, 0xcf, 0xef, 0x0d, 0x9c,
	0x25, 0x71, 0x5e, 0x70, 0x11, 0xed, 0x44, 0x56, 0x64, 0xc4, 0xdc, 0x54, 0xf4, 0xaf, 0x01, 0xde,
	0xc7, 0x6c, 0x9d, 0xbe, 0x17, 0xab, 0x9c, 0xbc, 0x02, 0x77, 0x2a, 0xb2, 0x72, 0x97, 0x87, 0xc6,
	0xd0, 0x1a, 0x05, 0xe3, 0x30, 0xda, 0x54, 0x91, 0x9e, 0x46, 0x6a, 0x74, 0x93, 0x16, 0x62, 0xcf,
	0xea, 0x7b, 0x64, 0x00, 0xde, 0x82, 0xff, 0x2a, 0x79, 0xfa, 0x8d, 0x87, 0xe6, 0xd0, 0x18, 0x11,
	0xd6, 0xf4, 0x72, 0x36, 0xd9, 0xae, 0x79, 0x5a, 0xcc, 0x96, 0xa1, 0xa5, 0x66, 0xba, 0x1f, 0x7c,
	0x80, 0xa0, 0x03, 0x47, 0xce, 0xc1, 0xda, 0xf0, 0x7d, 0x68, 0x0c, 0x8d, 0x91, 0xc3, 0x64, 0x49,
	0x5e, 0x80, 0x53, 0xc5, 0xdb, 0x52, 0xa1, 0x06, 0xe3, 0x40, 0x32, 0x59, 0x70, 0x51, 0x71, 0x91,
	0x33, 0x35, 0x79, 0x6b, 0xbe, 0x31, 0xe8, 0x15, 0xf8, 0x92, 0x1f, 0xe3, 0xbb, 0xed, 0x9e, 0x0c,
	0x21, 0xb8, 0x17, 0x59, 0xba, 0x9a, 0xf3, 0x78, 0xc9, 0x05, 0xa2, 0x79, 0xac, 0x7b, 0x24, 0xdf,
	0xb9, 0x11, 0x02, 0x31, 0x7d, 0x26, 0x4b, 0xfa, 0x12, 0x4e, 0x6b, 0x58, 0xf2, 0x1c, 0x40, 0x95,
	0xf3, 0x75, 0x5e, 0xa0, 0x02, 0x3e, 0xeb, 0x9c, 0xd0, 0x7b, 0xf0, 0xe7, 0x3c, 0xae, 0x38, 0x4a,
	0x45, 0xc0, 0x9e, 0xce, 0xae, 0x95, 0x50, 0x0e, 0xc3, 0xba, 0xaf, 0x18, 0xf4, 0x1d, 0x00, 0x02,
	0xf7, 0xff, 0x17, 0x3f, 0xc1, 0xbb, 0xcd, 0x6a, 0x66, 0xcf, 0xc0, 0x59, 0xfc, 0x88, 0xc5, 0xb2,
	0x56, 0x53, 0x35, 0xf2, 0x37, 0xd3, 0xd9, 0x35, 0xfe, 0xc6, 0x61, 0xb2, 0x3c, 0x60, 0x6b, 0x3d,
	0xc0, 0xd6, 0xfe, 0x8f, 0xed, 0x15, 0xf8, 0xf2, 0xad, 0xfe, 0x64, 0xbf, 0x80, 0xff, 0xb9, 0xe4,
	0x62, 0x8f, 0x6c, 0xcf, 0xc1, 0xba, 0x2b, 0x13, 0xfd, 0xe5, 0xef, 0xca, 0xa4, 0xb7, 0x8a, 0x4b,
	0x00, 0x84, 0xed, 0x4d, 0x8c, 0x50, 0x70, 0x27, 0x59, 0xfa, 0x7d, 0xbd, 0x42, 0xec, 0x60, 0x0c,
	0xd2, 0x74, 0xea, 0x84, 0xd5, 0x13, 0xfa, 0xc7, 0xd0, 0x97, 0x8e, 0x50, 0xbf, 0x00, 0x17, 0xd5,
	0xce, 0x43, 0x13, 0x6d, 0x51, 0x77, 0x24, 0x6a, 0x72, 0x65, 0x61, 0xae, 0x2e, 0x5a, 0xe0, 0x63,
	0xa9, 0x7a, 0xb4, 0x74, 0xfc, 0x36, 0xc1, 0xbb, 0xc5, 0xc4, 0x7f, 0xda, 0x91, 0x10, 0x4e, 0x27,
	0x59, 0x92, 0xc4, 0xa9, 0x72, 0x86, 0xcf, 0x74, 0x7b, 0xa0, 0xaa, 0x79, 0xa8, 0xea, 0x83, 0x2e,
	0x69, 0xd7, 0x85, 0xdd, 0xae, 0x0b, 0xfd, 0xde, 0xd1, 0x75, 0xa1, 0x53, 0xe3, 0x74, 0x52, 0xd3,
	0xf8, 0xd5, 0x3d, 0xe2, 0xd7, 0xd3, 0xd6, 0xaf, 0xb5, 0xdc, 0x5e, 0x23, 0xf7, 0x63, 0xc9, 0x34,
	0x96, 0xdf, 0x54, 0xd1, 0x26, 0x97, 0x60, 0xcb, 0x7d, 0x42, 0xce, 0xba, 0x9b, 0x6f, 0xf0, 0x44,
	0x77, 0xe8, 0x2d, 0x7a, 0x42, 0x46, 0xe0, 0x60, 0x62, 0x09, 0x4e, 0x9a, 0xad, 0x30, 0x78, 0xda,
	0xb4, 0xfa, 0xe6, 0x25, 0xd8, 0x32, 0x2d, 0x0a, 0x50, 0x67, 0x54, 0x01, 0x36, 0x29, 0x52, 0x80,
	0x68, 0x5e, 0x05, 0xd8, 0xc4, 0x43, 0x01, 0xb6, 0xb6, 0xa6, 0x27, 0x5f, 0x5d, 0xdc, 0xdd, 0xaf,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xb8, 0x39, 0xb1, 0xcb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	// RPC方法，向集群中加入一系列的replica group
	Join(ctx context.Context, in *JoinArgs, opts ...grpc.CallOption) (*JoinReply, error)
	// RPC方法，从集群中移除一些group
	Leave(ctx context.Context, in *LeaveArgs, opts ...grpc.CallOption) (*LeaveReply, error)
	// RPC方法，将shard从一个group移动到另一个
	Move(ctx context.Context, in *MoveArgs, opts ...grpc.CallOption) (*MoveReply, error)
	// RPC方法，查询指定的config
	Query(ctx context.Context, in *QueryArgs, opts ...grpc.CallOption) (*QueryReply, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Join(ctx context.Context, in *JoinArgs, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/kv.Master/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Leave(ctx context.Context, in *LeaveArgs, opts ...grpc.CallOption) (*LeaveReply, error) {
	out := new(LeaveReply)
	err := c.cc.Invoke(ctx, "/kv.Master/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Move(ctx context.Context, in *MoveArgs, opts ...grpc.CallOption) (*MoveReply, error) {
	out := new(MoveReply)
	err := c.cc.Invoke(ctx, "/kv.Master/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Query(ctx context.Context, in *QueryArgs, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/kv.Master/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	// RPC方法，向集群中加入一系列的replica group
	Join(context.Context, *JoinArgs) (*JoinReply, error)
	// RPC方法，从集群中移除一些group
	Leave(context.Context, *LeaveArgs) (*LeaveReply, error)
	// RPC方法，将shard从一个group移动到另一个
	Move(context.Context, *MoveArgs) (*MoveReply, error)
	// RPC方法，查询指定的config
	Query(context.Context, *QueryArgs) (*QueryReply, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) Join(ctx context.Context, req *JoinArgs) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedMasterServer) Leave(ctx context.Context, req *LeaveArgs) (*LeaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (*UnimplementedMasterServer) Move(ctx context.Context, req *MoveArgs) (*MoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (*UnimplementedMasterServer) Query(ctx context.Context, req *QueryArgs) (*QueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.Master/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Join(ctx, req.(*JoinArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.Master/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Leave(ctx, req.(*LeaveArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.Master/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Move(ctx, req.(*MoveArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.Master/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Query(ctx, req.(*QueryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kv.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Master_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Master_Leave_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _Master_Move_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Master_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}