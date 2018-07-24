package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"strconv"
	"net"
	"github.com/YasnaTeam/callbacker/storage"
	uuid2 "github.com/satori/go.uuid"
)

var log *logrus.Logger
var routes storage.RouteTable
var connections storage.RouteTable

func Initialize(address string, port uint, serverUrl string, logger *logrus.Logger) {
	log = logger
	routes = storage.NewMemoryTable(0)
	connections = storage.NewMemoryTable(0)

	//listenAndServeCallbacks(address)
	go listenAndServeClients(port, serverUrl)
	//go listenAndServeClients(port)
	listenAndServeCallbacks(address)
}

func listenAndServeCallbacks(address string) {
	r := muxRouter()

	log.Info("Server serves requests on " + address + "...")
	log.Fatal(http.ListenAndServe(address, r))
}

func listenAndServeClients(port uint, serverUrl string) {
	portAddress := ":" + strconv.Itoa(int(port))
	log.Info("Server serves clients on " + portAddress)
	ln, err := net.Listen("tcp", portAddress)
	if err != nil {
		log.Fatal("Error on connecting to clients, " + err.Error())
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error on receiving command, " + err.Error())
		}
		log.Debug(conn.RemoteAddr().String() + " connected to the server...")

		go getMessageFromClient(conn, serverUrl)
	}
}

func muxRouter() *mux.Router {
	log.Debug("Gorilla mux started...")

	r := mux.NewRouter()
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/callback/{route_token}", callbacker)

	return r
}

func getMessageFromClient(conn net.Conn, domain string) {
	for {
		var packet []byte = make([]byte, 512)
		_, err := conn.Read(packet)
		if err != nil {
			conn.Close()
			log.Warn("Could not receive packet...")
		}

		var packetStr string = string(packet)
		var packetType = string(packetStr[0])
		log.Debug("Received packet: " + packetStr)

		switch packetType {
		case "0":
			registerUserConnection(conn, packetStr[2:])
		case "1":
			registerCallback(conn, packetStr[2:], domain)
		}
	}
}

func registerUserConnection(conn net.Conn, username string) {
	log.Debug("Trying to add `" + username + "` as a user...")
	connections.Set(username, conn)

	if connections.Has(username) {
		log.Debug("`" + username + "` connection added successfully.")
	}
}

func registerCallback(conn net.Conn, route, domain string) {
	log.Debug("Trying to add `" + route + "` as a route...")

	uuid := uuid2.NewV4().String()
	strCallback := domain + "/" + uuid
	log.Debug("Callback url is: ", strCallback)

	routes.Set(uuid, strCallback)
	conn.Write([]byte(strCallback))
}