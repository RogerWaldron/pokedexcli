package pokeapi

import (
	"net/http"
	"time"

	"github.com/RogerWaldron/pokedexcli/internal/pokecache"
)

type Client struct {
	cache 	   	pokecache.Cache
	HttpClient 	http.Client
}

func NewClient(cacheInterval, timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}