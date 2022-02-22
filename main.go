package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
	Locations []string
	Dates     string
}

type Relations struct {
	ID             int
	DatesLocations interface{}
}

func main() {
	maintemp := OpenTemplate("index")
	var url []string
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		url = GetUrl(r)
		if len(url) > 1 {
			data := &Artist{}
			Artists := []Artist{}
			intUrl, _ := strconv.Atoi(url[1])
			searchInApi(fmt.Sprintf("artists/%d", intUrl), data)
			Artists = append(Artists, *data)
			maintemp.Execute(rw, Artists)
		} else {
			listOfArtist := []Artist{}
			data := &Artist{}
			for i := 1; i <= 52; i++ {
				searchInApi(fmt.Sprintf("artists/%d", i), data)
				listOfArtist = append(listOfArtist, *data)
			}
			maintemp.Execute(rw, listOfArtist)
		}
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

func GetUrl(r *http.Request) []string {
	path := strings.Split(r.URL.Path[1:], "/")
	return path
}

func OpenTemplate(fileName string) *template.Template {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("./templates/%s.html", fileName), "./templates/components/card.html"))
	return tmpl
}
