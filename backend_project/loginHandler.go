package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// password := "okudera"
	// fmt.Println(r)
	// fmt.Println("Body:", r.Body)
	// user_log := User_Log{}
	// mp := map[string]any{}
	// err := json.NewDecoder(r.Body).Decode(&user_log)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte("can not save data"))
	// 	return
	// }
	// if user_log.Email == "" || user_log.Password == "" {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte("Provide all credentials"))
	// 	return
	// }
	// if user_log.Password != password {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte("Wrong credentials"))
	// 	return
	// }
	// fmt.Println(user_log, user_log.Email, user_log.Password)

	// DataBase read
	rows, err := Db.Query("SELECT * FROM users")
	if err != nil {
		// http.Error(w, "Database error", http.StatusInternalServerError)
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Database error"))
		return
	}
	// fmt.Println(rows)
	var users []Users // slice of structures (will hold the rows of the users table)
	for rows.Next() {
		var user Users
		// scanning users data row by row one user at a time
		rows.Scan(&user.Id, &user.Email, &user.Password)
		// storing the users one by one in slice
		users = append(users, user)
	}
	rows.Close()
	// converting the data into json but in byte slice form
	// users_data, err := json.MarshalIndent(users, "", " ")
	// if err != nil {
	// 	// http.Error(w, "Database error", http.StatusInternalServerError)
	// 	fmt.Println(err)
	// 	w.WriteHeader(500)
	// 	w.Write([]byte("Database error"))
	// 	return
	// }
	// fmt.Println(string(users_data))

	// WriteHeader writes the status code to the connection as part of an HTTP reply.
	w.WriteHeader(http.StatusOK)
	// Write writes the data to the connection as part of an HTTP reply.
	// w.Write(users_data)

	// Directly converts and writes the json data (structures or slice of structures)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}
