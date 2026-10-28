package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("error: 'exit' doesn't take any arguments!")
	}

	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("your pokedex is empty, go catch some pokemons to add them here")
	}
	fmt.Println("Your pokedex:")
	for key := range cfg.pokedex {
		fmt.Printf("  -%s\n", key)
	}

	return nil
}
