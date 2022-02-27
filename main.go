package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	config "./config"
)

var artists_information, Data []config.Artist
var artists_locations []config.Locations
var artists_relation config.Relation_Index
var locations_Index config.Locations_Index
var date_Index config.Dates_Index
var AllLocations, AllLoc []string
var state, input1, input2, input3 string
var CountryStrData = map[string][]string{}
var input_backup = "Queen"

//Main code used to run the server setup
func main() {
	Artists_info()
	Artists_loc()
	Artists_Date()
	Artists_rela()

	fmt.Println("Please connect to\u001b[31m localhost", config.LocalhostPort, "\u001b[0m")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) // Join Assets Directory to the server
	http.HandleFunc("/", index)
	http.HandleFunc("/locations", locations)
	http.HandleFunc("/dates", dates)
	http.HandleFunc("/bestdate", BestDate)
	http.HandleFunc("/relation", relation)
	http.HandleFunc("/artists", artists)
	http.HandleFunc("/singleArtist", singleArtist)
	http.HandleFunc("/countryConcert", concertLocation)
	http.HandleFunc("/Loc&date", locAndDate)
	err := http.ListenAndServe(config.LocalhostPort, nil) // Set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//generate the main page when first loading the site
func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index-template")
	t = template.Must(t.ParseFiles("index.html", "./tmpl/header&footer.html"))
	fmt.Println(artists_information[0])
	t.ExecuteTemplate(w, "index", artists_information)
}

//	generate the page about the Locations
func locations(w http.ResponseWriter, r *http.Request) {
	Data = nil

	for index := range artists_information {
		artists_information[index].Locations.AllLocations = AllLocations
		Data = append(Data, artists_information[index])
	}

	t := template.New("locations-template")
	t = template.Must(t.ParseFiles("./tmpl/locations.html", "./tmpl/header&footer.html"))
	t.ExecuteTemplate(w, "locations", Data)
}

//	generate the page about the relation
func relation(w http.ResponseWriter, r *http.Request) {
	Data = nil
	for index := range artists_information {
		for keys := range artists_information[index].Relations.DatesLocations {
			AllLoc = append(AllLoc, keys)
		}
		artists_information[index].Locations.AllLocations = AllLocations
		Data = append(Data, artists_information[index])
	}
	AllLoc = removeDuplicateValues(AllLoc)

	AllLoc = sorting(AllLoc)
	sort.Strings(AllLoc)
	AllLoc = sorting(AllLoc)
	countrystr(AllLoc)
	Data[51].Relations.Countrystr = CountryStrData

	t, _ := template.New("relation-template").ParseFiles("./tmpl/relation.html", "./tmpl/header&footer.html")
	t.ExecuteTemplate(w, "relation", Data)
}

//range over all the location then remove unwanted string
func countrystr(strarr []string) {
	CountryStrData = map[string][]string{}
	var splited []string
	var splited_city string
	for _, v := range strarr {
		splited = strings.Split(v, "-")
		splited_city = splited[0]
		splited = splited[1:]
		CountryStrData[splited[0]] = append(CountryStrData[splited[0]], splited_city)
	}
}

//invert value between a "-"
//
// example:
//
//	"abc-def" ==> "def-abc"
func sorting(strarr []string) []string {
	var splited []string
	for index, v := range strarr {
		splited = strings.Split(v, "-")
		splited[0], splited[1] = splited[1], splited[0]
		strarr[index] = strings.Join(splited, "-")
	}
	return strarr
}

//generate the page about the dates
func dates(w http.ResponseWriter, r *http.Request) {
	t := template.New("dates-template")
	t = template.Must(t.ParseFiles("./tmpl/dates.html", "./tmpl/header&footer.html"))
	t.ExecuteTemplate(w, "dates", artists_information)
}

//generate the page about all the artists
func artists(w http.ResponseWriter, r *http.Request) {
	t := template.New("artist-template")
	t = template.Must(t.ParseFiles("./tmpl/artists.html", "./tmpl/header&footer.html", "./tmpl/content.html"))
	t.ExecuteTemplate(w, "artists", artists_information)
}

//generate the page about a single artist
func singleArtist(w http.ResponseWriter, r *http.Request) {
	Artists_info()
	Artists_rela()

	state = "singleArtist"
	Data = nil
	input1 = r.FormValue("searchArtist") // valeur de l'input text
	input2 = r.FormValue("idArtist")     // valeur de l'input pour "see more"
	InsertData(state)
	// cherche si ce qui est dans un des input correspond à un des artistes

	t := template.New("singleArtist-template")
	t = template.Must(t.ParseFiles("./tmpl/artists.html", "./tmpl/header&footer.html", "./tmpl/content.html"))
	t.ExecuteTemplate(w, "singleArtist", Data)
}

//generate the page about the concert location
func concertLocation(w http.ResponseWriter, r *http.Request) {
	input1 = r.FormValue("country")
	input2 = r.FormValue("artist")
	state = "concertLocation"
	Data = nil

	fmt.Println("Data (empty): ", Data)
	InsertData(state)
	fmt.Println("Data (filled): ", Data)

	t := template.New("singleArtist-template")
	t = template.Must(t.ParseFiles("./tmpl/countryConcert.html", "./tmpl/header&footer.html", "./tmpl/content.html"))
	t.ExecuteTemplate(w, "countryConcert", Data)
}

//insert an artist to Data then update information about
//D
//the research made by the user
func TabArtist(index int, a int, state string) {
	Data = append(Data, artists_information[index])
	if state == "concertLocation" {
		reg, _ := regexp.Compile("-.*")
		Data[a].Locations.Locations = nil
		for _, place := range artists_information[index].Locations.Locations {
			matched := reg.FindString(place)
			matched = strings.Title(matched[1:])
			if matched == input1 {
				Data[a].Locations.Locations = append(Data[a].Locations.Locations, place)
			}
		}
	} else if state == "locAndDate" {
		for _, place := range artists_information[index].Locations.Locations {
			if input1 != place {
				delete(Data[a].Relations.DatesLocations, place)
			}
		}
	}
}

//append value to the main struct "Artist"
func Artists_info() {
	response, err := http.Get(config.APIURL + "/artists")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &artists_information)
}

//append value to the main struct "Artist" about the relation
func Artists_rela() {
	response, err := http.Get(config.APIURL + "/relation")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &artists_relation)
	if len(artists_relation.Index) == 0 {
		fmt.Println("array is empty")
	} else {
		for index := range artists_relation.Index {
			artists_information[index].Relations = artists_relation.Index[index]
		}
	}
}

//append value to the main struct "Artist" about the Location
func Artists_loc() {
	response, err := http.Get(config.APIURL + "/locations")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &locations_Index)

	if len(locations_Index.Index) == 0 {
		fmt.Println("array is empty")
	} else {
		for index := range locations_Index.Index {
			artists_information[index].Locations = locations_Index.Index[index]
		}
	}

	reg, _ := regexp.Compile("-.*") // regex use for finding the name of the countries

	for index := range artists_information {
		for i := range artists_information[index].Locations.Locations {
			matched := reg.FindString(artists_information[index].Locations.Locations[i]) // found the regex (with "-")
			AllLocations = append(AllLocations, strings.Title(matched[1:]))              // array with all countries from an artist (sans le "-")
		}
		AllLocations = removeDuplicateValues(AllLocations) // remove duplicate
		sort.Strings(AllLocations)                         // Sorts countries in alphabetical order
	}
}

//remove all value which have duplicate then return an array
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//input all informations about the date to each artist
func Artists_Date() {
	response, err := http.Get(config.APIURL + "/dates")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &date_Index)

	if len(date_Index.Index) == 0 {
		fmt.Println("array is empty")
	} else {
		for index := range date_Index.Index {
			artists_information[index].ConcertDates = date_Index.Index[index]

		}

	}
}

//insert value to data depending of the state
//	"state" equal to either "singleArtist", "concertLocation" or "locAndDate"
//	InsertData(state)
func InsertData(state string) {
	for index, artist := range artists_information {
		if state == "singleArtist" {
			if strings.ToLower(artist.Name) == strings.ToLower(input1) || input2 == strconv.Itoa(artist.Id) {
				Data = append(Data, artists_information[index])
				break
			}
		} else if state == "concertLocation" {
			if input2 == "All Artists" && input1 == "All Countries" {
				Data = append(Data, artists_information[index])

			} else if strings.ToLower(artist.Name) == strings.ToLower(input2) && input1 == "All Countries" {
				// if found then insert the artist to Data which is the data send to the page
				Data = append(Data, artists_information[index])
				break

			} else if strings.ToLower(artist.Name) == strings.ToLower(input2) && input1 != "All Countries" {
				TabArtist(index, 0, state)
				break

			} else if input2 == "All Artists" && input1 != "All Countries" {
				TabArtist(index, index, state)
			}
		} else if state == "locAndDate" {
			if input1 == "All Locations" {
				Data = append(Data, artists_information[index])
			} else if input1 != "All Locations" {
				TabArtist(index, index, state)
			}
		}
	}
}

//give each date on which an artist participates between 2 dates
func BestDate(w http.ResponseWriter, r *http.Request) {
	input1 = r.FormValue("trip-start")
	input2 = r.FormValue("trip-end")
	input3 = r.FormValue("artist-date")

	for i := range artists_information {
		if input3 == artists_information[i].Name || input3 == "All Artists" {
			break
		} else if i == 51 && input3 != artists_information[i].Name {
			input3 = input_backup
		}
	}

	Data = nil
	if input1 >= input2 {
		input1, input2 = input2, input1
	}
	reverseDates(input1, input2)

	input_backup = input3

	t := template.New("dateTemplate")
	t = template.Must(t.ParseFiles("./tmpl/bestdate.html", "./tmpl/header&footer.html", "./tmpl/content.html"))
	t.ExecuteTemplate(w, "bestdate", Data)
}

//	reverse the type of the dates from " dd-mm-yyyy " to " yyyy-mm-dd "
func reverseDates(input1 string, input2 string) {
	layout := "2006-01-02"
	var dat []string
	time.Parse(layout, input1)
	time.Parse(layout, input2)

	for index, val := range artists_information {
		dat = nil
		for _, date := range artists_information[index].ConcertDates.Dates {

			if string(date[0]) == "*" {
				date = date[1:]
			}

			date = date[6:10] + "-" + date[3:5] + "-" + date[0:2]
			time.Parse(layout, date)

			if input1 <= date && input2 >= date {
				if input3 == "All Artists" {
					dat = append(dat, date)
				} else if strings.ToLower(val.Name) == strings.ToLower(input3) {
					dat = append(dat, date)
				}
			}
		}
		if dat != nil {
			Data = append(Data, artists_information[index])
			Data[len(Data)-1].ConcertDates.Dates = dat
		}
	}
}

//information page for places where concert attend to
func locAndDate(w http.ResponseWriter, r *http.Request) {
	Artists_rela()
	input1 = r.FormValue("loc")
	state = "locAndDate"
	Data = nil
	var boolart = false
	InsertData(state)
	var data_copy []config.Artist
	// loop over the artists to remove anyone without value in DatesLocations
	for index := range artists_information {
		for _, v := range artists_information[index].Relations.DatesLocations {
			if v != nil {
				boolart = true
				// reverse the date so we can work with it
				for i := 0; i < len(v); i++ {
					v[i] = v[i][6:10] + "-" + v[i][3:5] + "-" + v[i][0:2]
				}
			}
		}
		if boolart {
			data_copy = append(data_copy, Data[index])
			boolart = false
		}
	}

	// cherche si ce qui est dans un des input correspond à un des artistes

	t := template.New("loc&date-template")
	t = template.Must(t.ParseFiles("./tmpl/loc&date.html", "./tmpl/header&footer.html", "./tmpl/content.html"))
	t.ExecuteTemplate(w, "loc&date", data_copy)
}
