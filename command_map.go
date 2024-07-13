package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeApiClient.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}
	cfg.previousUrl = locationsResp.Previous
	cfg.nextUrl = locationsResp.Next
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousUrl == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := cfg.pokeApiClient.ListLocations(cfg.previousUrl)
	if err != nil {
		return err
	}
	cfg.previousUrl = locationsResp.Previous
	cfg.nextUrl = locationsResp.Next
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
