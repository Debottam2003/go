package main

import (
	"fmt"
	"log"
)

// Define a type for cleaner code
type Middleware map[string]func()

func main() {
	fmt.Println("This is the middleware function.")

	middleware := Middleware{
		"/users": func() {
			fmt.Println("This is the /users api endpoint handler.")
		},
		"/login": func() {
			fmt.Println("This is the /login api endpoint handler.")
		},
		"/admin": func() {
			fmt.Println("This is the /admin api endpoint handler.")
		},
	}

	// Execute handlers
	middleware["/users"]()
	middleware["/login"]()

	// Demonstrate with error handling
	executeHandler(middleware, "/users")
	executeHandler(middleware, "/login")
	executeHandler(middleware, "/nonexistent") // This will show error handling
}

// Helper function with error handling
func executeHandler(mw Middleware, route string) {
	if handler, exists := mw[route]; exists {
		handler()
	} else {
		log.Printf("Warning: No handler found for route '%s'", route)
	}
}
