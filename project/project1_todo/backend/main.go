package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the home handler."))
}

func main() {

	// anonymous function
	func() {
		fmt.Println("hello world!")
	}()

	fn := func() {
		fmt.Println("again hello world")
	}
	fn()

	fmt.Println("The server is running on http://localhost:3333")
	router := http.NewServeMux()
	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is the home handler."))
	})
	err := http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatal("Something went wrong")
	}
}
