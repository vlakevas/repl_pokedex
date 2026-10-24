package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	usage := ""
	for _, value := range cfg.commands {
		usage = fmt.Sprintf("%s%s: %s\n", usage, value.name, value.description)
	}

	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", usage)
	return nil
}
