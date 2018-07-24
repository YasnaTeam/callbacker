package main

import (
	"flag"
	"github.com/YasnaTeam/callbacker/server"
	"github.com/YasnaTeam/callbacker/client"
	"github.com/YasnaTeam/callbacker/logger"
	"github.com/sirupsen/logrus"
)

var isServer bool
var serverAddress string
var serverRepresentationalMainUrl string
var serverPort uint
var debugMode bool
var logLocation string
var log *logrus.Logger


func init() {
	flag.BoolVar(&isServer, "server", false, "make instance as a server")
	flag.BoolVar(&isServer, "s", false, "make instance as a server")

	flag.StringVar(&serverAddress, "host", "127.0.0.1:1616", "set server address in [http://]address[:port] form")
	flag.StringVar(&serverAddress, "a", "127.0.0.1:1616", "set server address in [http://]address[:port] form")

	flag.StringVar(&serverRepresentationalMainUrl, "url", "http://localhost/callback", "A representational url which will be used to client when getting callback url")
	flag.StringVar(&serverRepresentationalMainUrl, "u", "http://localhost/callback", "A representational url which will be used to client when getting callback url")

	flag.UintVar(&serverPort, "port", 2424, "if instance is a server, set the listening port")
	flag.UintVar(&serverPort, "p", 2424, "if instance is a server, set the listening port")

	flag.BoolVar(&debugMode, "debug", false, "enable debug mode")
	flag.BoolVar(&debugMode, "v", false, "enable debug mode")

	flag.StringVar(&logLocation, "log", "/var/log/callbacker/error.log", "set location of error log file")
	flag.StringVar(&logLocation, "l", "/var/log/callbacker/error.log", "set location of error log file")
}

func main() {
	flag.Parse()
	log = logger.New(debugMode)
	log.Debug("log module fired up!")


	if isServer {
		server.Initialize(serverAddress, serverPort, serverRepresentationalMainUrl, log)
	} else {
		client.Initialize(serverPort, log)
	}

	log.Info("Shutting down...")
}