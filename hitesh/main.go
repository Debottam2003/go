package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string
	Age      int16
}

func main() {
	fmt.Println("We will learn about conversion in to json")
	// Struct to json
	var u1 User = User{UserName: "Debottam Kar", Age: 22}
	var u2 User = User{UserName: "Sritama Kar", Age: 26}
	var u3 User = User{UserName: "Soma Kar", Age: 49}
	var u4 User = User{UserName: "Goutam Kar", Age: 49}
	var array []User = []User{u1, u2, u3, u4}
	data, err := json.MarshalIndent(array, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// data is byte slice
	fmt.Println(data)
	fmt.Println(string(data))
	// json to struct
	var jsonData string = `[
 {
  "UserName": "Debottam Kar",
  "Age": 22
 },
 {
  "UserName": "Sritama Kar",
  "Age": 26
 },
 {
  "UserName": "Soma Kar",
  "Age": 49
 },
 {
  "UserName": "Goutam Kar",
  "Age": 49
 }
]`
	err = json.Unmarshal(data, &array)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonData)
	fmt.Println(array)
}
