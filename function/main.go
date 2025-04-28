package main

import "fmt"

func add(x, y int) {
	fmt.Println("Sum:", x + y)
}

func main() {// Entry point of the program
	fmt.Println("Hello, World!")
	add(5, 3)
}