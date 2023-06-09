// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package game

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

type GameInfo struct {
	Players              []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Day                  int32     `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GameInfo) Reset()         { *m = GameInfo{} }
func (m *GameInfo) String() string { return proto.CompactTextString(m) }
func (*GameInfo) ProtoMessage()    {}
func (*GameInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *GameInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameInfo.Unmarshal(m, b)
}
func (m *GameInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameInfo.Marshal(b, m, deterministic)
}
func (m *GameInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameInfo.Merge(m, src)
}
func (m *GameInfo) XXX_Size() int {
	return xxx_messageInfo_GameInfo.Size(m)
}
func (m *GameInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GameInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GameInfo proto.InternalMessageInfo

func (m *GameInfo) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *GameInfo) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type VotingResult struct {
	Message              string    `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Players              []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	GameOver             bool      `protobuf:"varint,3,opt,name=gameOver,proto3" json:"gameOver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VotingResult) Reset()         { *m = VotingResult{} }
func (m *VotingResult) String() string { return proto.CompactTextString(m) }
func (*VotingResult) ProtoMessage()    {}
func (*VotingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *VotingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VotingResult.Unmarshal(m, b)
}
func (m *VotingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VotingResult.Marshal(b, m, deterministic)
}
func (m *VotingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingResult.Merge(m, src)
}
func (m *VotingResult) XXX_Size() int {
	return xxx_messageInfo_VotingResult.Size(m)
}
func (m *VotingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingResult.DiscardUnknown(m)
}

var xxx_messageInfo_VotingResult proto.InternalMessageInfo

func (m *VotingResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *VotingResult) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *VotingResult) GetGameOver() bool {
	if m != nil {
		return m.GameOver
	}
	return false
}

type PlayerVote struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	SessionId            string   `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerVote) Reset()         { *m = PlayerVote{} }
func (m *PlayerVote) String() string { return proto.CompactTextString(m) }
func (*PlayerVote) ProtoMessage()    {}
func (*PlayerVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *PlayerVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerVote.Unmarshal(m, b)
}
func (m *PlayerVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerVote.Marshal(b, m, deterministic)
}
func (m *PlayerVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerVote.Merge(m, src)
}
func (m *PlayerVote) XXX_Size() int {
	return xxx_messageInfo_PlayerVote.Size(m)
}
func (m *PlayerVote) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerVote.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerVote proto.InternalMessageInfo

func (m *PlayerVote) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerVote) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type GameSession struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameSession) Reset()         { *m = GameSession{} }
func (m *GameSession) String() string { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()    {}
func (*GameSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *GameSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameSession.Unmarshal(m, b)
}
func (m *GameSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameSession.Marshal(b, m, deterministic)
}
func (m *GameSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameSession.Merge(m, src)
}
func (m *GameSession) XXX_Size() int {
	return xxx_messageInfo_GameSession.Size(m)
}
func (m *GameSession) XXX_DiscardUnknown() {
	xxx_messageInfo_GameSession.DiscardUnknown(m)
}

var xxx_messageInfo_GameSession proto.InternalMessageInfo

func (m *GameSession) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Player struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Player) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type RegisterRequest struct {
	PlayerName           string   `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

type RegisterResponse struct {
	GameSessionId        string   `protobuf:"bytes,1,opt,name=gameSessionId,proto3" json:"gameSessionId,omitempty"`
	CurrentPlayer        *Player  `protobuf:"bytes,2,opt,name=currentPlayer,proto3" json:"currentPlayer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{6}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetGameSessionId() string {
	if m != nil {
		return m.GameSessionId
	}
	return ""
}

func (m *RegisterResponse) GetCurrentPlayer() *Player {
	if m != nil {
		return m.CurrentPlayer
	}
	return nil
}

type StartGame struct {
	GameStarted          bool     `protobuf:"varint,1,opt,name=gameStarted,proto3" json:"gameStarted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartGame) Reset()         { *m = StartGame{} }
func (m *StartGame) String() string { return proto.CompactTextString(m) }
func (*StartGame) ProtoMessage()    {}
func (*StartGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{7}
}

func (m *StartGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGame.Unmarshal(m, b)
}
func (m *StartGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGame.Marshal(b, m, deterministic)
}
func (m *StartGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGame.Merge(m, src)
}
func (m *StartGame) XXX_Size() int {
	return xxx_messageInfo_StartGame.Size(m)
}
func (m *StartGame) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGame.DiscardUnknown(m)
}

var xxx_messageInfo_StartGame proto.InternalMessageInfo

func (m *StartGame) GetGameStarted() bool {
	if m != nil {
		return m.GameStarted
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{8}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GameInfo)(nil), "game.GameInfo")
	proto.RegisterType((*VotingResult)(nil), "game.VotingResult")
	proto.RegisterType((*PlayerVote)(nil), "game.PlayerVote")
	proto.RegisterType((*GameSession)(nil), "game.GameSession")
	proto.RegisterType((*Player)(nil), "game.Player")
	proto.RegisterType((*RegisterRequest)(nil), "game.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "game.RegisterResponse")
	proto.RegisterType((*StartGame)(nil), "game.StartGame")
	proto.RegisterType((*Empty)(nil), "game.Empty")
}

func init() {
	proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769)
}

var fileDescriptor_38fc58335341d769 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x9b, 0xb4, 0xdb, 0x26, 0x37, 0xbb, 0x6b, 0xbd, 0xa0, 0x84, 0xa2, 0x12, 0x06, 0x91,
	0xbe, 0x6c, 0x75, 0x2b, 0xf8, 0xe2, 0x8b, 0xf8, 0x67, 0xa5, 0x2f, 0x2a, 0x53, 0x58, 0xc1, 0xb7,
	0xd1, 0x5e, 0x43, 0x20, 0xc9, 0xd4, 0x99, 0xa9, 0xd0, 0x8f, 0xe1, 0x37, 0x96, 0x99, 0x69, 0x9a,
	0x6c, 0x28, 0xfb, 0x76, 0xe7, 0x97, 0x73, 0xef, 0x39, 0x37, 0x93, 0x00, 0xe4, 0xa2, 0xa2, 0xc5,
	0x56, 0x49, 0x23, 0x71, 0x64, 0x6b, 0xf6, 0x11, 0xa2, 0xcf, 0xa2, 0xa2, 0x55, 0xfd, 0x5b, 0xe2,
	0x0b, 0x98, 0x6c, 0x4b, 0xb1, 0x27, 0xa5, 0xd3, 0x20, 0x1b, 0xce, 0x93, 0xe5, 0xf9, 0xc2, 0xe9,
	0xbf, 0x39, 0xc8, 0x9b, 0x87, 0x38, 0x85, 0xe1, 0x46, 0xec, 0xd3, 0x30, 0x0b, 0xe6, 0x67, 0xdc,
	0x96, 0xac, 0x84, 0xf3, 0x5b, 0x69, 0x8a, 0x3a, 0xe7, 0xa4, 0x77, 0xa5, 0xc1, 0x14, 0x26, 0x15,
	0x69, 0x2d, 0x72, 0x4a, 0x83, 0x2c, 0x98, 0xc7, 0xbc, 0x39, 0x76, 0x3d, 0xc2, 0xfb, 0x3c, 0x66,
	0x10, 0x59, 0xfe, 0xf5, 0x2f, 0xa9, 0x74, 0x98, 0x05, 0xf3, 0x88, 0x1f, 0xcf, 0xec, 0x06, 0xc0,
	0xcb, 0x6f, 0xa5, 0x21, 0xab, 0xf4, 0x4d, 0xab, 0xcd, 0xc1, 0xec, 0x78, 0xc6, 0x27, 0x10, 0x6b,
	0xd2, 0xba, 0x90, 0xf5, 0x6a, 0xe3, 0xf2, 0xc6, 0xbc, 0x05, 0xec, 0x29, 0x24, 0x76, 0xf7, 0xb5,
	0x07, 0x78, 0x09, 0x61, 0xd1, 0x8c, 0x08, 0x8b, 0x0d, 0x7b, 0x07, 0x63, 0x6f, 0x83, 0x08, 0xa3,
	0x5a, 0x54, 0xcd, 0x2e, 0xae, 0x3e, 0xa8, 0xc3, 0x46, 0x6d, 0x35, 0x4a, 0x96, 0xe4, 0xc2, 0xc6,
	0xdc, 0xd5, 0xec, 0x1a, 0x1e, 0x70, 0xca, 0x0b, 0x6d, 0x48, 0x71, 0xfa, 0xb3, 0x23, 0x6d, 0xf0,
	0x19, 0x80, 0x4f, 0xf7, 0xa5, 0x1d, 0xd8, 0x21, 0xac, 0x84, 0x69, 0xdb, 0xa2, 0xb7, 0xb2, 0xd6,
	0x84, 0xcf, 0xe1, 0x22, 0x6f, 0x73, 0x1e, 0xd7, 0xbc, 0x0b, 0x71, 0x09, 0x17, 0xbf, 0x76, 0x4a,
	0x51, 0x6d, 0x7c, 0x6a, 0x97, 0xad, 0xff, 0x7e, 0xef, 0x4a, 0xd8, 0x15, 0xc4, 0x6b, 0x23, 0x94,
	0xb1, 0xaf, 0x01, 0x33, 0x48, 0xdc, 0x44, 0x0b, 0xc8, 0x9b, 0x44, 0xbc, 0x8b, 0xd8, 0x04, 0xce,
	0x3e, 0x55, 0x5b, 0xb3, 0x5f, 0xfe, 0x0b, 0x61, 0xe4, 0x7a, 0xde, 0x42, 0xd4, 0xc4, 0xc5, 0x47,
	0xde, 0xa9, 0xb7, 0xf1, 0xec, 0x71, 0x1f, 0xfb, 0xad, 0xd8, 0x00, 0xaf, 0x21, 0xf9, 0x2e, 0x0a,
	0x73, 0x23, 0x95, 0x9b, 0xf5, 0xd0, 0x0b, 0x3b, 0x57, 0x32, 0x4b, 0x3c, 0x72, 0xa6, 0x6c, 0xf0,
	0x2a, 0xc0, 0x2b, 0x00, 0x7b, 0xe9, 0x87, 0x7b, 0x99, 0x76, 0x77, 0xb3, 0xbc, 0xd7, 0x80, 0x6f,
	0x20, 0x59, 0x1b, 0x91, 0xd3, 0xe1, 0xb3, 0x3c, 0xe1, 0x80, 0x1e, 0x75, 0xbf, 0x5e, 0x36, 0xc0,
	0x25, 0x24, 0x1f, 0x76, 0xea, 0xf8, 0x63, 0x9c, 0xe8, 0xbb, 0x6c, 0x91, 0x95, 0xb0, 0xc1, 0xfb,
	0xe8, 0xc7, 0x78, 0xf1, 0xd2, 0xc2, 0x9f, 0x63, 0xf7, 0x83, 0xbd, 0xfe, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xe5, 0x13, 0x07, 0xed, 0x6e, 0x03, 0x00, 0x00,
}
