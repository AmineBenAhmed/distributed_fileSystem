package main

import (
	"log"

	"github.com/AmineBenAhmed/foreverfs/p2p"
)

func main() {
	tcpOpts := p2p.TCPTRansortOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Printf("Returning ...")
		log.Fatal(err)
	}

	select {}
}
