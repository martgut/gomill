package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allCmds = [][2]string{
	{"apply", "Apply best move"},
	{"exit", "Exit the game"},
	{"print", "Print playing field"},
	{"help", "List all available commands"},
	{"level", "Change level to new value"},
	{"play", "Find best move and play"},
	{"quit", "Quit the game"},
	{"stoneA", "Set stone for A on field"},
	{"stoneB", "Set stone for B on field"},
}

type Command struct {
}

func (cmd *Command) print() {
	fmt.Printf("Available commands:\n")
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
			fmt.Fprintf(os.Stderr, "")
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

	fmt.Printf("\nWelcome to Mill\n\n")
	for {
		command := cmd.prompt("")
		switch command {
		case "exit":
			return
		case "quit":
			return
		case "help":
			cmd.print()
		case "print":
			game.print()
		case "level":
			value, err := cmd.promptInt("Enter new level")
			if err == nil {
				game.level = value
			}
		case "play":
			game.calcBestMove()
		case "apply":
			move := game.mo.perfectMove[0][0]
			game.stonesA = game.stonesA.applyMove(move)
		case "stoneA":
			value, err := cmd.promptInt("Enter field")
			if err == nil {
				game.stonesA = append(game.stonesA, value)
			}
		case "stoneB":
			value, err := cmd.promptInt("Enter field")
			if err == nil {
				game.stonesB = append(game.stonesB, value)
			}
		case "":
			cmd.print()
		default:
			fmt.Printf("*** Unknown command: %v\n", command)
		}
		fmt.Println()
	}
}
