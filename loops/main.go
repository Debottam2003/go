package main

import "fmt"

func main() {
	// This is a while loop.
	index := 1
	for index <= 5 {
		fmt.Println("This is while loop.")
		index++
	}
	// This is for loop.
	for i := 1; i <= 5; i++ {
		fmt.Println("This ia a for loop.")
	}
	// This is a for loop.
	for {
		fmt.Println("This is an infinite loop.")
		fmt.Println("This is an infinite loop.")
		fmt.Println("This is an infinite loop.")
		fmt.Println("This is an infinite loop.")
		fmt.Println("This is an infinite loop.")
		break // This will break the infinite loop
	}
	// for - range lopp.
	nums := []int{10, 20, 30, 40, 50}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	for _, value := range nums {
		fmt.Printf("Value: %d\n", value) // Only prints values
		// '_' this is blank identifier, it ignores the index value.
		// It is used when we don't need the value that we get.
	}
	for index := range nums {
		fmt.Println(index) // Only prints indexes
	}
	studentGrades := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 92,
	}

	for name, grade := range studentGrades {
		fmt.Printf("Name: %s, Grade: %d\n", name, grade)
	}
}
