// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamedata.proto

package gamedata

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Prefab int32

const (
	Prefab_Invalid Prefab = 0
	Prefab_DemoBox Prefab = 1
)

var Prefab_name = map[int32]string{
	0: "Invalid",
	1: "DemoBox",
}

var Prefab_value = map[string]int32{
	"Invalid": 0,
	"DemoBox": 1,
}

func (x Prefab) String() string {
	return proto.EnumName(Prefab_name, int32(x))
}

func (Prefab) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{0}
}

type Header_OpCode int32

const (
	Header_Invalid Header_OpCode = 0
	// (A.1) client informs game server of successful connection to
	// coordinator
	Header_ClientOnline      Header_OpCode = 100
	Header_GameServerOnline  Header_OpCode = 101
	Header_GameServerStart   Header_OpCode = 102
	Header_ClientGameRequest Header_OpCode = 200
	Header_ClientGameCancel  Header_OpCode = 201
	Header_ClientGameFound   Header_OpCode = 202
	// (A.5) (FixedUpdate) sim server sends existing client game object's
	// delta to game server
	// (B.3) (FixedUpdate) game server sends existing game object's delta to
	// all clients
	Header_ClientJoin Header_OpCode = 300
	Header_GameObject Header_OpCode = 301
	Header_Rigidbody  Header_OpCode = 302
)

var Header_OpCode_name = map[int32]string{
	0:   "Invalid",
	100: "ClientOnline",
	101: "GameServerOnline",
	102: "GameServerStart",
	200: "ClientGameRequest",
	201: "ClientGameCancel",
	202: "ClientGameFound",
	300: "ClientJoin",
	301: "GameObject",
	302: "Rigidbody",
}

var Header_OpCode_value = map[string]int32{
	"Invalid":           0,
	"ClientOnline":      100,
	"GameServerOnline":  101,
	"GameServerStart":   102,
	"ClientGameRequest": 200,
	"ClientGameCancel":  201,
	"ClientGameFound":   202,
	"ClientJoin":        300,
	"GameObject":        301,
	"Rigidbody":         302,
}

func (x Header_OpCode) String() string {
	return proto.EnumName(Header_OpCode_name, int32(x))
}

func (Header_OpCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{12, 0}
}

type Packet struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Packet) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type ClientJoin struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientJoin) Reset()         { *m = ClientJoin{} }
func (m *ClientJoin) String() string { return proto.CompactTextString(m) }
func (*ClientJoin) ProtoMessage()    {}
func (*ClientJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{1}
}

func (m *ClientJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientJoin.Unmarshal(m, b)
}
func (m *ClientJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientJoin.Marshal(b, m, deterministic)
}
func (m *ClientJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientJoin.Merge(m, src)
}
func (m *ClientJoin) XXX_Size() int {
	return xxx_messageInfo_ClientJoin.Size(m)
}
func (m *ClientJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientJoin.DiscardUnknown(m)
}

var xxx_messageInfo_ClientJoin proto.InternalMessageInfo

type ClientConnection struct {
	// id of client to accept packets from
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// address of the client packets should be coming from
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientConnection) Reset()         { *m = ClientConnection{} }
func (m *ClientConnection) String() string { return proto.CompactTextString(m) }
func (*ClientConnection) ProtoMessage()    {}
func (*ClientConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{2}
}

func (m *ClientConnection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConnection.Unmarshal(m, b)
}
func (m *ClientConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConnection.Marshal(b, m, deterministic)
}
func (m *ClientConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConnection.Merge(m, src)
}
func (m *ClientConnection) XXX_Size() int {
	return xxx_messageInfo_ClientConnection.Size(m)
}
func (m *ClientConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConnection.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConnection proto.InternalMessageInfo

func (m *ClientConnection) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ClientConnection) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ClientOnline struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientOnline) Reset()         { *m = ClientOnline{} }
func (m *ClientOnline) String() string { return proto.CompactTextString(m) }
func (*ClientOnline) ProtoMessage()    {}
func (*ClientOnline) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{3}
}

func (m *ClientOnline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientOnline.Unmarshal(m, b)
}
func (m *ClientOnline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientOnline.Marshal(b, m, deterministic)
}
func (m *ClientOnline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientOnline.Merge(m, src)
}
func (m *ClientOnline) XXX_Size() int {
	return xxx_messageInfo_ClientOnline.Size(m)
}
func (m *ClientOnline) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientOnline.DiscardUnknown(m)
}

var xxx_messageInfo_ClientOnline proto.InternalMessageInfo

func (m *ClientOnline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GameServerOnline struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Capacity             int32    `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServerOnline) Reset()         { *m = GameServerOnline{} }
func (m *GameServerOnline) String() string { return proto.CompactTextString(m) }
func (*GameServerOnline) ProtoMessage()    {}
func (*GameServerOnline) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{4}
}

func (m *GameServerOnline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServerOnline.Unmarshal(m, b)
}
func (m *GameServerOnline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServerOnline.Marshal(b, m, deterministic)
}
func (m *GameServerOnline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServerOnline.Merge(m, src)
}
func (m *GameServerOnline) XXX_Size() int {
	return xxx_messageInfo_GameServerOnline.Size(m)
}
func (m *GameServerOnline) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServerOnline.DiscardUnknown(m)
}

var xxx_messageInfo_GameServerOnline proto.InternalMessageInfo

func (m *GameServerOnline) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *GameServerOnline) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *GameServerOnline) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type GameServerStart struct {
	ID                   string              `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Clients              []*ClientConnection `protobuf:"bytes,2,rep,name=clients,proto3" json:"clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GameServerStart) Reset()         { *m = GameServerStart{} }
func (m *GameServerStart) String() string { return proto.CompactTextString(m) }
func (*GameServerStart) ProtoMessage()    {}
func (*GameServerStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{5}
}

func (m *GameServerStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServerStart.Unmarshal(m, b)
}
func (m *GameServerStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServerStart.Marshal(b, m, deterministic)
}
func (m *GameServerStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServerStart.Merge(m, src)
}
func (m *GameServerStart) XXX_Size() int {
	return xxx_messageInfo_GameServerStart.Size(m)
}
func (m *GameServerStart) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServerStart.DiscardUnknown(m)
}

var xxx_messageInfo_GameServerStart proto.InternalMessageInfo

func (m *GameServerStart) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GameServerStart) GetClients() []*ClientConnection {
	if m != nil {
		return m.Clients
	}
	return nil
}

type ClientGameFound struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientGameFound) Reset()         { *m = ClientGameFound{} }
func (m *ClientGameFound) String() string { return proto.CompactTextString(m) }
func (*ClientGameFound) ProtoMessage()    {}
func (*ClientGameFound) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{6}
}

func (m *ClientGameFound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientGameFound.Unmarshal(m, b)
}
func (m *ClientGameFound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientGameFound.Marshal(b, m, deterministic)
}
func (m *ClientGameFound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientGameFound.Merge(m, src)
}
func (m *ClientGameFound) XXX_Size() int {
	return xxx_messageInfo_ClientGameFound.Size(m)
}
func (m *ClientGameFound) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientGameFound.DiscardUnknown(m)
}

var xxx_messageInfo_ClientGameFound proto.InternalMessageInfo

func (m *ClientGameFound) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ClientGameRequest struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientGameRequest) Reset()         { *m = ClientGameRequest{} }
func (m *ClientGameRequest) String() string { return proto.CompactTextString(m) }
func (*ClientGameRequest) ProtoMessage()    {}
func (*ClientGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{7}
}

func (m *ClientGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientGameRequest.Unmarshal(m, b)
}
func (m *ClientGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientGameRequest.Marshal(b, m, deterministic)
}
func (m *ClientGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientGameRequest.Merge(m, src)
}
func (m *ClientGameRequest) XXX_Size() int {
	return xxx_messageInfo_ClientGameRequest.Size(m)
}
func (m *ClientGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientGameRequest proto.InternalMessageInfo

func (m *ClientGameRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type ClientGameCancel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientGameCancel) Reset()         { *m = ClientGameCancel{} }
func (m *ClientGameCancel) String() string { return proto.CompactTextString(m) }
func (*ClientGameCancel) ProtoMessage()    {}
func (*ClientGameCancel) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{8}
}

func (m *ClientGameCancel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientGameCancel.Unmarshal(m, b)
}
func (m *ClientGameCancel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientGameCancel.Marshal(b, m, deterministic)
}
func (m *ClientGameCancel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientGameCancel.Merge(m, src)
}
func (m *ClientGameCancel) XXX_Size() int {
	return xxx_messageInfo_ClientGameCancel.Size(m)
}
func (m *ClientGameCancel) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientGameCancel.DiscardUnknown(m)
}

var xxx_messageInfo_ClientGameCancel proto.InternalMessageInfo

type GameObject struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Prefab               Prefab   `protobuf:"varint,2,opt,name=Prefab,proto3,enum=gamedata.Prefab" json:"Prefab,omitempty"`
	Position             *Vector2 `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Rotation             float32  `protobuf:"fixed32,4,opt,name=rotation,proto3" json:"rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameObject) Reset()         { *m = GameObject{} }
func (m *GameObject) String() string { return proto.CompactTextString(m) }
func (*GameObject) ProtoMessage()    {}
func (*GameObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{9}
}

func (m *GameObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameObject.Unmarshal(m, b)
}
func (m *GameObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameObject.Marshal(b, m, deterministic)
}
func (m *GameObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameObject.Merge(m, src)
}
func (m *GameObject) XXX_Size() int {
	return xxx_messageInfo_GameObject.Size(m)
}
func (m *GameObject) XXX_DiscardUnknown() {
	xxx_messageInfo_GameObject.DiscardUnknown(m)
}

var xxx_messageInfo_GameObject proto.InternalMessageInfo

func (m *GameObject) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GameObject) GetPrefab() Prefab {
	if m != nil {
		return m.Prefab
	}
	return Prefab_Invalid
}

func (m *GameObject) GetPosition() *Vector2 {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *GameObject) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

type Rigidbody struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Prefab               Prefab   `protobuf:"varint,2,opt,name=Prefab,proto3,enum=gamedata.Prefab" json:"Prefab,omitempty"`
	Position             *Vector2 `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Rotation             float32  `protobuf:"fixed32,4,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Velocity             *Vector2 `protobuf:"bytes,5,opt,name=velocity,proto3" json:"velocity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rigidbody) Reset()         { *m = Rigidbody{} }
func (m *Rigidbody) String() string { return proto.CompactTextString(m) }
func (*Rigidbody) ProtoMessage()    {}
func (*Rigidbody) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{10}
}

func (m *Rigidbody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rigidbody.Unmarshal(m, b)
}
func (m *Rigidbody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rigidbody.Marshal(b, m, deterministic)
}
func (m *Rigidbody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rigidbody.Merge(m, src)
}
func (m *Rigidbody) XXX_Size() int {
	return xxx_messageInfo_Rigidbody.Size(m)
}
func (m *Rigidbody) XXX_DiscardUnknown() {
	xxx_messageInfo_Rigidbody.DiscardUnknown(m)
}

var xxx_messageInfo_Rigidbody proto.InternalMessageInfo

func (m *Rigidbody) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Rigidbody) GetPrefab() Prefab {
	if m != nil {
		return m.Prefab
	}
	return Prefab_Invalid
}

func (m *Rigidbody) GetPosition() *Vector2 {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Rigidbody) GetRotation() float32 {
	if m != nil {
		return m.Rotation
	}
	return 0
}

func (m *Rigidbody) GetVelocity() *Vector2 {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type Vector2 struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector2) Reset()         { *m = Vector2{} }
func (m *Vector2) String() string { return proto.CompactTextString(m) }
func (*Vector2) ProtoMessage()    {}
func (*Vector2) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{11}
}

func (m *Vector2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector2.Unmarshal(m, b)
}
func (m *Vector2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector2.Marshal(b, m, deterministic)
}
func (m *Vector2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector2.Merge(m, src)
}
func (m *Vector2) XXX_Size() int {
	return xxx_messageInfo_Vector2.Size(m)
}
func (m *Vector2) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector2.DiscardUnknown(m)
}

var xxx_messageInfo_Vector2 proto.InternalMessageInfo

func (m *Vector2) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector2) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Header struct {
	// this is a connection id that expires once the coordinator connection is
	// lost
	Cid                  string        `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	OpCode               Header_OpCode `protobuf:"varint,2,opt,name=opCode,proto3,enum=gamedata.Header_OpCode" json:"opCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d51a89ffb18ac9, []int{12}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Header) GetOpCode() Header_OpCode {
	if m != nil {
		return m.OpCode
	}
	return Header_Invalid
}

func init() {
	proto.RegisterEnum("gamedata.Prefab", Prefab_name, Prefab_value)
	proto.RegisterEnum("gamedata.Header_OpCode", Header_OpCode_name, Header_OpCode_value)
	proto.RegisterType((*Packet)(nil), "gamedata.Packet")
	proto.RegisterType((*ClientJoin)(nil), "gamedata.ClientJoin")
	proto.RegisterType((*ClientConnection)(nil), "gamedata.ClientConnection")
	proto.RegisterType((*ClientOnline)(nil), "gamedata.ClientOnline")
	proto.RegisterType((*GameServerOnline)(nil), "gamedata.GameServerOnline")
	proto.RegisterType((*GameServerStart)(nil), "gamedata.GameServerStart")
	proto.RegisterType((*ClientGameFound)(nil), "gamedata.ClientGameFound")
	proto.RegisterType((*ClientGameRequest)(nil), "gamedata.ClientGameRequest")
	proto.RegisterType((*ClientGameCancel)(nil), "gamedata.ClientGameCancel")
	proto.RegisterType((*GameObject)(nil), "gamedata.GameObject")
	proto.RegisterType((*Rigidbody)(nil), "gamedata.Rigidbody")
	proto.RegisterType((*Vector2)(nil), "gamedata.Vector2")
	proto.RegisterType((*Header)(nil), "gamedata.Header")
}

func init() { proto.RegisterFile("gamedata.proto", fileDescriptor_38d51a89ffb18ac9) }

var fileDescriptor_38d51a89ffb18ac9 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0xd6, 0x6d, 0x9d, 0x76, 0x52, 0xa5, 0xdb, 0xfd, 0x4a, 0x09, 0x3d, 0x95, 0x95, 0x90,
	0x2c, 0x10, 0xae, 0x14, 0x38, 0x72, 0x81, 0x44, 0x40, 0xb8, 0xa4, 0xda, 0x4a, 0x70, 0x41, 0x48,
	0x1b, 0x7b, 0x12, 0x0c, 0xce, 0x6e, 0x70, 0x9c, 0x28, 0xfe, 0x19, 0x9c, 0xf8, 0x13, 0xc0, 0x5f,
	0xe0, 0x4a, 0xf9, 0x55, 0x68, 0xd7, 0xeb, 0x24, 0x8d, 0xe1, 0xce, 0xcd, 0xef, 0xed, 0xf3, 0xf3,
	0xbc, 0xd9, 0x19, 0x43, 0x6b, 0x2c, 0x27, 0x18, 0xcb, 0x5c, 0x86, 0xd3, 0x4c, 0xe7, 0x9a, 0xed,
	0x57, 0xf8, 0xec, 0xce, 0x58, 0xeb, 0x71, 0x8a, 0x17, 0x96, 0x1f, 0xce, 0x47, 0x17, 0x52, 0x15,
	0xa5, 0x88, 0xbf, 0x05, 0xff, 0x52, 0x46, 0x1f, 0x31, 0x67, 0x01, 0xf8, 0xef, 0x51, 0xc6, 0x98,
	0xb5, 0xc9, 0x39, 0x09, 0x9a, 0x1d, 0x1a, 0xae, 0xfc, 0x5e, 0x5a, 0x5e, 0xb8, 0x73, 0x16, 0xc0,
	0xae, 0xa1, 0xdb, 0x9e, 0xd5, 0x9d, 0x84, 0xa5, 0x7b, 0x58, 0xb9, 0x87, 0x4f, 0x55, 0x21, 0xac,
	0x82, 0x1f, 0x02, 0x74, 0xd3, 0x04, 0x55, 0xfe, 0x4a, 0x27, 0x8a, 0x3f, 0x01, 0x5a, 0xa2, 0xae,
	0x56, 0x0a, 0xa3, 0x3c, 0xd1, 0x8a, 0xb5, 0xc0, 0xeb, 0xf7, 0xec, 0x17, 0x0f, 0x84, 0xd7, 0xef,
	0xb1, 0x36, 0x34, 0x64, 0x1c, 0x67, 0x38, 0x9b, 0x59, 0xfb, 0x03, 0x51, 0x41, 0xce, 0xe1, 0xb0,
	0x7c, 0x7b, 0xa0, 0xd2, 0x44, 0x21, 0x63, 0xb0, 0xab, 0xe4, 0x04, 0xdd, 0xbb, 0xf6, 0x99, 0xbf,
	0x03, 0xfa, 0x42, 0x4e, 0xf0, 0x0a, 0xb3, 0x05, 0x66, 0x4e, 0x77, 0x0a, 0xfe, 0x0c, 0xa3, 0x0c,
	0x73, 0xa7, 0x74, 0xc8, 0xf0, 0x19, 0x8e, 0x13, 0xad, 0xdc, 0x87, 0x1c, 0x62, 0x67, 0xb0, 0x1f,
	0xc9, 0xa9, 0x8c, 0x92, 0xbc, 0x68, 0xef, 0x9c, 0x93, 0x60, 0x4f, 0xac, 0x30, 0x7f, 0x03, 0x47,
	0x6b, 0xff, 0xab, 0x5c, 0x66, 0x79, 0x2d, 0xc0, 0x63, 0x68, 0x44, 0xb6, 0x4c, 0x13, 0x60, 0x27,
	0x68, 0x76, 0xce, 0xd6, 0x7d, 0xdc, 0x4e, 0x2f, 0x2a, 0x29, 0xbf, 0x0b, 0x47, 0xe5, 0xa1, 0xb1,
	0x7f, 0xae, 0xe7, 0x2a, 0xde, 0x36, 0xe6, 0x0f, 0xe0, 0x78, 0x2d, 0x11, 0xf8, 0x69, 0x8e, 0xb3,
	0xcd, 0x10, 0x64, 0x33, 0x04, 0x67, 0x55, 0xab, 0x8d, 0xb8, 0x2b, 0x55, 0x84, 0x29, 0xff, 0x4c,
	0x00, 0x0c, 0x1c, 0x0c, 0x3f, 0x60, 0x54, 0x2f, 0x3c, 0x00, 0xff, 0x32, 0xc3, 0x91, 0x1c, 0xda,
	0x7e, 0xb4, 0x36, 0xef, 0xbf, 0xe4, 0x85, 0x3b, 0x67, 0x0f, 0x61, 0x7f, 0xaa, 0x67, 0x89, 0x49,
	0x60, 0x3b, 0xd4, 0xec, 0x1c, 0xaf, 0xb5, 0xaf, 0x31, 0xca, 0x75, 0xd6, 0x11, 0x2b, 0x89, 0x69,
	0x68, 0xa6, 0x73, 0x69, 0xe5, 0xbb, 0xe7, 0x24, 0xf0, 0xc4, 0x0a, 0xf3, 0x1f, 0x04, 0x0e, 0x44,
	0x32, 0x4e, 0xe2, 0xa1, 0x8e, 0x8b, 0x7f, 0xa2, 0x24, 0x63, 0xb5, 0xc0, 0x54, 0xdb, 0xfb, 0xdf,
	0xfb, 0xab, 0x55, 0x25, 0xe1, 0xf7, 0xa0, 0xe1, 0x48, 0x76, 0x08, 0x64, 0x69, 0xab, 0xf7, 0x04,
	0x59, 0x1a, 0x54, 0xd8, 0xba, 0x3d, 0x41, 0x0a, 0xfe, 0xc5, 0x03, 0xbf, 0x5c, 0x23, 0x46, 0x61,
	0x27, 0x4a, 0x62, 0x17, 0xd3, 0x3c, 0xb2, 0x0b, 0xf0, 0xf5, 0xb4, 0xab, 0x63, 0x74, 0x39, 0x6f,
	0x6f, 0xaf, 0x5e, 0x38, 0xb0, 0xc7, 0xc2, 0xc9, 0xf8, 0x35, 0x01, 0xbf, 0xa4, 0x58, 0x13, 0x1a,
	0x7d, 0xb5, 0x90, 0x69, 0x12, 0xd3, 0xff, 0x18, 0xbd, 0xb9, 0x23, 0x34, 0x66, 0x27, 0xf5, 0x8d,
	0xa0, 0xc8, 0xfe, 0xaf, 0xcd, 0x31, 0x1d, 0xb1, 0xd3, 0x3f, 0x0c, 0x18, 0xfd, 0x49, 0xd8, 0xad,
	0xfa, 0x2c, 0xd1, 0x6b, 0xc2, 0x4e, 0x6a, 0x23, 0x4b, 0x7f, 0x11, 0x76, 0xb4, 0xb9, 0xf1, 0xf4,
	0xab, 0x67, 0x88, 0xf5, 0xd0, 0xd1, 0x6f, 0x1e, 0x6b, 0x6d, 0xdc, 0x38, 0xfd, 0xee, 0xdd, 0xe7,
	0xd5, 0x25, 0xdf, 0x8c, 0xd2, 0x84, 0x46, 0x0f, 0x27, 0xfa, 0x99, 0x5e, 0x52, 0x32, 0xf4, 0xed,
	0xbf, 0xe5, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x28, 0x43, 0x4b, 0xe3, 0x04, 0x00,
	0x00,
}
