package cpu_player

import (
	"github.com/christopherriley/3dttt/engine"
)

func GetNextMove(b engine.Board, c engine.Colour, maxDepth int) engine.PegLabel {
	bn := NewBoardNode(b)
	return bn.GetBestMove(c, maxDepth)
}
