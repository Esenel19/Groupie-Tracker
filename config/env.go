package config

// APIRURL
const APIURL = "https://groupietrackers.herokuapp.com/api"

// Initiliazing the port
const LocalhostPort = ":8080"

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstalbum"`
	Locations    Locations
	ConcertDates Dates
	Relations    Relation
}

type Locations struct {
	Id           int      `json:id`
	Locations    []string `json:locations`
	Dates        string   `json:dates`
	AllLocations []string
}

type Locations_Index struct {
	Index []Locations `json:index`
}

type Dates_Index struct {
	Index []Dates `json:index`
}

type Dates struct {
	Id    int      `json:id`
	Dates []string `json:dates`
}

type Relation_Index struct {
	Index []Relation `json:index`
}

type Relation struct {
	Id             int                 `json:id`
	DatesLocations map[string][]string `json:datesLocations`
	Countrystr     map[string][]string
	AllLocations   []string
}
