package client

import (
	"net"
	"strconv"
	"github.com/YasnaTeam/callbacker/common"
)

func syncCallbackInformationWithServer(conn net.Conn, route string) error {
	log.Debug("Prepare send `" + route + "` as a route...")

	b, err := common.GetByteFromTransferable(&common.TransferableString{"add_callback", route})
	if err != nil {
		log.Errorf("An error occurred during get bytes from TransferableData: `%s`", err)
		return err
	}

	bs, err := common.SendDataToConnection(conn, b)

	if err != nil {
		log.Errorf("An error occurred during sync callback with server: `%s`", err)
		return err
	}

	log.Debug("`" + route + "` has been sent (" + strconv.Itoa(bs) + ") to the server...")

	return nil
}
