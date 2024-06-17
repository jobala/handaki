package main

import (
	"fmt"
	"log"

	"github.com/songgao/water"
)

const (
	BUFFER_SIZE = 1500
	MTU         = "1300"
)

func main() {
	tun, err := water.New(water.Config{DeviceType: water.TUN})
	fmt.Println(tun)
	if err != nil {
		log.Fatal("Failed to open tun device")
	}

	for {
		buf := make([]byte, BUFFER_SIZE)

		if _, err = tun.Read(buf); err != nil {
			log.Println("failed to read packets")
		}

		fmt.Println(buf)
	}

}
