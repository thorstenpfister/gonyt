[![GitHub license](https://img.shields.io/badge/license-BSD%203--clause-brightgreen)](https://github.com/thorstenpfister/gonyt/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/release/thorstenpfister/gonyt)](https://github.com/thorstenpfister/gonyt/releases)
[![Go Report card](https://goreportcard.com/badge/github.com/thorstenpfister/gonyt?style=plastic)](https://goreportcard.com/report/github.com/thorstenpfister/gonyt)

# Disclaimer

`gonyt` is not endorsed by, directly affiliated with, maintained, authorized, or sponsored by the New York Times. All product and company names are the registered trademarks of their original owners. The use of any trade name or trademark is for identification and reference purposes only and does not imply any association with the trademark holder of their product brand.

# gonyt

`gonyt` is a lightweight wrapper and opinionated CLI tool for the New York Times API. 

**Library.** It's focus is on providing a library to access various features of the New York Times API in a convenient manner. 

**CLI.** In addition there's a CLI with opinionated formatting to quickly query the API for content. Besides there's also the option to gain access to the full JSON representation for inclusion in various scripts.

The CLI tool also acts as an example for the usage of the library.

# Features

The following APIs are currently implemented in both the library and CLI. Other available APIs are listed in alphabetical order and may or may not be implemented at a later date.

 - [x] Top stories
 - [ ] Archive
 - [ ] Article search
 - [ ] Books
 - [ ] Community
 - [ ] Most popular
 - [ ] Movie reviews
 - [ ] Semantic 
 - [ ] Times tags
 - [ ] Times wire

# Getting started

For any usage you will need an API key from the New York Times which you can register for via their [developer portal](https://developer.nytimes.com/).

Feel free to integrate the library via 
```bash
go get -u github.com/thorstenpfister/gonyt/nytapi
```
then include in your application via
```Go
import "github.com/thorstenpfister/gonyt/nytapi"
```

To build the CLI application, clone the repo and feel free to 
```bash
go build
```
and explore. You can provide the API key directly to the CLI tool on every command or persist it in a `.gonyt` file in your user folder like this
```bash
APIKEY=your_key
```


# Motivation

Play around with Go, provide a solid & thought out library, get a fun CLI tool in the process.

# Contributing

Feel free to fork and pose a PR. Please note the present structure loosely aligned to CQRS principles and that the library aims for a test coverage of 80%+.

# License

`gonyt` is released under the 3-clause BSD License. See [LICENSE](https://github.com/thorstenpfister/gonyt/blob/main/LICENSE).
