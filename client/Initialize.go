package client

import (
	"net"

	"github.com/YasnaTeam/callbacker/storage"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var username string
var routes storage.RouteTable // map callback url to route
var configuration *Configuration

func Initialize(address string, logger *logrus.Logger, notification func(title, text string)) {
	log = logger
	log.Info("Client started...")

	routes = storage.NewMemoryTable(0)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Could not connect to server, " + err.Error())
	}
	defer conn.Close()

	configuration, err = GetConfiguration()
	if err != nil {
		log.Errorf("There is an error on getting configuration file, `%s`.", err.Error())
	}
	defer configuration.Close()

	if configuration.Username == "" {
		username, err := getAndRegisterUserOnServer(conn)

		if err == nil {
			configuration.Username = username
			if err := configuration.Save(); err != nil {
				log.Errorf("There is an error on creating user on configurations, `%s`.", err.Error())
			}
		}
	} else {
		log.Debugf("User `%s` has been selected from configurations file.", configuration.Username)
		registerUserOnServer(conn, configuration.Username)
		setSavedCallbacksOnClientRouteTable()
	}

	// All receiving of data is handled by this function
	go doActionOnReceivingDataFromServer(conn, notification)

	// print menu and send information to server if needed
	getUserCommandsAndSendThemToServer(conn)
}
