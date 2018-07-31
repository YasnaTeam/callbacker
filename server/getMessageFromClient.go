package server

import (
	"net"
	"github.com/YasnaTeam/callbacker/common"
	"encoding/binary"
)

func getMessageFromClient(conn net.Conn, domain string) {
	for {
		// first read packet length
		var packetLength []byte = make([]byte, 4)
		_, err := conn.Read(packetLength)
		if err != nil {
			log.Errorf("Could not read packet length, `%s`", err.Error())
			return
		}

		// convert packet length to int
		packetLengthInt := binary.LittleEndian.Uint32(packetLength)

		// read main data
		var packet []byte = make([]byte, packetLengthInt)
		_, err = conn.Read(packet)
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

		switch packetStruct.GetCommand() {
		case "register_user":
			registerUserConnection(conn, packetStruct.GetData().(string))
		case "add_callback":
			registerCallback(conn, packetStruct.GetData().(string), domain)
		default:
			log.Debugf("No type defined for %v", packetStruct.GetType())
		}
	}
}
