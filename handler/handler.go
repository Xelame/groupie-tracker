package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var PATH = []string{}

type ArtistHandlerData struct {
	ListOfArtists []Artist
	PageNumber    []int
	SavedData     Cookies
}

type Cookies struct {
	Page      int
	SearchBar string
	Trie      string
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
		Trie      string
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

var Maintemp = OpenTemplate("index")
var ArtistTemp = OpenTemplate("artist")

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
	var isPaginated = false
	var pages []int
	var listmp []Artist
	var listOfArtist []Artist
	var artistName string
	var trieur string
	var page int

	SearchInApi("artists", &listOfArtist)

	if r.Method == "POST" {
		r.ParseForm()

		if r.FormValue("artists") == "" {
			artistName = r.FormValue("savedArtists")
		} else {
			artistName = r.FormValue("artists")
		}

		if r.FormValue("trie") == "" {
			trieur = r.FormValue("savedTrie")
		} else {
			trieur = r.FormValue("trie")
		}

		if r.FormValue("page") == "" {
			if r.FormValue("savedPage") == "" {
				page = 1
			} else {
				page, _ = strconv.Atoi(r.FormValue("savedPage"))
			}
		} else {
			page, _ = strconv.Atoi(r.FormValue("page"))
		}

		for i := 0; i < len(listOfArtist); i++ {
			if strings.Contains(strings.ToUpper(listOfArtist[i].Name), strings.ToUpper(artistName)) {
				listmp = append(listmp, listOfArtist[i])
			}
		}

		listOfArtist = listmp

		ArtistTrie(listOfArtist, trieur)
	}

	for i := 1; i <= len(listOfArtist)/12+1; i++ {
		pages = append(pages, i)
	}

	if r.Method == "POST" {
		isPaginated = true

		if page*12 > len(listOfArtist) {
			listmp = listOfArtist[12*(page-1):]
		} else {
			listmp = listOfArtist[12*(page-1) : 12*page]
		}

		fmt.Println("-----------------------------------------------")
		fmt.Println("\nLa recherche des artistes=", artistName, "!")
		fmt.Println("Avant c'était", r.FormValue("savedArtists"), "!")
		fmt.Printf("\nSur la page%d!\n", page)
		fmt.Println("Avant c'était", r.FormValue("savedPage"), "!")
		fmt.Println("\nTrie par ", r.FormValue("trie"), "!")
		fmt.Println("Avant c'était", r.FormValue("savedTrie"), "!")

	}

	if !isPaginated && len(listOfArtist) > 12 {
		listOfArtist = listOfArtist[:12]
	} else {
		listOfArtist = listmp
	}

	Maintemp.Execute(w, ArtistHandlerData{listOfArtist, pages, Cookies{page, artistName, trieur}})
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	SearchInApi(fmt.Sprintf("artists/%s", PATH[1]), &artist)
	GetWiki(&artist)
	ArtistTemp.Execute(w, artist) // En cours
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
