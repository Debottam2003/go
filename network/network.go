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

func handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
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
	w.Header().Set("Content-Type", "application/json")
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
