package api

const (
	Translator_Shakespeare = "shakespeare"
	Translator_Yoda = "yoda"
)

type Api interface {
	GetHabitat(pokemonName string) string
	GetDescription(pokemonName string) string
	GetLegendaryStatus(species string) bool
	GetTranslation(translator string, text string) string
}

type Pokemon struct {
	Name string
	Description string	// https://pokeapi.co/api/v2/characteristic/{id}/ descriptions
	Habitat string		// https://pokeapi.co/api/v2/pokemon-habitat/{id or name}/
	Legendary bool		// https://pokeapi.co/api/v2/pokemon-species/{id or name}/ is_legendary
}