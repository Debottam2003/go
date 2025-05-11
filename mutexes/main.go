package main

import (
	"fmt"
	"sync"
)

type user struct {
	views int
	mu sync.Mutex
}

// var mu sync.Mutex

func (u *user) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	u.mu.Lock()
	defer u.mu.Unlock()
	for i := 0; i < 5; i++ {
		u.views++
	}
}

func main() {
	var wg sync.WaitGroup
	u := user{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go u.inc(&wg)
	}
	wg.Wait()
	fmt.Println(u.views)
}
