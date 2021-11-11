package pokemon

const (
	languageEnglish = "en"
	HabitatCave     = "cave"
)

// Pokemon represents data as it is stored within the pokedex
type Pokemon struct {
	Name                string            `json:"name"`
	VariantDescriptions map[string]string `json:"-"`
	Description         string            `json:"description"`
	Habitat             string            `json:"habitat"`
	Legendary           bool              `json:"isLegendary"`
}

// PokemonRaw represents data as it arrives from pokeAPI
type PokemonRaw struct {
	Name         string        `json:"name"`
	FalvourTexts []FlavourText `json:"flavor_text_entries"`
	Habitat      Habitat       `json:"habitat"`
	Legendary    bool          `json:"is_legendary"`
}

// Contains pokemon description information and the language version of the description
type FlavourText struct {
	FalvourText string `json:"flavor_text"`
	Language    struct {
		Name string `json:"name"`
	} `json:"language"`
}

// Habitat contains information on where does a pokemon usually reside
type Habitat struct {
	Name string `json:"name"`
}

// ToPokemon converts raw representation of the pokemon data from pokeAPI
// into internal representation, which is to be exposed via this pokedex API
func (r PokemonRaw) ToPokemon() Pokemon {
	p := Pokemon{
		Name:                r.Name,
		Habitat:             r.Habitat.Name,
		Legendary:           r.Legendary,
		VariantDescriptions: map[string]string{},
	}

	for _, t := range r.FalvourTexts {
		if t.Language.Name == languageEnglish {
			p.Description = t.FalvourText
			break
		}
	}

	return p
}
