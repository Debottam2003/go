package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
)

func helloWorld() {
	fmt.Println("hello world")
}
func add(a int, b int) int {
	return a + b
}
func main() {
	helloWorld()
	// var name string = "Debottam"
	// var age int = 22
	// var isgarduated bool = true
	// fmt.Println("Name:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Graduated:", isgarduated)
	// if 1 > 0 {
	// 	fmt.Println("1 is greater than 0")
	// } else {
	// 	fmt.Println("earth is dying...")
	// }
	// nickname := "Rony"
	// fmt.Println("Nickname:", nickname)
	// day := "Monday"

	// switch day {
	// case "Monday":
	// 	fmt.Println("Start of the week.")
	// case "Friday":
	// 	fmt.Println("Weekend incoming!")
	// default:
	// 	fmt.Println("A regular day.")
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Count:", i)
	// }
	// var result int = add(10, 12)
	// fmt.Println("The addition result: ", result)
	// // var x int
	// // var y int
	// // fmt.Scan(&x)
	// // fmt.Scan(&y)
	// // fmt.Println(x, y);
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your full name: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Hello,", input)
	// input = strings.TrimSpace(input)
	// number, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println("Invalid number:", err)
	// }

	// fmt.Println("You entered:", number)
	// var arr [3]int = [3]int{1, 2, 3}
    // fmt.Println("Array:", arr)
	// var slice []int= []int{1, 2, 3}
    // slice = append(slice, 4, 5)
    // fmt.Println("Slice:", slice)
	// var matrix[10][10] int
	// for i := 0; i < 10; i++{
	// 	for j := 0; j < 10; j++{
	// 		matrix[i][j] = i * j
	// 	}
	// }
	// fmt.Println(matrix)
	fmt.Print("Debottam Kar's nick name is ")
	fmt.Println("...Rony...")
	// * pattern
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	var index uint16 = 0
	for {
		fmt.Println("I am a loop in go...");
		if index > 10{
			break
		}
		index++
	}
	iterator := 1
	for iterator < 10 {
		fmt.Println(iterator)
		iterator++
	}
}
