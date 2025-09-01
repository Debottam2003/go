package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"production/handlers"
	"production/db"
)

func main() {
	fmt.Println("This is the Production grade golang backend.")
	godotenv.Load(".env")
	var PORT string = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Can't load PORT from '.env'")
	} else {
		PORT = ":" + PORT
	}
	db.DB_connection()
	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.HomeHandler)
	router.HandleFunc("POST /login", handlers.LoginHandler)

	if router == nil {
		log.Fatal("Can not start the server...")
	} else {
		fmt.Println("The server is running and listening on PORt:", PORT)
	}
	http.ListenAndServe(PORT, router)
}
