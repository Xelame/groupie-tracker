package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Locations struct {
	ID        int
	Locations []string
	Dates     string
}

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

type Dates struct {
	ID    int
	Dates []string
}

type Relations struct {
	DunedinNewZealand []string
	GeorgiaUsa        []string
	LosAngelesUsa     []string
	NagoyaJapan       []string
	NorthCarolinaUsa  []string
	OsakaJapan        []string
	PenroseNewZealand []string
	SaitamaJapan      []string
}

func main() {
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data := &Locations{}
		for i := 1; i <= 52; i++ {
			searchInApi(fmt.Sprintf("locations/%d", i), data)
			fmt.Fprint(rw, data)
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
