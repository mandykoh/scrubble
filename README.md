# scrubble

[![GoDoc](https://godoc.org/github.com/mandykoh/scrubble?status.svg)](https://godoc.org/github.com/mandykoh/scrubble)
[![Go Report Card](https://goreportcard.com/badge/github.com/mandykoh/scrubble)](https://goreportcard.com/report/github.com/mandykoh/scrubble)
[![Build Status](https://travis-ci.org/mandykoh/scrubble.svg?branch=master)](https://travis-ci.org/mandykoh/scrubble)

`scrubble` is a Go library for modelling a letter/word board game.

Key features include:

  * Complete and flexible game logic implementation, including support for tournament-style player initiated challenges or automatic word validation.
  * Solitaire (single player) support in case you want to practice alone.
  * Replayable games—Turn history is automatically tracked, and all random number generation is externalised in case you want to use deterministic random sequences.
  * Customisable board layouts—Want weirdly long rectangular boards with multiple start positions and score multipliers everywhere? You got it.
  * Customisable tile sets—Why limit yourself to the English alphabet? You know you always wanted to play with emoji phrases.
  * Customisable tile bags—Just specify the distribution of tiles you’d like.
  * Customisable dictionaries (the public domain ENABLE word list is included by default).
  * Extensible rules—The rules for how tile placement is validated, how words are formed and scored, when word validation is done, when the game ends, etc can be extended or completely replaced.

See the [API documentation](https://godoc.org/github.com/mandykoh/scrubble) for more information.

This software is made available under an [MIT license](LICENSE).


## Getting started

This library uses [dep](https://github.com/golang/dep) for dependency management. After cloning/downloading, first update the dependencies:

```
$ dep ensure
```


## Demo

As a demonstration of how it might be used, this library includes `textscrubble`, a simple terminal-based game for one or more players that can be played on one computer.

![Screenshot of textscrubble game in progress](screenshots/textscrubble.png)

From the project location, `textscrubble` can be run as follows:

```
$ go run cmd/textscrubble.go [mode] [player1_name] ... [playerN_name]
```

`mode` can either be `simple` (where words are automatically validated and only valid words may be played) or `challenge` (where any words can be played but players may challenge a play to have it validated, at the risk of a penalty).


## Running tests

Tests can be run like any Go project:

```
$ go test
```


## Example usage

### Setting up a game

A Game first needs to be created and started to begin play. At minimum, a [Bag](https://godoc.org/github.com/mandykoh/scrubble#Bag) of tiles and a [Board](https://godoc.org/github.com/mandykoh/scrubble#Board) must be provided:

```go
bag := scrubble.BagWithStandardEnglishTiles()
board := scrubble.BoardWithStandardLayout()
game := scrubble.NewGame(bag, board)
```

The above can also be written more concisely as:

```go
game := scrubble.NewGameWithDefaults()
```

Once a game is created, it is in the Setup phase, and players can be added. `scrubble` doesn’t model players directly, preferring to let different usages model them as appropriate for the usage. Instead, we have [Seats](https://godoc.org/github.com/mandykoh/scrubble#Seat), which represent a player’s presence at a game (and all things relevant to that, such as the score and the rack of tiles). We can add seats as follows:

```go
seat, err := game.AddPlayer()
```

When some players have been added to the game, the game can be started, which begins the Main game phase:

```go
var rng *rand.Rand
...
err := game.Start(rng)
```


### Custom boards

Apart from using the standard board layout provided by [`BoardWithStandardLayout`](https://godoc.org/github.com/mandykoh/scrubble#BoardWithStandardLayout), custom board layouts are supported. A custom layout can be easily created as follows:

```go
__, st, dl, dw, tl, tw := scrubble.BoardPositionTypes()

board := scrubble.BoardWithLayout(scrubble.BoardLayout{
    {tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
    {__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
    {__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
    {dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
    {__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
    {__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
    {__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
    {tw, __, __, dl, __, __, __, st, __, __, __, dl, __, __, tw},
    {__, __, dl, __, __, __, dl, __, dl, __, __, __, dl, __, __},
    {__, tl, __, __, __, tl, __, __, __, tl, __, __, __, tl, __},
    {__, __, __, __, dw, __, __, __, __, __, dw, __, __, __, __},
    {dl, __, __, dw, __, __, __, dl, __, __, __, dw, __, __, dl},
    {__, __, dw, __, __, __, dl, __, dl, __, __, __, dw, __, __},
    {__, dw, __, __, __, tl, __, __, __, tl, __, __, __, dw, __},
    {tw, __, __, dl, __, __, __, tw, __, __, __, dl, __, __, tw},
})
```

with `__`, `st`, `dl`, `dw`, `tl`, and `tw` representing positions where regular, starting, double-letter score bonuses, double-word score bonuses, triple-letter score bonuses, and triple-word score bonuses should appear, respectively.


### Custom tile bags

Bags can also be created with tiles different to those provided by [`BagWithStandardEnglishTiles`](https://godoc.org/github.com/mandykoh/scrubble#BagWithStandardEnglishTiles) by specifying what tiles and how many of each tile a bag should contain:

```go
// Creates a Bag with 9 x A tiles, 2 x B tiles, 2 x C tiles, 4 x D tiles, and
// 1 x 😃 tile (that happens to be worth 50 points)
bag := scrubble.BagWithDistribution(scrubble.TileDistribution{
    {Tile{'A', 1}, 9},
    {Tile{'B', 3}, 2},
    {Tile{'C', 3}, 2},
    {Tile{'D', 2}, 4},
    {Tile{'😃', 50}, 1},
})
```
