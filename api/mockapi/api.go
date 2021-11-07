package mockapi

import "pokedex/pokemon"

type Api struct{}

func (a Api) GetPokemon(name string) (pokemon.Pokemon, error) {
	return pokemon.Pokemon{Name: name}, nil
}

func (a Api) GetTranslation(translator string, text string) string {
	return "[" + translator + "]" + text
}
