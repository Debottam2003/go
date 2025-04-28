package main

import (
// "fmt"
)

var memory [100]int

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	if memory[n] != 0 {
		return memory[n]
	}
	memory[n] = fibo(n-1) + fibo(n-2)
	return memory[n]
}
