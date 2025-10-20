package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(http.StatusOK)                  // 200
	fmt.Println(http.StatusInternalServerError) // 500
	fmt.Println(http.StatusBadRequest)          // 400
	fmt.Println(http.StatusNotFound)            // 404
	fmt.Println(http.StatusUnauthorized)        // 401
	fmt.Println(http.StatusCreated)             // 201
	fmt.Println(http.StatusForbidden)           // 403
	fmt.Println(http.StatusMovedPermanently)    // 301
}
