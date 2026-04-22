package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"empty":              {input: "  ", expected: []string{}},
		"hello world":        {input: " hello world ", expected: []string{"hello", "world"}},
		"starters":           {input: "bulbasaur charmander squirtle", expected: []string{"bulbasaur", "charmander", "squirtle"}},
		"one word":           {input: "hello", expected: []string{"hello"}},
		"surrounding spaces": {input: "  hello world  ", expected: []string{"hello", "world"}},
		"random capitals":    {input: "  HelLO WOrlD", expected: []string{"hello", "world"}},
	}

	for _, test := range tests {
		actual := cleanInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("lengths don't match: actual length: %v; expected length: %v", len(actual), len(test.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := test.expected[i]
			if word != expectedWord {
				t.Errorf("Word %v: %v, Expected Word %v: %v", i+1, word, i+1, expectedWord)
			}
		}
	}
}
