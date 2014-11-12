package protocol

import (
	"encoding/binary"
	"io"
)

func (pkt *Handshake) Read(c *Conn, v McVersion) (err error) {
	if pkt.ProtocolVersion, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.Address, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Port, err = decodeUint16(c, v); err != nil {
		return
	}
	if pkt.State, err = DecodeVarInt(c, v); err != nil {
		return
	}
	return
}
func (pkt *Handshake) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.ProtocolVersion, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Address, c, v); err != nil {
		return
	}
	if err = encodeUint16(pkt.Port, c, v); err != nil {
		return
	}
	if err = EncodeVarInt(pkt.State, c, v); err != nil {
		return
	}
	return
}
func (pkt *StatusResponse) Read(c *Conn, v McVersion) (err error) {
	if pkt.Data, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *StatusResponse) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Data, c, v); err != nil {
		return
	}
	return
}
func (pkt *StatusPing) Read(c *Conn, v McVersion) (err error) {
	if pkt.Time, err = decodeInt64(c, v); err != nil {
		return
	}
	return
}
func (pkt *StatusPing) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt64(pkt.Time, c, v); err != nil {
		return
	}
	return
}
func (pkt *StatusGet) Read(c *Conn, v McVersion) (err error) {
	return
}
func (pkt *StatusGet) Write(c *Conn, v McVersion) (err error) {
	return
}
func (pkt *ClientStatusPing) Read(c *Conn, v McVersion) (err error) {
	if pkt.Time, err = decodeInt64(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientStatusPing) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt64(pkt.Time, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientKeepAlive) Read(c *Conn, v McVersion) (err error) {
	if pkt.KeepAliveID, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientKeepAlive) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.KeepAliveID, c, v); err != nil {
		return
	}
	return
}
func (pkt *ChatMessage) Read(c *Conn, v McVersion) (err error) {
	if pkt.Message, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *ChatMessage) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Message, c, v); err != nil {
		return
	}
	return
}
func (pkt *UseEntity) Read(c *Conn, v McVersion) (err error) {
	if pkt.Target, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Mouse, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *UseEntity) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.Target, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Mouse, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayer) Read(c *Conn, v McVersion) (err error) {
	if pkt.OnGround, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayer) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeBool(pkt.OnGround, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerPosition) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Stance)
	TODO(pkt.Z)
	if pkt.OnGround, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerPosition) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Stance)
	TODO(pkt.Z)
	if err = EncodeBool(pkt.OnGround, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerLook) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Yaw)
	TODO(pkt.Pitch)
	if pkt.OnGround, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerLook) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Yaw)
	TODO(pkt.Pitch)
	if err = EncodeBool(pkt.OnGround, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerPositionLook) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Stance)
	TODO(pkt.Z)
	TODO(pkt.Yaw)
	TODO(pkt.Pitch)
	if pkt.OnGround, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerPositionLook) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Stance)
	TODO(pkt.Z)
	TODO(pkt.Yaw)
	TODO(pkt.Pitch)
	if err = EncodeBool(pkt.OnGround, c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerDigging) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Status)
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Face)
	return
}
func (pkt *PlayerDigging) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Status)
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	TODO(pkt.Face)
	return
}
func (pkt *PlayerBlockPlacement) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Direction, err = DecodeInt8(c, v); err != nil {
		return
	}
	if err = pkt.HeldItem.Read(c, v); err != nil {
		return
	}
	if pkt.CursorPositionX, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.CursorPositionY, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.CursorPositionZ, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerBlockPlacement) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Direction, c, v); err != nil {
		return
	}
	if err = pkt.HeldItem.Write(c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.CursorPositionX, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.CursorPositionY, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.CursorPositionZ, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientHeldItemChange) Read(c *Conn, v McVersion) (err error) {
	if pkt.SlotID, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientHeldItemChange) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt16(pkt.SlotID, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientAnimation) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Animation, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientAnimation) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Animation, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityAction) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.ActionID, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.JumpBoost, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityAction) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.ActionID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.JumpBoost, c, v); err != nil {
		return
	}
	return
}
func (pkt *SteerVehicle) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Sideways)
	TODO(pkt.Forward)
	if pkt.Jump, err = DecodeBool(c, v); err != nil {
		return
	}
	if pkt.Unmount, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *SteerVehicle) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Sideways)
	TODO(pkt.Forward)
	if err = EncodeBool(pkt.Jump, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.Unmount, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientWindowClose) Read(c *Conn, v McVersion) (err error) {
	if pkt.WindowID, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientWindowClose) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeInt8(pkt.WindowID, c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowClick) Read(c *Conn, v McVersion) (err error) {
	if pkt.WindowID, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Slot, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Button, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.ActionNumber, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Mode, err = DecodeInt8(c, v); err != nil {
		return
	}
	if err = pkt.Item.Read(c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowClick) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeInt8(pkt.WindowID, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Slot, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Button, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.ActionNumber, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Mode, c, v); err != nil {
		return
	}
	if err = pkt.Item.Write(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientWindowTransactionConfirm) Read(c *Conn, v McVersion) (err error) {
	if pkt.WindowID, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.ActionNumber, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Accepted, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientWindowTransactionConfirm) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeInt8(pkt.WindowID, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.ActionNumber, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.Accepted, c, v); err != nil {
		return
	}
	return
}
func (pkt *CreativeInventoryAction) Read(c *Conn, v McVersion) (err error) {
	if pkt.Slot, err = decodeInt16(c, v); err != nil {
		return
	}
	if err = pkt.Item.Read(c, v); err != nil {
		return
	}
	return
}
func (pkt *CreativeInventoryAction) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt16(pkt.Slot, c, v); err != nil {
		return
	}
	if err = pkt.Item.Write(c, v); err != nil {
		return
	}
	return
}
func (pkt *EnchantItem) Read(c *Conn, v McVersion) (err error) {
	if pkt.WindowID, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Enchantment, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EnchantItem) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeInt8(pkt.WindowID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Enchantment, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientUpdateSign) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Line1, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Line2, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Line3, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Line4, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientUpdateSign) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line1, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line2, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line3, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line4, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientPlayerAbilities) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Flags)
	TODO(pkt.FlyingSpeed)
	TODO(pkt.WalkingSpeed)
	return
}
func (pkt *ClientPlayerAbilities) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Flags)
	TODO(pkt.FlyingSpeed)
	TODO(pkt.WalkingSpeed)
	return
}
func (pkt *ClientTabComplete) Read(c *Conn, v McVersion) (err error) {
	if pkt.Text, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientTabComplete) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Text, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientSettings) Read(c *Conn, v McVersion) (err error) {
	if pkt.Locale, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.ViewDistance, err = DecodeInt8(c, v); err != nil {
		return
	}
	TODO(pkt.ChatFlags)
	if pkt.Difficulty, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.ShowCape, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientSettings) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Locale, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.ViewDistance, c, v); err != nil {
		return
	}
	TODO(pkt.ChatFlags)
	if err = EncodeInt8(pkt.Difficulty, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.ShowCape, c, v); err != nil {
		return
	}
	return
}
func (pkt *ClientStatuses) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Payload)
	return
}
func (pkt *ClientStatuses) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Payload)
	return
}
func (pkt *ClientPluginMessage) Read(c *Conn, v McVersion) (err error) {
	if pkt.Channel, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *ClientPluginMessage) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Channel, c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *Slot) Read(c *Conn, v McVersion) (err error) {
	if pkt.ID, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Count, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Damage, err = decodeInt16(c, v); err != nil {
		return
	}
	TODO(pkt.Tag)
	return
}
func (pkt *Slot) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt16(pkt.ID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Count, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Damage, c, v); err != nil {
		return
	}
	TODO(pkt.Tag)
	return
}
func (pkt *KeepAlive) Read(c *Conn, v McVersion) (err error) {
	if pkt.KeepAliveID, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *KeepAlive) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.KeepAliveID, c, v); err != nil {
		return
	}
	return
}
func (pkt *JoinGame) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Gamemode)
	if pkt.Dimension, err = DecodeInt8(c, v); err != nil {
		return
	}
	TODO(pkt.Difficulty)
	TODO(pkt.MaxPlayers)
	if pkt.LevelType, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.ReducedDebug, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *JoinGame) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	TODO(pkt.Gamemode)
	if err = EncodeInt8(pkt.Dimension, c, v); err != nil {
		return
	}
	TODO(pkt.Difficulty)
	TODO(pkt.MaxPlayers)
	if err = EncodeString(pkt.LevelType, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.ReducedDebug, c, v); err != nil {
		return
	}
	return
}
func (pkt *ServerMessage) Read(c *Conn, v McVersion) (err error) {
	if pkt.Message, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *ServerMessage) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Message, c, v); err != nil {
		return
	}
	return
}
func (pkt *TimeUpdate) Read(c *Conn, v McVersion) (err error) {
	if pkt.AgeOfTheWorld, err = decodeInt64(c, v); err != nil {
		return
	}
	if pkt.TimeOfDay, err = decodeInt64(c, v); err != nil {
		return
	}
	return
}
func (pkt *TimeUpdate) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt64(pkt.AgeOfTheWorld, c, v); err != nil {
		return
	}
	if err = encodeInt64(pkt.TimeOfDay, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityEquipment) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Slot, err = decodeInt16(c, v); err != nil {
		return
	}
	if err = pkt.Item.Read(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityEquipment) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Slot, c, v); err != nil {
		return
	}
	if err = pkt.Item.Write(c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnPosition) Read(c *Conn, v McVersion) (err error) {
	return
}
func (pkt *SpawnPosition) Write(c *Conn, v McVersion) (err error) {
	return
}
func (pkt *UpdateHealth) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Health)
	if pkt.Food, err = decodeInt16(c, v); err != nil {
		return
	}
	TODO(pkt.FoodSaturation)
	return
}
func (pkt *UpdateHealth) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Health)
	if err = encodeInt16(pkt.Food, c, v); err != nil {
		return
	}
	TODO(pkt.FoodSaturation)
	return
}
func (pkt *Respawn) Read(c *Conn, v McVersion) (err error) {
	if pkt.Dimension, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Difficulty)
	TODO(pkt.Gamemode)
	if pkt.LevelType, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *Respawn) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.Dimension, c, v); err != nil {
		return
	}
	TODO(pkt.Difficulty)
	TODO(pkt.Gamemode)
	if err = EncodeString(pkt.LevelType, c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerPositionLook) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	TODO(pkt.Yaw)
	TODO(pkt.Pitch)
	if pkt.OnGround, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerPositionLook) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	TODO(pkt.Yaw)
	TODO(pkt.Pitch)
	if err = EncodeBool(pkt.OnGround, c, v); err != nil {
		return
	}
	return
}
func (pkt *HeldItemChange) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.SlotID)
	return
}
func (pkt *HeldItemChange) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.SlotID)
	return
}
func (pkt *UseBed) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *UseBed) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	return
}
func (pkt *Animation) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	TODO(pkt.Animation)
	return
}
func (pkt *Animation) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	TODO(pkt.Animation)
	return
}
func (pkt *SpawnPlayer) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.PlayerUUID, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.PlayerName, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Yaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Pitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.CurrentItem, err = decodeInt16(c, v); err != nil {
		return
	}
	TODO(pkt.Metadata)
	return
}
func (pkt *SpawnPlayer) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.PlayerUUID, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.PlayerName, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Yaw, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Pitch, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.CurrentItem, c, v); err != nil {
		return
	}
	TODO(pkt.Metadata)
	return
}
func (pkt *CollectItem) Read(c *Conn, v McVersion) (err error) {
	if pkt.CollectedEntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.CollectorEntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *CollectItem) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.CollectedEntityID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.CollectorEntityID, c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnObject) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.Type, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Pitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Yaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.ExtraData, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.SpeedX, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.SpeedY, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.SpeedZ, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnObject) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Type, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Pitch, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Yaw, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.ExtraData, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.SpeedX, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.SpeedY, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.SpeedZ, c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnMob) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	TODO(pkt.Type)
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Pitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.HeadPitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Yaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.VelocityX, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.VelocityY, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.VelocityZ, err = decodeInt16(c, v); err != nil {
		return
	}
	TODO(pkt.Metadata)
	return
}
func (pkt *SpawnMob) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	TODO(pkt.Type)
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Pitch, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.HeadPitch, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Yaw, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.VelocityX, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.VelocityY, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.VelocityZ, c, v); err != nil {
		return
	}
	TODO(pkt.Metadata)
	return
}
func (pkt *SpawnPainting) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.Title, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Direction, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnPainting) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Title, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Direction, c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnExperienceOrb) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Count, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnExperienceOrb) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Count, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityVelocity) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.VelocityX, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.VelocityY, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.VelocityZ, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityVelocity) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.VelocityX, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.VelocityY, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.VelocityZ, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityDestroy) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.EntityIDs)
	return
}
func (pkt *EntityDestroy) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.EntityIDs)
	if err = EncodeInt8(int8(len(pkt.EntityIDs)), c, v); err != nil {
		return
	}
	return
}
func (pkt *Entity) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *Entity) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityMove) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.DX, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.DY, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.DZ, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityMove) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DX, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DY, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DZ, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityLook) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Yaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Pitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityLook) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Yaw, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Pitch, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityLookMove) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.DX, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.DY, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.DZ, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Yaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Pitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityLookMove) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DX, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DY, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DZ, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Yaw, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Pitch, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityTeleport) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Yaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Pitch, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityTeleport) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Yaw, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Pitch, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityHeadLook) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.HeadYaw, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityHeadLook) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.HeadYaw, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityStatus) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Status, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityStatus) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Status, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityAttach) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.VehicleID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Leash, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityAttach) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.VehicleID, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.Leash, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityMetadata) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Metadata)
	return
}
func (pkt *EntityMetadata) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	TODO(pkt.Metadata)
	return
}
func (pkt *EntityEffect) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.EffectID, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Amplifier, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Duration, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityEffect) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.EffectID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Amplifier, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Duration, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityEffectRemove) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.EffectID, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityEffectRemove) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.EffectID, c, v); err != nil {
		return
	}
	return
}
func (pkt *SetExperience) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.ExperienceBar)
	if pkt.Level, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.TotalExperience, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *SetExperience) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.ExperienceBar)
	if err = encodeInt16(pkt.Level, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.TotalExperience, c, v); err != nil {
		return
	}
	return
}
func (pkt *EntityProperties) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Properties)
	return
}
func (pkt *EntityProperties) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	TODO(pkt.Properties)
	if err = encodeInt32(int32(len(pkt.Properties)), c, v); err != nil {
		return
	}
	return
}
func (pkt *Property) Read(c *Conn, v McVersion) (err error) {
	if pkt.Key, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.Value)
	TODO(pkt.Modifiers)
	return
}
func (pkt *Property) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Key, c, v); err != nil {
		return
	}
	TODO(pkt.Value)
	TODO(pkt.Modifiers)
	if err = encodeInt16(int16(len(pkt.Modifiers)), c, v); err != nil {
		return
	}
	return
}
func (pkt *Modifier) Read(c *Conn, v McVersion) (err error) {
	if pkt.UUIDHigh, err = decodeInt64(c, v); err != nil {
		return
	}
	if pkt.UUIDLow, err = decodeInt64(c, v); err != nil {
		return
	}
	TODO(pkt.Amount)
	if pkt.Operation, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *Modifier) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt64(pkt.UUIDHigh, c, v); err != nil {
		return
	}
	if err = encodeInt64(pkt.UUIDLow, c, v); err != nil {
		return
	}
	TODO(pkt.Amount)
	if err = EncodeInt8(pkt.Operation, c, v); err != nil {
		return
	}
	return
}
func (pkt *ChunkData) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.GroundUp, err = DecodeBool(c, v); err != nil {
		return
	}
	if pkt.PrimaryBitMap, err = decodeUint16(c, v); err != nil {
		return
	}
	if pkt.AddBitMap, err = decodeUint16(c, v); err != nil {
		return
	}
	TODO(pkt.CompressedData)
	return
}
func (pkt *ChunkData) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.GroundUp, c, v); err != nil {
		return
	}
	if err = encodeUint16(pkt.PrimaryBitMap, c, v); err != nil {
		return
	}
	if err = encodeUint16(pkt.AddBitMap, c, v); err != nil {
		return
	}
	TODO(pkt.CompressedData)
	return
}
func (pkt *MultiBlockChange) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.RecordCount, err = decodeInt16(c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *MultiBlockChange) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.RecordCount, c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *BlockChange) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Type, err = DecodeVarInt(c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *BlockChange) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeVarInt(pkt.Type, c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *BlockAction) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Byte1)
	TODO(pkt.byte2)
	if pkt.BlockID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	return
}
func (pkt *BlockAction) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	TODO(pkt.Byte1)
	TODO(pkt.byte2)
	if err = EncodeVarInt(pkt.BlockID, c, v); err != nil {
		return
	}
	return
}
func (pkt *BlockBreakAnimation) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.DestroyStage, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *BlockBreakAnimation) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.DestroyStage, c, v); err != nil {
		return
	}
	return
}
func (pkt *MapChunkBulk) Read(c *Conn, v McVersion) (err error) {
	if pkt.ChunkCount, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.DataLength, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.SkyLight)
	TODO(pkt.Data)
	TODO(pkt.Meta)
	return
}
func (pkt *MapChunkBulk) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt16(pkt.ChunkCount, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.DataLength, c, v); err != nil {
		return
	}
	TODO(pkt.SkyLight)
	TODO(pkt.Data)
	TODO(pkt.Meta)
	return
}
func (pkt *ChunkMeta) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.PrimaryBit, err = decodeUint16(c, v); err != nil {
		return
	}
	if pkt.AddBitmap, err = decodeUint16(c, v); err != nil {
		return
	}
	return
}
func (pkt *ChunkMeta) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = encodeUint16(pkt.PrimaryBit, c, v); err != nil {
		return
	}
	if err = encodeUint16(pkt.AddBitmap, c, v); err != nil {
		return
	}
	return
}
func (pkt *Explosion) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	TODO(pkt.Radius)
	TODO(pkt.Records)
	TODO(pkt.MotionX)
	TODO(pkt.MotionY)
	TODO(pkt.MotionZ)
	return
}
func (pkt *Explosion) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	TODO(pkt.Radius)
	TODO(pkt.Records)
	if err = encodeInt32(int32(len(pkt.Records)), c, v); err != nil {
		return
	}
	TODO(pkt.MotionX)
	TODO(pkt.MotionY)
	TODO(pkt.MotionZ)
	return
}
func (pkt *Record) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	return
}
func (pkt *Record) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	return
}
func (pkt *Effect) Read(c *Conn, v McVersion) (err error) {
	if pkt.EffectID, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Data, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.DisableRelative, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *Effect) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.EffectID, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	TODO(pkt.Y)
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Data, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.DisableRelative, c, v); err != nil {
		return
	}
	return
}
func (pkt *SoundEffect) Read(c *Conn, v McVersion) (err error) {
	if pkt.Name, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Volume)
	TODO(pkt.Pitch)
	return
}
func (pkt *SoundEffect) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Name, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	TODO(pkt.Volume)
	TODO(pkt.Pitch)
	return
}
func (pkt *Particle) Read(c *Conn, v McVersion) (err error) {
	if pkt.Name, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	TODO(pkt.OffsetX)
	TODO(pkt.OffsetY)
	TODO(pkt.OffsetZ)
	TODO(pkt.ParticleSpeed)
	if pkt.Count, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *Particle) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Name, c, v); err != nil {
		return
	}
	TODO(pkt.X)
	TODO(pkt.Y)
	TODO(pkt.Z)
	TODO(pkt.OffsetX)
	TODO(pkt.OffsetY)
	TODO(pkt.OffsetZ)
	TODO(pkt.ParticleSpeed)
	if err = encodeInt32(pkt.Count, c, v); err != nil {
		return
	}
	return
}
func (pkt *GameState) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Reason)
	TODO(pkt.Value)
	return
}
func (pkt *GameState) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Reason)
	TODO(pkt.Value)
	return
}
func (pkt *SpawnGlobalEntity) Read(c *Conn, v McVersion) (err error) {
	if pkt.EntityID, err = DecodeVarInt(c, v); err != nil {
		return
	}
	if pkt.Type, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *SpawnGlobalEntity) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.EntityID, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Type, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowOpen) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	TODO(pkt.Type)
	if pkt.Title, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.Slots)
	if pkt.UseTitle, err = DecodeBool(c, v); err != nil {
		return
	}
	if pkt.EntityID, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowOpen) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	TODO(pkt.Type)
	if err = EncodeString(pkt.Title, c, v); err != nil {
		return
	}
	TODO(pkt.Slots)
	if err = EncodeBool(pkt.UseTitle, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.EntityID, c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowClose) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	return
}
func (pkt *WindowClose) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	return
}
func (pkt *WindowSetSlot) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	if pkt.Slot, err = decodeInt16(c, v); err != nil {
		return
	}
	if err = pkt.Item.Read(c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowSetSlot) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	if err = encodeInt16(pkt.Slot, c, v); err != nil {
		return
	}
	if err = pkt.Item.Write(c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowItems) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	TODO(pkt.Slots)
	return
}
func (pkt *WindowItems) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	TODO(pkt.Slots)
	if err = encodeInt16(int16(len(pkt.Slots)), c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowUpdateProperty) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	if pkt.Property, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Value, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowUpdateProperty) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	if err = encodeInt16(pkt.Property, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Value, c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowTransactionConfirm) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	if pkt.ActionNumber, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Accepted, err = DecodeBool(c, v); err != nil {
		return
	}
	return
}
func (pkt *WindowTransactionConfirm) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.WindowID)
	if err = encodeInt16(pkt.ActionNumber, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.Accepted, c, v); err != nil {
		return
	}
	return
}
func (pkt *UpdateSign) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Line1, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Line2, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Line3, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Line4, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *UpdateSign) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line1, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line2, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line3, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Line4, c, v); err != nil {
		return
	}
	return
}
func (pkt *Maps) Read(c *Conn, v McVersion) (err error) {
	if pkt.ItemData, err = DecodeVarInt(c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *Maps) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeVarInt(pkt.ItemData, c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *UpdateBlockEntity) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt16(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	TODO(pkt.Action)
	TODO(pkt.Data)
	return
}
func (pkt *UpdateBlockEntity) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	TODO(pkt.Action)
	TODO(pkt.Data)
	return
}
func (pkt *SignEditorOpen) Read(c *Conn, v McVersion) (err error) {
	if pkt.X, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Y, err = decodeInt32(c, v); err != nil {
		return
	}
	if pkt.Z, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *SignEditorOpen) Write(c *Conn, v McVersion) (err error) {
	if err = encodeInt32(pkt.X, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Y, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Z, c, v); err != nil {
		return
	}
	return
}
func (pkt *Statistics) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Statistics)
	return
}
func (pkt *Statistics) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Statistics)
	if err = EncodeVarInt(VarInt(len(pkt.Statistics)), c, v); err != nil {
		return
	}
	return
}
func (pkt *Statistic) Read(c *Conn, v McVersion) (err error) {
	if pkt.Name, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Amount, err = DecodeVarInt(c, v); err != nil {
		return
	}
	return
}
func (pkt *Statistic) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Name, c, v); err != nil {
		return
	}
	if err = EncodeVarInt(pkt.Amount, c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerListItem) Read(c *Conn, v McVersion) (err error) {
	if pkt.PlayerName, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Online, err = DecodeBool(c, v); err != nil {
		return
	}
	if pkt.Ping, err = decodeInt16(c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerListItem) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.PlayerName, c, v); err != nil {
		return
	}
	if err = EncodeBool(pkt.Online, c, v); err != nil {
		return
	}
	if err = encodeInt16(pkt.Ping, c, v); err != nil {
		return
	}
	return
}
func (pkt *PlayerAbilities) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Flags)
	TODO(pkt.FlyingSpeed)
	TODO(pkt.WalkingSpeed)
	return
}
func (pkt *PlayerAbilities) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Flags)
	TODO(pkt.FlyingSpeed)
	TODO(pkt.WalkingSpeed)
	return
}
func (pkt *TabComplete) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.Completions)
	return
}
func (pkt *TabComplete) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.Completions)
	if err = EncodeVarInt(VarInt(len(pkt.Completions)), c, v); err != nil {
		return
	}
	return
}
func (pkt *ScoreboardObjective) Read(c *Conn, v McVersion) (err error) {
	if pkt.Name, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Value, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Mode, err = DecodeInt8(c, v); err != nil {
		return
	}
	return
}
func (pkt *ScoreboardObjective) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Name, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Value, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Mode, c, v); err != nil {
		return
	}
	return
}
func (pkt *UpdateScore) Read(c *Conn, v McVersion) (err error) {
	if pkt.ObjectiveName, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Mode, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.Name, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Value, err = decodeInt32(c, v); err != nil {
		return
	}
	return
}
func (pkt *UpdateScore) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.ObjectiveName, c, v); err != nil {
		return
	}
	if err = EncodeInt8(pkt.Mode, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Name, c, v); err != nil {
		return
	}
	if err = encodeInt32(pkt.Value, c, v); err != nil {
		return
	}
	return
}
func (pkt *DisplayScoreboard) Read(c *Conn, v McVersion) (err error) {
	if pkt.Position, err = DecodeInt8(c, v); err != nil {
		return
	}
	if pkt.ObjectiveName, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *DisplayScoreboard) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeInt8(pkt.Position, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.ObjectiveName, c, v); err != nil {
		return
	}
	return
}
func (pkt *Teams) Read(c *Conn, v McVersion) (err error) {
	if pkt.Name, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.Mode)
	if pkt.DisplayName, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Prefix, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Suffix, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.Flags)
	TODO(pkt.Players)
	return
}
func (pkt *Teams) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Name, c, v); err != nil {
		return
	}
	TODO(pkt.Mode)
	if err = EncodeString(pkt.DisplayName, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Prefix, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Suffix, c, v); err != nil {
		return
	}
	TODO(pkt.Flags)
	TODO(pkt.Players)
	if err = encodeInt16(int16(len(pkt.Players)), c, v); err != nil {
		return
	}
	return
}
func (pkt *PluginMessage) Read(c *Conn, v McVersion) (err error) {
	if pkt.Channel, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *PluginMessage) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Channel, c, v); err != nil {
		return
	}
	TODO(pkt.Data)
	return
}
func (pkt *Disconnect) Read(c *Conn, v McVersion) (err error) {
	if pkt.Reason, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *Disconnect) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Reason, c, v); err != nil {
		return
	}
	return
}
func (pkt *LoginDisconnect) Read(c *Conn, v McVersion) (err error) {
	if pkt.Data, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *LoginDisconnect) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Data, c, v); err != nil {
		return
	}
	return
}
func (pkt *EncryptionKeyRequest) Read(c *Conn, v McVersion) (err error) {
	if pkt.ServerID, err = DecodeString(c, v); err != nil {
		return
	}
	TODO(pkt.PublicKey)
	TODO(pkt.VerifyToken)
	return
}
func (pkt *EncryptionKeyRequest) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.ServerID, c, v); err != nil {
		return
	}
	TODO(pkt.PublicKey)
	TODO(pkt.VerifyToken)
	return
}
func (pkt *LoginSuccess) Read(c *Conn, v McVersion) (err error) {
	if pkt.UUID, err = DecodeString(c, v); err != nil {
		return
	}
	if pkt.Username, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *LoginSuccess) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.UUID, c, v); err != nil {
		return
	}
	if err = EncodeString(pkt.Username, c, v); err != nil {
		return
	}
	return
}
func (pkt *LoginStart) Read(c *Conn, v McVersion) (err error) {
	if pkt.Username, err = DecodeString(c, v); err != nil {
		return
	}
	return
}
func (pkt *LoginStart) Write(c *Conn, v McVersion) (err error) {
	if err = EncodeString(pkt.Username, c, v); err != nil {
		return
	}
	return
}
func (pkt *EncryptionKeyResponse) Read(c *Conn, v McVersion) (err error) {
	TODO(pkt.SharedSecret)
	TODO(pkt.VerifyToken)
	return
}
func (pkt *EncryptionKeyResponse) Write(c *Conn, v McVersion) (err error) {
	TODO(pkt.SharedSecret)
	TODO(pkt.VerifyToken)
	return
}
func encodeInt32(i int32, c *Conn, v McVersion) (err error) {
	buf := c.wb[:4]
	binary.BigEndian.PutUint32(buf, uint32(i))
	_, err = c.Out.Write(buf)
	return
}
func decodeInt32(c *Conn, v McVersion) (i int32, err error) {
	buf := c.rb[:4]
	_, err = io.ReadFull(c.In, buf)
	if err != nil {
		return
	}
	i = int32(binary.BigEndian.Uint32(buf))
	return
}
func encodeInt16(i int16, c *Conn, v McVersion) (err error) {
	buf := c.wb[:2]
	binary.BigEndian.PutUint16(buf, uint16(i))
	_, err = c.Out.Write(buf)
	return
}
func decodeInt16(c *Conn, v McVersion) (i int16, err error) {
	buf := c.rb[:2]
	_, err = io.ReadFull(c.In, buf)
	if err != nil {
		return
	}
	i = int16(binary.BigEndian.Uint16(buf))
	return
}
func encodeUint64(i uint64, c *Conn, v McVersion) (err error) {
	buf := c.wb[:8]
	binary.BigEndian.PutUint64(buf, i)
	_, err = c.Out.Write(buf)
	return
}
func decodeUint64(c *Conn, v McVersion) (i uint64, err error) {
	buf := c.rb[:8]
	_, err = io.ReadFull(c.In, buf)
	if err != nil {
		return
	}
	i = binary.BigEndian.Uint64(buf)
	return
}
func encodeInt64(i int64, c *Conn, v McVersion) (err error) {
	buf := c.wb[:8]
	binary.BigEndian.PutUint64(buf, uint64(i))
	_, err = c.Out.Write(buf)
	return
}
func decodeInt64(c *Conn, v McVersion) (i int64, err error) {
	buf := c.rb[:8]
	_, err = io.ReadFull(c.In, buf)
	if err != nil {
		return
	}
	i = int64(binary.BigEndian.Uint64(buf))
	return
}
func encodeUint16(i uint16, c *Conn, v McVersion) (err error) {
	buf := c.wb[:2]
	binary.BigEndian.PutUint16(buf, i)
	_, err = c.Out.Write(buf)
	return
}
func decodeUint16(c *Conn, v McVersion) (i uint16, err error) {
	buf := c.rb[:2]
	_, err = io.ReadFull(c.In, buf)
	if err != nil {
		return
	}
	i = binary.BigEndian.Uint16(buf)
	return
}
func encodeUint32(i uint32, c *Conn, v McVersion) (err error) {
	buf := c.wb[:4]
	binary.BigEndian.PutUint32(buf, i)
	_, err = c.Out.Write(buf)
	return
}
func decodeUint32(c *Conn, v McVersion) (i uint32, err error) {
	buf := c.rb[:4]
	_, err = io.ReadFull(c.In, buf)
	if err != nil {
		return
	}
	i = binary.BigEndian.Uint32(buf)
	return
}
