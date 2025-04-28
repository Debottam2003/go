package main

import "fmt"

func swapRef(a *int32, b *int32) {
	// call by reference
	temp := *a
	*a = *b
	*b = temp
}

func swapVal(a int32, b int32) {
	// call by value
	temp := a
	a = b
	b = temp
	fmt.Println("After Swap in 'swapVal' function:", a, b)
}

func print(arr []int) {
	fmt.Println(arr) // Print he slice
}

func show(arr *[]int) {
	fmt.Println(arr)
	fmt.Println(*arr)
	fmt.Println((*arr)[0])
	fmt.Println((*arr)[2])
}

func main() {
	// Declare an int32 variable and assign it a value
	var a int32 = 42
	fmt.Println(a)  // Print the value of 'a'
	fmt.Println(&a) // Print the memory address of 'a'

	// Declare a pointer to int32 and assign it the address of 'a'
	var ptr *int32 = &a
	fmt.Println(ptr)  // Print the memory address stored in 'ptr'
	fmt.Println(*ptr) // Dereference 'ptr' to get the value of 'a'

	// Compare the dereferenced value of 'ptr' with 'a'
	fmt.Println(*ptr == a) // This will print true

	// Modify the value of 'a' through the pointer
	*ptr = 100
	fmt.Println(a)    // Print the updated value of 'a'
	fmt.Println(*ptr) // Print the updated value through the pointer
	fmt.Println(&ptr) // Print the memory address stored of 'ptr'

	var pointer **int32 = &ptr // Declare a pointer to a pointer and assign it the address of 'ptr'
	fmt.Println(pointer)       // Print the memory address stored in 'pointer'
	fmt.Println(*pointer)      // Dereference 'pointer' to get the value of 'ptr'

	var x int32 = 10
	var y int32 = 20
	fmt.Println("Before Swap:", x, y)
	swapRef(&x, &y)
	fmt.Println("After Swap:", x, y)

	var p int32 = 30
	var q int32 = 20
	fmt.Println("Before Swap:", p, q)
	swapVal(p, q)
	fmt.Println("After Swap in 'main' function:", p, q)

	name := "rony"
	var nameptr *string = &name
	fmt.Println(&name)
	fmt.Println(nameptr)

	var arr []int = []int{10, 20, 30, 40, 50}
	fmt.Println(&arr, &arr[0])
	print(arr)
	show(&arr)
}
