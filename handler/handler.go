package handler

import (
	"fmt"
	"net/http"
	"strings"
)

var PATH = []string{}

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
	Description  string
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

type Data struct {
	ListOfArtists []Artist
	PageNumber    []int
}

var Maintemp = OpenTemplate("index")
var ArtistTemp = OpenTemplate("artist")
var FormRoute = []string{"pages"}
var ListOfArtist []Artist

func RoutingHandler(rw http.ResponseWriter, r *http.Request) {
	PATH = GetUrl(r)
	if PATH[0] == "artists" {
		if len(PATH) > 1 {
			ArtistHandler(rw, r)
		} else {
			AllArtistsHandler(rw, r)
		}
	}
}

func AllArtistsHandler(w http.ResponseWriter, r *http.Request) {
	isPaginated := false
	page := []int{}
	listmp := []Artist{}

	if len(r.URL.Query()["page"]) == 0 {
		SearchInApi("artists", &ListOfArtist)
	}

	if r.Method == "POST" {
		for i := 0; i < len(ListOfArtist); i++ {
			if strings.Contains(strings.ToUpper(ListOfArtist[i].Name), strings.ToUpper(r.FormValue("artists"))) {
				listmp = append(listmp, ListOfArtist[i])
			}
		}
		ListOfArtist = listmp
	}

	for i := 1; i <= len(ListOfArtist)/12+1; i++ {
		if len(r.URL.Query()["page"]) > 0 && fmt.Sprintf("%d", i) == r.URL.Query()["page"][0] {
			isPaginated = true
			if i*12 > len(ListOfArtist) {
				listmp = ListOfArtist[12*(i-1):]
			} else {
				listmp = ListOfArtist[12*(i-1) : 12*i]
			}
		}
		page = append(page, i)
	}

	if !isPaginated && len(ListOfArtist) > 12 {
		ListOfArtist = ListOfArtist[:12]
	} else {
		ListOfArtist = listmp
	}

	Maintemp.Execute(w, Data{ListOfArtist, page})
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	SearchInApi(fmt.Sprintf("artists/%s", PATH[1]), &artist)
	GetWiki(&artist)
	ArtistTemp.Execute(w, artist)
}
