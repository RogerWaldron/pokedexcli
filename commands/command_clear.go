package commands

import (
	"os"
	"os/exec"

	"github.com/RogerWaldron/pokedexcli/types"
)

func commandClear(cfg *types.ApiConfig) error {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()

	return nil
}