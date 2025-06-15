package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

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

func main() {

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not found")
	} else {
		port = ":" + port
	}
	pg_url := os.Getenv("DATABASE_URL")
	if pg_url == "" {
		log.Fatal("port not found")
	}

	Db, err = sql.Open("postgres", pg_url)
	Error_Reaction(err)

	err = Db.Ping()
	Error_Reaction(err)
	fmt.Println("✅ Connected to PostgreSQL")

	rows, err := Db.Query("SELECT * FROM users")
	Error_Reaction(err)
	fmt.Println(rows)
	var users []Users // slice of structures (will hold the columns of the users table)
	for rows.Next() {
		var user Users
		// scanning users data row by row one user at a time
		rows.Scan(&user.Id, &user.Email, &user.Password)
		// storing the users one by one in slice
		users = append(users, user)
	}
	rows.Close()
	// converting the data into json but in byte slice form
	users_data, err := json.MarshalIndent(users, "", " ")
	Error_Reaction(err)
	fmt.Println(string(users_data))
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /home", HomeHandler)
	router.HandleFunc("POST /login", LoginHandler)

	fmt.Println("Server started running on port:", port)
	err = http.ListenAndServe(port, router)
	Error_Reaction(err)
}
