package main

import (
	"fmt"
	"net/http"
	"strings"
)

func homeHandler (w http.ResponseWriter, r *http.Request) {
	
	fmt.Println(r.Header.Get("User-Agent"))
	agent := r.Header.Get("User-Agent")
	
	// Anti-Bot - Basic
	if strings.Contains(agent, "curl") {
		http.Error(w, "You are Hacker?", http.StatusForbidden)
		return
	}

	// Method checker
	switch r.Method {
		case http.MethodGet : 
			fmt.Fprintln(w, "You're in GET Method")
		case http.MethodPost :
			fmt.Fprintln(w, "Have you POST something?")
		default :
			http.Error(w, "Don't have this method", http.StatusMethodNotAllowed)
	}
}

func main () {
	// Handler
	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}