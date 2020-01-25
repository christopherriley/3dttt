package state

import (
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
