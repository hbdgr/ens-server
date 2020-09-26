package main

import (
	"ens_feed/config"
	"ens_feed/server"
	"log"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	log.SetFlags(log.Flags() | log.Lmicroseconds | log.Lshortfile)

	cfgPath := "./config.yml"
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// dbg
	log.Printf("Config: %#v", cfg)

	client, err := ethclient.Dial(cfg.Eth.InfuraURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Eth connection ready")
	_ = client

	srv := server.NewServer(cfg.Server.Host, cfg.Server.Port)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Starting http server..")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(quit, os.Interrupt)

	// Block until we receive our signal.
	<-quit

	log.Println("Shutting down")
}
