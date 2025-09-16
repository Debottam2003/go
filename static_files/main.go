package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Serve ./public at root
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/", fs)

	// Strip the prefix if serving under a subpath (optional)
	// router.Handle("/static/", http.StripPrefix("/static/", fs))

	// about route → about.html
	router.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/about.html")
	})

	log.Println("🚀 Server running at http://localhost:3333")
	http.ListenAndServe(":3333", router)
}
