package client

import "fmt"

func printHelp() {
	fmt.Println("Use this commands:")
	fmt.Println("    a (or add)\tAdd a route")
	fmt.Println("    l (or list)\tPrint route table")
	fmt.Println("    t (or truncate)\tTruncate route table and configuration file")
	fmt.Println("    x (or exit)\tExit from program")
}
