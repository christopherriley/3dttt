package state

import (
	"fmt"

	"github.com/christopherriley/3dttt/engine"
)

type GlobalState struct {
	gameMap map[string]*engine.Game
}

func (state *GlobalState) Initialize() {
	state.gameMap = make(map[string]*engine.Game)
}

func (state *GlobalState) AddGame(guid string, g *engine.Game) {
	state.gameMap[guid] = g
}

func (state GlobalState) FetchGame(guid string) (*engine.Game, error) {
	var game *engine.Game
	var found bool

	if game, found = state.gameMap[guid]; !found {
		return nil, fmt.Errorf("game with id '%s' not found", guid)
	}

	return game, nil
}
