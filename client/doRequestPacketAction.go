package client

import (
	"github.com/YasnaTeam/callbacker/common"
	"net"
)

func doRequestPacketAction(conn net.Conn, tr *common.TransferableRequest) {
	data := tr.GetData().(*common.Request)
	switch tr.GetCommand() {
	case "callback_forward":
		sendRequestToLocalRoute(data)
	default:
		log.Debugf("No client request action found for `%s`.", tr.GetCommand())
	}
}
