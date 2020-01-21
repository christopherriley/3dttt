package engine

type BoardState int

const (
	RedToMove BoardState = iota
	BlueToMove
	RedWins
	BlueWins
	Draw
)

type GameState struct {
	RedLines   int
	BlueLines  int
	BoardState BoardState
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

func (g Game) GetGameState() GameState {
	var state GameState
	state.RedLines = g.board.CountCompleteLines(Red)
	state.BlueLines = g.board.CountCompleteLines(Blue)

	if g.board.IsFull() {
		if state.RedLines > state.BlueLines {
			state.BoardState = RedWins
		} else if state.BlueLines > state.RedLines {
			state.BoardState = BlueWins
		} else {
			state.BoardState = Draw
		}
	} else if g.nextMove == Red {
		state.BoardState = RedToMove
	} else {
		state.BoardState = BlueToMove
	}

	return state
}
