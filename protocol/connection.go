package protocol

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type Conn struct {
	In             io.Reader
	Out            io.Writer
	bIn            io.ByteReader
	compress       bool
	wb             [10]byte
	rb             [10]byte
	ReadDirection  Direction
	WriteDirection Direction
	State          State
}

// func (c *Conn) Init() {
// 	if bIn, ok := c.In.(io.ByteReader); ok {
// 		c.bIn = bIn
// 	} else {
// 		c.bIn = &byteReader{c.In, [1]byte{}}
// 	}
// }

func (c *Conn) ReadPacket() (pkt Packet, err error) {
	var pktLen, code VarInt
	v := McVersion(0)
	if pktLen, err = DecodeVarInt(c, v); err != nil {
		return nil, err
	}
	if code, err = DecodeVarInt(c, v); err != nil {
		return nil, err
	}

	st := packets[c.State][c.ReadDirection]
	if code < 0 || int(code) >= len(st) {
		return nil, fmt.Errorf("Invalid packet %02X", code)
	}
	ty := st[code]
	if ty == nil {
		return nil, fmt.Errorf("Invalid packet %02X", code)
	}

	pkt, _ = reflect.New(ty).Interface().(Packet)
	if err = pkt.Read(c, v); err != nil {
		return nil, err
	}

	fmt.Printf("packet [%d] #0x%02X\n", pktLen, code)

	return

}

func (c *Conn) WritePacket(pkt Packet) error {

	var buf bytes.Buffer
	temp := c.Out
	c.Out = &buf

	ty := reflect.TypeOf(pkt).Elem()

	code, ok := packetsToID[c.WriteDirection][ty]
	if !ok {
		panic("Invalid Packet")
	}
	v := McVersion(0)

	// write packet to buffer
	EncodeVarInt(VarInt(code), c, v)
	pkt.Write(c, v)

	// write length + packet to network
	c.Out = temp
	EncodeVarInt(VarInt(buf.Len()), c, v)
	buf.WriteTo(c.Out)
	return nil
}

// func (src *Conn) ProxyPacket(dst *Conn) (err error) {

// }

// func (c *Conn) SetCompression() {
// 	panic("fasd")
// }

// type byteReader struct {
// 	io.Reader
// 	buf [1]byte
// }

// func (b byteReader) ReadByte() (byte, error) {
// 	bs := b.buf[:]
// 	_, err := b.Read(bs)
// 	return bs[0], err
// }
