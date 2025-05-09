package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/RogerWaldron/pokedexcli/commands"
	"github.com/RogerWaldron/pokedexcli/internal/pokeapi"
	"github.com/RogerWaldron/pokedexcli/utils"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &commands.ApiConfig{
		ApiClient: pokeClient,
	}

	reader := bufio.NewScanner(os.Stdin)
	commands := commands.GetCommands()
	printPrompt()
	for reader.Scan() {
		text := utils.CleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}

        if command, exists := commands[text]; exists {
            // Call a hardcoded function
            err := command.Callback(cfg) // TODO: use later to pass Pokeapi
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