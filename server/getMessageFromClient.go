package server

import (
	"net"

	"github.com/YasnaTeam/callbacker/common"
)

func getMessageFromClient(conn net.Conn, domain string) {
	for {
		packet, err := common.ReceiveDataFromConnection(conn)
		if err != nil {
			conn.Close()
			log.Warn("Could not receive packet...")
			username, err := connections.GetKey(conn)
			if err == nil {
				log.Debug("Connection of `" + username + "` has been closed.")
				connections.Set(username, nil)
			}
			return
		}

		log.Debug("Received packet: " + string(packet))
		packetStruct, err := common.GetTransferableFromByte(packet)
		if err != nil {
			log.Errorf("There are some errors on receiving packet from client: `%s`", err)
			return
		}

		switch packetStruct.Type {
		case "string":
			doStringPacketAction(conn, packetStruct.Data.(*common.TransferableString))
		case "route_callback":
			doRouteCallbackPacketAction(conn, packetStruct.Data.(*common.TransferableRouteCallback), domain)
		case "request":
			doRequestPacketAction(conn, packetStruct.Data.(*common.TransferableRequest))
		default:
			log.Debugf("No type defined for %s.", packetStruct.Type)
		}
	}
}
