package client

import (
	"net"
)

func registerUserOnServer(conn net.Conn, username string) (string, error) {
	if err := registerUserConnectionOnServer(conn, username); err != nil {
		log.Fatal("Could not register username on server, " + err.Error())
		return "", err
	}

	return username, nil
}
