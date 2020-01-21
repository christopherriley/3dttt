package testsamples

import "github.com/christopherriley/3dttt/engine"

func CreateOneHorizontalRedRow() engine.Board {
	var board engine.Board

	board.Peg[engine.A].Add(engine.Red)
	board.Peg[engine.B].Add(engine.Red)
	board.Peg[engine.C].Add(engine.Red)

	return board
}

func CreateOneVerticalBlueRow() engine.Board {
	var board engine.Board

	board.Peg[engine.A].Add(engine.Blue)
	board.Peg[engine.A].Add(engine.Blue)
	board.Peg[engine.A].Add(engine.Blue)

	return board
}

// RED LINE 1: vertical on peg A
// RED LINE 2: diagonal on peg A, B, C
//
// BLUE LINE 1: horizontal on peg F, G, H
// BLUE LINE 2: diagonal on peg F, D, B
// BLUE LINE 2: horizontal on peg F, D, B
//
// [R]   [B]   [R]
// [R]   [R]   [B]
// [R]   [B]   [B]
//    [R]   [B]
//    [B]   [R]
//    [B]   [R]
// [R]   [B]   [R]
// [R]   [R]   [B]
// [B]   [B]   [B]
func CreateBlueWinsThreeToTwoBoard() engine.Board {
	var board engine.Board

	board.Peg[engine.A].Add(engine.Red)
	board.Peg[engine.A].Add(engine.Red)
	board.Peg[engine.A].Add(engine.Red)

	board.Peg[engine.B].Add(engine.Blue)
	board.Peg[engine.B].Add(engine.Red)
	board.Peg[engine.B].Add(engine.Blue)

	board.Peg[engine.C].Add(engine.Red)
	board.Peg[engine.C].Add(engine.Blue)
	board.Peg[engine.C].Add(engine.Blue)

	board.Peg[engine.D].Add(engine.Red)
	board.Peg[engine.D].Add(engine.Blue)
	board.Peg[engine.D].Add(engine.Blue)

	board.Peg[engine.E].Add(engine.Blue)
	board.Peg[engine.E].Add(engine.Red)
	board.Peg[engine.E].Add(engine.Red)

	board.Peg[engine.F].Add(engine.Red)
	board.Peg[engine.F].Add(engine.Red)
	board.Peg[engine.F].Add(engine.Blue)

	board.Peg[engine.G].Add(engine.Blue)
	board.Peg[engine.G].Add(engine.Red)
	board.Peg[engine.G].Add(engine.Blue)

	board.Peg[engine.H].Add(engine.Red)
	board.Peg[engine.H].Add(engine.Blue)
	board.Peg[engine.H].Add(engine.Blue)

	return board
}
