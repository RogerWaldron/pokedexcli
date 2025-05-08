package commands

import "github.com/RogerWaldron/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}