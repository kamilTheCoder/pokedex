package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokedex/pokemon"
	"strings"
)

const (
	pokemonApiPath = "https://pokeapi.co/api/v2/pokemon-species/"
)

type Api struct{}

func (a Api) GetPokemon(name string) (pokemon.PokemonRaw, error) {
	resp, err := http.Get(pokemonApiPath + name + "/")
	if err != nil {
		return pokemon.PokemonRaw{}, fmt.Errorf("error fetching %v: %w", name, err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pokemon.PokemonRaw{}, fmt.Errorf("error reading response for %v: %w", name, err)
	}

	if strings.HasPrefix(string(respData), "Not Found") {
		return pokemon.PokemonRaw{}, fmt.Errorf("error fetching %v: %w", name, errors.New("not found"))
	}

	var raw pokemon.PokemonRaw
	err = json.Unmarshal(respData, &raw)
	if err != nil {
		return pokemon.PokemonRaw{}, fmt.Errorf("error unmarshalling response for %v: %w", name, err)
	}

	return raw, nil
}

func (a Api) GetTranslation(translator string, text string) string {
	return "[" + translator + "]" + text
}
