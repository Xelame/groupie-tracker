package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

func main() {
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data := &Artist{}
		for i := 1; i <= 52; i++ {
			searchInApi(fmt.Sprintf("artists/%d", i), data)
			fmt.Fprint(rw, data.Id)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func searchInApi(endOfUrl string, target interface{}) error {
	var url string
	if endOfUrl == "" {
		url = "https://groupietrackers.herokuapp.com/api"
	} else {
		url = fmt.Sprintf("https://groupietrackers.herokuapp.com/api/%s", endOfUrl)
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
