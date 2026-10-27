package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("error: 'exit' doesn't take any arguments!")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil //i put this to make the compiler shut up (even tho it will never reach)
}
