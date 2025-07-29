package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {

	url := baseURL + "/pokemon/" + name

	if e, ok := c.cache.Get(url); ok {
		fmt.Println("Loading data from cache...")
		details := Pokemon{}

		err := json.Unmarshal(e, &details)
		if err != nil {
			return Pokemon{}, err
		}

		return details, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	details := Pokemon{}

	err = json.Unmarshal(data, &details)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	fmt.Printf("Saved data to cache (%v)...\n", c.cache.Count())

	return details, nil
}
