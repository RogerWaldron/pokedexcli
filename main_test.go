package main

import (
	"testing"

	"github.com/RogerWaldron/pokedexcli/utils"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input string
		expected string
	}{
		{
			input: " hello world ",
			expected: "hello world",
		},
		{
			input: "helloworld ",
			expected: "helloworld",
		},
		{
			input: " hello-world ",
			expected: "hello-world",
		},
	}

	for _, c := range cases {
		actual := utils.CleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected: %v, Got: %v", expectedWord, word)
			}
		}
	}
}