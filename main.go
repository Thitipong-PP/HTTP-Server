package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler
func testHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Test" {
		http.Error(
			w, 
			"404 not found.", 
			http.StatusNotFound,
		)
		return
	}

	if r.Method != "GET" {
		http.Error(
			w,
			"Method is not supperted.",
			http.StatusNotFound,
		)
		return
	}

	fmt.Fprintln(w, "Testing Response")
	fmt.Println("Testing Terminal")
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL : ", r.URL.Path)
	fmt.Fprintf(w, "Hello this is home handler using net/http")
}

func statusHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([] byte("System Normal: All systems operational"))
}

func main () {
	// Handle
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/Test", testHandler)
	
	// Set Server Port
	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}