package main

import (
	"log"

	"github.com/sid995/godfs/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":4344")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
