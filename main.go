package main

import (
	"fmt"
	"gosample/config"
	"gosample/tun"
)

func someFn(cfg config.Config) {
	err := cfg.Download()
	if err != nil {
		fmt.Println("Error!!!")
	}
}

func main() {
	fmt.Println("HiHiHi....")
	cfg := config.ConfigFS{}
	someFn(&cfg)
	err := tun.CreateTUN()
	fmt.Println("CreateTUN err:", err)
	tun.LocalInterfaces()
	fmt.Println("-----------------------")
}
