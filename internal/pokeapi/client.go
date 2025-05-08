package pokeapi

import (
	"net/http"
	"time"

	"github.com/RogerWaldron/pokedexcli/types"
)

func NewClient(timeout time.Duration) types.Client{
	return types.Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}