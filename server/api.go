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
	router.HandleFunc("/resolve/{name}", a.handleResolve).Methods(http.MethodGet)
	router.HandleFunc("/reverse-resolve/{address}", a.handleReverseResolve).Methods(http.MethodGet)

	return router
}

func loggingMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

		inner.ServeHTTP(w, r)
	})
}

func (a *api) name(r *http.Request) string {
	return mux.Vars(r)["name"]
}

func (a *api) address(r *http.Request) string {
	return mux.Vars(r)["address"]
}

// test handler
func (a *api) test(w http.ResponseWriter, r *http.Request) {
	testDomain := "hbdgr1234.eth"
	addr, err := a.ens.Resolve(testDomain)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err)
	}

	msg := resolveMsg{Name: testDomain, EthAddr: addr}

	respondWithJSON(w, http.StatusOK, msg)
}

func (a *api) handleResolve(w http.ResponseWriter, r *http.Request) {
	name := a.name(r)

	addr, err := a.ens.Resolve(name)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, errorMsg{err.Error()})
		return
	}

	msg := resolveMsg{Name: name, EthAddr: addr}

	respondWithJSON(w, http.StatusOK, msg)
}

func (a *api) handleReverseResolve(w http.ResponseWriter, r *http.Request) {
	address := a.address(r)

	name, err := a.ens.ReverseResolve(address)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, errorMsg{err.Error()})
		return
	}

	msg := resolveMsg{Name: name, EthAddr: address}

	respondWithJSON(w, http.StatusOK, msg)
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
