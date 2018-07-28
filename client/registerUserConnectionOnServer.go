package client

import "net"

func registerUserConnectionOnServer(conn net.Conn, username string) error {
	log.Debug("Registering user (" + username + ") on server...")

	command := "0:" + username
	if _, err := conn.Write([]byte(command)); err != nil {
		return err
	}

	return nil
}
