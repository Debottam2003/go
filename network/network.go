package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message  string
	UserName string
}

type User struct {
	UserName string `json:"userName"`
	Age      int    `json:"age"`
}

// type Header map[string][]string

// type ResponseWriter interface {
//     Header() Header
//     Write([]byte) (n int, err error)
//     WriteHeader(statusCode int)
// }

// type response struct {
//     statusCode int                     // HTTP status code (e.g., 200, 404)
//     header     map[string][]string     // Map holding headers: map[string][]string
//     body       []byte                  // Body of the response
// }

// // Header returns the map of headers for the HTTP response.
// func (r *response) Header() map[string][]string {
//     if r.header == nil {
//         r.header = make(map[string][]string)
//     }
//     return r.header
// }

// func (h Header) Set(key, value string) {
//     h[key] = []string{value}
// }

// func (h Header) Add(key, value string) {
//     h[key] = append(h[key], value)
// }



var handlerGet = func (w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprintf(w, "Hello!, %s", r.URL)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
	// w.Write([]byte("this is a post method"))
	var m map[string][]string = w.Header()
	m["Content-Type"] = []string{"application/json"}
	res := Response{Message: "this is a post method", UserName: "Debottam Kar"}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
	// w.Write([]byte(`{"message": "this is a post method"}`))
}

// func main() {
//     http.HandleFunc("/", handler)
//     http.ListenAndServe(":8080", nil)
// }

func main() {
	// Fetching Data (Client Side)
	// resp, err := http.Get("http://localhost:5000/jokes")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(*resp)
	// fmt.Println(resp.Status)
	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(body)
	// fmt.Println(string(body))

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlerGet)
	router.HandleFunc("POST /", handlerPost)
	fmt.Println("Server is Running on Port: 5000")
	http.ListenAndServe(":5000", router)
}
