package handler

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var PATH = []string{}

//var Locationtemp = OpenTemplate("locations")
//var Listtemp = OpenTemplate("index")
//var ArtistTemp = OpenTemplate("artist")
//var HomeTemp = OpenTemplate("home")
var FormRoute = []string{"pages"}
var ListOfArtist []Artist
var forbiddenInput = []string{"{", "=>", "}", ";", ">", "<", "&", "+", "-", "%", "%q", "\n", "#", "~", "`", "^", "(", ")", "[", "]", "|"}

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
	} else {
		Error404Handler(rw, r)
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
	var err error

	SearchInApi("artists", &listOfArtist)

	if r.Method == "POST" {
		r.ParseForm()

		if r.FormValue("artists") == "" {
			artistName = r.FormValue("savedArtists")
		} else if CheckForbiddenInput(r.FormValue("artists")) {
			errortmpl, erR := OpenTemplate("err400")
			if erR != nil {
				fmt.Fprint(w, "Not working")
				return
			}
			errortmpl.Execute(w, nil)
			return
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
				page, err = strconv.Atoi(r.FormValue("savedPage"))
				if err != nil {
					errortmpl, erR := OpenTemplate("err500")
					if erR != nil {
						fmt.Fprint(w, "Not working")
						return
					}
					errortmpl.Execute(w, nil)
					return
				}
			}
		} else {
			page, err = strconv.Atoi(r.FormValue("page"))
			if err != nil {
				errortmpl, erR := OpenTemplate("err500")
				if erR != nil {
					fmt.Fprint(w, "Not working")
					return
				}
				errortmpl.Execute(w, nil)
				return
			}
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
			intNumber, err := strconv.Atoi(strNumber)
			print(strNumber)
			memberNumbers = append(memberNumbers, intNumber)
			if err != nil {
				errortmpl, erR := OpenTemplate("err500")
				if erR != nil {
					fmt.Fprint(w, "Not working")
					return
				}
				errortmpl.Execute(w, nil)
				return
			}
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
	Listtemp, erR := OpenTemplate("index")
	if erR != nil {
		fmt.Fprint(w, "Not working")
		return
	}
	Listtemp.Execute(w, ArtistHandlerData{listOfArtist, pages, members, Cookies{page, artistName, trieur, memberNumbers}})
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	err := SearchInApi(fmt.Sprintf("artists/%s", PATH[1]), &artist)
	ArtistTemp, erR := OpenTemplate("artist")
	if erR != nil {
		fmt.Fprint(w, "Not working")
		return
	}
	if err != nil {
		errortmpl, erR := OpenTemplate("err500")
		if erR != nil {
			fmt.Fprint(w, "Not working")
			return
		}
		errortmpl.Execute(w, nil)
		return
	}
	GetWiki(&artist)
	ArtistTemp.Execute(w, artist) // En cours
}

func LocationsHandler(rw http.ResponseWriter, r *http.Request) { //used to get all concert locations and display artists in area and their date(s) of concert
	var locations Locations
	var location string
	var indexes []int
	var ArtistsinArea []string
	var listOfListsOfLocations [][]string
	var listOfLocations []string

	err1 := SearchInApi("locations", &locations)
	listOfArtist := &[]Artist{}
	if err1 != nil {
		errortmpl, erR := OpenTemplate("err500")
		if erR != nil {
			fmt.Fprint(rw, "Not working")
			return
		}
		errortmpl.Execute(rw, nil)
		return
	}
	err2 := SearchInApi("artists", listOfArtist)
	var listOfDates Dates
	if err2 != nil {
		errortmpl, erR := OpenTemplate("err500")
		if erR != nil {
			fmt.Fprint(rw, "Not working")
			return
		}
		errortmpl.Execute(rw, nil)
		return
	}
	err3 := SearchInApi("dates", &listOfDates)
	if err3 != nil {
		errortmpl, erR := OpenTemplate("err500")
		if erR != nil {
			fmt.Fprint(rw, "Not working")
			return
		}
		errortmpl.Execute(rw, nil)
		return
	}
	var listOfRelations Relations
	err4 := SearchInApi("relation", &listOfRelations)
	if err4 != nil {
		errortmpl, erR := OpenTemplate("err500")
		if erR != nil {
			fmt.Fprint(rw, "Not working")
			return
		}
		errortmpl.Execute(rw, nil)
		return
	}
	for i := 0; i < len(locations.Index); i++ {
		listOfListsOfLocations = append(listOfListsOfLocations, locations.Index[i].Locations) //in the locations struct, the locations is a list of list of strings, i.e each artist has a list of locations
	}
	for j := 0; j < 51; j++ {
		for s := 0; s < len(listOfListsOfLocations[j]); s++ {
			listOfLocations = append(listOfLocations, listOfListsOfLocations[j][s]) //put all the concert locations in this list, later we will remove the duplicates
		}
	}
	sort.Sort(sort.StringSlice(listOfLocations)) //Sort List of Locations alphabetically
	// fmt.Println("1:", len(listOfLocations))
	// fmt.Println("2:", len(RemoveDuplicateStr(listOfLocations))) --> some useful prints to check if we really remove the duplicate strings
	if r.Method == "POST" {
		fmt.Println(r.FormValue("locations"))
		for i := 0; i <= 51; i++ {
			for j := 0; j < len(locations.Index[i].Locations); j++ {
				if strings.Contains(strings.ToUpper(locations.Index[i].Locations[j]), strings.ToUpper(strings.ReplaceAll(r.FormValue("locations"), " ", "_"))) {
					location = "https:www.google.com/maps/embed/v1/place?key=AIzaSyAXXPpGp3CYZDcUSiE2YRlNID4ybzoZa7o&q=" + locations.Index[i].Locations[j] //to be able to display the google map
					//date = append(date, listOfDates.Index[i].Dates[j])
					indexes = append(indexes, i)
					for s := 0; s < len(listOfRelations.Index[i].DatesLocations[locations.Index[i].Locations[j]]); s++ {
						ArtistsinArea = append(ArtistsinArea, (*listOfArtist)[i].Name+" in "+listOfRelations.Index[i].DatesLocations[locations.Index[i].Locations[j]][s]) //we get the concert dates from relations because some artist have multiple dates on the same area (example : SOJA in playa del carmen)
					}
				}
			}
		}
	}
	start := Loc{ArtistsinArea, location, RemoveDuplicateStr(listOfLocations)}
	Locationtemp, erR := OpenTemplate("locations")
	if erR != nil {
		fmt.Fprint(rw, "Not working")
		return
	}
	Locationtemp.Execute(rw, start)
}

func RemoveDuplicateStr(strSlice []string) []string { //We use this function to remove duplicate strings in an array of strings
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Error404Handler(rw http.ResponseWriter, r *http.Request) {
	errortmpl, erR := OpenTemplate("err404")
	if erR != nil {
		fmt.Fprint(rw, "Not working")
		return
	}
	errortmpl.Execute(rw, nil)
}

func CheckForbiddenInput(str string) bool { //range through the form value to check if it contains a dangerous character
	result := false
	for i := 0; i < len(forbiddenInput); i++ {
		if strings.Contains(str, forbiddenInput[i]) {
			result = true
			break
		}
	}
	return result
}

//func HomeHandler(w http.ResponseWriter, r *http.Request) {
