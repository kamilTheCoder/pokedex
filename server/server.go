package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pokedex/api/pokeapi"
	"pokedex/cache"

	"github.com/gorilla/mux"
)

// HandleRequests serves as the main API request handler, setting up routes and starting up the server
func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleMain)
	router.HandleFunc("/{name}", handlePokemon)
	router.HandleFunc("/translated/{name}", handleTranslated)

	fmt.Println("serving")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	fmt.Fprint(w, "hello")
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

	api := pokeapi.Api{}
	raw, err := api.GetPokemon(pokeName)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprint(w, err.Error())
		return
	}

	p := raw.ToPokemon()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
	fmt.Println("adding " + p.Name + " to the cache")
	cache.Access().AddNew(p)
}

func handleTranslated(w http.ResponseWriter, r *http.Request) {
	pokeName := mux.Vars(r)["name"]

	fmt.Println("GET /translated/" + pokeName)
	fmt.Fprint(w, "GET /translated/"+pokeName)

}
