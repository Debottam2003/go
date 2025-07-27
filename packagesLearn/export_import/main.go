package main

import (
	"fmt" 
	"myapp/inner"
)

func main() {
	fmt.Println("Hello world")
	// add()
	// Sub()
	inner.AddNested()
	inner.SubNested()
}