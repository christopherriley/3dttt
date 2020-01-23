package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/christopherriley/3dttt/cpu_player"
	"github.com/christopherriley/3dttt/engine"
)

func main() {
	var game engine.Game
	var humanColour, cpuColour engine.Colour
	var humanFirst bool
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Would you like to be [R]ed or [B]lue: ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \n\t")
		if strings.EqualFold(text, "r") {
			humanColour = engine.Red
			cpuColour = engine.Blue
			break
		} else if strings.EqualFold(text, "b") {
			humanColour = engine.Blue
			cpuColour = engine.Red
			break
		}
	}

	for {
		fmt.Printf("Would you like to go first? [Y/N]: ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \n\t")
		if strings.EqualFold(text, "y") {
			humanFirst = true
			break
		} else if strings.EqualFold(text, "n") {
			humanFirst = false
			break
		}
	}

	if humanFirst {
		game = engine.NewGame(humanColour)
	} else if humanColour == engine.Red {
		game = engine.NewGame(engine.Blue)
	} else {
		game = engine.NewGame(engine.Red)
	}

	for {
		game.GetBoard().Print()
		gameState := game.GetGameState()
		fmt.Printf(" Red: %d\n", gameState.RedLines)
		fmt.Printf("Blue: %d\n", gameState.BlueLines)
		if gameState.BoardState == engine.RedWins {
			fmt.Println("Red Wins!")
			return
		} else if gameState.BoardState == engine.BlueWins {
			fmt.Println("Blue Wins!")
			return
		} else if gameState.BoardState == engine.Draw {
			fmt.Println("It's a draw!")
			return
		} else if (gameState.BoardState == engine.RedToMove && humanColour == engine.Red) ||
			(gameState.BoardState == engine.BlueToMove && humanColour == engine.Blue) {
			fmt.Printf("Enter peg no move to [A-H]: ")
			text, _ := reader.ReadString('\n')
			text = strings.ToUpper(strings.Trim(text, " \n\t"))
			var err error
			var move engine.PegLabel
			if move, err = engine.StringToPeg(text); err != nil {
				fmt.Println("Please enter a peg letter to move to from 'A' to 'H'.")
				move = engine.NoPeg
			}
			if move != engine.NoPeg {
				if err := game.Move(move); err != nil {
					fmt.Println("Invalid move.")
				} else {
					game.GetBoard().Print()
				}
			}
		} else {
			move := cpu_player.GetNextMove(game.GetBoard(), cpuColour, 6)
			fmt.Printf("CPU moves to peg %s\n", engine.PegToString(move))
			game.Move(move)
		}
	}
}
