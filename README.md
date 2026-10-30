# Pokédex CLI

An interactive command-line Pokédex (REPL) written in Go. It fetches data from the public [PokéAPI](https://pokeapi.co/) over HTTP, lets you browse the Pokémon world, and catch Pokémon into your own Pokédex.

> **Status:** work in progress. The commands below work; more features are still being added. The Pokédex is kept in memory only, so it is empty again every time you start the program.

## What works so far

- **REPL.** Reads commands in a loop, ignores case and extra spaces, and prints a clear error when a command gets the wrong number of arguments.
- **REST API client.** A small client package built on `net/http`, with a request timeout and JSON decoding into typed structs.
- **In-memory cache.** API responses are cached by URL in a map protected by a mutex. A background goroutine removes entries that have not been used for 5 minutes, so repeated commands do not hit the network again.
- **Paging.** `map` and `mapb` move forward and backward through the location areas, 20 at a time.
- **Catching.** The chance to catch a Pokémon scales linearly with its base experience, from 98% for the weakest to 2% for the strongest.

## Commands

| Command | Description |
| --- | --- |
| `help` | Show all commands |
| `map` | Show the next 20 location areas |
| `mapb` | Show the previous 20 location areas |
| `explore <location-area>` | List the Pokémon found in a location area |
| `catch <pokemon>` | Try to catch a Pokémon (the result is random) |
| `inspect <pokemon>` | Show name, height, weight, species and types of a caught Pokémon |
| `pokedex` | List the Pokémon you have caught |
| `exit` | Quit |

Example session (the result of `catch` is random):

```
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
Pokedex > pokedex
Your pokedex:
  -pikachu
```

## Project structure

```
main.go                       Sets up the config, API client and cache, starts the REPL
repl.go                       Command registry, REPL loop and input cleaning
command_*.go                  One file per command
internal/pokeapi/             PokéAPI client and response types
internal/pokecache/           Thread-safe cache with a background cleanup goroutine
repl_test.go                  Tests for input cleaning
internal/pokecache/*_test.go  Tests for the cache
```

## Usage

Requires Go (the version is set in `go.mod`) and an internet connection.

```bash
go run .
```

Or build a binary:

```bash
go build -o pokedex
./pokedex
```

## Development

```bash
go test ./...
```
