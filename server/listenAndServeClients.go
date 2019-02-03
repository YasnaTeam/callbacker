package server

import (
	"net"
)

func listenAndServeClients(port string, serverUrl string) {
	portAddress := addressAndPortGenerator(port)
	log.Info("Server serves clients on " + portAddress)
	ln, err := net.Listen("tcp4", portAddress)
	if err != nil {
		log.Fatal("Error on connecting to clients, " + err.Error())
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error on receiving command, " + err.Error())
		}
		log.Debug(conn.RemoteAddr().String() + " connected to the server...")

		go getMessageFromClient(conn, serverUrl)
	}
}
