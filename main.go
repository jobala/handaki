package main

import (
	"log"

	"github.com/jobala/handaki/client"
	"github.com/jobala/handaki/server"
	"github.com/songgao/water"
)

func main() {

	tun, err := water.New(water.Config{DeviceType: water.TUN})
	if err != nil {
		log.Fatal("Failed to open tun device")
	}

	isServer := false
	if isServer {
		server.Start(tun, "")
	} else {
		client.Start(tun, "")
	}
}
