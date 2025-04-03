package main

import (
	"fmt"
	"proj/controller" // Import your package
)

func main() {
	fmt.Println("Starting the program...")

	// Call functions from the controller package
	controller.Hello()
	controller.Greet("Debottam")
}
