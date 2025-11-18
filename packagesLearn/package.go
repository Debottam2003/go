package main

import (
	"fmt"
	"github.com/Debottam2003/my_go_app/inner"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Time package features
	t1 := time.Now()
	fmt.Println("Current time:", t1)
	fmt.Println("Current Year:", t1.Year())
	fmt.Println("Current Month:", t1.Month())
	fmt.Println("Current Day:", t1.Day())
	fmt.Println("Current Hour:", t1.Hour())
	fmt.Println("Current Minute:", t1.Minute())
	fmt.Println("Current Second:", t1.Second())
	fmt.Println("Unix Timestamp:", t1.Unix())
	fmt.Println("Formatted Time:", t1.Format("2003-01-02 15:04:05"))

	// time.Sleep(time.Second * 3)

	str := "debottam kar it is"
	// for _, v := range str {
	// 	fmt.Print(string(v))
	// 	time.Sleep(time.Second)
	// }

	fmt.Println(str)

	// for i := 0; i < len(str); i++ {
	// 	fmt.Print(string(str[i]))
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println()

	// Time package features
	n1 := math.Abs(-3)
	n2 := math.Floor(3.2)
	n3 := math.Ceil(3.1)
	n4 := math.Pow(2, 3)
	fmt.Println(n1, n2, n3, n4)
	// rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 99
	randomNumber := rand.Intn(100)

	fmt.Println("Random number:", randomNumber)

	inner.Utils()

}
