package testsamples

import (
	"github.com/christopherriley/3dttt/engine"
)

type Move struct {
	Peg    engine.PegLabel
	Colour engine.Colour
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
//
var BlueWinsThreeToTwo = []Move{
	Move{engine.A, engine.Red},
	Move{engine.B, engine.Blue},
	Move{engine.C, engine.Red},
	Move{engine.C, engine.Blue},
	Move{engine.A, engine.Red},
	Move{engine.C, engine.Blue},
	Move{engine.D, engine.Red},
	Move{engine.D, engine.Blue},
	Move{engine.B, engine.Red},
	Move{engine.B, engine.Blue},
	Move{engine.A, engine.Red},
	Move{engine.D, engine.Blue},
	Move{engine.F, engine.Red},
	Move{engine.E, engine.Blue},
	Move{engine.E, engine.Red},
	Move{engine.G, engine.Blue},
	Move{engine.G, engine.Red},
	Move{engine.G, engine.Blue},
	Move{engine.H, engine.Red},
	Move{engine.H, engine.Blue},
	Move{engine.F, engine.Red},
	Move{engine.F, engine.Blue},
	Move{engine.E, engine.Red},
	Move{engine.H, engine.Blue},
}
