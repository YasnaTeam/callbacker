package client

import (
	"net"

	"github.com/YasnaTeam/callbacker/common"
)

func doRouteCallbackPacketAction(conn net.Conn, rc *common.TransferableRouteCallback, domain string, notification func(title, text string)) {
	data := rc.GetData().(*common.RouteCallback)
	switch rc.GetCommand() {
	case "callback_information":
		addCallbackInformation(data.Callback, data.Route, notification)
	default:
		log.Debugf("No client route-callback action defined for `%s`.", rc.GetCommand())
	}
}
