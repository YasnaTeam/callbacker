package client

import (
	"github.com/YasnaTeam/callbacker/common"
	"net"
)

func doRouteCallbackPacketAction(conn net.Conn, rc *common.TransferableRouteCallback, domain string) {
	data := rc.GetData().(*common.RouteCallback)
	switch rc.GetCommand() {
	case "callback_information":
		addCallbackInformation(data.Callback, data.Route)
	default:
		log.Debugf("No client route-callback action defined for `%s`.", rc.GetCommand())
	}
}
