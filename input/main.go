package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"strconv"
)

// Notes:

//     fmt.Scanln() reads input until a newline.

//     fmt.Scan() reads space-separated inputs.

//     fmt.Scanf() allows formatted input like fmt.Scanf("%d %s", &age, &city).

func main() {
	// var age int
	// var city string
	// var number uint32
	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)
	// fmt.Print("Enter your city: ")
	// // fmt.Scanln(&city)
	// fmt.Scan(&city)
	// fmt.Scanln()
	// fmt.Print("Enter your number: ")
	// fmt.Scanf("%d", &number)
	// fmt.Scanln()
	// fmt.Println("Age:", age)
	// fmt.Println("City", city)
	// fmt.Println("Number:", number)
	// fmt.Scan(&age, &city, &number)
	// fmt.Println("Age:", age)
	// fmt.Println("City:", city)
	// fmt.Println("Number:", number)
	// var num1, num2 int8
	// fmt.Scanf("%d%d", &num1, &num2)
	// fmt.Println("Num1:", num1)
	// fmt.Println("Num2:", num2)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove the newline character
	// name = name[:len(name)-1] // Remove the newline character
	fmt.Println("Name:", name)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your age: ")
	scanner.Scan()
	ageinput := scanner.Text()
	age, _ := strconv.Atoi(ageinput)
	fmt.Println("Age:", age)
}
