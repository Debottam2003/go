package main

import (
	"fmt"
)

func add1(x, y int) {
	fmt.Println("Sum:", x+y)
}

func add2(x, y float32) {
	fmt.Println("Sum:", x+y)
}

func swap(x, y int) (int, int) {
	return y, x
}

func errorHandling() (string, string) {
	if false {
		return "", "Error occurred"
	} else {
		return "Success", ""
	}
}

func callback(fn func(int) (string, string)) {
	fmt.Println(fn(5))
}

func returnFunction() func(int) string {
	return func(x int) string {
		fmt.Println("Return Function:", x)
		return "Done"
	}
}

func main() { // Entry point of the program
	fmt.Println("Hello, World!")
	add1(2, 7)
	add2(5.0, 3.0)
	var x int = 5
	var y int = 10
	fmt.Println("Before Swapped:", x, y)
	x, y = swap(x, y)
	fmt.Println("After Swapped:", x, y)
	res, err := errorHandling()
	if err != "" {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
	for i := range 10 {
		// fmt.Println(i)
		fmt.Println(fibo(i))
	}
	callback(func(x int) (string, string) {
		fmt.Println("Callback:", x)
		return "Done", "success"
	})
	var rn func(int) string = returnFunction()
	// rn := returnFucntion()
	fmt.Println(rn(10))
}
