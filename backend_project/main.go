package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var err error

func main() {

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not found")
	} else {
		port = ":" + port
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /home", HomeHandler)
	router.HandleFunc("GET /login", LoginHandler)
	router.HandleFunc("POST /write", HandleWrite);

	fmt.Println("Server started running on port:", port)
	DB_connect()
	err = http.ListenAndServe(port, router)
	Error_Reaction(err)
}
