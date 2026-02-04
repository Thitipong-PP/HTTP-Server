package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	// Our server api and handler for CRUD
	"github.com/Thitipong-PP/HTTP-Server/server"
)

// Basic Respons Json data
type ResponseData struct {
	Message string `json:"message"`
	Success bool `json:"success"`
	Data []string `json:"data_list"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Method Checker
	if r.Method != http.MethodGet {
		http.Error(w, "Don't have this method", http.StatusMethodNotAllowed)
		return
	}

	// MIME Type
	w.Header().Set("Content-type", "application/json")

	res := ResponseData{
		Message: "Say Hi Json",
		Success: true,
		Data: []string{"Server", "Database"},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Basic Handler Request
func homeHandler(w http.ResponseWriter, r *http.Request) {
	
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

func main() {
	// Handler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/api", server.ServerHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}