package client

import (
	"net"

	"github.com/YasnaTeam/callbacker/common"
)

func doStringPacketAction(conn net.Conn, ts *common.TransferableString, notification func(title, text string)) {
	//data := ts.GetData().(string)
	switch ts.GetCommand() {
	case "register_user":
	default:
		log.Debugf("No client string action defined for `%s`.", ts.GetCommand())
	}
}
