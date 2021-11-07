package mockapi

type MockApi struct{}

func (a MockApi) GetHabitat(pokemonName string) string {
	return "cave"
}

func (a MockApi) GetDescription(pokemonName string) string {
	return pokemonName + "'s description"
}

func (a MockApi) GetLegendaryStatus(species string) bool {
	return false
}

func (a MockApi) GetTranslation(translator string, text string) string {
	return "[" + translator + "]" + text
}
