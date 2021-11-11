package pokeapi

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	mock_pokeapi "pokedex/mocks/pokeapi"
	"pokedex/pokemon"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPokemon(t *testing.T) {
	Convey("Given mock http client", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClient := mock_pokeapi.NewMockHttpClient(ctrl)
		api := Api{Client: mockClient}

		Convey("When a request returns an error", func() {
			mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("do error"))
			_, err := api.GetPokemon("testmon")

			Convey("I expect the API to return same error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "error fetching testmon: do error")
			})
		})

		Convey("When a pokemon was not found", func() {
			// response from pokeapi.co when a pokemon is not found is simply, an unformatted:
			// Not Found

			rsp := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(pokemonSpeciesError)))}
			mockClient.EXPECT().Do(gomock.Any()).Return(rsp, nil)
			_, err := api.GetPokemon("testmon")

			Convey("I expect the API to return error not found", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "error fetching testmon: pokemon not found")
			})
		})

		Convey("When a pokemon was successfully found", func() {
			rsp := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(pokemonSpeciesJSON)))}
			mockClient.EXPECT().Do(gomock.Any()).Return(rsp, nil)
			p, err := api.GetPokemon("testmon")

			Convey("I expect the API to return the pokemon and no errors", func() {
				So(err, ShouldBeNil)
				So(p, ShouldResemble, pokemon.PokemonRaw{
					Name: "wormadam",
					FalvourTexts: []pokemon.FlavourText{
						{
							FalvourText: "When the bulb on\nits back grows\nlarge, it appears\fto lose the\nability to stand\non its hind legs.",
							Language: struct {
								Name string "json:\"name\""
							}{"en"},
						},
					},
					Habitat:   pokemon.Habitat{Name: ""},
					Legendary: false,
				})

			})
		})
	})
}

func TestGetTranslation(t *testing.T) {
	Convey("Given mock http client", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClient := mock_pokeapi.NewMockHttpClient(ctrl)
		api := Api{Client: mockClient}

		Convey("When a request returns an error", func() {
			mockClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("do error"))
			_, err := api.GetTranslation("testTranslator", "testmon")

			Convey("I expect the API to return same error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "error fetching testTranslator translation: do error")
			})
		})

		Convey("When a translation failed", func() {
			// see translationResponseError below for expected error syntax
			rsp := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(translationResponseError)))}
			mockClient.EXPECT().Do(gomock.Any()).Return(rsp, nil)
			_, err := api.GetTranslation("translator2", "test")

			Convey("I expect the API to return error not found", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "error calling translation API: Too Many Requests")
			})
		})

		Convey("When a translation was successful", func() {
			rsp := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(translationResponse)))}
			mockClient.EXPECT().Do(gomock.Any()).Return(rsp, nil)
			t, err := api.GetTranslation("yoda", "starwars")

			Convey("I expect the API to return the pokemon and no errors", func() {
				So(err, ShouldBeNil)
				So(t, ShouldEqual, "Lost a planet,  master obiwan has.")

			})
		})
	})
}

var (
	translationResponseError = `{
	"error": {
		"code": 429,
		"message": "Too Many Requests: Rate limit of 5 requests per hour exceeded. Please wait for 57 minutes and 23 seconds."
	}
}`
	translationResponse = `{
    "success": {
        "total": 1
    },
    "contents": {
        "translated": "Lost a planet,  master obiwan has.",
        "text": "Master Obiwan has lost a planet.",
        "translation": "yoda"
    }
}`
	pokemonSpeciesError = "Not Found"
	pokemonSpeciesJSON  = `{
	"id": 413,
	"name": "wormadam",
	"order": 441,
	"is_baby": false,
	"is_legendary": false,
	"habitat": null,
	"names": [
	  {
		"name": "Wormadam",
		"language": {
		  "name": "en",
		  "url": "https://pokeapi.co/api/v2/language/9/"
		}
	  }
	],
	"flavor_text_entries": [
	  {
		"flavor_text": "When the bulb on\nits back grows\nlarge, it appears\fto lose the\nability to stand\non its hind legs.",
		"language": {
		  "name": "en",
		  "url": "https://pokeapi.co/api/v2/language/9/"
		},
		"version": {
		  "name": "red",
		  "url": "https://pokeapi.co/api/v2/version/1/"
		}
	  }
	]
  }
  `
)
