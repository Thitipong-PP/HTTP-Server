package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
	// Handle
	http.HandleFunc("/Test", testHandler)
	
	// Set Server Port
	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func testHandler (res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/Test" {
		http.Error(
			res, 
			"404 not found.", 
			http.StatusNotFound,
		)
		return
	}

	if req.Method != "GET" {
		http.Error(
			res,
			"Method is not supperted.",
			http.StatusNotFound,
		)
		return
	}

	fmt.Fprintln(res, "Testing Response")
	fmt.Println("Testing Terminal")
}