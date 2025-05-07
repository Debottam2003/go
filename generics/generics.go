package main

import "fmt"

func print[T comparable](arr []T) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

type UserProfile[T comparable] struct {
	data T
}

func main() {
	print([]int{10, 30, 20})
	print([]string{"hello", "golang"})
	// user := UserProfile[string]{data: "debottam"}
	user := UserProfile[int32]{data: 10}
	fmt.Println(user)
}