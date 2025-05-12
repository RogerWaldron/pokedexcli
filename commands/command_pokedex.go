package commands

import "fmt"

func commandPokedex(cfg *ApiConfig) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Captured {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}