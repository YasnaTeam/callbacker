package client

import (
	"net"
	"github.com/YasnaTeam/callbacker/common"
	"encoding/binary"
)

func registerUserConnectionOnServer(conn net.Conn, username string) error {
	log.Debug("Registering user (" + username + ") on server...")

	b, err := common.GetByteFromTransferable(&common.TransferableString{"register_user", username})
	if err != nil {
		log.Errorf("An error occurred during get bytes from TransferableData: `%s`", err)
		return err
	}


	packetLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(packetLength, uint32(len(b)))
	_, err = conn.Write(packetLength)
	if err != nil {
		log.Errorf("An error occurred during send request length to server: `%s`", err)
		return err
	}
	if _, err := conn.Write(b); err != nil {
		log.Errorf("An error occurred register username on server: `%s`", err)
		return err
	}

	return nil
}
