package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your pokédex:")
	for _, p := range cfg.pokedex {
		fmt.Println("  ", p.Name)
	}

	return nil
}
