package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, names ...string) error {
	if len(names) < 1 {
		return errors.New("you must provide a location name\n")
	}
	var url string
	url = "https://pokeapi.co/api/v2/location-area/" + names[0]

	var bodyBytes []byte
	if cached, ok := cfg.cache.Get(url); ok {
		bodyBytes = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		bodyBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, bodyBytes)
	}

	var data locationData
	errUnmarsh := json.Unmarshal(bodyBytes, &data)
	if errUnmarsh != nil {
		return errUnmarsh
	}
	fmt.Printf("Exploring %v...\n", names[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

type locationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
