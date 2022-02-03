package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Locations struct {
	ID        int
	Locations []string
	Dates     string
}

func main() {
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, searchInApi(""))
	})

	// Open the server (let's go)
	fmt.Println("Open server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchInApi(endOfUrl string) string {
	json, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api%s", endOfUrl))
	if err != nil {
		return err.Error()
	}

	content, err := ioutil.ReadAll(json.Body)
	if err != nil {
		return err.Error()
	}

	return string(content)
}
