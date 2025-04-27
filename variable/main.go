package main

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello, World!")
}

type userProfile struct {
	Name    string
	Age     int
	Email   string
	Address string
	Phone   int64
}

func main() {
	// Variable declaration and initialization and data types in go
	// Integer types
	var a int = 10
	var b int8 = 20
	var c int16 = 30
	var d int32 = 40
	var e int64 = 50
	fmt.Println("This is Integer:", a, b, c, d, e)

	// Unsigned integer types
	var f uint = 10
	var g uint8 = 20
	var h uint16 = 30
	var i uint32 = 40
	var j uint64 = 50
	fmt.Println("This is unsigned Integer:", f, g, h, i, j)

	// Floating point types
	var k float32 = 10.5
	var l float64 = 20.5
	fmt.Println("This is floating:", k, l)

	//Rune type (Unicode code point) Alias for int32
	var char1 rune = 'A'
	var char2 rune = '界'
	var char3 rune = 'অ'
	var char4 rune = '🙂'
	fmt.Println("These are runes:", char1, char2, char3, char4)

	// Complex types
	var m complex64 = 1 + 2i
	var n complex128 = 3 + 2i
	fmt.Println("This is complex:", m, n)

	// Boolean type
	var loggedin bool = true
	fmt.Println("This is boolean:", loggedin)

	// String type
	var name string = "debottam kar"
	fmt.Println("This is String:", name)
	str := "Debottam"
	fmt.Println(str[0:3]) // "Deb"
	fmt.Println(str[4:])  // "ttam"
	fmt.Println(str[:5])  // "Debot"

	// Array type
	var marks [7]int = [7]int{99, 96, 95, 91, 90, 85, 84}
	fmt.Println("This is an Array:", marks)

	// Slice type
	var fruits []string = []string{"apple", "banana", "orange"}
	fmt.Println("This is a Slice:", fruits)

	// Map type
	var student map[string]int = map[string]int{"debottam": 1, "kar": 2}
	fmt.Println(student)
	student["age"] = 22
	fmt.Println(student)
	delete(student, "kar")
	fmt.Println("This is a Map:", student)

	Fruits := []string{"apple", "banana", "orange", "grape"}
	// Removing the element at index 1 (banana)
	Fruits = append(Fruits[:1], Fruits[2:]...)
	fmt.Println("This is a slice after deletion:", Fruits)

	// var newArr []string = Fruits[:]
	// newArr := make([]string, len(Fruits))
	// copy(newArr, Fruits)
	newArr := Fruits[:]
	fmt.Println("Original:", Fruits)
	fmt.Println("Copy:", newArr)

	// Pointer type
	var ptr *int = &a
	fmt.Println("This is a pointer to an integer:", ptr)
	fmt.Println("This is a pointer value:", *ptr)

	// Function type
	Hello()

	// Strcut type
	var user1 = userProfile{Name: "Debottam", Age: 22, Email: "debottamkar2003@gmail.com", Address: "Kolkata-India-743165", Phone: 9073584850}
	fmt.Println("This is a structure:", user1)
}
