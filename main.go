package main

import (
	"fmt"
	"log"

	"github.com/jobala/handaki/client"
	"github.com/jobala/handaki/server"
	"github.com/songgao/water"
	"golang.org/x/net/ipv4"
)

const (
	BUFFER_SIZE = 1500
	MTU         = "1300"
)

func main() {
	isServer := false
	if isServer {
		server.Start()
	} else {
		client.Start()
	}

	tun, err := water.New(water.Config{DeviceType: water.TUN})
	fmt.Println(tun)
	if err != nil {
		log.Fatal("Failed to open tun device")
	}

	for {
		buf := make([]byte, BUFFER_SIZE)

		n, err := tun.Read(buf)
		if err != nil {
			log.Println("failed to read packets")
		}

		header, _ := ipv4.ParseHeader(buf[:n])
		fmt.Printf("read %d bytes from device %s\n", header.TotalLen, tun.Name())
	}
}
