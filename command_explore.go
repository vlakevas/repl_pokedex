package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("error: 'explore' takes 1 location argument, none were given")
	}

	if len(args) > 1 {
		return fmt.Errorf("error: 'explore' takes only 1 argument, %d were given", len(args))
	}

	pokemonResponse, err := cfg.pokeapiClient.GetPokemons(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range pokemonResponse.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
