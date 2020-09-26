package main

import (
	"ens_feed/config"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

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

	// run http server
	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods(http.MethodGet)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	srv := &http.Server{
		Handler: router,
		Addr:    addr,

		// Good practice: enforce timeouts for servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Starting http server")
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
