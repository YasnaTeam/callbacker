package client

import (
	"github.com/YasnaTeam/callbacker/common"
	"net"
)

func doStringPacketAction(conn net.Conn, ts *common.TransferableString) {
	//data := ts.GetData().(string)
	switch ts.GetCommand() {
	case "register_user":
	default:
		log.Debugf("No client string action defined for `%s`.", ts.GetCommand())
	}
}
