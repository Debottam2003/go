package main

import (
	"fmt"
	"maps"
	"slices"
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
	char4 := '🙂' // warlus operator
	fmt.Println("These are runes:", char1, char2, char3, char4)

	// Byte type (Alias for uint8)
	var char5 byte = 'A'
	fmt.Println("This is a byte:", char5)

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
	m1 := make(map[string]string)
	m1["name"] = "deb kar"
	m1["address"] = "naihati"
	fmt.Println(m1, len(m1))

	val, ok := m1["name"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("No such key.")
	}

	m2 := m1
	fmt.Println(m2)
	fmt.Println(maps.Equal(m1, m2))

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

	// Constant type
	const pi float32 = 3.14
	fmt.Println("This is a constant:", pi)

	// Multiple assignment and swap
	var ageNew, marksNew int16 = 10, 15
	ageNew, marksNew = marksNew, ageNew
	fmt.Println(ageNew, marksNew)

	// constant grouping
	const (
		port int32  = 5000
		host string = "localhost"
	)
	fmt.Println("PORT:", port, "HOST:", host)

	// 2d array
	matrix := [2][2]int{{2, 3}, {4, 1}}
	matrix[1][1] = 10
	fmt.Println(matrix)

	sl1 := []int32{1, 2, 3}
	// sl := make([]int32, 2, 5)
	// fmt.Println(sl, len(sl), cap(sl))
	// sl2 := make([]int32, len(sl1))
	// copy(sl2, sl1)
	sl2 := sl1[:]
	fmt.Println(sl2)
	fmt.Println(slices.Equal(sl1, sl2))

	type Map map[string]interface{}
	var jsonData map[string]interface{} = map[string]interface{}{
		"name": "Debottam Kar",
		"age":  22,
		"address": map[string]interface{}{
			"city":  "Kolkata",
			"state": "West bengal",
		},
	}
	fmt.Println(jsonData)
}
