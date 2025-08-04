// socket = ip + port

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// type user struct {
// 	name string
// }

// type head *user // custom type head is reference of user structure instance

func do(connection net.Conn) {
	time.Sleep(5 * time.Second)
	buff := make([]byte, 1024)
	_, err := connection.Read(buff)
	ErrorHandler(err)
	connection.Write([]byte("hello world!"))
	connection.Close()
}

func main() {

	// var ref head = &user{name: "debottam"}
	// fmt.Println(ref)

	// TCP connection
	listener, err := net.Listen("tcp", ":3333")
	ErrorHandler(err)
	// fmt.Println(listener)
	fmt.Println("Server is running and listening on port: 3333")

	for {
		connect, err := listener.Accept()
		ErrorHandler(err)
		// fmt.Println(connect)
		go do(connect)

		fmt.Println("request came")
	}
}
