package server

import (
	"net"
	"github.com/YasnaTeam/callbacker/common"
	"fmt"
)

func getMessageFromClient(conn net.Conn, domain string) {
	for {
		var packet []byte = make([]byte, 1024)
		_, err := conn.Read(packet)
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
		packetStruct, err := common.GetTransferableDataFromByte(packet)
		fmt.Println(packetStruct)
		if err != nil {
			log.Errorf("There are some errors on receiving packet from client: `%s`", err)
			return
		}

		switch packetStruct.Type {
		case "register_user":
			registerUserConnection(conn, packetStruct.Data.(string))
		case "add_callback":
			registerCallback(conn, packetStruct.Data.(string), domain)
		}
	}
}
