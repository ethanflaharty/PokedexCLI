package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		lowerInput := strings.ToLower(input)
		words := strings.Fields(lowerInput)
		fmt.Printf("Your command was: %v\n", words[0])
	}
}
