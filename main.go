package main

import (
	"fmt"
)

func main() {
	// var name string
	// var city string
	// fmt.Scan(&age) // do not consumes the white spaces
	// fmt.Scanf("%s", &city) // \n stays in the buffer for fmt.Scan too but Scanln flushes the \n
	// fmt.Println(city)
	//fmt.Scanln() // flushes \n
	// fmt.Scanln(&name) // consumes the first whitespace
	// fmt.Println(name)
	var a, b int
	// fmt.Scanf("%d%d", &a, &b)
	fmt.Scanf("%d\n", &a) // 3\n  The newline (\n) is consumed (because the format expects a newline after the integer).
	fmt.Scanf("%d\n", &b) // 4\n
	// fmt.Scanf("%d", &a)
	// fmt.Scanln()
	// fmt.Scanf("%d", &b)
	fmt.Println(a, b)
}
