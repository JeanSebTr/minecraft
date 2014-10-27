package protocol

import (
	//"encoding/binary"
	"io"
)

type MCVersion int
type VarInt int

type Packet interface {
	Write(*Conn, MCVersion) error
	Read(*Conn, MCVersion) error
}

type Bool bool

func (b Bool) V() bool {
	return bool(b)
}

func (b Bool) Write(c *Conn, v MCVersion) (err error) {
	bs := c.wb[:1]
	if b {
		bs[0] = 1
	} else {
		bs[0] = 0
	}
	_, err = c.Out.Write(bs)
	return
}

func ReadBool(c *Conn, v MCVersion) (b Bool, err error) {
	bs := c.rb[:1]
	_, err = io.ReadFull(c.In, bs)
	if err != nil {
		return
	}
	if bs[0] == 1 {
		b = true
	} else {
		b = false
	}
	return
}

type Position struct {
	X int32
	Y int32
	Z int32
}

type Buffer []byte

type Strings []string

type EIDs []int32

type Properties []Property

type Modifiers []Modifier

type ChunkMetas []ChunkMeta

type Records []Record

type Slots []Slot

type Statistics []Statistic

type MetaData map[byte]interface{}
