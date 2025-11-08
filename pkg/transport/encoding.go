package transport

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(reader io.Reader, msg *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(reader io.Reader, msg *RPC) error {
	return gob.NewDecoder(reader).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(reader io.Reader, msg *RPC) error {
	buf := make([]byte, 1024)

	n, err := reader.Read(buf)

	if err != nil {
		return err
	}

	msg.Payload = buf[:n]

	return nil
}
