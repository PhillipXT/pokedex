package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a Pokémon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := (1000 - pokemon.BaseExperience) / 10
	if chance < 25 {
		chance = 25
	} else if chance > 80 {
		chance = 80
	}

	fmt.Printf("Throwing a Pokéball at %s...\n", name)
	fmt.Printf("You have a %v%% chance to catch %v\n", chance, pokemon.Name)

	roll := rand.Intn(100)

	if roll <= chance {
		fmt.Printf("You rolled %v and caught %v!\n", roll, pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("You rolled %v and %v escaped!\n", roll, pokemon.Name)
	}

	return nil
}
