package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokemonLocations, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if e, ok := c.cache.Get(url); ok {
		fmt.Println("Loading data from cache...")
		locations := PokemonLocations{}

		err := json.Unmarshal(e, &locations)
		if err != nil {
			return PokemonLocations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonLocations{}, err
	}

	c.cache.Add(url, data)
	fmt.Printf("Saved data to cache (%v)...\n", c.cache.Count())

	locations := PokemonLocations{}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return PokemonLocations{}, err
	}

	return locations, nil
}
