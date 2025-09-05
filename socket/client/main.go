package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3333")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send a message
	fmt.Fprintf(conn, "Hello from client.")

	// Read server reply
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Server replied:", string(buf[:n]))
}
