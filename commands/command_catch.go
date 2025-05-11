package commands

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *ApiConfig) error {
	if len(cfg.Text) < 2 {
		return nil
	}
	pokemon, err := cfg.ApiClient.PokemonDetails(cfg.Text[1])
	if err != nil {
		return err
	}
	// TODO: FIX - above should error is pokemon doesn't exist
	// Pokedex >catch picachu
	// Throwing a Pokeball at s..
    // panic: invalid argument to Intn

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	// The higher the base experience, the harder it should be to catch.
	chance := rand.Intn(pokemon.BaseExperience) + 1
	threshold := int(math.Floor(float64(pokemon.BaseExperience) * 0.8))
	if chance > threshold {
		fmt.Printf("caught %s\n", pokemon.Name)
		cfg.Captured[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}