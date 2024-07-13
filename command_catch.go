package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonResp, err := cfg.pokeApiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...", pokemonResp.Name)

	res := rand.Intn(pokemonResp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResp.Name)

	cfg.caught[pokemonResp.Name] = pokemonResp

	return nil
}
