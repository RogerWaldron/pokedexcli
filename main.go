package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/RogerWaldron/pokedexcli.git/utils"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)

	// printPrompt()
	for reader.Scan() {
		// text := CleanInput(reader.Text())
        // if command, exists := commands[text]; exists {
        //     // Call a hardcoded function
		// 	config := config{}
        //     err := command.callback(&config)
		// 	if err != nil {
		// 		fmt.Println(fmt.Errorf("Error: %w", err))
		// 	}
        // } else {
		// 	fmt.Println("Unknown command")
		// }
		printPrompt()
	}
}

func printPrompt() {
	fmt.Print("Pokedex >")
}