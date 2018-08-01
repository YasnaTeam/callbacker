package client

import (
	"net"
	"github.com/YasnaTeam/callbacker/common"
	"os"
)

func doActionOnReceivingDataFromServer(conn net.Conn) {
	for {
		packet, err := common.ReceiveDataFromConnection(conn)
		if err != nil {
			log.Errorf("Some errors on receiving data from server occurred: `%s`", err.Error())
			os.Exit(1)
		}

		log.Debug("Received packet: " + string(packet))
		packetStruct, err := common.GetTransferableFromByte(packet)
		if err != nil {
			log.Errorf("There are some errors on receiving packet from server: `%s`", err)
			return
		}

		switch packetStruct.Type {
		case "string":
			doStringPacketAction(conn, packetStruct.Data.(*common.TransferableString))
		case "route_callback":
			doRouteCallbackPacketAction(conn, packetStruct.Data.(*common.TransferableRouteCallback), "")
		case "request":
			doRequestPacketAction(conn, packetStruct.Data.(*common.TransferableRequest))
		default:
			log.Debugf("No type defined on client for %s.", packetStruct.Type)
		}
	}
}
