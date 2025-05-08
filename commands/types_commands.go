package commands

import "github.com/RogerWaldron/pokedexcli/types"

type cliCommand struct {
	name        string
	description string
	Callback    func(*types.ApiConfig) error
}