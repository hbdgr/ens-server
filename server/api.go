package server

import (
	"encoding/json"
	"ens_feed/eth"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type api struct {
	ens *eth.NameService
}

// NewServer creates and returns http server with api handlers
func NewServer(host, port string, ens *eth.NameService) *http.Server {
	a := &api{ens: ens}
	handlers := newRouter(a)

	addr := host + ":" + port
	return &http.Server{
		Handler: handlers,
		Addr:    addr,

		// Good practice: enforce timeouts for servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func newRouter(a *api) http.Handler {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.HandleFunc("/test", a.test).Methods(http.MethodGet)

	return router
}

func loggingMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

		inner.ServeHTTP(w, r)
	})
}

// test handler
func (a *api) test(w http.ResponseWriter, r *http.Request) {
	type hello struct {
		Msg     string `json:"msg"`
		TestENS string `json:"test_ens"`
	}
	addr, err := a.ens.ResolveTest()
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
	}

	h := hello{Msg: "test ens", TestENS: addr}

	respondWithJSON(w, http.StatusOK, h)
}

// helper
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
