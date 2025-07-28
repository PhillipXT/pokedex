package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokemonLocations, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	return locations, nil
}
