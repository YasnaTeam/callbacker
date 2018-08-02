package server

import (
	"github.com/YasnaTeam/callbacker/common"
	"net"
)

func doStringPacketAction(conn net.Conn, ts *common.TransferableString) {
	data := ts.GetData().(string)
	switch ts.GetCommand() {
	case "register_user":
		registerUserConnection(conn, data)
	default:
		log.Debugf("No string action defined for `%s`.", ts.GetCommand())
	}
}
