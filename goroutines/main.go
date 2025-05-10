package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Printf("Task %d is running\n", id)
	for j := 1; j <= 1000; j++ {
		fmt.Printf("Task %d: %d\n", id, j)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	// time.Sleep(time.Second * 2)
	wg.Wait()
}
