package server

import (
	"net"
	"net/http"
	"github.com/YasnaTeam/callbacker/common"
)

func forwardRequestToUser(conn net.Conn, r *http.Request) {
	log.Debug("Send request to user connection...")
	marshalledRequest, err := common.PrepareDataToSend(r)
	if err == nil {
		n, err := conn.Write(marshalledRequest)
		if err != nil {
			log.Error("There are some errors on forward callback to server", err)
			return
		}

		log.Debugf("%d bytes of request, has been forwarded to the client", n)
	}
}
