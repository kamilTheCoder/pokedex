package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"pokedex/cache"
	mock_api "pokedex/mocks/api"
	"pokedex/pokemon"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandleRequest(t *testing.T) {
	cache.Access().Wipe()
	Convey("Given mock api and server", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockApi := mock_api.NewMockApi(ctrl)
		pokeApi = mockApi

		Convey("When I request data on a pokemon", func() {
			req, err := http.NewRequest(http.MethodGet, "/testmon", nil)
			So(err, ShouldBeNil)

			rec := httptest.NewRecorder()

			router := mux.NewRouter().StrictSlash(true)
			router.HandleFunc("/{name}", handlePokemon)

			Convey("And the API returns an error", func() {
				req, err := http.NewRequest(http.MethodGet, "/testmon", nil)
				So(err, ShouldBeNil)

				mockApi.EXPECT().GetPokemon(gomock.Any()).Return(pokemon.PokemonRaw{}, errors.New("error"))

				Convey("I expect the API to be hit and return response", func() {
					router.ServeHTTP(rec, req)
					So(rec.Code, ShouldEqual, http.StatusNotFound)
					So(rec.Body.String(), ShouldContainSubstring, "error")
				})
			})

			Convey("I expect the API to be hit and return response", func() {
				mockApi.EXPECT().GetPokemon(gomock.Any()).Return(pokemon.PokemonRaw{Name: "testmon"}, nil)

				router.ServeHTTP(rec, req)
				So(rec.Code, ShouldEqual, http.StatusOK)
				So(rec.Body.String(), ShouldContainSubstring, "testmon")
			})

			Convey("When I request it again and it is in the cache already", func() {
				router.ServeHTTP(rec, req)
				Convey("I expect it not to hit the API, but get the data from the cache", func() {
					So(rec.Code, ShouldEqual, http.StatusOK)
					So(rec.Body.String(), ShouldContainSubstring, "testmon")
				})
			})
		})

	})
}

func TestHandleTranslated(t *testing.T) {
	cache.Access().Wipe()
	Convey("Given mock api and server", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockApi := mock_api.NewMockApi(ctrl)
		pokeApi = mockApi

		Convey("When I request data on a pokemon", func() {
			req, err := http.NewRequest(http.MethodGet, "/translated/testmon", nil)
			So(err, ShouldBeNil)

			rec := httptest.NewRecorder()

			router := mux.NewRouter().StrictSlash(true)
			router.HandleFunc("/translated/{name}", handleTranslated)

			Convey("And the pokeAPI returns an error", func() {
				req, err := http.NewRequest(http.MethodGet, "/translated/testmonErr", nil)
				So(err, ShouldBeNil)

				mockApi.EXPECT().GetPokemon("testmonErr").Return(pokemon.PokemonRaw{}, errors.New("error"))

				Convey("I expect the API to be hit and return response", func() {
					router.ServeHTTP(rec, req)
					So(rec.Code, ShouldEqual, http.StatusNotFound)
					So(rec.Body.String(), ShouldContainSubstring, "error")
				})
			})

			Convey("And the pokeAPI to be hit and return response", func() {
				mockApi.EXPECT().GetPokemon("testmon").Return(
					pokemon.PokemonRaw{
						Name: "testmon",
						FalvourTexts: []pokemon.FlavourText{{
							FalvourText: "untranslated",
							Language: struct {
								Name string "json:\"name\""
							}{"en"},
						}},
						Habitat:   pokemon.Habitat{},
						Legendary: true,
					}, nil).AnyTimes()

				Convey("And the translation API returns an errror", func() {
					mockApi.EXPECT().GetTranslation(gomock.Any(), gomock.Any()).Return("", errors.New("error"))

					Convey("I expect to receive untranslated description", func() {
						router.ServeHTTP(rec, req)
						So(rec.Code, ShouldEqual, http.StatusOK)
						So(rec.Body.String(), ShouldContainSubstring, "untranslated")
					})
				})

				Convey("And the translation API returns a valid answer", func() {
					mockApi.EXPECT().GetTranslation(gomock.Any(), gomock.Any()).Return("translated", nil)

					Convey("I expect to receive a response", func() {
						router.ServeHTTP(rec, req)
						So(rec.Code, ShouldEqual, http.StatusOK)
						So(rec.Body.String(), ShouldContainSubstring, "translated")
						So(rec.Body.String(), ShouldNotContainSubstring, "untranslated")

						Convey("And I expect no more calls to API when I call it again", func() {
							router.ServeHTTP(rec, req)
							So(rec.Code, ShouldEqual, http.StatusOK)
							So(rec.Body.String(), ShouldContainSubstring, "translated")
							So(rec.Body.String(), ShouldNotContainSubstring, "untranslated")
						})
					})
				})
			})

			Convey("When I request it again and it is in the cache already", func() {
				router.ServeHTTP(rec, req)
				Convey("I expect it not to hit the API, but get the data from the cache", func() {
					So(rec.Code, ShouldEqual, http.StatusOK)
					So(rec.Body.String(), ShouldContainSubstring, "testmon")
				})
			})
		})

	})
}
