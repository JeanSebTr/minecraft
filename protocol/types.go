package protocol

import (
	"encoding/binary"
	"io"
	"math"
)

type McVersion int
type varint int

type Packet interface {
	Write(*Conn, McVersion) error
	Read(*Conn, McVersion) error
}

func TODO(_ interface{}) {
	panic("Not implemented")
}

func EncodeBool(b bool, c *Conn, v McVersion) (err error) {
	bs := c.wb[:1]
	if b {
		bs[0] = 1
	} else {
		bs[0] = 0
	}
	_, err = c.Out.Write(bs)
	return
}

func DecodeBool(c *Conn, v McVersion) (bool, error) {
	bs := c.rb[:1]
	if _, err := io.ReadFull(c.In, bs); err != nil {
		return false, err
	}
	if bs[0] == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

type Byte byte

func (b Byte) Write(c *Conn, v McVersion) (err error) {
	bs := c.wb[:1]
	bs[0] = byte(b)
	_, err = c.Out.Write(bs)
	return
}

func (b *Byte) Read(c *Conn, v McVersion) (err error) {
	bs := c.rb[:1]
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	*b = Byte(bs[0])
	return
}

func EncodeInt8(i int8, c *Conn, v McVersion) (err error) {
	bs := c.wb[:1]
	bs[0] = byte(i)
	_, err = c.Out.Write(bs)
	return
}

func DecodeInt8(c *Conn, v McVersion) (i int8, err error) {
	bs := c.rb[:1]
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	i = int8(bs[0])
	return
}

type Float32 float32

func (f Float32) Write(c *Conn, v McVersion) (err error) {
	bs := c.wb[:4]
	binary.BigEndian.PutUint32(bs, math.Float32bits(float32(f)))
	_, err = c.Out.Write(bs)
	return
}

func (f *Float32) Read(c *Conn, v McVersion) (err error) {
	bs := c.rb[:4]
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	*f = Float32(math.Float32frombits(binary.BigEndian.Uint32(bs)))
	return
}

type Float64 float64

func (f Float64) Write(c *Conn, v McVersion) (err error) {
	bs := c.wb[:8]
	binary.BigEndian.PutUint64(bs, math.Float64bits(float64(f)))
	_, err = c.Out.Write(bs)
	return
}

func (f *Float64) Read(c *Conn, v McVersion) (err error) {
	bs := c.rb[:8]
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	*f = Float64(math.Float64frombits(binary.BigEndian.Uint64(bs)))
	return
}

type VarInt int

func EncodeVarInt(i VarInt, c *Conn, v McVersion) (err error) {
	bs := c.wb[:]
	n := binary.PutUvarint(bs, uint64(i))
	_, err = c.Out.Write(bs[:n])
	return
}

func DecodeVarInt(c *Conn, v McVersion) (i VarInt, err error) {
	x, err := binary.ReadUvarint(byteReader{c.In, c.rb[:1]})
	if err != nil {
		return
	}
	i = VarInt(uint32(x))
	return
}

func ReadVarInt(c *Conn, v McVersion) (b VarInt, err error) {
	x, err := binary.ReadUvarint(byteReader{c.In, c.rb[:1]})
	return VarInt(int32(uint32(x))), err
}

func EncodeString(s string, c *Conn, v McVersion) (err error) {
	str := []byte(s)
	err = EncodeVarInt(VarInt(len(str)), c, v)
	if err != nil {
		return
	}
	_, err = c.Out.Write(str)
	return
}

func DecodeString(c *Conn, v McVersion) (s string, err error) {
	l, err := ReadVarInt(c, v)
	if err != nil {
		return
	}
	bs := make([]byte, l)
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	s = string(bs)
	return
}

type Position struct {
	X int32
	Y int32
	Z int32
}

type Buffer []byte

func (s Buffer) Write(c *Conn, v McVersion) (err error) {
	str := []byte(s)
	err = EncodeVarInt(VarInt(len(str)), c, v)
	if err != nil {
		return
	}
	_, err = c.Out.Write(str)
	return
}

func (b *Buffer) Read(c *Conn, v McVersion) (err error) {
	l, err := ReadVarInt(c, v)
	if err != nil {
		return
	}
	bs := make([]byte, l)
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	*b = Buffer(bs)
	return
}

type Strings []string

type EIDs []int32

type MetaData map[byte]interface{}

func (m MetaData) Write(c *Conn, v McVersion) (err error) {
	panic("no MetaData.Write :(")
	return
}

func (m MetaData) Read(c *Conn, v McVersion) (err error) {
	panic("no MetaData.Read :(")
	return
}
