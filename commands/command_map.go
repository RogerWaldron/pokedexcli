package commands

func commandMap(config *config) error {
	endpointLocAreas := "https://pokeapi.co/api/v2/location-area/"
	req, err := http.NewRequest(http.MethodGet, endpointLocAreas, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Problem requesting Location Areas: ", err)
	}	
	defer resp.Body.Close()
	Body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("")
	}
	decoder := json.NewDecoder(&body)
	data := decoder()
	for _, loc := range  {
		fmt.Println(loc.name)
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}