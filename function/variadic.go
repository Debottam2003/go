package main

import (
	"fmt"
)

// func variadic(nums ...int32) {
// 	fmt.Println("This is a variadic function:")
// 	for _, val := range nums{
// 	fmt.Println(val);
// 	}
// }

func variadic(nums ...any) { // interface{} -> aby (modern)
	fmt.Println("This is a variadic function:")
	for _, val := range nums {
		fmt.Println(val)
	}
}

func math(n ...int32) {
	fmt.Println(n)
}
