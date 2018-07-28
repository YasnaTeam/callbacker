package server

import (
	"net"
	uuid2 "github.com/satori/go.uuid"
)

func registerCallback(conn net.Conn, route, domain string) {
	log.Debug("Trying to add `" + route + "` as a route...")

	uuid := uuid2.NewV4().String()
	strCallback := domain + "/" + uuid
	log.Debug("Callback url is: ", strCallback)

	username, err := connections.GetKey(conn)

	if err != nil {
		log.Debug("`" + username + "` is not available now.")
		return
	}

	routes.Set(uuid, username)
	conn.Write([]byte(strCallback))
}
