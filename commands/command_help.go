package commands

import (
	"fmt"
)

func commandHelp(cfg *ApiConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("explore location_name: get list of Pokemon located in area")
	fmt.Println("catch: throw a pokeball to try and catch Pokemon")
	fmt.Println("inspect: display details for pokemon you have caught")
	fmt.Println("pokedex: list pokemon you have caught")
	fmt.Println("map: display more location areas")
	fmt.Println("mapb: display previous location areas")

	return nil
}