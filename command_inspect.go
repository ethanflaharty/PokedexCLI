package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, names ...string) error {
	if len(names) < 1 {
		return errors.New("you must write a Pokemon name")
	}
	_, ok := cfg.pokedex[names[0]]
	if !ok {
		fmt.Println("you have not caught that Pokemon")
	} else {
		fmt.Printf("Name: %v\n", cfg.pokedex[names[0]].Name)
		fmt.Printf("Height: %v\n", cfg.pokedex[names[0]].Height)
		fmt.Printf("Weight: %v\n", cfg.pokedex[names[0]].Weight)
		fmt.Println("Stats:")
		fmt.Printf("  -hp: %v\n", cfg.pokedex[names[0]].Stats[0].BaseStat)
		fmt.Printf("  -attack: %v\n", cfg.pokedex[names[0]].Stats[1].BaseStat)
		fmt.Printf("  -defense: %v\n", cfg.pokedex[names[0]].Stats[2].BaseStat)
		fmt.Printf("  -special-attack: %v\n", cfg.pokedex[names[0]].Stats[3].BaseStat)
		fmt.Printf("  -special-defense: %v\n", cfg.pokedex[names[0]].Stats[4].BaseStat)
		fmt.Printf("  -speed: %v\n", cfg.pokedex[names[0]].Stats[5].BaseStat)
		fmt.Println("Types:")
		fmt.Printf("  - %v\n", cfg.pokedex[names[0]].Types[0].Type.Name)
		if len(cfg.pokedex[names[0]].Types) > 1 {
			fmt.Printf("  - %v\n", cfg.pokedex[names[0]].Types[1].Type.Name)

		}
	}
	return nil
}
