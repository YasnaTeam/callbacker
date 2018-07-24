package server

import (
	"net/http"
	"github.com/YasnaTeam/callbacker/logger"
	"fmt"
)

func callbacker(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handler `callbacker` has been ran.")
	requestlog := logger.WithRequest(log, r)

	token := r.URL.Path[1:]
	requestlog.Debug(token + " has been browsed.")
	fmt.Fprintf(w, "Hi there, I love %s!", token)
}