package handler

type ArtistHandlerData struct {
	ListOfArtists []Artist
	PageNumber    []int
	MembersFilter []int
	SavedData     Cookies
}

type Cookies struct {
	Page       int
	SearchBar  string
	Trie       string
	PinChecked []int
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

type Relations struct {
	Index []struct {
		ID             int
		DatesLocations map[string][]string
	}
}

type Loc struct {
	Artists         []string
	Location        string
	ListOfLocations []string
	//Dates    []string
}
