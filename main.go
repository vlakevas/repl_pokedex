package main

import (
	"time"

	"github.com/vlakevas/repl_pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		next:          pokeapi.GetLocationAreaUrl(),
		previous:      nil,
	}
	startRepl(cfg)
}
