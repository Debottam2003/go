package main

import (
	"errors"
	"fmt"
)

func doSomething(flag bool) error {
	if flag {
		return nil // no error
	}
	return errors.New("something went wrong")
}

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	err := doSomething(false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success!")
	}
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		// pointer to an error interface
		fmt.Printf("Error type: %T\n", err)
	} else {
		fmt.Println("result of the division:", res)
	}
}
