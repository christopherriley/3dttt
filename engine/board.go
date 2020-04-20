package engine

import (
	"fmt"
)

type Colour int
type PegLabel int

const (
	None Colour = iota
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
	NoPeg
)

type columnSet [3]PegLabel

// [A]   [B]   [C]
//    [D]   [E]
// [F]   [G]   [H]
var boardColumnSet = []columnSet{
	{A, B, C},
	{A, D, G},
	{F, G, H},
	{F, D, B},
	{G, E, C},
	{B, E, H},
}

type Peg struct {
	Slot [3]Colour
}
type Board struct {
	Peg [8]Peg
}

func StringToPeg(s string) (PegLabel, error) {
	switch s[0] {
	case 'A':
		return A, nil
	case 'B':
		return B, nil
	case 'C':
		return C, nil
	case 'D':
		return D, nil
	case 'E':
		return E, nil
	case 'F':
		return F, nil
	case 'G':
		return G, nil
	case 'H':
		return H, nil
	default:
		return NoPeg, fmt.Errorf("invalid peg '%s'", s)
	}
}

func PegToString(p PegLabel) string {
	switch p {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case H:
		return "H"
	default:
		return "NoPeg"
	}
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

func (p Peg) isFull() bool {
	return p.Slot[2] != None
}

func (p *Peg) Add(c Colour) error {
	if p.Slot[0] == None {
		p.Slot[0] = c
	} else if p.Slot[1] == None {
		p.Slot[1] = c
	} else if p.Slot[2] == None {
		p.Slot[2] = c
	} else {
		return fmt.Errorf("peg is full")
	}

	return nil
}

func (c Colour) String() string {
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

func (b Board) IsFull() bool {
	for _, peg := range b.Peg {
		if !peg.isFull() {
			return false
		}
	}

	return true
}

func (b Board) Print() {
	for slot := 2; slot >= 0; slot-- {
		for peg := A; peg <= C; peg++ {
			if peg == B || peg == C {
				fmt.Print("   ")
			}
			fmt.Printf("[%s]", b.Peg[peg].Slot[slot])
		}
		fmt.Println()
	}

	for slot := 2; slot >= 0; slot-- {
		for peg := D; peg <= E; peg++ {
			fmt.Printf("   [%s]", b.Peg[peg].Slot[slot])
		}
		fmt.Println()
	}

	for slot := 2; slot >= 0; slot-- {
		for peg := F; peg <= H; peg++ {
			if peg == G || peg == H {
				fmt.Print("   ")
			}
			fmt.Printf("[%s]", b.Peg[peg].Slot[slot])
		}
		fmt.Println()
	}
}

func (p Peg) completed() Colour {
	if (p.Slot[0]) == p.Slot[1] && p.Slot[1] == p.Slot[2] {
		return p.Slot[0]
	}

	return None
}

// count vertical lines on one peg each
//
// [R] [ ] [ ]
// [R] [ ] [ ]
// [R] [ ] [ ]
func (b Board) countCompleteVerticalLines(c Colour) int {
	completeLines := 0

	for _, peg := range b.Peg {
		if peg.completed() == c {
			completeLines++
		}
	}

	return completeLines
}

// count horizontal lines across three pegs
//
// [ ] [ ] [ ]
// [R] [R] [R]
// [ ] [ ] [ ]
func (b Board) countCompleteHorizontalLines(c Colour, cs columnSet) int {
	completeLines := 0

	for slot := 0; slot < 3; slot++ {
		colour := 0
		for peg := 0; peg < 3; peg++ {
			if b.Peg[cs[peg]].Slot[slot] == c {
				colour++
			}
		}
		if colour == 3 {
			completeLines++
		}
	}

	return completeLines
}

// count diagonal lines across three pegs
//
// [R] [ ] [ ]
// [ ] [R] [ ]
// [ ] [ ] [R]
func (b Board) countCompleteDiagonalLines(c Colour, cs columnSet) int {
	completeLines := 0

	if b.Peg[cs[0]].Slot[0] == c && b.Peg[cs[1]].Slot[1] == c && b.Peg[cs[2]].Slot[2] == c {
		completeLines++
	}

	if b.Peg[cs[0]].Slot[2] == c && b.Peg[cs[1]].Slot[1] == c && b.Peg[cs[2]].Slot[0] == c {
		completeLines++
	}

	return completeLines
}

func (b Board) CountCompleteLines(c Colour) int {
	completeLines := 0

	for _, bcs := range boardColumnSet {
		completeLines += b.countCompleteDiagonalLines(c, bcs)
		completeLines += b.countCompleteHorizontalLines(c, bcs)
	}

	completeLines += b.countCompleteVerticalLines(c)

	return completeLines
}
