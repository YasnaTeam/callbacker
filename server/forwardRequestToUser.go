package server

import (
	"net"
	"net/http"

	"github.com/YasnaTeam/callbacker/common"
)

func forwardRequestToUser(conn net.Conn, r *http.Request, domain string) error {
	log.Debug("Forward request to user connection...")

	b, err := common.PrepareRequestToSend(r, domain)
	if err != nil {
		log.Errorf("An error occurred during get bytes from TransferableRequest: `%s`", err)
		return err
	}

	bs, err := common.SendDataToConnection(conn, b)
	if err != nil {
		log.Errorf("An error occurred on forwarding request to user: `%s`", err)
		return err
	}

	log.Debugf("%d bytes as a forwarded request has been sent to user.", bs)

	return nil
}
