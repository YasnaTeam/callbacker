package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func muxRouter(domain string) *mux.Router {
	log.Debug("Gorilla mux started...")

	r := mux.NewRouter()
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/callback/{route_token}", func (w http.ResponseWriter, r *http.Request) {
		callbacker(w, r, domain)
	})

	return r
}
