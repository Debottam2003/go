package main

import (
	"fmt"
)

func main() {
	// x := 10
	// if-else if-else statement
	if x := 10; x > 15 || x > 100 {
		fmt.Println("x is greater than 15")
	} else if x > 5 && x <= 15 {
		fmt.Println("x is greater than 5 but less than or equal to 15")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// switch statement
	switch x := 6; x {
	case 1:
		fmt.Println("x is 1")
		// break // no need go compiler under the hood done this
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	case 4:
		fmt.Println("x is 4")
	case 5:
		fmt.Println("x is 5")
	case 6:
		// os.Exit(0)
		break
	default:
		fmt.Println("x is greater than 5")
	}

	var loggedin bool = true
	fmt.Println("User logged in:", !loggedin)
}
