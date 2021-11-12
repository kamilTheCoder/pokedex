package main

import (
	"net/http"
	"pokedex/api/pokeapi"
	"pokedex/server"
)

func main() {
	api := pokeapi.Api{Client: &http.Client{}}
	server.HandleRequests(api)
}
