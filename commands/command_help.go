package commands

import (
	"fmt"

	"github.com/RogerWaldron/pokedexcli/types"
)

func commandHelp(cfg *types.ApiConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}