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
	Index []struct {
		ID        int
		Locations []string
		Dates     string
	}
}

type Loc struct {
	Artists  []string
	Location string
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
var HomeTemp = OpenTemplate("home")
var FormRoute = []string{"pages"}
var ListOfArtist []Artist

func RoutingHandler(rw http.ResponseWriter, r *http.Request) {
	PATH = GetUrl(r)
	fmt.Println(PATH[0])
	if PATH[0] == "locations" {
		LocationsHandler(rw, r)
	} else if PATH[0] == "artists" {
		if len(PATH) > 1 {
			ArtistHandler(rw, r)
		} else {
			AllArtistsHandler(rw, r)
		}
	} else if PATH[0] == "home" {
		HomeHandler(rw, r)
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

func LocationsHandler(rw http.ResponseWriter, r *http.Request) {
	Maintemp = OpenTemplate("locations")
	var locations Locations
	var listOfLocations string
	var indexes []int
	var ArtistsinArea []string

	SearchInApi("locations", &locations)
	listOfArtist := &[]Artist{}
	SearchInApi("artists", listOfArtist)

	if r.Method == "POST" {
		for i := 0; i <= 51; i++ {
			for j := 0; j < len(locations.Index[i].Locations); j++ {
				if strings.Contains(strings.ToUpper(locations.Index[i].Locations[j]), strings.ToUpper(strings.ReplaceAll(r.FormValue("locations"), " ", "_"))) {
					listOfLocations = "https:www.google.com/maps/embed/v1/place?key=AIzaSyAXXPpGp3CYZDcUSiE2YRlNID4ybzoZa7o&q=" + locations.Index[i].Locations[j]
					indexes = append(indexes, i)
					ArtistsinArea = append(ArtistsinArea, (*listOfArtist)[i].Name)
				}
			}
		}
	}
	start := Loc{ArtistsinArea, listOfLocations}
	fmt.Println(start)
	Maintemp.Execute(rw, start)
	fmt.Println(ArtistsinArea)
	fmt.Println(listOfLocations)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	HomeTemp.Execute(w, nil)
}
