package commands

import (
	"fmt"
	"os"
)

func commandExit(cfg *ApiConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
