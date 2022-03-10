package handler

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var PATH = []string{}

var Locationtemp = OpenTemplate("locations")
var Listtemp = OpenTemplate("index")
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
	var members []int
	var memberNumbers []int

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

	listmp = []Artist{}

	for _, artist := range listOfArtist {
		members = append(members, len(artist.Members))
	}

	members = RemoveDuplicateInt(members)
	sort.Ints(members)

	if r.Method == "POST" {
		r.ParseForm()
		for _, strNumber := range r.Form["members"] {
			intNumber, _ := strconv.Atoi(strNumber)
			memberNumbers = append(memberNumbers, intNumber)
		}

		if len(memberNumbers) > 0 {
			for _, number := range memberNumbers {
				for _, artist := range listOfArtist {
					if len(artist.Members) == number {
						listmp = append(listmp, artist)
					}
				}
			}

			listOfArtist = listmp
		}
	}

	for i := 1; i <= len(listOfArtist)/12+1; i++ {
		pages = append(pages, i)
	}

	if len(pages) < 2 {
		pages = []int{}
	}

	if r.Method == "POST" {
		isPaginated = true

		if page*12 > len(listOfArtist) {
			listmp = listOfArtist[12*(page-1):]
		} else {
			listmp = listOfArtist[12*(page-1) : 12*page]
		}
	}

	if !isPaginated && len(listOfArtist) > 12 {
		listOfArtist = listOfArtist[:12]
	} else {
		listOfArtist = listmp
	}
	fmt.Println(memberNumbers)
	fmt.Println(r.FormValue("savedMembers"))

	Listtemp.Execute(w, ArtistHandlerData{listOfArtist, pages, members, Cookies{page, artistName, trieur, memberNumbers}})
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	SearchInApi(fmt.Sprintf("artists/%s", PATH[1]), &artist)
	GetWiki(&artist)
	ArtistTemp.Execute(w, artist) // En cours
}

func LocationsHandler(rw http.ResponseWriter, r *http.Request) {
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
	Locationtemp.Execute(rw, start)
	fmt.Println(ArtistsinArea)
	fmt.Println(listOfLocations)

}

//func HomeHandler(w http.ResponseWriter, r *http.Request) {
