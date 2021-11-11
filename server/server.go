package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pokedex/api"
	"pokedex/api/pokeapi"
	"pokedex/cache"
	poke "pokedex/pokemon"

	"github.com/gorilla/mux"
)

var pokeApi pokeapi.Api

type ApiError struct {
	Error string `json:"error"`
}

// HandleRequests serves as the main API request handler, setting up routes and starting up the server
func HandleRequests() {
	pokeApi = pokeapi.Api{}
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/{name}", handlePokemon)
	router.HandleFunc("/translated/{name}", handleTranslated)

	fmt.Println("serving")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handlePokemon(w http.ResponseWriter, r *http.Request) {
	pokeName := mux.Vars(r)["name"]
	fmt.Println("received GET /" + pokeName)
	defer fmt.Println("processed GET /" + pokeName)

	if p, ok := cache.Access().Get(pokeName); ok {
		fmt.Println("getting " + pokeName + " from the cache")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
		return
	}

	p, err := createNewPokemon(pokeName)
	if err != nil {
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(ApiError{"error fetching pokemon information for " + pokeName})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func handleTranslated(w http.ResponseWriter, r *http.Request) {
	pokeName := mux.Vars(r)["name"]

	fmt.Println("received GET /translated/" + pokeName)
	defer fmt.Println("processed GET /translated/" + pokeName)

	var err error
	p, ok := cache.Access().Get(pokeName)
	if !ok {
		p, err = createNewPokemon(pokeName)
		if err != nil {
			fmt.Println(err.Error())
			json.NewEncoder(w).Encode(ApiError{"error fetching pokemon information for " + pokeName})
			return
		}
	}

	translator := getTranslator(p)
	if desc, ok := p.VariantDescriptions[translator]; ok {
		fmt.Printf("getting %s's %s description from the cache\n", pokeName, translator)
		p.Description = desc

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
		return
	}

	fmt.Printf("fetching %s's %s description from the API\n", pokeName, translator)
	translated, err := pokeApi.GetTranslation(translator, p.Description)
	if err != nil {
		// error at this point indicates we could not get a translation for some reason,
		// quite possibly breaking the limit rate.
		// as per specs, if we can't translate for any reason, return standard description.
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
		return
	}

	p.VariantDescriptions[translator] = translated
	cache.Access().Update(p)

	p.Description = translated
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func createNewPokemon(pokeName string) (poke.Pokemon, error) {
	fmt.Println("creating new pokemon: " + pokeName)
	raw, err := pokeApi.GetPokemon(pokeName)
	if err != nil {
		return poke.Pokemon{}, err
	}

	p := raw.ToPokemon()
	fmt.Println("adding " + p.Name + " to the cache")
	cache.Access().AddNew(p)

	return p, nil
}

func getTranslator(pokemon poke.Pokemon) string {
	if pokemon.Habitat == poke.HabitatCave ||
		pokemon.Legendary {

		return api.TranslatorYoda
	}
	return api.TranslatorShakespeare
}
