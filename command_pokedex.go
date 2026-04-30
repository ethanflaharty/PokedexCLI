package main

import "fmt"

func commandPokedex(cfg *config, names ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.pokedex {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}
