package common

import (
	"net"
	"encoding/binary"
)

func ReceiveDataFromConnection(conn net.Conn) ([]byte, error) {
	// first read packet length
	var packetLength []byte = make([]byte, 4)
	_, err := conn.Read(packetLength)
	if err != nil {
		return nil, err
	}

	// convert packet length to int
	packetLengthInt := binary.LittleEndian.Uint32(packetLength)

	// read main data
	var packet []byte = make([]byte, packetLengthInt)
	_, err = conn.Read(packet)
	if err != nil {
		return nil, err
	}

	return packet, nil
}
