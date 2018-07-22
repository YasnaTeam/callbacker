package main

import (
	"flag"
	"github.com/YasnaTeam/callbacker/server"
	"github.com/YasnaTeam/callbacker/client"
	"github.com/sirupsen/logrus"
	"github.com/YasnaTeam/callbacker/logger"
)

var isServer bool
var serverAddress string
var serverPort uint
var debugMode bool
var logLocation string
var log *logrus.Logger


func init() {
	flag.BoolVar(&isServer, "server", false, "make instance as a server")
	flag.BoolVar(&isServer, "s", false, "make instance as a server")

	flag.StringVar(&serverAddress, "host", "http://127.0.0.1:1616", "set server address in [http://]address[:port] form")
	flag.StringVar(&serverAddress, "a", "http://127.0.0.1:1616", "set server address in [http://]address[:port] form")

	flag.UintVar(&serverPort, "port", 1616, "if instance is a server, set the listening port")
	flag.UintVar(&serverPort, "p", 1616, "if instance is a server, set the listening port")

	flag.BoolVar(&debugMode, "debug", false, "enable debug mode")
	flag.BoolVar(&debugMode, "v", false, "enable debug mode")

	flag.StringVar(&logLocation, "log", "/var/log/callbacker/error.log", "set location of error log file")
	flag.StringVar(&logLocation, "l", "/var/log/callbacker/error.log", "set location of error log file")
}

func main() {
	flag.Parse()
	log = logger.New(debugMode)


	if isServer {
		server.Initialize(serverPort, log)
	} else {
		client.Initialize(serverAddress, log)
	}

	log.Info("Shutting down...")
}