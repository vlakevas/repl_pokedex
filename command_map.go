package main

import "fmt"

func commandMapf(cfg *config) error {
	if cfg.next == nil {
		fmt.Println("You are on the last page!")
		return nil
	}

	locationResponse, err := cfg.pokeapiClient.GetLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationResponse.Next
	cfg.previous = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You are on the first page!")
		return nil
	}

	locationResponse, err := cfg.pokeapiClient.GetLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locationResponse.Next
	cfg.previous = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
