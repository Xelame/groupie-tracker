package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		for i := 0; i <= 52; i++ {
			fmt.Fprint(rw, searchInApi(fmt.Sprintf("artists/%d", i)))
		}
	})

	// Open the server (let's go)
	fmt.Println("Open server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchInApi(endOfUrl string) string {
	var url string
	if endOfUrl == "" {
		url = "https://groupietrackers.herokuapp.com/api"
	} else {
		url = fmt.Sprintf("https://groupietrackers.herokuapp.com/api/%s", endOfUrl)
	}

	json, err := http.Get(url)

	fmt.Print(json)
	if err != nil {
		return err.Error()
	}

	content, err := ioutil.ReadAll(json.Body)
	if err != nil {
		return err.Error()
	}

	return string(content)
}
