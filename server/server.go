package server

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Println("GET /{" + "???" + "}")
	fmt.Fprint(w, "GET /{"+"???"+"}")
}

func handleTranslated(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /translated/{" + "???" + "}")
	fmt.Fprint(w, "GET /translated/{"+"???"+"}")

}
