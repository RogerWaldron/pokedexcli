package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RogerWaldron/pokedexcli/commands"
	"github.com/RogerWaldron/pokedexcli/internal/pokeapi"
	"github.com/RogerWaldron/pokedexcli/utils"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 60 * time.Second)
	cfg := &commands.ApiConfig{
		ApiClient: pokeClient,
		Text: []string{},
	}
	reader := bufio.NewScanner(os.Stdin)
	commands := commands.GetCommands()
	printPrompt()
	for reader.Scan() {
		text := utils.CleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}

		query := strings.Split(text, " ")
		
		if len(query) > 1 {
			if command, exists := commands[query[0]]; exists {
				cfg.Text = query
            	err := command.Callback(cfg) // TODO: use later to pass Pokeapi
				if err != nil {
					fmt.Println(fmt.Errorf("Error: %w", err))
				}			
			} else {
				fmt.Println("Unknown command")
			}
		} else if command, exists := commands[text]; exists {
            // Call a hardcoded function
			cfg.Text = query
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