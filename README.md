# pokedex

## Usage
### Requirements
Requires at least version 1.16.3 of Golang to run. See https://golang.org/doc/install for installation details.

### Build
To build the Pokedex, run `go build ./...`

#### Mock file generation
To generate mock files install `mockgen`:

`go install github.com/golang/mock/mockgen@v1.6.0`

And then run:

`mockgen -source=api/pokeapi/pokeapi.go -destination=mocks/pokeapi/mock_pokeapi.go`

See [gomock](https://github.com/golang/mock) for more information.

### Running
To run the programme, execute `go run main.go`

## Design decisions

### Testing methodology
Most testing during development was done using Postman and requests. This was deemed more useful in the experimental stages of the API, when it wasn't really clear what are we even expecting to get from the API.

As the API evolved, Postman tests have turned into an integration test set, consisting of to get following pokemon:
- Existing, non-legendary pokemon (charmander)
- Existing, legendary pokemon (zapdos)
- Non-existing pokemon (charmander2)
- Poor pokemon without a habitat to its name :( (wormadam)

Additionally, following pokemon were translated:
- Existing, non-legendary, non-cave pokemon (charmander)
- Existing, non-legendary, cave pokemon (diglett)
- Existing, legendary, non-cave pokemon (zapdos)
- Non-existing pokemon (digletttt)

All of the tests above were repeated a few times, in order to test caching mechanism.

[Go Convey](https://github.com/smartystreets/goconvey) was used for unit tests, since it provides a variety of assertions and a `given-when-then` syntax support, making it better option than the default golang tests.

Additionally, [gomock](https://github.com/golang/mock) was used for interface mocking.

### Rate limitation
One of the APIs used, funtranslations.com, allows limited rate of usages for free: no more than 5 calls per hour or 60 calls per day. With that in mind, error handling was made to gracefully handle errors from that API, as well as a caching system was put in place, to reduce number of potential requests (see below).

### Cache
The `cache` package (`/cache`) is adopted from my earlier implementation https://github.com/kamilTheCoder/web-crawler/tree/master/crawler/cache, which in the current form is not flexible enough to re-use as a stand-alone import.

In the current form caching saves all information from APIs. The scope of this project implies it will be no more than 898 entries, as there are only that many pokemon as of Nov '21.
Less than 1000 entries represent a very small memory footprint, even including cached fun translation of descriptions. Keeping them in memory might also help overcome translation API call limitation, since repeated requests would be served by cache and not retrieved from the real, limited API.

### Error handling
For the most part, errors are not exposed to the API consumer, in order not to reveal too much about API implementation. Alternative route could have pre-defined errors which user could encounter.

## Potential improvements
### Config
Some of the settings, such as port to listen on and URLs for third-party APIs could be extracted to a config file, for example using [Viper](https://github.com/spf13/viper), or a manually built solution.

### Fun translations API key
Currently, the API only ever uses the free version of fun translations API. A potential feature could accommodate users who have API key for paid subscription, most likely through a config parameter. 

### Better errors and logging
This implementation uses simple logging via console printing and free-text error messages. It could be improved by implementing actual logging and pre-defined errors, to better reflect a real-world system.

### Add automated integration tests
Using [Godog](https://github.com/cucumber/godog) implement gherkin tests to automate postman requests from section above.