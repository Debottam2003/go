package main

import (
	"fmt"
)

func main() {
	concat := StringConcat() // Now it can find the function
	concat("Hello")
	concat("world")
	concat("from")
	fmt.Println(concat("Go"))
}
