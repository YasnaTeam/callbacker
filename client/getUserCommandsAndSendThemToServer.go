package client

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func getUserCommandsAndSendThemToServer(conn net.Conn) {
	log.Debug("Client is ready to send commands...")

	var command string = ""
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please select a command (Press h for help): ")
	scanner.Scan()
	command = scanner.Text()

	for command != "x" && command != "exit" {

		switch command {
		case "a", "add":
			route := callbackCommand(scanner)
			syncCallbackInformationWithServer(conn, route)
		case "l", "list":
			listUserRoutesAndCallbacksInATable()
		case "t", "truncate":
			truncateRouteTableAndFile()
		default:
			printHelp()
		}

		fmt.Printf("\nPlease select a command (Press h for help): ")
		scanner.Scan()
		command = scanner.Text()
	}
}
