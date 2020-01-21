package cpu_player

import (
	"github.com/christopherriley/3dttt/engine"
)

const RedVictoryScore = 9999
const BlueVictoryScore = -9999
const DrawScore = 0

func evaluateBoard(b engine.Board) int {
	redLines := b.CountCompleteLines(engine.Red)
	blueLines := b.CountCompleteLines(engine.Blue)

	if b.IsFull() {
		if redLines > blueLines {
			return RedVictoryScore
		} else if blueLines > redLines {
			return BlueVictoryScore
		} else {
			return DrawScore
		}
	}

	return redLines - blueLines
}

type BoardNode struct {
	board      engine.Board
	children   []*BoardNode
	parent     *BoardNode
	boardScore int
	nodeScore  int
	lastMove   engine.PegLabel
	bestMove   engine.PegLabel
}

func NewBoardNode(b engine.Board, parent *BoardNode, move engine.PegLabel) *BoardNode {
	var bn BoardNode
	bn.board = b
	bn.parent = parent
	bn.lastMove = move
	bn.boardScore = evaluateBoard(b)

	return &bn
}

func (bn BoardNode) getDepth() int {
	if bn.parent == nil {
		return 0
	}

	return 1 + bn.parent.getDepth()
}

func (bn *BoardNode) AddChildren(c engine.Colour, maxDepth int) {
	if bn.getDepth() < maxDepth {
		for peg := engine.A; peg <= engine.H; peg++ {
			childBoard := bn.board
			if err := childBoard.Peg[peg].Add(c); err == nil {
				newChildNode := NewBoardNode(childBoard, bn, peg)
				bn.children = append(bn.children, newChildNode)
				if c == engine.Red {
					newChildNode.AddChildren(engine.Blue, maxDepth)
				} else {
					newChildNode.AddChildren(engine.Red, maxDepth)
				}
			}
		}
	}

	if len(bn.children) == 0 {
		bn.bestMove = engine.NoPeg
		bn.nodeScore = bn.boardScore
	} else {
		bn.nodeScore = bn.children[0].nodeScore
		bn.bestMove = bn.children[0].lastMove
		for i := 1; i < len(bn.children); i++ {
			if c == engine.Red {
				if bn.children[i].nodeScore > bn.nodeScore {
					bn.nodeScore = bn.children[i].nodeScore
					bn.bestMove = bn.children[i].lastMove
				}
			} else {
				if bn.children[i].nodeScore < bn.nodeScore {
					bn.nodeScore = bn.children[i].nodeScore
					bn.bestMove = bn.children[i].lastMove
				}
			}
		}
	}
}

func (bn BoardNode) GetBestMove(c engine.Colour, maxDepth int) engine.PegLabel {
	bn.AddChildren(c, maxDepth)
	return bn.bestMove
}
