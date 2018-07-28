package client

import (
	"bufio"
	"os"
	"fmt"
	"net"
)

func registerUserOnServer(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your username: ")
	scanner.Scan()
	username = scanner.Text()

	if err := registerUserConnectionOnServer(conn, username); err != nil {
		log.Fatal("Could not register username on server, " + err.Error())
	}
}
