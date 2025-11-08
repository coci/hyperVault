package transport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	options := TCPTransportOption{
		ListenAddr:    ":3000",
		Decoder:       DefaultDecoder{},
		HandShakeFunc: NOPHandShaker,
	}

	tr := NewTCPTransport(options)

	assert.Equal(t, tr.ListenAddr, ":3000")

	// Server
	assert.Nil(t, tr.ListenAndAccept())
}
