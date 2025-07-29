package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokémon name")
	}

	name := args[0]

	pokemon, exists := cfg.pokedex[name]
	if !exists {
		return errors.New("you have not caught that pokémon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, info := range pokemon.Types {
		fmt.Printf("  %v\n", info.Type.Name)
	}

	return nil
}
