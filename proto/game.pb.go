// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: game.proto

package gamepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Status int32

const (
	Status_UNKNOWN_STATUS       Status = 0
	Status_GAME_CREATED         Status = 1
	Status_WAITING_FOR_OPPONENT Status = 2
	Status_GAME_STARTED         Status = 3
	Status_HOST_HIT             Status = 4
	Status_HOST_MISS            Status = 5
	Status_OPPONENT_HIT         Status = 6
	Status_OPPONENT_MISS        Status = 7
	Status_HOST_WON             Status = 8
	Status_OPPONENT_WON         Status = 9
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "GAME_CREATED",
		2: "WAITING_FOR_OPPONENT",
		3: "GAME_STARTED",
		4: "HOST_HIT",
		5: "HOST_MISS",
		6: "OPPONENT_HIT",
		7: "OPPONENT_MISS",
		8: "HOST_WON",
		9: "OPPONENT_WON",
	}
	Status_value = map[string]int32{
		"UNKNOWN_STATUS":       0,
		"GAME_CREATED":         1,
		"WAITING_FOR_OPPONENT": 2,
		"GAME_STARTED":         3,
		"HOST_HIT":             4,
		"HOST_MISS":            5,
		"OPPONENT_HIT":         6,
		"OPPONENT_MISS":        7,
		"HOST_WON":             8,
		"OPPONENT_WON":         9,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

type Ship int32

const (
	Ship_UNKNOWN_SHIP      Ship = 0
	Ship_SINGLE_DECK       Ship = 1
	Ship_DOUBLE_DECK_DOWN  Ship = 2
	Ship_THREE_DECK_DOWN   Ship = 3
	Ship_FOUR_DECK_DOWN    Ship = 4
	Ship_DOUBLE_DECK_RIGHT Ship = 5
	Ship_THREE_DECK_RIGHT  Ship = 6
	Ship_FOUR_DECK_RIGHT   Ship = 7
)

// Enum value maps for Ship.
var (
	Ship_name = map[int32]string{
		0: "UNKNOWN_SHIP",
		1: "SINGLE_DECK",
		2: "DOUBLE_DECK_DOWN",
		3: "THREE_DECK_DOWN",
		4: "FOUR_DECK_DOWN",
		5: "DOUBLE_DECK_RIGHT",
		6: "THREE_DECK_RIGHT",
		7: "FOUR_DECK_RIGHT",
	}
	Ship_value = map[string]int32{
		"UNKNOWN_SHIP":      0,
		"SINGLE_DECK":       1,
		"DOUBLE_DECK_DOWN":  2,
		"THREE_DECK_DOWN":   3,
		"FOUR_DECK_DOWN":    4,
		"DOUBLE_DECK_RIGHT": 5,
		"THREE_DECK_RIGHT":  6,
		"FOUR_DECK_RIGHT":   7,
	}
)

func (x Ship) Enum() *Ship {
	p := new(Ship)
	*p = x
	return p
}

func (x Ship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ship) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[1].Descriptor()
}

func (Ship) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[1]
}

func (x Ship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ship.Descriptor instead.
func (Ship) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostNickname string `protobuf:"bytes,1,opt,name=host_nickname,json=hostNickname,proto3" json:"host_nickname,omitempty"`
	HostField    string `protobuf:"bytes,2,opt,name=host_field,json=hostField,proto3" json:"host_field,omitempty"`
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGameRequest) GetHostNickname() string {
	if x != nil {
		return x.HostNickname
	}
	return ""
}

func (x *CreateGameRequest) GetHostField() string {
	if x != nil {
		return x.HostField
	}
	return ""
}

type CreateGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostUuid string `protobuf:"bytes,1,opt,name=host_uuid,json=hostUuid,proto3" json:"host_uuid,omitempty"`
}

func (x *CreateGameResponse) Reset() {
	*x = CreateGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameResponse) ProtoMessage() {}

func (x *CreateGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameResponse.ProtoReflect.Descriptor instead.
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGameResponse) GetHostUuid() string {
	if x != nil {
		return x.HostUuid
	}
	return ""
}

type GetGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGamesRequest) Reset() {
	*x = GetGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamesRequest) ProtoMessage() {}

func (x *GetGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamesRequest.ProtoReflect.Descriptor instead.
func (*GetGamesRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

type GetGamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []string `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
}

func (x *GetGamesResponse) Reset() {
	*x = GetGamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamesResponse) ProtoMessage() {}

func (x *GetGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamesResponse.ProtoReflect.Descriptor instead.
func (*GetGamesResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *GetGamesResponse) GetGames() []string {
	if x != nil {
		return x.Games
	}
	return nil
}

type JoinGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostNickname     string `protobuf:"bytes,1,opt,name=host_nickname,json=hostNickname,proto3" json:"host_nickname,omitempty"`
	OpponentNickname string `protobuf:"bytes,2,opt,name=opponent_nickname,json=opponentNickname,proto3" json:"opponent_nickname,omitempty"`
}

func (x *JoinGameRequest) Reset() {
	*x = JoinGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameRequest) ProtoMessage() {}

func (x *JoinGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameRequest.ProtoReflect.Descriptor instead.
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *JoinGameRequest) GetHostNickname() string {
	if x != nil {
		return x.HostNickname
	}
	return ""
}

func (x *JoinGameRequest) GetOpponentNickname() string {
	if x != nil {
		return x.OpponentNickname
	}
	return ""
}

type JoinGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpponentUuid string `protobuf:"bytes,1,opt,name=opponent_uuid,json=opponentUuid,proto3" json:"opponent_uuid,omitempty"`
}

func (x *JoinGameResponse) Reset() {
	*x = JoinGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGameResponse) ProtoMessage() {}

func (x *JoinGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGameResponse.ProtoReflect.Descriptor instead.
func (*JoinGameResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *JoinGameResponse) GetOpponentUuid() string {
	if x != nil {
		return x.OpponentUuid
	}
	return ""
}

type StartGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostNickname  string `protobuf:"bytes,1,opt,name=host_nickname,json=hostNickname,proto3" json:"host_nickname,omitempty"`
	OpponentField string `protobuf:"bytes,2,opt,name=opponent_field,json=opponentField,proto3" json:"opponent_field,omitempty"`
	OpponentUuid  string `protobuf:"bytes,3,opt,name=opponent_uuid,json=opponentUuid,proto3" json:"opponent_uuid,omitempty"`
}

func (x *StartGameRequest) Reset() {
	*x = StartGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameRequest) ProtoMessage() {}

func (x *StartGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameRequest.ProtoReflect.Descriptor instead.
func (*StartGameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *StartGameRequest) GetHostNickname() string {
	if x != nil {
		return x.HostNickname
	}
	return ""
}

func (x *StartGameRequest) GetOpponentField() string {
	if x != nil {
		return x.OpponentField
	}
	return ""
}

func (x *StartGameRequest) GetOpponentUuid() string {
	if x != nil {
		return x.OpponentUuid
	}
	return ""
}

type StartGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartGameResponse) Reset() {
	*x = StartGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGameResponse) ProtoMessage() {}

func (x *StartGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGameResponse.ProtoReflect.Descriptor instead.
func (*StartGameResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

type ShootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostNickname string `protobuf:"bytes,1,opt,name=host_nickname,json=hostNickname,proto3" json:"host_nickname,omitempty"`
	X            uint32 `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y            uint32 `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Uuid         string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ShootRequest) Reset() {
	*x = ShootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShootRequest) ProtoMessage() {}

func (x *ShootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShootRequest.ProtoReflect.Descriptor instead.
func (*ShootRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{8}
}

func (x *ShootRequest) GetHostNickname() string {
	if x != nil {
		return x.HostNickname
	}
	return ""
}

func (x *ShootRequest) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ShootRequest) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *ShootRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ShootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        Status `protobuf:"varint,1,opt,name=status,proto3,enum=game.Status" json:"status,omitempty"`
	DestroyedShip Ship   `protobuf:"varint,2,opt,name=destroyed_ship,json=destroyedShip,proto3,enum=game.Ship" json:"destroyed_ship,omitempty"`
	X             uint32 `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y             uint32 `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ShootResponse) Reset() {
	*x = ShootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShootResponse) ProtoMessage() {}

func (x *ShootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShootResponse.ProtoReflect.Descriptor instead.
func (*ShootResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{9}
}

func (x *ShootResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN_STATUS
}

func (x *ShootResponse) GetDestroyedShip() Ship {
	if x != nil {
		return x.DestroyedShip
	}
	return Ship_UNKNOWN_SHIP
}

func (x *ShootResponse) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ShootResponse) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type WaitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *WaitRequest) Reset() {
	*x = WaitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitRequest) ProtoMessage() {}

func (x *WaitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitRequest.ProtoReflect.Descriptor instead.
func (*WaitRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{10}
}

func (x *WaitRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type WaitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        Status `protobuf:"varint,1,opt,name=status,proto3,enum=game.Status" json:"status,omitempty"`
	X             uint32 `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y             uint32 `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	DestroyedShip Ship   `protobuf:"varint,4,opt,name=destroyed_ship,json=destroyedShip,proto3,enum=game.Ship" json:"destroyed_ship,omitempty"`
	DestroyedX    uint32 `protobuf:"varint,5,opt,name=destroyed_x,json=destroyedX,proto3" json:"destroyed_x,omitempty"`
	DestroyedY    uint32 `protobuf:"varint,6,opt,name=destroyed_y,json=destroyedY,proto3" json:"destroyed_y,omitempty"`
	Message       string `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WaitResponse) Reset() {
	*x = WaitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitResponse) ProtoMessage() {}

func (x *WaitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitResponse.ProtoReflect.Descriptor instead.
func (*WaitResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{11}
}

func (x *WaitResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN_STATUS
}

func (x *WaitResponse) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *WaitResponse) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *WaitResponse) GetDestroyedShip() Ship {
	if x != nil {
		return x.DestroyedShip
	}
	return Ship_UNKNOWN_SHIP
}

func (x *WaitResponse) GetDestroyedX() uint32 {
	if x != nil {
		return x.DestroyedX
	}
	return 0
}

func (x *WaitResponse) GetDestroyedY() uint32 {
	if x != nil {
		return x.DestroyedY
	}
	return 0
}

func (x *WaitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f,
	0x73, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x68, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x31, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0f, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x70,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37,
	0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x63, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x01, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x6f,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x31, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x73, 0x68, 0x69,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x53, 0x68,
	0x69, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x22, 0x21,
	0x0a, 0x0b, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x01, 0x79, 0x12, 0x31, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65,
	0x64, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x65, 0x64, 0x53, 0x68, 0x69, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x65, 0x64, 0x5f, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x58, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x59, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0xbc, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x4f, 0x50, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x04, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x10, 0x05, 0x12, 0x10, 0x0a,
	0x0c, 0x4f, 0x50, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x06, 0x12,
	0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53,
	0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x57, 0x4f, 0x4e, 0x10, 0x08,
	0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x4f, 0x4e,
	0x10, 0x09, 0x2a, 0xaa, 0x01, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4b, 0x5f, 0x44, 0x4f,
	0x57, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x44, 0x45,
	0x43, 0x4b, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4f, 0x55,
	0x52, 0x5f, 0x44, 0x45, 0x43, 0x4b, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4b, 0x5f, 0x52, 0x49, 0x47,
	0x48, 0x54, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x44, 0x45,
	0x43, 0x4b, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f,
	0x55, 0x52, 0x5f, 0x44, 0x45, 0x43, 0x4b, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x07, 0x32,
	0xe8, 0x03, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x51,
	0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x69,
	0x6e, 0x12, 0x55, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x49, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x6f,
	0x74, 0x12, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x68, 0x6f,
	0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x73, 0x68,
	0x6f, 0x6f, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x57, 0x61, 0x69, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x69, 0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x68, 0x75, 0x6c, 0x33, 0x2f, 0x73, 0x65, 0x61, 0x2d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x67, 0x61, 0x6d,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_game_proto_goTypes = []interface{}{
	(Status)(0),                // 0: game.Status
	(Ship)(0),                  // 1: game.Ship
	(*CreateGameRequest)(nil),  // 2: game.CreateGameRequest
	(*CreateGameResponse)(nil), // 3: game.CreateGameResponse
	(*GetGamesRequest)(nil),    // 4: game.GetGamesRequest
	(*GetGamesResponse)(nil),   // 5: game.GetGamesResponse
	(*JoinGameRequest)(nil),    // 6: game.JoinGameRequest
	(*JoinGameResponse)(nil),   // 7: game.JoinGameResponse
	(*StartGameRequest)(nil),   // 8: game.StartGameRequest
	(*StartGameResponse)(nil),  // 9: game.StartGameResponse
	(*ShootRequest)(nil),       // 10: game.ShootRequest
	(*ShootResponse)(nil),      // 11: game.ShootResponse
	(*WaitRequest)(nil),        // 12: game.WaitRequest
	(*WaitResponse)(nil),       // 13: game.WaitResponse
}
var file_game_proto_depIdxs = []int32{
	0,  // 0: game.ShootResponse.status:type_name -> game.Status
	1,  // 1: game.ShootResponse.destroyed_ship:type_name -> game.Ship
	0,  // 2: game.WaitResponse.status:type_name -> game.Status
	1,  // 3: game.WaitResponse.destroyed_ship:type_name -> game.Ship
	2,  // 4: game.GameService.CreateGame:input_type -> game.CreateGameRequest
	4,  // 5: game.GameService.GetGames:input_type -> game.GetGamesRequest
	6,  // 6: game.GameService.JoinGame:input_type -> game.JoinGameRequest
	8,  // 7: game.GameService.StartGame:input_type -> game.StartGameRequest
	10, // 8: game.GameService.Shoot:input_type -> game.ShootRequest
	12, // 9: game.GameService.Wait:input_type -> game.WaitRequest
	3,  // 10: game.GameService.CreateGame:output_type -> game.CreateGameResponse
	5,  // 11: game.GameService.GetGames:output_type -> game.GetGamesResponse
	7,  // 12: game.GameService.JoinGame:output_type -> game.JoinGameResponse
	9,  // 13: game.GameService.StartGame:output_type -> game.StartGameResponse
	11, // 14: game.GameService.Shoot:output_type -> game.ShootResponse
	13, // 15: game.GameService.Wait:output_type -> game.WaitResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameRequest); i {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameResponse); i {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamesRequest); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamesResponse); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGameRequest); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGameResponse); i {
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
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameRequest); i {
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
		file_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGameResponse); i {
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
		file_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShootRequest); i {
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
		file_game_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShootResponse); i {
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
		file_game_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitRequest); i {
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
		file_game_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitResponse); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
