package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 10
	// if-else if-else statement
	// logical operations &&, ||, !
	if x := 10; x > 15 || x > 100 {
		fmt.Println("x is greater than 15")
	} else if x > 5 && x <= 15 {
		fmt.Println("x is greater than 5 but less than or equal to 15")
	} else {
		fmt.Println("x is less than or equal to 5")
	}
	var loggedin bool = true
	fmt.Println("User logged in:", !loggedin)

	// switch statement
	switch x := 6; x {
	case 1:
		fmt.Println("x is 1")
		// break // no need go compiler under the hood done this
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	case 4:
		fmt.Println("x is 4")
	case 5:
		fmt.Println("x is 5")
	case 6:
		// os.Exit(0)
		break
	default:
		fmt.Println("x is greater than 5")
	}

	// multiple condition switch case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is weekend.")
	default:
		fmt.Println("This is working day.")
	}

	// type switch
	// anonymous function 
	test := func(i any) { // intefcae{} == any
		switch i.(type) {
		case int:
			fmt.Println("i is of type int")
		case string:
			fmt.Println("i is of type string")
		case bool:
			fmt.Println("i is of type bool")
		case float64:
			fmt.Println("i is of type float64")
		default:
			fmt.Println("i is of an unknown type")
		}
	}
	test(true)

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
	// fmt.Println("Is it before now?:", t1.Before(time.Now().Add(1*time.Hour)))
	// fmt.Println("Is it after now?:", t1.After(time.Now().Add(-1*time.Hour)))

}
