package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func Fetch(wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get("http://localhost:3333")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	fmt.Println(stringData)
	defer response.Body.Close()
}