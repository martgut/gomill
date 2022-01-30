package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allCmds = [][2]string{
	{"exit", "Exit the game"},
	{"field", "Show playing field"},
	{"help", "List all available commands"},
	{"level", "Change level to new value"},
	{"play", "Find best move and play"},
	{"quit", "Quit the game"},
	{"stone", "Set stone on field"},
}

type Command struct {
}

func (cmd *Command) print() {
	fmt.Printf("\nAvailable commands:\n")
	for _, value := range allCmds {
		fmt.Printf("   %-8s: %v\n", value[0], value[1])
	}
}

func (cmd *Command) isValid(command string) bool {
	for _, value := range allCmds {
		if value[0] == command {
			return true
		}
	}
	return false
}

func (cmd *Command) prompt(label string) string {
	var command string
	reader := bufio.NewReader(os.Stdin)
	for {
		if len(label) > 0 {
			fmt.Fprintf(os.Stderr, "%v: ", label)
		} else {
			fmt.Fprint(os.Stderr, "mill> ")
		}
		command, _ = reader.ReadString('\n')
		if command != "" {
			break
		}
	}
	return strings.TrimSpace(command)
}

func (cmd *Command) promptInt(label string) (int, error) {
	input := cmd.prompt(label)
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("*** This is not a valid integer")
	}
	return value, err
}

func (cmd *Command) process(game *Game) {

	for {
		command := cmd.prompt("")
		switch command {
		case "exit":
			return
		case "quit":
			return
		case "help":
			cmd.print()
		case "field":
			game.stonesA.printPlayingField()
			game.stonesB.printPlayingField()
		case "level":
			value, err := cmd.promptInt("Enter new level")
			if err != nil {
				game.level = value
			}
		case "play":
			game.calcBestMove()
		case "":
			cmd.print()
		default:
			fmt.Printf("*** Unknown command: %v\n", command)
		}
	}
}
