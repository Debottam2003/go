package main

import "fmt"

type shape interface {
	area() float32
	perimeter() float32
}

type square struct {
	width  float32
	height float32
}

func (s square) area() float32 {
	return s.width * s.height
}
func (s square) perimeter() float32 {
	return 2 * (s.width + s.height)
}

func inf() {
	var s1 shape = square{8.32, 8.32}
	areaRes := s1.area()
	perimeterRes := s1.perimeter()
	fmt.Println("Area of the square is:", areaRes)
	fmt.Println("Perimeter of the square is:", perimeterRes)
	// var sptr *shape = new(shape)
	// *sptr = s1
	sptr := &s1
	fmt.Println("Pointer to the shape interface:", sptr)
	fmt.Println("Memory address of the shape interface", &s1)
	fmt.Println(s1)
	fmt.Println(*sptr)
	if sq, ok := s1.(square); ok {
		fmt.Println("Type assertion successful. Square:", sq)
	} else {
		fmt.Println("Type assertion failed.")
	}
}
