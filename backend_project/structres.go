package main

// For json response
type Json_Response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// For users table
type Users struct {
	Id       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// For users login
type User_Log struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}