package main

import (
	"log"
	"os"

	"github.com/jobala/handaki/client"
	"github.com/jobala/handaki/server"
	"github.com/songgao/water"
)

func main() {

	flag := os.Args[len(os.Args)-1]
	isServer := flag == "server"

	tun, err := water.New(water.Config{DeviceType: water.TUN})
	if err != nil {
		log.Fatal("Failed to open tun device")
	}

	if isServer {
		server.Start(tun, ":4000")
	} else {
		client.Start(tun, "")
	}
}
