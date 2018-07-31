package client

import (
	"net"
	"github.com/YasnaTeam/callbacker/common"
)

func syncCallbackInformationWithServer(conn net.Conn, route string) error {
	log.Debug("Prepare send `" + route + "` as a route...")

	b, err := common.GetByteFromTransferable(&common.TransferableRouteCallback{"add_callback", common.RouteCallback{Route: route}})
	if err != nil {
		log.Errorf("An error occurred during get bytes from TransferableData: `%s`", err)
		return err
	}

	bs, err := common.SendDataToConnection(conn, b)

	if err != nil {
		log.Errorf("An error occurred during sync callback with server: `%s`", err)
		return err
	}

	log.Debugf("`%s` has been sent (%d bytes) to the server...", route, bs)

	return nil
}
