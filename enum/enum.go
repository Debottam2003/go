package main

import "fmt"

// type status string

// const (
// 	Received status = "received"
// 	Pending status = "pending"
// 	Delayed status = "delayed"
// )

const (
	Received1 string = "Received"
	Pending1  string = "Pending"
	Delayed1  string = "Delayed"
)

type Status int

const (
	Received Status = iota
	Pending
	Delayed
)

func (s Status) String() string {
	names := [3]string{"Received", "Pending", "Delayed"}
	if int(s) < 0 || int(s) >= len(names) {
		return "Unknown"
	}
	return names[s]
}

func main() {
	var s Status = Received
	fmt.Println(s.String())
	fmt.Println(Received1)
}
