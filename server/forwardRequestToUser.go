package server

import (
	"net"
	"net/http"
)

func forwardRequestToUser(conn net.Conn, r *http.Request) {
	conn.Write([]byte("Hellooooooo"))
}
