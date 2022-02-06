package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Game struct {
	// Stones of player A
	stonesA Fields

	// Stones of player B
	stonesB Fields

	// Think level
	level int

	// Number of stones to place
	freeStones int

	// Optimizer to calculate the best move
	mo MoveOptimizer
}

// Constructor to create new game
func newGame() *Game {
	game := new(Game)
	game.mo.rater = MultiplexRater{raters: []RateField{CountStonesRater{100}, BestCrossingRaster{}}}
	game.level = 4
	game.freeStones = 18
	return game
}

// Calculate next best move
func (game *Game) calcBestMove() {
	game.mo.calcBestMove(game.stonesA, game.stonesB, game.freeStones, game.level)
	game.mo.printPerfectMove()
}

// Pretty print player field
func (game *Game) print() {
	pf := playFieldT{stonesA: game.stonesA, stonesB: game.stonesB, printLarge: true}
	pf.printField()
}

// Save game to disk as JSON
func (game *Game) writeToFile(fileName string) {

	data := map[string]Fields{
		"stonesA": game.stonesA,
		"stonesB": game.stonesB,
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)

}

// Save game to disk as JSON
func (game *Game) readFromFile(fileName string) {

	data := map[string]Fields{
		"stonesA": game.stonesA,
		"stonesB": game.stonesB,
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	game.stonesA = data["stonesA"]
	game.stonesB = data["stonesB"]
}
