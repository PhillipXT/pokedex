package main

import (
	"time"

	"github.com/PhillipXT/pokedex/internal/pokeapi"
)

// To run the program and log the output to a file instead of stdout:
// > go run . | tee repl.log

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
