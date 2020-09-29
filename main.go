package main

import (
	"ens_feed/config"
	"ens_feed/eth"
	"ens_feed/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.Flags() | log.Lmicroseconds | log.Lshortfile)

	cfgPath := "./config.yml"
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config: %#v", cfg) // dbg

	ens, err := eth.NewNameService(cfg.Eth.InfuraURL, cfg.Eth.EnsContractAddr)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(cfg.Server.Host, cfg.Server.Port, ens)
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
