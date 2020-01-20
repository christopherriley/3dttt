package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/christopherriley/3dttt/engine"
)

func main() {
	var game engine.Game
	var humanColour engine.Colour
	var humanFirst bool
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Would you like to be [R]ed or [B]lue: ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \n\t")
		if strings.EqualFold(text, "r") {
			humanColour = engine.Red
			break
		} else if strings.EqualFold(text, "b") {
			humanColour = engine.Blue
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

	fmt.Println(game)
}
