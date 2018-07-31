package server

import (
	"net"
	"net/http"
	"github.com/YasnaTeam/callbacker/common"
)

func forwardRequestToUser(conn net.Conn, r *http.Request) error {
	log.Debug("Send request to user connection...")

	b, err := common.PrepareRequestToSend(r)
	if err != nil {
		log.Errorf("An error occurred during get bytes from TransferableRequest: `%s`", err)
		return err
	}

	if _, err := common.SendDataToConnection(conn, b); err != nil {
		log.Errorf("An error occurred on forwarding request to user: `%s`", err)
		return err
	}

	return nil
}
