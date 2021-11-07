package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pokedex/api/pokeapi"

	"github.com/gorilla/mux"
)

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
	fmt.Println("receivedGET /" + pokeName)
	defer fmt.Println("processed GET /" + pokeName)

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
}

func handleTranslated(w http.ResponseWriter, r *http.Request) {
	pokeName := mux.Vars(r)["name"]

	fmt.Println("GET /translated/" + pokeName)
	fmt.Fprint(w, "GET /translated/"+pokeName)

}
