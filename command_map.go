package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type defLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request", err)
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return err
	}

	locations := defLocations{}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println("Error unmarshaling response: ", err)
		return err
	}

	fmt.Printf("Total count: %v\n", locations.Count)
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
