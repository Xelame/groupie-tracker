package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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

func main() {
	maintemp := OpenTemplate("index")
	var url []string
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	data := &Artist{}
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

func GetUrl(r *http.Request) []string {
	path := strings.Split(r.URL.Path[1:], "/")
	return path
}

// GET DESCRIPTION PART _______________________________________________________________________________________________________________

func GetWiki(target Artist) {
	url := ""
	switch target.Name {
	case "Green Day":
		target.Name = "<span class=\"lang-en\" lang=\"en\">Green Day</span>"
		url = "https://fr.wikipedia.org/wiki/Green_Day"
	case "Alec Benjamin":
		target.Name = "Alec Shane Benjamin"
		url = "https://fr.wikipedia.org/wiki/Alec_Benjamin"
	case "Bee Gees":
		target.Name = "The Bee Gees"
		url = "https://fr.wikipedia.org/wiki/Bee_Gees"
	case "ACDC":
		target.Name = "AC/DC"
		url = "https://fr.wikipedia.org/wiki/AC/DC"
	case "SOJA":
		target.Name = "Soldiers of Jah Army"
		url = "https://fr.wikipedia.org/wiki/Soldiers_of_Jah_Army"
	case "Bobby McFerrins":
		target.Name = "Bobby McFerrin"
		url = "https://fr.wikipedia.org/wiki/Bobby_McFerrin"
	case "R3HAB":
		target.Name = "R3hab"
		url = "https://fr.wikipedia.org/wiki/R3hab"
	case "Genesis":
		url = "https://fr.wikipedia.org/wiki/Genesis_(groupe)"
	case "Muse":
		url = "https://fr.wikipedia.org/wiki/Muse_(groupe)"
	case "NWA":
		url = "https://fr.wikipedia.org/wiki/NWA_(groupe)"
	default:
		url = fmt.Sprintf("https://fr.wikipedia.org/wiki/%s", target.Name)
		url = strings.ReplaceAll(url, " ", "_")
	}

	fmt.Println(url)
	res, err := http.Get(url)

	if err != nil {
		res.StatusCode = 404
	}
	defer res.Body.Close()
	contentBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		res.StatusCode = 404
	}

	if res.StatusCode == 200 {
		pageContent := string(contentBytes)
		begin := strings.Index(pageContent, fmt.Sprintf("<b>%s</b>", target.Name))
		if begin == -1 {
			fmt.Println("Aie")
		}

		regex := regexp.MustCompile(`<div id="toc" class="toc" role="navigation" aria-labelledby="mw-toc-heading">|<h2>`)
		end := regex.FindStringIndex(string(pageContent[begin:]))
		end[0] += begin
		description := string([]byte(pageContent[begin:end[0]]))

		description = RegexTag(description)

		target.Description = description
	}
}

func GetTagName(tag string) string {
	end := 1
	for i := 0; i < len(tag) && end == 1; i++ {
		if tag[i] == ' ' || tag[i] == '>' {
			end = i
		}
	}
	return tag[1:end]
}

func RegexTag(content string) string {
	regex := regexp.MustCompile(`<(\"[^\"]*\"|'[^']*'|[^'\">])*>`)
	tags := regex.FindAllString(content, 1000)
	for len(tags) > 0 {
		tag := tags[0]
		if (GetTagName(tag) == "span" && !(tag == "<span class=\"lang-en\" lang=\"en\">" || tag == "<span class=\"nowrap\">")) || GetTagName(tag) == "style" || GetTagName(tag) == "sup" || GetTagName(tag) == "small" {
			for i := 1; i < len(tag) && "/"+GetTagName(tag) != GetTagName(tags[i]); i++ {
				if GetTagName(tag) == GetTagName(tags[i]) {
					tag = tags[i]
				}
			}
			begin := strings.Index(content, tag)
			end := strings.Index(content, fmt.Sprintf("</%s>", GetTagName(tag))) + len(fmt.Sprintf("</%s>", GetTagName(tag)))
			content = strings.Replace(content, content[begin:end], "", 1)
		} else {
			content = strings.Replace(content, tag, "", 1)
		}
		tags = regex.FindAllString(content, 1000)
	}
	return strings.ReplaceAll(content, "&#160;", "")
}

//______________________________________________________________________________________________________________________________
