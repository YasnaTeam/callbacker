package client

import (
	"bufio"
	"fmt"
)

func addCommand(scanner *bufio.Scanner) string {
	var cmd string = ""

	fmt.Println("")
	fmt.Println("In this section you can add a local callback url. After adding the route, you")
	fmt.Println("will recieve a unique callback url which you can use it as a valid callback.")
	fmt.Print("Enter route: ")
	if scanner.Scan() {
		cmd = scanner.Text()
	}

	return cmd
}
