package api

import "pokedex/pokemon"

const (
	Translator_Shakespeare = "shakespeare"
	Translator_Yoda        = "yoda"
)

type Api interface {
	GetPokemon(name string) (pokemon.Pokemon, error)
	GetTranslation(translator string, text string) (string, error)
}
