package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerInput := strings.ToLower(text)
	words := strings.Fields(lowerInput)
	return words
}
