package engine

import "fmt"

type NextMove int

const (
	RedToMove NextMove = iota
	BlueToMove
	RedWins
	BlueWins
	Draw
)

func NextMoveToString(b NextMove) (string, error) {
	switch b {
	case RedToMove:
		return "RedToMove", nil
	case BlueToMove:
		return "BlueToMove", nil
	case RedWins:
		return "RedWins", nil
	case BlueWins:
		return "BlueWins", nil
	case Draw:
		return "Draw", nil
	default:
		return "", fmt.Errorf("unknown NextMove '%d'", b)
	}
}

type GameState struct {
	RedLines  int
	BlueLines int
	NextMove  NextMove
}

type Game struct {
	nextMove Colour
	board    Board
}

func NewGame(firstPlayer Colour) Game {
	var game Game
	game.nextMove = firstPlayer
	game.board = NewBoard()

	return game
}

func (g Game) GetBoard() Board {
	return g.board
}

func (g Game) GetGameState() GameState {
	var state GameState
	state.RedLines = g.board.CountCompleteLines(Red)
	state.BlueLines = g.board.CountCompleteLines(Blue)

	if g.board.IsFull() {
		if state.RedLines > state.BlueLines {
			state.NextMove = RedWins
		} else if state.BlueLines > state.RedLines {
			state.NextMove = BlueWins
		} else {
			state.NextMove = Draw
		}
	} else if g.nextMove == Red {
		state.NextMove = RedToMove
	} else {
		state.NextMove = BlueToMove
	}

	return state
}

func (g *Game) Move(p PegLabel) error {
	if err := g.board.Peg[p].Add(g.nextMove); err != nil {
		return err
	}

	if g.nextMove == Red {
		g.nextMove = Blue
	} else {
		g.nextMove = Red
	}

	return nil
}
