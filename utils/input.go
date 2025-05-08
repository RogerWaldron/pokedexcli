package utils

import "strings"

func cleanInput(text string) string {
	trimmed := strings.TrimSpace(text)
	words := strings.Split(strings.ToLower(trimmed), " ")

	return strings.Join(words, " ")
}
