package main

import (
	"ens_feed/config"
	"fmt"
	"log"
)

func main() {
	cfgPath := "./config.yml"
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// dbg
	fmt.Printf("Config: %#v\n", cfg)
}
