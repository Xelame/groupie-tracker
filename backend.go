package main

import (
	"fmt"
	"net/http"

	// Import Backend functions
	handler "handler/handler"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	// Apply a function in this page (don't worry i display every time a html template ^^)
	http.HandleFunc("/", handler.RoutingHandler)
	http.HandleFunc("/locations", handler.LocationsHandler)
	fmt.Println("Server Open In http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
