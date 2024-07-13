package main

import (
	"time"

	"github.com/carterdr/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokemon := make(map[string]pokeapi.Pokemon)
	cfg := &config{pokeApiClient: pokeClient, caught: pokemon}
	startRepl(cfg)
}
