// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: ProtoBufPacket.proto

package ProtoBufPacket

import (
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

// PayloadType 열거형 정의
type PayloadType int32

const (
	// 어플 관련
	PayloadType_AppInit         PayloadType = 0
	PayloadType_AppVersionCheck PayloadType = 1
	PayloadType_AppCurrentUTC   PayloadType = 2
	PayloadType_AppReconnect    PayloadType = 3
	PayloadType_AppDisconnect   PayloadType = 4
	PayloadType_AppQuit         PayloadType = 5
	// 유저 관련
	PayloadType_UserEnter PayloadType = 1001
	// 방 관련
	PayloadType_RoomsSearch PayloadType = 2001
	PayloadType_RoomCreate  PayloadType = 2002
	PayloadType_RoomJoin    PayloadType = 2003
	PayloadType_RoomQuit    PayloadType = 2004
	// RPS게임 관련
	PayloadType_RPSGamePlayersInfo PayloadType = 3001
	PayloadType_RPSGameTableInfo   PayloadType = 3002
	PayloadType_RPSGameGo          PayloadType = 3003
	PayloadType_RPSGamePlaying     PayloadType = 3004
	PayloadType_RPSGameProgress    PayloadType = 3005
)

// Enum value maps for PayloadType.
var (
	PayloadType_name = map[int32]string{
		0:    "AppInit",
		1:    "AppVersionCheck",
		2:    "AppCurrentUTC",
		3:    "AppReconnect",
		4:    "AppDisconnect",
		5:    "AppQuit",
		1001: "UserEnter",
		2001: "RoomsSearch",
		2002: "RoomCreate",
		2003: "RoomJoin",
		2004: "RoomQuit",
		3001: "RPSGamePlayersInfo",
		3002: "RPSGameTableInfo",
		3003: "RPSGameGo",
		3004: "RPSGamePlaying",
		3005: "RPSGameProgress",
	}
	PayloadType_value = map[string]int32{
		"AppInit":            0,
		"AppVersionCheck":    1,
		"AppCurrentUTC":      2,
		"AppReconnect":       3,
		"AppDisconnect":      4,
		"AppQuit":            5,
		"UserEnter":          1001,
		"RoomsSearch":        2001,
		"RoomCreate":         2002,
		"RoomJoin":           2003,
		"RoomQuit":           2004,
		"RPSGamePlayersInfo": 3001,
		"RPSGameTableInfo":   3002,
		"RPSGameGo":          3003,
		"RPSGamePlaying":     3004,
		"RPSGameProgress":    3005,
	}
)

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}

func (x PayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_ProtoBufPacket_proto_enumTypes[0].Descriptor()
}

func (PayloadType) Type() protoreflect.EnumType {
	return &file_ProtoBufPacket_proto_enumTypes[0]
}

func (x PayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadType.Descriptor instead.
func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{0}
}

// PayloadClass 메시지 정의
type PayloadClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// App에 관한 정보
	AppRunOS   int32  `protobuf:"varint,1,opt,name=appRunOS,proto3" json:"appRunOS,omitempty"`    // 1. 디바이스 정보
	AppName    string `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`       // 2. 애플리케이션 이름
	AppVersion string `protobuf:"bytes,3,opt,name=appVersion,proto3" json:"appVersion,omitempty"` // 3. 애플리케이션 버전 .. 버전 정보는 일반적으로 "메이저.마이너.패치"와 같은 형식
	// Time에 관한 정보
	CurrentUTC int64 `protobuf:"varint,4,opt,name=currentUTC,proto3" json:"currentUTC,omitempty"` // 4. 현재 시간
	// User에 관한 정보
	UserUUID string `protobuf:"bytes,5,opt,name=userUUID,proto3" json:"userUUID,omitempty"` // 5. 사용자 UUID // 로그인 전엔 정보없음
	// 데이터에 관한 정보
	RequestCode    PayloadType `protobuf:"varint,6,opt,name=requestCode,proto3,enum=PayloadType" json:"requestCode,omitempty"` // 6. 코드 번호
	RequestData    []byte      `protobuf:"bytes,7,opt,name=requestData,proto3" json:"requestData,omitempty"`                   // 7. 요청 데이터 (바이트 배열)
	ResultMessage  string      `protobuf:"bytes,8,opt,name=resultMessage,proto3" json:"resultMessage,omitempty"`               // 8. 결과 메세지
	BroadCastGroup int32       `protobuf:"varint,9,opt,name=broadCastGroup,proto3" json:"broadCastGroup,omitempty"`            // 9. 방 번호 // 방 입장 전엔 0
}

func (x *PayloadClass) Reset() {
	*x = PayloadClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClass) ProtoMessage() {}

func (x *PayloadClass) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClass.ProtoReflect.Descriptor instead.
func (*PayloadClass) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadClass) GetAppRunOS() int32 {
	if x != nil {
		return x.AppRunOS
	}
	return 0
}

func (x *PayloadClass) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *PayloadClass) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *PayloadClass) GetCurrentUTC() int64 {
	if x != nil {
		return x.CurrentUTC
	}
	return 0
}

func (x *PayloadClass) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

func (x *PayloadClass) GetRequestCode() PayloadType {
	if x != nil {
		return x.RequestCode
	}
	return PayloadType_AppInit
}

func (x *PayloadClass) GetRequestData() []byte {
	if x != nil {
		return x.RequestData
	}
	return nil
}

func (x *PayloadClass) GetResultMessage() string {
	if x != nil {
		return x.ResultMessage
	}
	return ""
}

func (x *PayloadClass) GetBroadCastGroup() int32 {
	if x != nil {
		return x.BroadCastGroup
	}
	return 0
}

// /////////////////////////
// ! User
// /////////////////////////
// PayloadClassUserEnder 메시지 정의 (1001)
type PayloadClassUserEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *PayloadClassUserEnter) Reset() {
	*x = PayloadClassUserEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassUserEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassUserEnter) ProtoMessage() {}

func (x *PayloadClassUserEnter) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassUserEnter.ProtoReflect.Descriptor instead.
func (*PayloadClassUserEnter) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadClassUserEnter) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// /////////////////////////
// ! Room
// /////////////////////////
// RoomsSearch 메시지 정의 (2001)
type PayloadClassRoomsSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomSeq          int32 `protobuf:"varint,1,opt,name=roomSeq,proto3" json:"roomSeq,omitempty"`                   // 1. 방 번호
	RoomType         int32 `protobuf:"varint,2,opt,name=roomType,proto3" json:"roomType,omitempty"`                 // 2. 방 게임 타입
	RoomNowPersonnel int32 `protobuf:"varint,3,opt,name=roomNowPersonnel,proto3" json:"roomNowPersonnel,omitempty"` // 3. 현재 방에 몇명 있는지
	RoomMaxPersonnel int32 `protobuf:"varint,4,opt,name=roomMaxPersonnel,proto3" json:"roomMaxPersonnel,omitempty"` // 4. 현재 방에 몇명 참여 가능한지
}

func (x *PayloadClassRoomsSearch) Reset() {
	*x = PayloadClassRoomsSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRoomsSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRoomsSearch) ProtoMessage() {}

func (x *PayloadClassRoomsSearch) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRoomsSearch.ProtoReflect.Descriptor instead.
func (*PayloadClassRoomsSearch) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{2}
}

func (x *PayloadClassRoomsSearch) GetRoomSeq() int32 {
	if x != nil {
		return x.RoomSeq
	}
	return 0
}

func (x *PayloadClassRoomsSearch) GetRoomType() int32 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

func (x *PayloadClassRoomsSearch) GetRoomNowPersonnel() int32 {
	if x != nil {
		return x.RoomNowPersonnel
	}
	return 0
}

func (x *PayloadClassRoomsSearch) GetRoomMaxPersonnel() int32 {
	if x != nil {
		return x.RoomMaxPersonnel
	}
	return 0
}

// RoomJoin 메시지 정의 (2002)
type PayloadClassRoomCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomSeq  int32 `protobuf:"varint,1,opt,name=roomSeq,proto3" json:"roomSeq,omitempty"`   // 1. 방 번호
	RoomType int32 `protobuf:"varint,2,opt,name=roomType,proto3" json:"roomType,omitempty"` // 2. 방 게임 타입
}

func (x *PayloadClassRoomCreate) Reset() {
	*x = PayloadClassRoomCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRoomCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRoomCreate) ProtoMessage() {}

func (x *PayloadClassRoomCreate) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRoomCreate.ProtoReflect.Descriptor instead.
func (*PayloadClassRoomCreate) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadClassRoomCreate) GetRoomSeq() int32 {
	if x != nil {
		return x.RoomSeq
	}
	return 0
}

func (x *PayloadClassRoomCreate) GetRoomType() int32 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

// RoomJoin 메시지 정의 (2002)
type PayloadClassRoomJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomSeq  int32 `protobuf:"varint,1,opt,name=roomSeq,proto3" json:"roomSeq,omitempty"`   // 1. 방 번호
	RoomType int32 `protobuf:"varint,2,opt,name=roomType,proto3" json:"roomType,omitempty"` // 2. 방 게임 타입
}

func (x *PayloadClassRoomJoin) Reset() {
	*x = PayloadClassRoomJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRoomJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRoomJoin) ProtoMessage() {}

func (x *PayloadClassRoomJoin) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRoomJoin.ProtoReflect.Descriptor instead.
func (*PayloadClassRoomJoin) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{4}
}

func (x *PayloadClassRoomJoin) GetRoomSeq() int32 {
	if x != nil {
		return x.RoomSeq
	}
	return 0
}

func (x *PayloadClassRoomJoin) GetRoomType() int32 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

// RoomQuit 메시지 정의 (2003)
type PayloadClassRoomQuit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomSeq  int32 `protobuf:"varint,1,opt,name=roomSeq,proto3" json:"roomSeq,omitempty"`   // 1. 방 번호
	RoomType int32 `protobuf:"varint,2,opt,name=roomType,proto3" json:"roomType,omitempty"` // 2. 방 게임 타입
}

func (x *PayloadClassRoomQuit) Reset() {
	*x = PayloadClassRoomQuit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRoomQuit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRoomQuit) ProtoMessage() {}

func (x *PayloadClassRoomQuit) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRoomQuit.ProtoReflect.Descriptor instead.
func (*PayloadClassRoomQuit) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{5}
}

func (x *PayloadClassRoomQuit) GetRoomSeq() int32 {
	if x != nil {
		return x.RoomSeq
	}
	return 0
}

func (x *PayloadClassRoomQuit) GetRoomType() int32 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

// * PayloadClassRPSGamePlayersInfo 메시지 정의 (3001)
type PayloadClassRPSGamePlayersInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpsGamePlayerNames []string `protobuf:"bytes,1,rep,name=rpsGamePlayerNames,proto3" json:"rpsGamePlayerNames,omitempty"` // 자리를 모두 ""으로 표시, 자리에 사람이 있으면 사람의 닉네임으로 대체
}

func (x *PayloadClassRPSGamePlayersInfo) Reset() {
	*x = PayloadClassRPSGamePlayersInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRPSGamePlayersInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRPSGamePlayersInfo) ProtoMessage() {}

func (x *PayloadClassRPSGamePlayersInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRPSGamePlayersInfo.ProtoReflect.Descriptor instead.
func (*PayloadClassRPSGamePlayersInfo) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{6}
}

func (x *PayloadClassRPSGamePlayersInfo) GetRpsGamePlayerNames() []string {
	if x != nil {
		return x.RpsGamePlayerNames
	}
	return nil
}

// * PayloadClassRPSGameTableInfo 메시지 정의 (3002)
type PayloadClassRPSGameTableInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpsTableStatus int32   `protobuf:"varint,1,opt,name=rpsTableStatus,proto3" json:"rpsTableStatus,omitempty"`    // 게임이 시작상태인지 대기 상태인지
	RpsGamePlays   []int32 `protobuf:"varint,2,rep,packed,name=rpsGamePlays,proto3" json:"rpsGamePlays,omitempty"` // 가위바위보 뭐냇는지 확인?
}

func (x *PayloadClassRPSGameTableInfo) Reset() {
	*x = PayloadClassRPSGameTableInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRPSGameTableInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRPSGameTableInfo) ProtoMessage() {}

func (x *PayloadClassRPSGameTableInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRPSGameTableInfo.ProtoReflect.Descriptor instead.
func (*PayloadClassRPSGameTableInfo) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{7}
}

func (x *PayloadClassRPSGameTableInfo) GetRpsTableStatus() int32 {
	if x != nil {
		return x.RpsTableStatus
	}
	return 0
}

func (x *PayloadClassRPSGameTableInfo) GetRpsGamePlays() []int32 {
	if x != nil {
		return x.RpsGamePlays
	}
	return nil
}

// * PayloadClassRPSGamePlaying 메시지 정의 (3004)
type PayloadClassRPSGamePlaying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpsGamePlay int32 `protobuf:"varint,1,opt,name=rpsGamePlay,proto3" json:"rpsGamePlay,omitempty"` // 가위바위보 뭐냇는지 확인?
}

func (x *PayloadClassRPSGamePlaying) Reset() {
	*x = PayloadClassRPSGamePlaying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBufPacket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadClassRPSGamePlaying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadClassRPSGamePlaying) ProtoMessage() {}

func (x *PayloadClassRPSGamePlaying) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBufPacket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadClassRPSGamePlaying.ProtoReflect.Descriptor instead.
func (*PayloadClassRPSGamePlaying) Descriptor() ([]byte, []int) {
	return file_ProtoBufPacket_proto_rawDescGZIP(), []int{8}
}

func (x *PayloadClassRPSGamePlaying) GetRpsGamePlay() int32 {
	if x != nil {
		return x.RpsGamePlay
	}
	return 0
}

var File_ProtoBufPacket_proto protoreflect.FileDescriptor

var file_ProtoBufPacket_proto_rawDesc = []byte{
	0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x52, 0x75,
	0x6e, 0x4f, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x70, 0x70, 0x52, 0x75,
	0x6e, 0x4f, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x54, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x54, 0x43, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x33, 0x0a, 0x15, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa7,
	0x01, 0x0a, 0x17, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x53, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x6d,
	0x4e, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x10,
	0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x78, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x78, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x4e, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x1e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x50, 0x53, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x70, 0x73, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x70, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x6a, 0x0a, 0x1c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x50, 0x53, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x70, 0x73, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x72, 0x70, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x70, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x70, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x73, 0x22, 0x3e, 0x0a, 0x1a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x50, 0x53, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x70, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x70, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x2a, 0xb0, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x69, 0x74, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x54, 0x43, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x70, 0x70,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x70, 0x70, 0x51, 0x75, 0x69, 0x74, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0xe9, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x10, 0xd1, 0x0f, 0x12, 0x0f, 0x0a, 0x0a, 0x52,
	0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0xd2, 0x0f, 0x12, 0x0d, 0x0a, 0x08,
	0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x10, 0xd3, 0x0f, 0x12, 0x0d, 0x0a, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x51, 0x75, 0x69, 0x74, 0x10, 0xd4, 0x0f, 0x12, 0x17, 0x0a, 0x12, 0x52, 0x50,
	0x53, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x10, 0xb9, 0x17, 0x12, 0x15, 0x0a, 0x10, 0x52, 0x50, 0x53, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xba, 0x17, 0x12, 0x0e, 0x0a, 0x09, 0x52, 0x50,
	0x53, 0x47, 0x61, 0x6d, 0x65, 0x47, 0x6f, 0x10, 0xbb, 0x17, 0x12, 0x13, 0x0a, 0x0e, 0x52, 0x50,
	0x53, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x10, 0xbc, 0x17, 0x12,
	0x14, 0x0a, 0x0f, 0x52, 0x50, 0x53, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x10, 0xbd, 0x17, 0x42, 0x33, 0x5a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75,
	0x66, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0xaa, 0x02, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x75, 0x66, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ProtoBufPacket_proto_rawDescOnce sync.Once
	file_ProtoBufPacket_proto_rawDescData = file_ProtoBufPacket_proto_rawDesc
)

func file_ProtoBufPacket_proto_rawDescGZIP() []byte {
	file_ProtoBufPacket_proto_rawDescOnce.Do(func() {
		file_ProtoBufPacket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtoBufPacket_proto_rawDescData)
	})
	return file_ProtoBufPacket_proto_rawDescData
}

var file_ProtoBufPacket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ProtoBufPacket_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ProtoBufPacket_proto_goTypes = []interface{}{
	(PayloadType)(0),                       // 0: PayloadType
	(*PayloadClass)(nil),                   // 1: PayloadClass
	(*PayloadClassUserEnter)(nil),          // 2: PayloadClassUserEnter
	(*PayloadClassRoomsSearch)(nil),        // 3: PayloadClassRoomsSearch
	(*PayloadClassRoomCreate)(nil),         // 4: PayloadClassRoomCreate
	(*PayloadClassRoomJoin)(nil),           // 5: PayloadClassRoomJoin
	(*PayloadClassRoomQuit)(nil),           // 6: PayloadClassRoomQuit
	(*PayloadClassRPSGamePlayersInfo)(nil), // 7: PayloadClassRPSGamePlayersInfo
	(*PayloadClassRPSGameTableInfo)(nil),   // 8: PayloadClassRPSGameTableInfo
	(*PayloadClassRPSGamePlaying)(nil),     // 9: PayloadClassRPSGamePlaying
}
var file_ProtoBufPacket_proto_depIdxs = []int32{
	0, // 0: PayloadClass.requestCode:type_name -> PayloadType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ProtoBufPacket_proto_init() }
func file_ProtoBufPacket_proto_init() {
	if File_ProtoBufPacket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtoBufPacket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClass); i {
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
		file_ProtoBufPacket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassUserEnter); i {
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
		file_ProtoBufPacket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRoomsSearch); i {
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
		file_ProtoBufPacket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRoomCreate); i {
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
		file_ProtoBufPacket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRoomJoin); i {
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
		file_ProtoBufPacket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRoomQuit); i {
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
		file_ProtoBufPacket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRPSGamePlayersInfo); i {
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
		file_ProtoBufPacket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRPSGameTableInfo); i {
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
		file_ProtoBufPacket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadClassRPSGamePlaying); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ProtoBufPacket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProtoBufPacket_proto_goTypes,
		DependencyIndexes: file_ProtoBufPacket_proto_depIdxs,
		EnumInfos:         file_ProtoBufPacket_proto_enumTypes,
		MessageInfos:      file_ProtoBufPacket_proto_msgTypes,
	}.Build()
	File_ProtoBufPacket_proto = out.File
	file_ProtoBufPacket_proto_rawDesc = nil
	file_ProtoBufPacket_proto_goTypes = nil
	file_ProtoBufPacket_proto_depIdxs = nil
}
