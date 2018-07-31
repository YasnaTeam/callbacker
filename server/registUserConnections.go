package server

import "net"

func registerUserConnection(conn net.Conn, username string) {
	log.Debugf("Trying to add `%s` as a user...", username)
	connections.Set(username, conn)

	if connections.Has(username) {
		log.Debugf("`%s` connection added successfully.", username)
	}
}
