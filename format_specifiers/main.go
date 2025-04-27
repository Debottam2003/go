package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b float64 = 8.145
	var name string = "debottam"
	var log bool = true
	var c complex64 = 3 + 2i
	var char1 rune = 'a'
	var char2 byte = 12 // byte is an alias for uint8
	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%d\n%f\n%T\n%T\n", a, b, a, b)
	fmt.Printf("%s\n", name)
	fmt.Printf("%t\n", log)
	fmt.Printf("%v\n", c)
	fmt.Printf("%c\n", char1)
	fmt.Printf("%d\n", char2)
	fmt.Printf("%v\n", arr) 
}
