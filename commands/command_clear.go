package commands

import (
	"os"
	"os/exec"
)

func commandClear(cfg *ApiConfig) error {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()

	return nil
}