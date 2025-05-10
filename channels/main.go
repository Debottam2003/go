package main

import "fmt"

func main() {
	messsageChan := make(chan string)
	messsageChan <- "ping"
	msg := <- messsageChan
	fmt.Println(msg)
}