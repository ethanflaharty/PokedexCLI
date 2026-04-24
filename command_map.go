package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	var url string
	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
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

type locationResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
