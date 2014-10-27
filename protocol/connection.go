package protocol

import (
	"io"
)

type Conn struct {
	In       io.Reader
	Out      io.Writer
	bIn      io.ByteReader
	compress bool
	wb       [10]byte
	rb       [10]byte
}

// func (c *Conn) Init() {
// 	if bIn, ok := c.In.(io.ByteReader); ok {
// 		c.bIn = bIn
// 	} else {
// 		c.bIn = &byteReader{c.In, [1]byte{}}
// 	}
// }

// func (c *Conn) ReadPacket() (pkt *Packet, err error) {
// 	var pktLen, code uint64
// 	if pktLen, err = binary.ReadUvarint(c.bIn); err != nil {
// 		return nil, err
// 	}
// 	if code, err = binary.ReadUvarint(c.bIn); err != nil {
// 		return nil, err
// 	}
// 	log.Printf("packet [%d] #%X", pktLen, code)
// 	return nil, nil

// }

// func (c *Conn) WritePacket(pkt *Packet) error {
// 	return nil
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
