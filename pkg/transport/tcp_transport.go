package transport

import (
	"fmt"
	"net"
	"sync"
)

// TCPNode represents tcp Node interface to establish connection
type TCPNode struct {
	// conn is the underlying connection of nodes
	conn net.Conn

	// if we dial a connection => outbound = true
	// if we accept a connection => outbound = false
	outbound bool
}

func NewTCPNode(conn net.Conn, outbound bool) *TCPNode {
	return &TCPNode{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOption struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	Transport
	TCPTransportOption
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Node
}

func NewTCPTransport(ops TCPTransportOption) *TCPTransport {
	return &TCPTransport{
		TCPTransportOption: ops,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)

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

		fmt.Printf("new incoming connection %+v\n", conn)
		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	node := NewTCPNode(conn, true)

	if err := t.HandShakeFunc(node); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error %s\n", err)
		return
	}

	rpc := &RPC{}

	for {
		err := t.Decoder.Decode(conn, rpc)

		if err != nil {
			fmt.Printf("TCP error %s\n", err)
			continue
		}

		rpc.From = conn.RemoteAddr()

		fmt.Printf("TCP message %+v\n", rpc)
	}

}
