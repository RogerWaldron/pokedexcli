package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GET https://pokeapi.co/api/v2/pokemon/{id or name}/
func(c *Client) PokemonDetails(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	cachedData, ok := c.cache.Get(url)
	if ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
	
		return pokemonResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}
	
	return pokemon, nil
}