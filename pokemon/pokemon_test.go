package pokemon

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToPokemon(t *testing.T) {
	Convey("Given raw pokemon with English flavour text", t, func() {
		raw := PokemonRaw{
			Name: "testmon",
			FalvourTexts: []FlavourText{
				{
					FalvourText: "ñññ",
					Language: struct {
						Name string "json:\"name\""
					}{"ñ"},
				},
				{
					FalvourText: "english",
					Language: struct {
						Name string "json:\"name\""
					}{"en"},
				},
				{
					FalvourText: "?",
					Language: struct {
						Name string "json:\"name\""
					}{"??"},
				},
			},
			Habitat:   Habitat{Name: "under your bed"},
			Legendary: true,
		}

		Convey("When I convert it to a normal pokemon", func() {
			p := raw.ToPokemon()

			Convey("I expect relevant information to be there", func() {
				So(p.Description, ShouldEqual, raw.FalvourTexts[1].FalvourText)
				So(p.Habitat, ShouldEqual, raw.Habitat.Name)
				So(p.Legendary, ShouldEqual, raw.Legendary)
				So(p.Name, ShouldEqual, raw.Name)
				So(p.VariantDescriptions, ShouldNotBeNil)
				So(p.VariantDescriptions, ShouldBeEmpty)
			})
		})
	})

	Convey("Given raw pokemon with no English flavour text", t, func() {
		raw := PokemonRaw{
			Name: "testmon2",
			FalvourTexts: []FlavourText{
				{
					FalvourText: "ñññ",
					Language: struct {
						Name string "json:\"name\""
					}{"ñ"},
				},
				{
					FalvourText: "?",
					Language: struct {
						Name string "json:\"name\""
					}{"??"},
				},
			},
			Habitat:   Habitat{Name: "under your bed"},
			Legendary: true,
		}

		Convey("When I convert it to a normal pokemon", func() {
			p := raw.ToPokemon()

			Convey("I expect relevant information to be there", func() {
				So(p.Description, ShouldEqual, "")
				So(p.Habitat, ShouldEqual, raw.Habitat.Name)
				So(p.Legendary, ShouldEqual, raw.Legendary)
				So(p.Name, ShouldEqual, raw.Name)
				So(p.VariantDescriptions, ShouldNotBeNil)
				So(p.VariantDescriptions, ShouldBeEmpty)
			})
		})
	})
}
