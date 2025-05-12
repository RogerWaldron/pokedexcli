package commands

import "fmt"

func commandPokedex(cfg *ApiConfig) error {
	for _, pokemon := range cfg.Captured {
		fmt.Println(pokemon.Name)
	}
	return nil
}