package commands

import (
	"errors"
	"fmt"
)

func commandMapNext(cfg *ApiConfig) error {
	resp, err := cfg.ApiClient.LocationList(cfg.NextLocationsURL)
	if err != nil {
		return err
	}
	cfg.NextLocationsURL = &resp.Next
	cfg.PrevLocationsURL = &resp.Previous
	
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapPrev(cfg *ApiConfig) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("You are on the first page")
	}

	resp, err := cfg.ApiClient.LocationList(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}
	cfg.NextLocationsURL = &resp.Next
	cfg.PrevLocationsURL = &resp.Previous
	
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}	