package client

import (
	"net"

	"github.com/YasnaTeam/callbacker/common"
)

func registerUserConnectionOnServer(conn net.Conn, username string) error {
	log.Debug("Registering user (" + username + ") on server...")

	b, err := common.GetByteFromTransferable(&common.TransferableString{"register_user", username})
	if err != nil {
		log.Errorf("An error occurred during get bytes from TransferableData: `%s`", err)
		return err
	}

	if _, err := common.SendDataToConnection(conn, b); err != nil {
		log.Errorf("An error occurred register username on server: `%s`", err)
		return err
	}

	return nil
}
