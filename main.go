package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

type Dates struct {
	Id    int
	Dates []string
}

type Locations struct {
	Id        int
	Locations interface{}
	Dates     string
}

type Relations struct {
	ID             int
	DatesLocations interface{}
}

func main() {
	maintemp := OpenTemplate("index")
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data := &Artist{}
		data1 := &Locations{}
		data2 := &Relations{}
		data3 := &Dates{}
		listOfDates := []Dates{}
		listOfRelations := []Relations{}
		listOfLocations := []Locations{}
		listOfArtist := []Artist{}
		for i := 1; i <= 52; i++ {
			searchInApi(fmt.Sprintf("artists/%d", i), data)
			searchInApi(fmt.Sprintf("locations/%d", i), data1)
			searchInApi(fmt.Sprintf("relation/%d", i), data2)
			searchInApi(fmt.Sprintf("dates/%d", i), data3)
			listOfDates = append(listOfDates, *data3)
			listOfLocations = append(listOfLocations, *data1)
			listOfRelations = append(listOfRelations, *data2)
			listOfArtist = append(listOfArtist, *data)
		}
		fmt.Println(listOfDates)
		maintemp.Execute(rw, listOfLocations)
	})

	fmt.Println("Server Open In http://localhost:8080")
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

	return json.NewDecoder(res.Body).Decode(target)
}

func OpenTemplate(fileName string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("./templates/%s.html", fileName))
	if err != nil {
		fmt.Println(err.Error())
	}
	return tmpl
}
