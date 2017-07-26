package main

import (
	"fmt"
	"net"
	"os"
)

const (
	defaultPort = "3333"
)

func handleFatalError(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		handleFatalError(err)
	}
	defer l.Close()

	fmt.Printf("Started TCP server on port %s\n", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			handleFatalError(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 2048)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Closed connection:", err)
			return
		}
		conn.Write([]byte("Received message: "))
		conn.Write(buf)
	}
}
