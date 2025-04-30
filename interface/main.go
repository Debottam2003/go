package main

import "fmt"

type Car struct {
	brand string
	model string
	year  int
}
type Bmw struct {
	Car   // embedding the Car struct
	color string
}

func main() {
	// fmt.Println("Hello, World!")
	// // anonymous struct
	// // anonymous struct is a struct without a name
	// myuser := struct {
	// 	name    string
	// 	age     int8
	// 	address struct { // added a field for the nested struct
	// 		city string
	// 	}
	// }{
	// 	"debottam",
	// 	22,
	// 	struct {
	// 		city string
	// 	}{"kolkata"},
	// }
	// fmt.Printf("User Name: %s, Age: %d, Address: %s\n", myuser.name, myuser.age, myuser.address.city)
	var mycar Bmw = Bmw{Car: Car{brand: "BMW", model: "x5", year: 2023}, color: "black"}
	fmt.Println(mycar.brand, mycar.model, mycar.year, mycar.color)
	fmt.Println(mycar)
}
