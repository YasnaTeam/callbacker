package client

import (
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"fmt"
	"os"
	"bufio"
)

var log *logrus.Logger
var username string

func Initialize(port uint, logger *logrus.Logger) {
	log = logger
	log.Info("Client started...")

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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your username: ")
	scanner.Scan()
	username = scanner.Text()

	if err := registerUserConnectionOnServer(conn, username); err != nil {
		log.Fatal("Could not register username on server, " + err.Error())
	}

	getCommands(conn)
}

func registerUserConnectionOnServer(conn net.Conn, username string) error {
	log.Debug("Registering user (" + username + ") on server...")

	command := "0:" + username
	if _, err := conn.Write([]byte(command)); err != nil {
		return err
	}

	return nil
}

func getCommands(conn net.Conn) {
	log.Debug("Client is ready to send commands...")

	var command string = ""
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please select a command (Press h for help): ")
	scanner.Scan()
	command = scanner.Text()

	for command != "x" && command != "exit" {

		switch command {
		case "a", "add":
			route := addCommand(scanner)
			result, err := sendCommandToServer(conn, route)
			if err != nil {
				log.Fatal("On sending commands to server, an error occurred! " + err.Error())
			}

			fmt.Println("Callback url: " + result)
		default:
			printHelp()
		}

		fmt.Println("")
		fmt.Print("Please select a command (Press h for help): ")
		scanner.Scan()
		command = scanner.Text()
	}
}

func addCommand(scanner *bufio.Scanner) string {
	var cmd string = ""

	fmt.Println("")
	fmt.Println("In this section you can add a local callback url. After adding the route, you")
	fmt.Println("will recieve a unique callback url which you can use it as a valid callback.")
	fmt.Print("Enter route: ")
	if (scanner.Scan()) {
		cmd = scanner.Text()
	}

	return cmd
}

func sendCommandToServer(conn net.Conn, route string) (string, error) {
	log.Debug("Prepare send `" + route + "` as a route...")
	command := "1:" + username + ":" + route
	var callback []byte

	bs, err := conn.Write([]byte(command))
	if err != nil {
		return "", err
	}
	log.Debug("`" + route + "` has been sent (" + strconv.Itoa(bs) + ") to the server...")

	br, err := conn.Read(callback)
	if err != nil {
		return "", nil
	}
	log.Debug("A response received (" + strconv.Itoa(br) + ") from the server...")

	return string(callback), nil
}

func printHelp() {
	fmt.Println("Use this commands:")
	fmt.Println("    a (or add)\tAdd a route")
	fmt.Println("    x (or exit)\tExit from program")
}
