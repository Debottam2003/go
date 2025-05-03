package main

import (
	"fmt"
)

func yoyo() {
	fmt.Println("Defer function executed")
}

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

func test() (int32, string) {
	return 5, "deb"
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

func safeDivision(a, b int) int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Printf("%T\n", r)
		}
	}()
	return a / b // If b is 0, it will panic
}

func main() { // Entry point of the program
	fmt.Println("Hello, World!")
	defer yoyo() // Defer function to be executed at the end of main
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

	safeDivision(10, 0)

	// An anonymous function is a function without a name. You can:
	// a := func() { ... }	Anonymous function assigned to a variable
	// func() { ... }()	Anonymous function immediately called

	anonymous := func() {
		fmt.Println("Anonymous function executed")
	}
	anonymous()
	func() {
		fmt.Println(("yo yo"))
	}()
	var score int32
	var name string
	score, name = test()
	fmt.Println(score, name)
	var age, marks int16 = 10, 15
	age, marks = marks, age
	fmt.Println(age, marks)
}
