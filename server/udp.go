package server

import (
	"fmt"
	"log"
	"net"
)

func Start(port string) {
	p, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		log.Fatalf("address resolution failed: %v", err)
	}

	conn, err := net.ListenUDP("udp4", p)
	if err != nil {
		log.Fatalf("failed to connect to remote: %v", conn)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		if _, _, err := conn.ReadFromUDP(buf); err != nil {
			fmt.Println("failed to read from UDP")
		}
	}
}
