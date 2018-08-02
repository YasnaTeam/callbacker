package client

import (
	"bufio"
	"os"
	"fmt"
	"net"
)

func getAndRegisterUserOnServer(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your username: ")
	scanner.Scan()
	username = scanner.Text()

	return registerUserOnServer(conn, username)
}
