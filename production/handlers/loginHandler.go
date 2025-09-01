package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"production/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user utils.Users
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
}
