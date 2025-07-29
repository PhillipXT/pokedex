package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Location(locationName string) (LocationDetails, error) {

	url := baseURL + "/location-area/" + locationName

	if e, ok := c.cache.Get(url); ok {
		fmt.Println("Loading data from cache...")
		details := LocationDetails{}

		err := json.Unmarshal(e, &details)
		if err != nil {
			return LocationDetails{}, err
		}

		return details, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	details := LocationDetails{}

	err = json.Unmarshal(data, &details)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Add(url, data)
	fmt.Printf("Saved data to cache (%v)...\n", c.cache.Count())

	return details, nil
}

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

	locations := PokemonLocations{}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return PokemonLocations{}, err
	}

	c.cache.Add(url, data)
	fmt.Printf("Saved data to cache (%v)...\n", c.cache.Count())

	return locations, nil
}
