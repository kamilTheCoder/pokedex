# pokedex

## Usage
### Requirements
Requires at least version 1.16.3 of Golang to run. See https://golang.org/doc/install for installation details.

### Build
To build the Pokedex, run `go build ./...`

### Running
To run the programme, execute `go run main.go`


## Testing
Most testing during development was done using Postman and requests. This was deemed more useful in the experimental stages of the API, when it wasn't really clear what are we even expecting to get from the API.

[Go Convey](https://github.com/smartystreets/goconvey) was used for unit tests, since it provides a variety of assertions and a `given-when-then` syntax support, making it better option than the default golang tests.

## Design decisions

### Rate limitation
One of the APIs used, funtranslations.com, allows limited rate of usages for free: no more than 5 calls per hour or 60 calls per day. 

### Cache
The `cache` package (`/cache`) is adopted from my earlier implementation https://github.com/kamilTheCoder/web-crawler/tree/master/crawler/cache, which in the current form is not flexible enough to re-use as a stand-alone import.

In the current form caching saves all information from APIs. The scope of this project implies it will be no more than 898 entries, as there are only that many pokemon as of of Nov '21.
Less than 1000 entries represent a very small memory footprint, even including cached fun translation of descriptions. Keeping them in memory might also help overcome translation API call limitation, since repeated requests would be served by cache and not retreived from the real, limited API.

### Error handling
For the most part, errors are not exposed to the API consumer, in order not to reveal too much about API implementation. Alternative route could have pre-defined errors which user could encounter.

## Potential improvements
### Config
Some of the settings, such as port to listen on and URLs for thrid-party APIs could be extracted to a config file, for example using [Viper](https://github.com/spf13/viper), or a manually built solution.

### Fun translations API key
Currently, the API only ever uses the free version of fun translations API. A potential feature could accommodate users who have API key for paid subscription, most likely through a config parameter. 

### Better errors and logging
This implementation uses simple logging via console printing and free-text error messages. It could be improved by implementing actual logging and pre-defined errors, to better reflect a real-world system.