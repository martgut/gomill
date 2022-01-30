package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	DebugLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	DebugLogger = log.New(io.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	fmt.Println("Welcome to Mill!")
	allMoves.print()
	allMills.print()

	// Evaluate with one stone
	mo := MoveOptimizer{rater: HighestFieldsRater{}}
	mo.calcBestMove(Fields{0, 3}, Fields{}, 0, 1)
	move := mo.perfectMove[0][0]
	fmt.Println(move)
	fmt.Printf("moves: %d\n", mo.moveCounter)

	stones := Fields{1, 2, 3}
	stones.printPlayingField()

}
