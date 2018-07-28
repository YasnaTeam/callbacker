package client

import (
	"net"
	"strconv"
)

func sendCommandToServer(conn net.Conn, route string) (string, error) {
	log.Debug("Prepare send `" + route + "` as a route...")
	command := "1:" + username + ":" + route
	var callback = make([]byte, 512)

	bs, err := conn.Write([]byte(command))
	if err != nil {
		return "", err
	}
	log.Debug("`" + route + "` has been sent (" + strconv.Itoa(bs) + ") to the server...")

	br, err := conn.Read(callback)
	if err != nil {
		return "", nil
	}
	log.Debug("A response received (" + strconv.Itoa(br) + ") from the server...")

	return string(callback), nil
}
