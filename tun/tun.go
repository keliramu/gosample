package tun

import (
	"fmt"
	"net"

	"golang.zx2c4.com/wireguard/tun"
)

func CreateTUN() error {
	_, err := tun.CreateTUN("lxtun", 1420)
	if err != nil {
		return fmt.Errorf("creating tun device: %w", err)
	}
	return nil
}

func LocalInterfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("%+v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		fmt.Println("Network interface:", i.Name)
	}
}
