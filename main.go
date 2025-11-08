package main

import (
	"fmt"
	"log"

	"github.com/coci/hyperVault/pkg/transport"
)

func main() {
	fmt.Println("we are running")

	tcpOps := transport.TCPTransportOption{
		ListenAddr:    ":3000",
		Decoder:       transport.DefaultDecoder{},
		HandShakeFunc: transport.NOPHandShaker,
	}

	tr := transport.NewTCPTransport(tcpOps)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
