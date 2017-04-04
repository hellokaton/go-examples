package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input:\"chat [server|client] [:port|IP address:port]\"")
		os.Exit(-1)
	}

	if os.Args[1] == "server" {
		startServer(os.Args[2])
	} else if os.Args[1] == "client" {
		startClient(os.Args[2])
	} else {
		fmt.Println("Wrong param")
		os.Exit(-1)
	}
	fmt.Println(os.Args[1])
	fmt.Println("Hello World!")
}
