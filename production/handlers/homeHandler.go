package handlers

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"production/db"
	"production/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Request came")
	var authors []utils.Authors = []utils.Authors{}
	rows, err := db.Db.Query("select * from authors")
	if err != nil {
		w.Header().Set("Content-Type", "text/html")   //  sets the header
		w.WriteHeader(http.StatusInternalServerError) // writes the header with a status
		w.Write([]byte("Internal server error"))      // writes the buffer in the tcp connection
	} else {
		for rows.Next() {
			var author utils.Authors
			rows.Scan(&author.ID, &author.Email, &author.Password, &author.FirstName, &author.LastName)
			authors = append(authors, author)
		}
		rows.Close()
		w.Header().Set("Content-Type", "application/json") //  sets the header
		w.WriteHeader(http.StatusOK)                       // writes the header with a status
		json.NewEncoder(w).Encode(authors)                 // writes the buffer in the tcp connection
	}
}
