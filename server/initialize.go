package server

import (
	"github.com/sirupsen/logrus"
	"github.com/YasnaTeam/callbacker/storage"
)

var log *logrus.Logger
var routes storage.RouteTable // map routes to users
var connections storage.RouteTable // map users to connections

func Initialize(address string, port uint, serverUrl string, logger *logrus.Logger) {
	log = logger
	routes = storage.NewMemoryTable(0)
	connections = storage.NewMemoryTable(0)

	go listenAndServeClients(port, serverUrl)
	listenAndServeCallbacks(address)
}
