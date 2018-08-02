package client

import (
	"net"
	"github.com/YasnaTeam/callbacker/common"
	"fmt"
)

func callCallbackOnRecieveDataFromServer(conn net.Conn) {
	var d []byte = make([]byte, 512)
	n, err := conn.Read(d)
	if err != nil {
		log.Errorf("There are some errors on reading callbacks from server: `%s`", err)
	}

	log.Debugf("%d bytes read from server.", n)

	receivedData, err := common.PrepareReceivedData(d)
	if err == nil {
		fmt.Println(receivedData)
	}
}
