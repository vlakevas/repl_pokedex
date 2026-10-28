package main

import (
	"time"

	"github.com/vlakevas/repl_pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		next:          pokeapi.GetLocationAreaUrl(),
		previous:      nil,
		pokedex:       map[string]pokeapi.IndividualPokemon{},
	}
	startRepl(cfg)
}
