package engine

import "fmt"

type GameState int

const (
	NeedNextMove GameState = iota
	RedWins
	BlueWins
	Tie
)

type Game struct {
	firstPlayer Colour
	nextMove    Colour
	board       Board
}

func NewGame(firstPlayer Colour) Game {
	var game Game
	game.firstPlayer = firstPlayer
	game.board = NewBoard()

	return game
}

func (g Game) GetNextMoveColour() Colour {
	return g.nextMove
}

func (g Game) NextMove(pl PegLabel) (GameState, error) {
	if err := g.board.Peg[pl].Add(g.nextMove); err != nil {
		return NeedNextMove, fmt.Errorf("Illegal move")
	}

	if g.board.IsFull() {
		score := g.board.Evaluate()
		if score == RedWinScore {
			return RedWins, nil
		} else if score == BlueWinScore {
			return BlueWins, nil
		}

		return Tie, nil
	}

	if g.nextMove == Blue {
		g.nextMove = Red
	} else {
		g.nextMove = Blue
	}

	return NeedNextMove, nil
}
