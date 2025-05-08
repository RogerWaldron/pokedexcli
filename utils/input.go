package utils

import "strings"

func CleanInput(text string) string {
	trimmed := strings.TrimSpace(text)
	words := strings.Split(strings.ToLower(trimmed), " ")

	return strings.Join(words, " ")
}
