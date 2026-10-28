package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("error: 'inspect' takes a pokemon name argument, none were given")
	}

	if len(args) > 1 {
		return fmt.Errorf("error: 'inspect' takes only 1 argument, %d were given", len(args))
	}

	poke, exists := cfg.pokedex[args[0]]
	if !exists {
		return fmt.Errorf("you haven't caught that one yet")
	}

	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Printf("Species: %s\n", poke.Species.Name)
	fmt.Println("Types: ")
	for _, pokeType := range poke.Types {
		fmt.Printf("  -%s\n", pokeType.Type.Name)
	}

	return nil
}
