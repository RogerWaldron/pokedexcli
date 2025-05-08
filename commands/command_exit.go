package commands

import (
	"fmt"
	"os"

	"github.com/RogerWaldron/pokedexcli/types"
)

func commandExit(cfg *types.ApiConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
