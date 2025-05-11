package main

import (
	"fmt"
	"io"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Method)
//     fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

// func main() {
//     http.HandleFunc("/", handler)
//     http.ListenAndServe(":8080", nil)
// }

func main() {
    // Fetching Data (Client Side)
	resp, err := http.Get("http://localhost:5000/jokes")
	if err != nil {
		panic(err)
	}
    fmt.Println(*resp)
    fmt.Println(resp.Status)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
    fmt.Println(body)   
    fmt.Println(string(body))
}
