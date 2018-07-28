package server

import "github.com/gorilla/mux"

func muxRouter() *mux.Router {
	log.Debug("Gorilla mux started...")

	r := mux.NewRouter()
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/callback/{route_token}", callbacker)

	return r
}
