package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(PokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + PokemonName

	if val, ok := c.cache.Get(url); ok {
		PokemonResp := Pokemon{}
		err := json.Unmarshal(val, &PokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return PokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	PokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &PokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return PokemonResp, nil
}
