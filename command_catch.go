package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("error: 'catch' takes a pokemon name argument, none were given")
	}

	if len(args) > 1 {
		return fmt.Errorf("error: 'catch' takes only 1 argument, %d were given", len(args))
	}

	if _, exists := cfg.pokedex[args[0]]; exists {
		return fmt.Errorf("%s is already in your pokedex!", args[0])
	}

	poke, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", poke.Name)
	if getProbabilty(poke.BaseExperience) > rand.Float64() {
		fmt.Printf("%s was caught!\n", poke.Name)
		cfg.pokedex[poke.Name] = poke
	} else {
		fmt.Printf("%s escaped!\n", poke.Name)
	}

	return nil
}

func getProbabilty(experience int) (probability float64) {
	// i will use a linear scaling formula for determining the probabilty
	// also for our convience i will use max and min rates
	MINRATE := 0.02
	MAXRATE := 0.98

	// the pokeapi base_experience ranges beetween 20 and 608
	INFSET := 20
	SUPSET := 608

	baseExp := max(INFSET, min(SUPSET, experience))
	scale := float64(baseExp-INFSET) / float64(SUPSET-INFSET)
	probability = MAXRATE - (float64(scale) * (MAXRATE - MINRATE))
	return

}
