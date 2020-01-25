package command

import (
	"fmt"
	"strings"

	"github.com/christopherriley/3dttt/engine"
	"github.com/christopherriley/3dttt/server/state"
	"github.com/segmentio/ksuid"
)

type NewGameCommand struct {
	humanColour engine.Colour
	humanFirst  bool
}

func (ngc *NewGameCommand) Create(params map[string]interface{}) error {
	var playerColourStr, playerFirstStr string
	var ok bool

	playerColourStr, ok = params["colour"].(string)
	if !ok {
		return fmt.Errorf("parameter 'colour' missing")
	}
	pc := strings.ToUpper(playerColourStr)
	switch pc {
	case "RED":
		ngc.humanColour = engine.Red
	case "BLUE":
		ngc.humanColour = engine.Blue
	default:
		return fmt.Errorf("unknown value '%s' for key 'colour'", playerColourStr)
	}

	playerFirstStr, ok = params["move_first"].(string)
	if !ok {
		return fmt.Errorf("parameter 'move_first' missing")
	}
	pf := strings.ToUpper(playerFirstStr)
	switch pf {
	case "TRUE":
		ngc.humanFirst = true
	case "FALSE":
		ngc.humanFirst = false
	default:
		return fmt.Errorf("unknown value '%s' for key 'move_first'", playerFirstStr)
	}

	return nil
}

func (ngc NewGameCommand) Execute(s *state.GlobalState) (CommandResponse, error) {
	var game engine.Game
	id := ksuid.New()

	if ngc.humanFirst {
		game = engine.NewGame(ngc.humanColour)
	} else if ngc.humanColour == engine.Red {
		game = engine.NewGame(engine.Blue)
	} else {
		game = engine.NewGame(engine.Red)
	}

	s.AddGame(id.String(), &game)
	cr := CreateCommandResponse()
	cr.Add("id", id.String())

	return *cr, nil
}
