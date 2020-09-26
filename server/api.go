package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// NewServer creates and returns http server
func NewServer(host, port string) *http.Server {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.HandleFunc("/hello", hello).Methods(http.MethodGet)

	addr := host + ":" + port
	return &http.Server{
		Handler: router,
		Addr:    addr,

		// Good practice: enforce timeouts for servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func loggingMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

		inner.ServeHTTP(w, r)
	})
}

// test handler
func hello(w http.ResponseWriter, r *http.Request) {
	type hello struct {
		Msg string `json:"msg"`
	}

	respondWithJSON(w, http.StatusOK, hello{Msg: "hello"})
}

func respondWithJSON(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if status != http.StatusOK {
			w.WriteHeader(status)
		}
		w.Write(body)
	}
}
