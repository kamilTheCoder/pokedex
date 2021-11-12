package api

import "pokedex/pokemon"

const (
	TranslatorShakespeare = "shakespeare"
	TranslatorYoda        = "yoda"
)

type Api interface {
	GetPokemon(name string) (pokemon.PokemonRaw, error)
	GetTranslation(translator string, text string) (string, error)
}
