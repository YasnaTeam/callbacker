package server

import (
	"net/http"
	"github.com/YasnaTeam/callbacker/logger"
	"net"
	"strings"
)

func callbacker(w http.ResponseWriter, r *http.Request, domain string) {
	log.Debug("Handler `callbacker` has been ran.")
	requestlog := logger.WithRequest(log, r)
	url := r.URL.Path[1:]
	token := strings.Split(url, "/")

	requestlog.Debug(url + " has been browsed.")

	username, err := routes.Get(token[1])
	if err != nil {
		log.Debug("Related user to `" + token[1] + "` not found.")
		return
	}

	conn, err := connections.Get(username.(string))
	if err != nil {
		log.Debug("No alive connection found related to `" + username.(string) +"`.")
	}

	forwardRequestToUser(conn.(net.Conn), r, domain)
}