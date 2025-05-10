package commands

import (
	"github.com/RogerWaldron/pokedexcli/internal/pokeapi"
)

type ApiConfig struct {
	ApiClient        pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	Text             []string
}