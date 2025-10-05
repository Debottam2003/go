package main

import (
	//"bufio"
	"fmt"
	//"os"
)

type userProfile struct {
	name string
	age  int
}

// The user here is the receiver variable, representing the instance of userProfile on which the method is called.
func (user userProfile) Display(id string) string {
	return user.name + " " + fmt.Sprintf("%d", user.age) + " " + "id: " + id
}
func helloWorld() {
	fmt.Println("hello world")
}
func main() {
	//variables
	// fmt.Println("Hello, World!")
	// var name string = "debottam kar"
	// fmt.Println(name)
	// var loggedin bool = true
	// fmt.Println(loggedin)
	// var marks float32 = 81.143
	// fmt.Println(marks)
	// var comp complex64 = 2+4i
	// fmt.Println(comp)
	// name = "rony"
	// fmt.Println(name)
	// price := 1234 // walrus operator
	// fmt.Println(price)
	// const pi float32 = 3.141
	// fmt.Println(pi)
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n');
	// fmt.Println(input)
	// var arr = [4]int{7, 8, 9, 2}
	// fmt.Println(arr)
	helloWorld()
	var user1 = userProfile{name: "Debottam", age: 21}
	var user2 = userProfile{name: "Alice", age: 25}
	var user3 = userProfile{name: "Bob", age: 30}

	user1.name = "rony"
	fmt.Println(user1)
	fmt.Println(user1.Display("1"))

	fmt.Println(user2)
	fmt.Println(user2.Display("2"))

	fmt.Println(user3)
	fmt.Println(user3.Display("3"))
	main2()
}
