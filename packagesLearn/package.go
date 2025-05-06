package main

import (
	"fmt"
	"math"
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
	fmt.Println("Formatted Time:", t1.Format("2006-01-02 15:04:05"))
	// Time package features
	n1 := math.Abs(-3)
	n2 := math.Floor(3.2)
	n3 := math.Ceil(3.1)
	n4 := math.Pow(2, 3)
	n5 := math.IsNaN(9)
	n6 := math.IsNaN(1.34)
	fmt.Println(n1, n2, n3, n4, n5, n6)
}
