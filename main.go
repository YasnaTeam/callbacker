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
var serverLocalAddress string
var serverRepresentationalMainUrl string
var serverPort string
var debugMode bool
var logLocation string
var showNotification bool
var log *logrus.Logger


func init() {
	flag.BoolVar(&isServer, "server", false, "make instance as a server")
	flag.BoolVar(&isServer, "s", false, "make instance as a server")

	flag.StringVar(&serverLocalAddress, "host", "127.0.0.1:1616", "set server local address in [http://]address[:port] format")
	flag.StringVar(&serverLocalAddress, "a", "127.0.0.1:1616", "set server local address in [http://]address[:port] format")

	flag.StringVar(&serverAddress, "callbacker-server", "callback.site:2424", "set server local address in address:port format")
	flag.StringVar(&serverAddress, "c", "callback.site:2424", "set server local address in address:port format")

	flag.StringVar(&serverRepresentationalMainUrl, "url", "http://callback.site/callback", "A representational url which will be used to client when getting callback url")
	flag.StringVar(&serverRepresentationalMainUrl, "u", "http://callback.site/callback", "A representational url which will be used to client when getting callback url")

	flag.StringVar(&serverPort, "port", ":2424", "if instance is a server, set the listening port in [callback.site|x.x.x.x:]port format")
	flag.StringVar(&serverPort, "p", ":2424", "if instance is a server, set the listening port in [callback.site|x.x.x.x:]port format")

	flag.BoolVar(&debugMode, "debug", false, "enable debug mode")
	flag.BoolVar(&debugMode, "v", false, "enable debug mode")

	flag.BoolVar(&showNotification, "n", false, "Show notification on client for events")
	flag.BoolVar(&showNotification, "notification", false, "Show notification on client for events")

	flag.StringVar(&logLocation, "log", "/var/log/callbacker/error.log", "set location of error log file")
	flag.StringVar(&logLocation, "l", "/var/log/callbacker/error.log", "set location of error log file")
}

func main() {
	flag.Parse()
	log = logger.New(debugMode)
	log.Debug("log module fired up!")


	if isServer {
		server.Initialize(serverLocalAddress, serverPort, serverRepresentationalMainUrl, log)
	} else {
		notification := client.NewNotification(showNotification, log)
		client.Initialize(serverAddress, log, notification)
	}

	log.Info("Shutting down...")
}