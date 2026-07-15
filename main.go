package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/intro", DisplayIntro)
	http.HandleFunc("/home", DisplayHome)
	fmt.Println("server running currently on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
