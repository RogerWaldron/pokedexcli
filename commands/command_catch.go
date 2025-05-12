package commands

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"

	"github.com/RogerWaldron/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *ApiConfig) error {
	if len(cfg.Text) < 2 {
		return nil
	}

	_, ok := cfg.Captured[cfg.Text[1]]
	if ok {
		return fmt.Errorf("you have already caught that pokemon")
	}

	pokemon, err := cfg.ApiClient.PokemonDetails(cfg.Text[1])
	if err != nil {
		return err
	}

	isEmpty := pokeapi.Pokemon{}

	if reflect.DeepEqual(pokemon, isEmpty) {
		return fmt.Errorf("%s does not exist", cfg.Text[1])
	}

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