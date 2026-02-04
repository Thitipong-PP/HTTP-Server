package main

import (
	"fmt"
	"net/http"
)

func homeHandler (w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Header.Get("User-Agent"))

	// agent := r.Header.Get("User-Agent")

	// if strings.Contains(agent, "Chrome") {
	// 	fmt.Fprintln(w, "This is run in Chrome")
	// }else if strings.Contains(agent, "curl") {
	// 	fmt.Fprintln(w, "You are Hacker?")
	// }else {
	// 	fmt.Fprintln(w, "Hello guys!")
	// }

	switch r.Method {
		case http.MethodGet : 
			fmt.Fprintln(w, "We sent it")
		case http.MethodPost :
			fmt.Fprintln(w, "We get it")
		default :
			fmt.Fprintln(w, "Don't have this method")
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