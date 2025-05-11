package main

import (
	"fmt"
	"sync"
)

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum:", sum);
}

func multiply(wg *sync.WaitGroup) {
	defer wg.Done()
	mul := 1
	for i := 1; i <= 10; i++ {
		mul *= i
	}
	fmt.Println("multiply:", mul);
}

func main() {
	// messsageChan := make(chan string)
	// messsageChan <- "ping"
	// msg := <- messsageChan
	// fmt.Println(msg)
	var wg sync.WaitGroup
	wg.Add(1)
	go add(&wg)
	wg.Add(1)
	go multiply(&wg)
	wg.Wait()
}