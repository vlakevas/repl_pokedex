package main

import (
	"fmt"
)

func commandHelp(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("error: 'help' doesn't take any arguments!")
	}

	usage := ""
	for _, value := range cfg.commands {
		usage = fmt.Sprintf("%s%s: %s\n", usage, value.name, value.description)
	}

	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", usage)
	return nil
}
