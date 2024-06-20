package server

import (
	"fmt"
	"log"
	"net"

	"github.com/songgao/water"
)

func Start(iface *water.Interface, port string) {
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
		log.Println("received packet")

		if _, err := iface.Write(buf); err != nil {
			log.Println("failed to write to interface")
		}
		log.Println()
	}
}
