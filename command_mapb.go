package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(cfg *config) error {
	var url string
	if cfg.Previous == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.Previous
	}

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

	var data locationResponse
	errUnmarsh := json.Unmarshal(bodyBytes, &data)
	if errUnmarsh != nil {
		return errUnmarsh
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
