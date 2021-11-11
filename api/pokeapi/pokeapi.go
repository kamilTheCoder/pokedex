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
)

type Api struct{}

// TranslatorRequest is a base struct for fun translations API request
type TranslatorRequest struct {
	Text string `json:"text"`
}

// TranslatorResponse is a base struct for fun translations API response
type TranslatorResponse struct {
	Contents *struct {
		Translated string `json:"translated"`
	} `json:"contents"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

// GetPokemon makes request to pokeapi.co trying to get information on a pokemon
// If successful, returns a normalised pokemon struct
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

// GetTranslation makes a request for translation from a third-party funtranslations.com API
// If succesful, returns translated string. Otherwise, returns an empty string and an error
func (a Api) GetTranslation(translator string, text string) (string, error) {
	body, err := json.Marshal(TranslatorRequest{text})
	if err != nil {
		return "", fmt.Errorf("error masharlling %v translation: %w", translator, err)
	}

	url := translatorApiPath + translator + ".json"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error fetching %v translation: %w", translator, err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response for %v translation: %w", translator, err)
	}

	var translation TranslatorResponse
	err = json.Unmarshal(respData, &translation)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response for %v: %w", translator, err)
	}

	if translation.Error != nil {
		return "", fmt.Errorf("error calling translation API: %v", translation.Error.Message)
	}

	return translation.Contents.Translated, nil
}

// Example of an erro message from the translation API after reaching the hourly rate limit
// {
//     "error": {
//         "code": 429,
//         "message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 57 minutes and 23 seconds."
//     }
// }
