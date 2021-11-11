package cache

import (
	"pokedex/pokemon"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTotal(t *testing.T) {
	Convey("Given empty cache", t, func() {
		Access().Wipe()
		So(Access().Size(), ShouldEqual, 0)

		Convey("When I add a new pokemon", func() {
			Access().AddNew(pokemon.Pokemon{Name: "testmon"})

			Convey("I expect one pokemon to be in the cache", func() {
				So(Access().Size(), ShouldEqual, 1)

				Convey("When I add three more pokemon", func() {
					Access().AddNew(pokemon.Pokemon{Name: "testmon2"})
					Access().AddNew(pokemon.Pokemon{Name: "testmon3"})
					Access().AddNew(pokemon.Pokemon{Name: "testmon4"})

					Convey("I expect four pokemon to be in the cache", func() {
						So(Access().Size(), ShouldEqual, 4)
					})
				})
			})
		})
	})
}

func TestWipe(t *testing.T) {
	Convey("Given a cache with a pokemon", t, func() {
		p := pokemon.Pokemon{Name: "testmon"}
		Access().AddNew(p)

		Convey("When I wipe it", func() {
			Access().Wipe()

			Convey("I expect nothing to be there", func() {
				_, exists := Access().Get("testmon")
				So(exists, ShouldBeFalse)
				So(Access().Size(), ShouldEqual, 0)
			})
		})
	})
}

func TestAddNew(t *testing.T) {
	Convey("Given empty cache", t, func() {
		Access().Wipe()
		So(Access().Size(), ShouldEqual, 0)

		Convey("When I add a new pokemon", func() {
			p := pokemon.Pokemon{Name: "testmon"}
			added := Access().AddNew(p)

			Convey("I expect one pokemon to be in the cache", func() {
				So(added, ShouldBeTrue)
				So(Access().Size(), ShouldEqual, 1)

				Convey("When I add it again", func() {
					added = Access().AddNew(p)
					Convey("I expect only one of them to be in the cache", func() {
						So(added, ShouldBeFalse)
						So(Access().Size(), ShouldEqual, 1)

						Convey("When I add another pokemon", func() {
							added = Access().AddNew(pokemon.Pokemon{Name: "testmon2"})
							Convey("I expect both of them to be in the cache", func() {
								So(added, ShouldBeTrue)
								So(Access().Size(), ShouldEqual, 2)
							})
						})
					})
				})
			})
		})
	})
}

func TestUpdate(t *testing.T) {
	Convey("Given a cache with no pokemon", t, func() {
		Access().Wipe()
		So(Access().Size(), ShouldEqual, 0)


		Convey("When I update a non-existing pokemon", func() {
			p := pokemon.Pokemon{Name: "testmon"}
			Access().Update(p)

			Convey("I expect one pokemon to be in the cache with no description", func() {
				So(Access().Size(), ShouldEqual, 1)
				cached, exists := Access().Get("testmon")
				So(exists, ShouldBeTrue)
				So(cached.Description, ShouldEqual, "")

				Convey("When I update it with a description", func() {
					desc := "description"
					p1 := pokemon.Pokemon{Name: "testmon", Description: desc}
					Access().Update(p1)

					Convey("I expect only one pokemon in the cache with updated description", func() {
						So(Access().Size(), ShouldEqual, 1)
						cached, exists := Access().Get("testmon")
						So(exists, ShouldBeTrue)
						So(cached.Description, ShouldEqual, desc)
					})
				})
			})
		})
	})
}
