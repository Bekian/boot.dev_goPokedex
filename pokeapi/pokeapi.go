package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// takes ID or name as input
func GetLocationViaInput(input *string) (mapResponse LocationIDResponse) {
	// set the base uri
	uri := "https://pokeapi.co/api/v2/location"
	// append input
	if input != nil {
		uri += "/" + *input
	}
	// request to uri
	res, err := http.Get(uri)
	check(err)
	defer res.Body.Close()
	// decode to the mapResponse pointer
	err = json.NewDecoder(res.Body).Decode(&mapResponse)
	check(err)
	fmt.Println(mapResponse)
	return
}

// returns a paginated list of the input
// takes a page number, defaults to 0
func GetLocation(page *int) (mapResponse NamedAPIResourceList) {
	var uri string
	// determine uri based on provided page number
	if page == nil {
		uri = "https://pokeapi.co/api/v2/location?offset=0"
	} else {
		pageNumber := *page
		if pageNumber < 0 {
			panic("Provided page is less than 0")
		}
		uri = "https://pokeapi.co/api/v2/location?offset=" + strconv.Itoa(pageNumber*20)
	}
	// request to uri
	res, err := http.Get(uri)
	check(err)
	defer res.Body.Close()
	// decode to the mapResponse pointer
	err = json.NewDecoder(res.Body).Decode(&mapResponse)
	check(err)
	return
}

// need to save the next location and set it up as the next uri

func QueryPokemon(pokemon string) (pokemonResponse Pokemon) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon))
	check(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(body))
}
