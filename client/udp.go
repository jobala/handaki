package client

import (
	"log"
	"net"

	"github.com/songgao/water"
)

const (
	BUFFER_SIZE = 1500
	MTU         = "1300"
)

func Start(iface *water.Interface, address string) {
	remote, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatalf("address resolution failed: %v", err)
	}

	conn, err := net.DialUDP("udp4", nil, remote)
	if err != nil {
		log.Fatalf("failed to connect to remote: %v", conn)
	}
	defer conn.Close()

	log.Printf("UDP server is listening on %s\n", conn.RemoteAddr().String())

	for {
		buf := make([]byte, BUFFER_SIZE)

		plen, err := iface.Read(buf)
		if err != nil {
			log.Println("failed to read packets")
		}
		log.Println("received packet")

		log.Println("sending packet via udp")
		if _, err = conn.Write(buf[:plen]); err != nil {
			log.Println("failed to send packet")
		}
	}
}
