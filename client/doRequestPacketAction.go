package client

import (
	"github.com/YasnaTeam/callbacker/common"
	"net"
)

func doRequestPacketAction(conn net.Conn, tr *common.TransferableRequest, notification func (title, text string)) {
	data := tr.GetData().(*common.Request)
	switch tr.GetCommand() {
	case "callback_forward":
		sendRequestToLocalRoute(data, notification)
	default:
		log.Debugf("No client request action found for `%s`.", tr.GetCommand())
	}
}
