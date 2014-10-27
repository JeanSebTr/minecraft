/*
   Copyright 2013 Matthew Collins (purggames@gmail.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package protocol

type ClientKeepAlive struct {
	KeepAliveID int32
}

type ChatMessage struct {
	Message string
}

type UseEntity struct {
	Target int32
	Mouse  int8
}

type ClientPlayer struct {
	OnGround Bool
}

type ClientPlayerPosition struct {
	X        float64
	Y        float64
	Stance   float64
	Z        float64
	OnGround Bool
}

type ClientPlayerLook struct {
	Yaw      float32
	Pitch    float32
	OnGround Bool
}

type ClientPlayerPositionLook struct {
	X        float64
	Y        float64
	Stance   float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround Bool
}

type PlayerDigging struct {
	Status byte
	X      int32
	Y      byte
	Z      int32
	Face   byte
}

type PlayerBlockPlacement struct {
	X               int32
	Y               byte
	Z               int32
	Direction       int8
	HeldItem        Slot
	CursorPositionX int8
	CursorPositionY int8
	CursorPositionZ int8
}

type ClientHeldItemChange struct {
	SlotID int16
}

type ClientAnimation struct {
	EntityID  int32
	Animation int8
}

type EntityAction struct {
	EntityID  int32
	ActionID  int8
	JumpBoost int32
}

type SteerVehicle struct {
	Sideways float32
	Forward  float32
	Jump     Bool
	Unmount  Bool
}

type ClientWindowClose struct {
	WindowID int8
}

type WindowClick struct {
	WindowID     int8
	Slot         int16
	Button       int8
	ActionNumber int16
	Mode         int8
	Item         Slot
}

type ClientWindowTransactionConfirm struct {
	WindowID     int8
	ActionNumber int16
	Accepted     Bool
}

type CreativeInventoryAction struct {
	Slot int16
	Item Slot
}

type EnchantItem struct {
	WindowID    int8
	Enchantment int8
}

type ClientUpdateSign struct {
	X     int32
	Y     int16
	Z     int32
	Line1 string
	Line2 string
	Line3 string
	Line4 string
}

type ClientPlayerAbilities struct {
	Flags        byte
	FlyingSpeed  float32
	WalkingSpeed float32
}

type ClientTabComplete struct {
	Text string
}

type ClientSettings struct {
	Locale       string
	ViewDistance int8
	ChatFlags    byte
	Difficulty   int8
	ShowCape     Bool
}

type ClientStatuses struct {
	Payload byte
}

type ClientPluginMessage struct {
	Channel string
	Data    Buffer `ltype:"int16"`
}

//Contains infomation on an item
type Slot struct {
	ID     int16
	Count  int8   `if:"ID,!=,-1"`
	Damage int16  `if:"ID,!=,-1"`
	Tag    Buffer `if:"ID,!=,-1" nil:"-1" ltype:"int16"`
}

type KeepAlive struct {
	KeepAliveID int32
}

type JoinGame struct {
	EntityID     int32
	Gamemode     byte
	Dimension    int8
	Difficulty   byte
	MaxPlayers   byte
	LevelType    string
	ReducedDebug Bool
}

type ServerMessage struct {
	Message string
}

type TimeUpdate struct {
	AgeOfTheWorld int64
	TimeOfDay     int64
}

type EntityEquipment struct {
	EntityID int32
	Slot     int16
	Item     Slot
}

type SpawnPosition struct {
	Position
}

type UpdateHealth struct {
	Health         float32
	Food           int16
	FoodSaturation float32
}

type Respawn struct {
	Dimension  int32
	Difficulty byte
	Gamemode   byte
	LevelType  string
}

type PlayerPositionLook struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround Bool
}

type HeldItemChange struct {
	SlotID byte
}

type UseBed struct {
	EntityID int32
	X        int32
	Y        byte
	Z        int32
}

type Animation struct {
	EntityID  VarInt
	Animation byte
}

type SpawnPlayer struct {
	EntityID    VarInt
	PlayerUUID  string
	PlayerName  string
	X           int32
	Y           int32
	Z           int32
	Yaw         int8
	Pitch       int8
	CurrentItem int16
	Metadata    MetaData `metadata:"true"`
}

type CollectItem struct {
	CollectedEntityID int32
	CollectorEntityID int32
}

type SpawnObject struct {
	EntityID  VarInt
	Type      int8
	X         int32
	Y         int32
	Z         int32
	Pitch     int8
	Yaw       int8
	ExtraData int32
	SpeedX    int16 `if:"ExtraData,!=,0"`
	SpeedY    int16 `if:"ExtraData,!=,0"`
	SpeedZ    int16 `if:"ExtraData,!=,0"`
}

type SpawnMob struct {
	EntityID  VarInt
	Type      byte
	X         int32
	Y         int32
	Z         int32
	Pitch     int8
	HeadPitch int8
	Yaw       int8
	VelocityX int16
	VelocityY int16
	VelocityZ int16
	Metadata  MetaData `metadata:"true"`
}

type SpawnPainting struct {
	EntityID  VarInt
	Title     string
	X         int32
	Y         int32
	Z         int32
	Direction int32
}

type SpawnExperienceOrb struct {
	EntityID VarInt
	X        int32
	Y        int32
	Z        int32
	Count    int16
}

type EntityVelocity struct {
	EntityID  int32
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

type EntityDestroy struct {
	EntityIDs EIDs `ltype:"int8"`
}

type Entity struct {
	EntityID int32
}

type EntityMove struct {
	EntityID int32
	DX       int8
	DY       int8
	DZ       int8
}

type EntityLook struct {
	EntityID int32
	Yaw      int8
	Pitch    int8
}

type EntityLookMove struct {
	EntityID int32
	DX       int8
	DY       int8
	DZ       int8
	Yaw      int8
	Pitch    int8
}

type EntityTeleport struct {
	EntityID int32
	X        int32
	Y        int32
	Z        int32
	Yaw      int8
	Pitch    int8
}

type EntityHeadLook struct {
	EntityID int32
	HeadYaw  int8
}

type EntityStatus struct {
	EntityID int32
	Status   int8
}

type EntityAttach struct {
	EntityID  int32
	VehicleID int32
	Leash     Bool
}

type EntityMetadata struct {
	EntityID int32
	Metadata MetaData `metadata:"true"`
}

type EntityEffect struct {
	EntityID  int32
	EffectID  int8
	Amplifier int8
	Duration  int16
}

type EntityEffectRemove struct {
	EntityID int32
	EffectID int8
}

type SetExperience struct {
	ExperienceBar   float32
	Level           int16
	TotalExperience int16
}

type EntityProperties struct {
	EntityID   int32
	Properties Properties `ltype:"int32"`
}

//Part of Entity Properties
type Property struct {
	Key       string
	Value     float64
	Modifiers Modifiers `ltype:"int16"`
}

//Part of Entity Properties
type Modifier struct {
	UUIDHigh  int64
	UUIDLow   int64
	Amount    float64
	Operation int8
}

type ChunkData struct {
	X              int32
	Z              int32
	GroundUp       Bool
	PrimaryBitMap  uint16
	AddBitMap      uint16
	CompressedData Buffer `ltype:"int32"`
}

type MultiBlockChange struct {
	X           int32
	Z           int32
	RecordCount int16
	Data        Buffer `ltype:"int32"`
}

type BlockChange struct {
	X    int32
	Y    byte
	Z    int32
	Type VarInt
	Data byte
}

type BlockAction struct {
	X            int32
	Y            int16
	Z            int32
	Byte1, Byte2 byte
	BlockID      VarInt
}

type BlockBreakAnimation struct {
	EntityID     VarInt
	X            int32
	Y            int32
	Z            int32
	DestroyStage int8
}

type MapChunkBulk struct {
	ChunkCount int16
	DataLength int32
	SkyLight   byte
	Data       Buffer     `ltype:"nil"`
	Meta       ChunkMetas `ltype:"nil"`
}

//Part of MapChunkBulk
type ChunkMeta struct {
	X, Z       int32
	PrimaryBit uint16
	AddBitmap  uint16
}

type Explosion struct {
	X       float32
	Y       float32
	Z       float32
	Radius  float32
	Records Records `ltype:"int32"`
	MotionX float32
	MotionY float32
	MotionZ float32
}

//Part of Explosion
type Record struct {
	X byte
	Y byte
	Z byte
}

type Effect struct {
	EffectID        int32
	X               int32
	Y               byte
	Z               int32
	Data            int32
	DisableRelative Bool
}

type SoundEffect struct {
	Name   string
	X      int32
	Y      int32
	Z      int32
	Volume float32
	Pitch  byte
}

type Particle struct {
	Name          string
	X             float32
	Y             float32
	Z             float32
	OffsetX       float32
	OffsetY       float32
	OffsetZ       float32
	ParticleSpeed float32
	Count         int32
}

type GameState struct {
	Reason byte
	Value  float32
}

type SpawnGlobalEntity struct {
	EntityID VarInt
	Type     int8
	X        int32
	Y        int32
	Z        int32
}

type WindowOpen struct {
	WindowID byte
	Type     byte
	Title    string
	Slots    byte
	UseTitle Bool
	EntityID int32 `if:"Type,==,11"`
}

type WindowClose struct {
	WindowID byte
}

type WindowSetSlot struct {
	WindowID byte
	Slot     int16
	Item     Slot
}

type WindowItems struct {
	WindowID byte
	Slots    Slots `ltype:"int16"`
}

type WindowUpdateProperty struct {
	WindowID byte
	Property int16
	Value    int16
}

type WindowTransactionConfirm struct {
	WindowID     byte
	ActionNumber int16
	Accepted     Bool
}

type UpdateSign struct {
	X     int32
	Y     int16
	Z     int32
	Line1 string
	Line2 string
	Line3 string
	Line4 string
}

type Maps struct {
	ItemData VarInt
	Data     Buffer `ltype:"int16"`
}

type UpdateBlockEntity struct {
	X      int32
	Y      int16
	Z      int32
	Action byte
	Data   Buffer `ltype:"int16"`
}

type SignEditorOpen struct {
	X int32
	Y int32
	Z int32
}

type Statistics struct {
	Statistics Statistics `ltype:"varint"`
}

type Statistic struct {
	Name   string
	Amount VarInt
}

type PlayerListItem struct {
	PlayerName string
	Online     Bool
	Ping       int16
}

type PlayerAbilities struct {
	Flags        byte
	FlyingSpeed  float32
	WalkingSpeed float32
}

type TabComplete struct {
	Completions Strings `ltype:"varint"`
}

type ScoreboardObjective struct {
	Name  string
	Value string
	Mode  int8
}

type UpdateScore struct {
	ObjectiveName string
	Mode          int8
	Name          string
	Value         int32
}

type DisplayScoreboard struct {
	Position      int8
	ObjectiveName string
}

type Teams struct {
	Name        string
	Mode        byte
	DisplayName string  `if:"Mode,==,0|2"`
	Prefix      string  `if:"Mode,==,0|2"`
	Suffix      string  `if:"Mode,==,0|2"`
	Flags       byte    `if:"Mode,==,0|2"`
	Players     Strings `if:"Mode,==,0|3|4" ltype:"int16"`
}

type PluginMessage struct {
	Channel string
	Data    Buffer `ltype:"int16"`
}

type Disconnect struct {
	Reason string
}
