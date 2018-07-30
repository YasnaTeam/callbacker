package client

import (
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"github.com/YasnaTeam/callbacker/storage"
)

var log *logrus.Logger
var username string
var routes storage.RouteTable

func Initialize(port uint, logger *logrus.Logger) {
	log = logger
	log.Info("Client started...")

	routes = storage.NewMemoryTable(0)

	var serverAddress string = "127.0.0.1:" + strconv.Itoa(int(port))
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		log.Fatal("There is an error on resolving address, " + err.Error())
	}

	log.Debug("Client tries to connect to server on " + serverAddress)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Could not connect to server, " + err.Error())
	}
	defer conn.Close()

	registerUserOnServer(conn)

	// All receiving of data is handled by this function
	go doActionOnReceivingDataFromServer(conn)

	// print menu and send information to server if needed
	getUserCommandsAndSendThemToServer(conn)
}
