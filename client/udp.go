package client

import (
	"log"
	"net"
)

func Start(address string) {
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

	if _, err = conn.Write([]byte("some data")); err != nil {
		log.Println("failed to write some data")
	}
}
