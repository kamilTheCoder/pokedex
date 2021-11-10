package pokeapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokedex/pokemon"
	"strings"
)

const (
	pokemonApiPath    = "https://pokeapi.co/api/v2/pokemon-species/"
	translatorApiPath = "https://api.funtranslations.com/translate/"
	useRealTranslator = false
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

func (a Api) GetTranslation(translator string, text string) (string, error) {
	if !useRealTranslator {
		return "[" + translator + "]" + text, nil
	}

	var body = []byte(`{"title":"` + text + `"}`)
	resp, err := http.Post(translatorApiPath+translator+"/", "text", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error fetching %v translation: %w", translator, err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response for %v translation: %w", translator, err)
	}

	if strings.HasPrefix(string(respData), "Not Found") {
		return "", fmt.Errorf("error fetching %v translation: %w", translator, err)
	}

	var raw pokemon.PokemonRaw
	err = json.Unmarshal(respData, &raw)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response for %v translation: %w", translator, err)
	}

	return "", nil
}
