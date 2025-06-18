package main

import (
	"fmt"
	"net/http"
)

func HandleWrite(w http.ResponseWriter, r *http.Request) {
	// r.FormValue("email")
	// r.FormValue("password")
	_, err := Db.Exec("INSERT INTO users(email, password) VALUES($1, $2)", []any{"debottam@gmail.com", "ichinose2003"}...)
	if err != nil {
		// http.Error(w, "Database error", http.StatusInternalServerError)
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Database error"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}