package main

import (
	"fmt"
	"net/http"

	// Import Backend functions
	handler "handler/handler"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// Apply a function in this page (don't worry i display every time a html template ^^)
	http.HandleFunc("/", handler.RoutingHandler)
	fmt.Println("Server Open In http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
