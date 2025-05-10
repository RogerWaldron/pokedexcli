package commands

import (
	"fmt"
)

/*
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
*/
func commandExplore(cfg *ApiConfig) error {
	if len(cfg.Text) < 2 {
		return nil
	}

	resp, err := cfg.ApiClient.PokemonInAreaList(cfg.Text[1])
	if err != nil {
		return err
	}
	
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}