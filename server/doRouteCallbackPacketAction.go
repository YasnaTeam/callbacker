package server

import (
	"net"

	"github.com/YasnaTeam/callbacker/common"
)

func doRouteCallbackPacketAction(conn net.Conn, rc *common.TransferableRouteCallback, domain string) {
	data := rc.GetData().(*common.RouteCallback)
	switch rc.GetCommand() {
	case "add_callback":
		registerCallback(conn, data, domain)
	default:
		log.Debugf("No route-callback action defined for `%s`.", rc.GetCommand())
	}
}
