package transport

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	Transport

	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Node
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept err %s\n", err)
		}
		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	fmt.Printf("new incoming connection %+v\n", conn)
}
