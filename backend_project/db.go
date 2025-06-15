package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Db *sql.DB

func DB_connect() {
	godotenv.Load(".env")

	pg_url := os.Getenv("DATABASE_URL")
	if pg_url == "" {
		log.Fatal("DATABASE_URL not found")
	}

	Db, err = sql.Open("postgres", pg_url)
	Error_Reaction(err)

	err = Db.Ping()
	Error_Reaction(err)
	fmt.Println("✅ Connected to PostgreSQL")

	rows, err := Db.Query("SELECT * FROM users")
	Error_Reaction(err)
	fmt.Println(rows)// reference of a Rows strcuture instance
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
}
