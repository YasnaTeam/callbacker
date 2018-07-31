package common

import (
	"net"
	"encoding/binary"
)

func SendDataToConnection(conn net.Conn, b []byte) (int, error) {
	packetLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(packetLength, uint32(len(b)))

	// send packet information to server
	_, err := conn.Write(packetLength)
	if err != nil {
		return 0, err
	}

	return conn.Write(b)
}
