package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
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

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	data := &Artist{}
	for i := 1; i <= 52; i++ {
		searchInApi(fmt.Sprintf("artists/%d", i), data)
		GetWiki(*data)
	}
	// Apply a function in this page (don't worry i diplay every time a html template ^^)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		PATH = GetUrl(r)
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
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("./templates/%s.html", fileName), "./templates/components/card.html"))
	return tmpl
}

func GetUrl(r *http.Request) []string {
	path := strings.Split(r.URL.Path[1:], "/")
	return path
}

func GetWiki(target Artist) {
	switch target.Name {
	case "SOJA":
		target.Name = "Soldiers of Jah Army"
	case "Bobby McFerrins":
		target.Name = "Bobby McFerrin"
	case "R3HAB":
		target.Name = "R3hab"
	case "Genesis":
		target.Name = "Genesis (groupe)"
	}

	url := fmt.Sprintf("https://fr.wikipedia.org/wiki/%s", target.Name)
	url = strings.ReplaceAll(url, " ", "_")
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
		begin := strings.Index(pageContent, "<b>")
		if begin == -1 {
			fmt.Println("Aie")
		}
		begin += len("<b>")

		end := strings.Index(string(pageContent[begin:]), "</p>")
		if end == -1 {
			fmt.Println("Aie")
		}
		end += begin

		description := string([]byte(pageContent[begin:end]))
		regex1 := regexp.MustCompile(`<(\"[^\"]*\"|'[^']*'|[^'\">])*>`)
		test := regex1.FindAllString(description, 1000)
		for _, v := range test {
			if len(v) >= 6 {
				if v[:6] == "<style" {
					regex2 := regexp.MustCompile(fmt.Sprintf(`%s(.*?)%s`, v, fmt.Sprintf("</%s>", v[1:6])))
					test2 := regex2.FindAllString(description, 1000)
					for _, w := range test2 {
						description = strings.Replace(description, w, "", 1)
					}
				}
			}
			if len(v) >= 5 {
				if v[:5] == "<span" {
					regex2 := regexp.MustCompile(`<span .*?(<span .*?</span>.*?</span>|.*?</span>)`)
					test2 := regex2.FindAllString(description, 1000)
					for _, w := range test2 {
						description = strings.Replace(description, w, "", 1)
					}
				}
			}
			if len(v) > 3 {
				if v[:4] == "<sup" {
					regex2 := regexp.MustCompile(fmt.Sprintf(`%s(.*?)%s`, v, fmt.Sprintf("</%s>", v[1:4])))
					test2 := regex2.FindAllString(description, 1000)
					for _, w := range test2 {
						description = strings.Replace(description, w, "", 1)
					}
				}
			}
			description = strings.Replace(description, v, "", 1)
		}
		// fmt.Println(description)
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

func DetectBadTag(tag string) {
	
}