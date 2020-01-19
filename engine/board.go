package engine

import "fmt"

type Color int
type PegLabel int

const (
	None Color = iota
	Red
	Blue
)

const (
	A PegLabel = iota
	B
	C
	D
	E
	F
	G
	H
)

type Peg struct {
	Slot [3]Color
}
type Board struct {
	Peg [8]Peg
}

func NewBoard() Board {
	var board Board
	for peg := A; peg <= H; peg++ {
		for slot := 0; slot < 3; slot++ {
			board.Peg[peg].Slot[slot] = None
		}
	}

	return board
}

func (c Color) String() string {
	if c == Red {
		return "R"
	} else if c == Blue {
		return "B"
	} else if c == None {
		return " "
	} else {
		return "ERROR"
	}
}

func (b Board) Print() {
	for slot := 0; slot < 3; slot++ {
		for peg := A; peg <= C; peg++ {
			if peg == B || peg == C {
				fmt.Print("   ")
			}
			fmt.Printf("[%s]", b.Peg[peg].Slot[slot])
		}
		fmt.Println()
	}

	for slot := 0; slot < 3; slot++ {
		for peg := D; peg <= E; peg++ {
			fmt.Printf("   [%s]", b.Peg[peg].Slot[slot])
		}
		fmt.Println()
	}

	for slot := 0; slot < 3; slot++ {
		for peg := F; peg <= H; peg++ {
			if peg == G || peg == H {
				fmt.Print("   ")
			}
			fmt.Printf("[%s]", b.Peg[peg].Slot[slot])
		}
		fmt.Println()
	}
}
