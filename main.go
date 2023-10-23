package main

import (
	"fmt"
	"gosample/config"
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
}
