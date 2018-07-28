package server

import "net"

func getMessageFromClient(conn net.Conn, domain string) {
	for {
		var packet []byte = make([]byte, 512)
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

		var packetStr string = string(packet)
		var packetType = string(packetStr[0])
		log.Debug("Received packet: " + packetStr)

		switch packetType {
		case "0":
			registerUserConnection(conn, packetStr[2:])
		case "1":
			registerCallback(conn, packetStr[2:], domain)
		}
	}
}
