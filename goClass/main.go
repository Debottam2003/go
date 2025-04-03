package main

import (
	"fmt"
)
type Message struct {
	Text string
	Sender User
	Receiver User
}
type User struct {
	Username string
}
func main(){
	fmt.Println("hello world")
	var msg Message = Message{"This is a message", User{"Alice"}, User{"Bob"}}
	fmt.Println(msg);
}