package main

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// Setting "Content-Type" as "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Setting Status code 200
	w.WriteHeader(http.StatusOK)

	jr := Json_Response{Name: "debottam kar", Age: 22}
	// Convert the Json_Response structure instance into json and returns a byte slice form []byte()
	json_data, err := json.MarshalIndent(jr, "", "")
	Error_Reaction(err)
	// fmt.Println(string(json_data))

	w.Write(json_data)
}
