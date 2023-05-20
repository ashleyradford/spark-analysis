package main

import (
	"fmt"
	"net"
	"os"
)

func start(filename string) {

}

func main() {
	buffer := make(chan string, 32)

	fmt.Printf("buffer: %v\n", buffer)

	listener, err := net.Listen("tcp", ":25009")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on port: 25009")
	for {
		if conn, err := listener.Accept(); err == nil {
			fmt.Println(conn)
		}
	}

}
