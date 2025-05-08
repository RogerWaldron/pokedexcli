package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	// user types something and presses enter, 
	// the code continues and the input is available 
	// in the returned bufio.Scanner
	// commands := map[string]interface{}{
	commands := 

	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
        if command, exists := commands[text]; exists {
            // Call a hardcoded function
			config := config{}
            err := command.callback(&config)
			if err != nil {
				fmt.Println(fmt.Errorf("Error: %w", err))
			}
        } else {
			fmt.Println("Unknown command")
		}
		printPrompt()
	}
}

func printPrompt() {
	fmt.Print("Pokedex >")
}