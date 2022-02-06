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
	{"changeA", "Set stone for A on field"},
	{"stone", "Set stone for B on field"},
	{"write", "Save game to disk"},
	{"read", "Read game from disk"},
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

	var commands []string
	reader := bufio.NewReader(os.Stdin)
	for {
		if len(label) > 0 {
			fmt.Fprintf(os.Stderr, "%v: ", label)
		} else {
			fmt.Fprint(os.Stderr, "mill> ")
		}
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		idx := 0
		if input != "" {
			for _, cmd := range allCmds {
				if strings.HasPrefix(cmd[0], input) {
					fmt.Fprintf(os.Stderr, "mill> %v\n", cmd[0])
					idx += 1
					commands = append(commands, cmd[0])
				}
			}
		}
		if idx == 1 {
			break
		}
		fmt.Fprintf(os.Stderr, "mill> ? %v\n", commands)
	}
	return strings.TrimSpace(commands[0])
}

func (cmd *Command) promptInt(label string) (int, error) {
	fmt.Fprintf(os.Stderr, "%v: ", label)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
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
			game.print()
		case "changeA":
			value, err := cmd.promptInt("Enter field")
			if err == nil {
				game.stonesA = append(game.stonesA, value)
			}
		case "stone":
			value, err := cmd.promptInt("Enter field")
			if err == nil {
				game.stonesB = append(game.stonesB, value)
			}
			game.print()
		case "write":
			fileName := "game.json"
			fmt.Printf("Saving game to file: %v\n", fileName)
			game.writeToFile(fileName)
		case "read":
			fileName := "game.json"
			fmt.Printf("Reading game from file: %v\n", fileName)
			game.readFromFile(fileName)
		case "":
			cmd.print()
		default:
			fmt.Printf("*** Unknown command: %v\n", command)
		}
		fmt.Println()
	}
}
