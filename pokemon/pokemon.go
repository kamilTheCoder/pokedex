package pokemon

const (
	languageEnglish = "en"
)

type Pokemon struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Habitat     string `json:"habitat"`
	Legendary   bool   `json:"isLegendary"`
}

type PokemonRaw struct {
	Name         string        `json:"name"`
	FalvourTexts []FlavourText `json:"flavor_text_entries"`
	Habitat      Habitat       `json:"habitat"`
	Legendary    bool          `json:"is_legendary"`
}

type FlavourText struct {
	FalvourText string `json:"flavor_text"`
	Language    struct {
		Name string `json:"name"`
	} `json:"language"`
}

type Habitat struct {
	Name string `json:"name"`
}

func (r PokemonRaw) ToPokemon() Pokemon {
	p := Pokemon{
		Name:      r.Name,
		Habitat:   r.Habitat.Name,
		Legendary: r.Legendary,
	}

	for _, t := range r.FalvourTexts {
		if t.Language.Name == languageEnglish {
			p.Description = t.FalvourText
			break
		}
	}

	return p
}
