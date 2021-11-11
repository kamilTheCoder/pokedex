package cache

import (
	"pokedex/pokemon"
	"sync"
)

var once sync.Once
var c *cache

type cache struct {
	pokeCache   map[string]pokemon.Pokemon
	mutex       *sync.Mutex
	rejcetMutex *sync.Mutex
}

// Access returns an instance of cache. Only one instance per program execution will be created
func Access() *cache {
	once.Do(func() {
		c = &cache{}
		c.Wipe()
	})

	return c
}

// Get tries to return pokemon from the cache
// If it exists, returns the pokemon and true, otherwise empty pokestruct and false
func (c *cache) Get(name string) (pokemon.Pokemon, bool) {
	c.mutex.Lock()
	p, ok := c.pokeCache[name]
	c.mutex.Unlock()
	return p, ok
}

// AddNew tries a new pokemon to the cache and returns whether it was possible (true) or not (false)
// False returned from the function signifies that the pokemon already exists in the cache
func (c *cache) AddNew(pokemon pokemon.Pokemon) bool {
	c.mutex.Lock()
	if _, ok := c.pokeCache[pokemon.Name]; ok {
		c.mutex.Unlock()
		return false
	}

	c.pokeCache[pokemon.Name] = pokemon
	c.mutex.Unlock()

	return true
}

// Wipe resets the cache
func (c *cache) Wipe() {
	c.pokeCache = map[string]pokemon.Pokemon{}
	c.mutex = &sync.Mutex{}
	c.rejcetMutex = &sync.Mutex{}
}

// Update allows amending information on an existing pokemon in the cache. Adds a new one, if there wasn't one already there
func (c *cache) Update(pokemon pokemon.Pokemon) {
	c.mutex.Lock()
	c.pokeCache[pokemon.Name] = pokemon
	c.mutex.Unlock()
}

// Size returns total number of entries stored in the cache
func (c *cache) Size() int {
	return len(c.pokeCache)
}
