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
			result, err := syncCallbackInformationWithServer(conn, route)
			if err != nil {
				log.Fatal("On sending commands to server, an error occurred! " + err.Error())
			}
			log.Debug("Callback url: " + result)
			addCallbackRunner(result, route)
		default:
			printHelp()
		}

		fmt.Println("")
		fmt.Print("Please select a command (Press h for help): ")
		scanner.Scan()
		command = scanner.Text()
	}
}
