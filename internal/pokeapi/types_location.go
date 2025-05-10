package pokeapi

// https://pokeapi.co/api/v2/location-area/{id or name}/
// By adding a name or id, the API will return a lot more information about the location area.
// without a name or id, the API will return a list of locations

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// https://pokeapi.co/api/v2/location-area/{id or name}/

type PokemonInArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}