package utils
//
//import (
//	"net"
//	"strconv"
//)
//
//func Read(conn net.Conn) (string, error) {
//	var packetLength = make([]byte, 4)
//	_, err := conn.Read(packetLength)
//	if err != nil {
//		return "", err
//	}
//	length := strconv.Atoi(string(packetLength))
//}
