package common

import (
	"encoding/binary"
	"net"
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
	var packet []byte = make([]byte, 0, packetLengthInt)
	var buf []byte = make([]byte, 64)
	var sum uint32 = 0

	for sum != packetLengthInt {
		n, err := conn.Read(buf)
		if err != nil {
			return nil, err
		}

		sum += uint32(n)
		packet = append(packet, buf[:n]...)
	}

	return packet, nil
}
