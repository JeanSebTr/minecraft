package protocol

import (
	"io"
)

type byteReader struct {
	io.Reader
	buf []byte
}

func (b byteReader) ReadByte() (byte, error) {
	// bs := b.buf[:]
	// _, err := b.Read(bs)
	_, err := b.Read(b.buf)
	return b.buf[0], err
}
