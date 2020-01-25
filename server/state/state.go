package state

import (
	"fmt"

	"github.com/christopherriley/3dttt/engine"
)

type GameState1P struct {
	Game         *engine.Game
	PlayerColour engine.Colour
}

type GlobalState struct {
	gameMap1P map[string]*GameState1P
}

func (state *GlobalState) Initialize() {
	state.gameMap1P = make(map[string]*GameState1P)
}

func (state *GlobalState) Add1PGame(guid string, playerColour engine.Colour, g *engine.Game) {
	var state1P GameState1P
	state1P.Game = g
	state1P.PlayerColour = playerColour
	state.gameMap1P[guid] = &state1P
}

func (state GlobalState) Fetch1PGame(guid string) (*GameState1P, error) {
	var state1P *GameState1P
	var found bool

	if state1P, found = state.gameMap1P[guid]; !found {
		return nil, fmt.Errorf("1p game state with id '%s' not found", guid)
	}

	return state1P, nil
}
