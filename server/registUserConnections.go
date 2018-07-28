package server

import "net"

func registerUserConnection(conn net.Conn, username string) {
	log.Debug("Trying to add `" + username + "` as a user...")
	connections.Set(username, conn)

	if connections.Has(username) {
		log.Debug("`" + username + "` connection added successfully.")
	}
}
