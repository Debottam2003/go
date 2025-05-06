package main

import (
	"fmt"
)

type Car struct {
	brand string
	model string
	year  int32
}

func newCar(brand string, model string, year int32) *Car {
	car := Car{
		brand: brand,
		model: model,
		year: year,
	}
	return &car
}

func (c *Car) setData(data int32) {
	c.year = data
}
type Bmw struct {
	Car   // embedding the Car struct
	color string
}

type rect struct {
	width  float32
	height float32
}

func (rectange rect) area() float32 {
	return rectange.width * rectange.height
}

type Address struct {
	city    string
	state   string
	country string
	zip     uint64
}

type userProfile struct {
	name    string
	age     uint8
	address Address
}

func main() {
	fmt.Println("Hello, World!")
	// anonymous struct
	// anonymous struct is a struct without a name
	myuser := struct {
		name    string
		age     int8
		address struct { // added a field for the nested struct
			city string
		}
	}{
		"debottam",
		22,
		struct {
			city string
		}{"kolkata"},
	}
	fmt.Printf("User Name: %s, Age: %d, Address: %s\n", myuser.name, myuser.age, myuser.address.city)
	var mycar Bmw = Bmw{Car: Car{brand: "BMW", model: "x5", year: 2023}, color: "black"}
	fmt.Println(mycar.brand, mycar.model, mycar.year, mycar.color)
	fmt.Println(mycar)
	var r1 rect = rect{width: 10, height: 10}
	fmt.Println(r1.area())
	fmt.Println(r1.width, r1.height)
	var r2 rect = rect{width: 20, height: 20}
	fmt.Println(r2.area())
	fmt.Println(r2.width, r2.height)
	inf()
	user1 := userProfile{
		name: "deb kar",
		age:  22,
		address: Address{
			city:    "kolkata",
			state:   "west bengal",
			country: "india",
			zip:     700001,
		},
	}
	fmt.Println(user1)
	// var c100 Car = Car{
	// 	brand: "bmw",
	// 	model: "xuv",
	// 	year:  2020,
	// }
	var c100 *Car = newCar("bmw", "xuv", 2020)
	c100.setData(2025)
	fmt.Println(*c100)
}
