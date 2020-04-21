package command

import (
	"fmt"
	"strings"

	"github.com/christopherriley/3dttt/engine"
	"github.com/christopherriley/3dttt/server/state"
	"github.com/segmentio/ksuid"
)

type New1PGameCommand struct {
	humanColour engine.Colour
	humanFirst  bool
}

func (ngc *New1PGameCommand) Create(p Params) error {
	var playerColourStr, playerFirstStr string
	var err error

	if playerColourStr, err = p.Get("colour"); err != nil {
		return err
	}
	if playerFirstStr, err = p.Get("move_first"); err != nil {
		return err
	}

	switch playerColourStr {
	case "RED":
		ngc.humanColour = engine.Red
	case "BLUE":
		ngc.humanColour = engine.Blue
	default:
		return fmt.Errorf("unknown value '%s' for key 'colour'", playerColourStr)
	}

	switch playerFirstStr {
	case "TRUE":
		ngc.humanFirst = true
	case "FALSE":
		ngc.humanFirst = false
	default:
		return fmt.Errorf("unknown value '%s' for key 'move_first'", playerFirstStr)
	}

	return nil
}

func (ngc New1PGameCommand) Execute(s *state.GlobalState) (Response, error) {
	var game engine.Game
	var firstPlayer engine.Colour
	id := strings.ToUpper(ksuid.New().String())

	if ngc.humanFirst {
		firstPlayer = ngc.humanColour
	} else if ngc.humanColour == engine.Red {
		firstPlayer = engine.Blue
	} else {
		firstPlayer = engine.Red
	}

	game = engine.NewGame(firstPlayer)

	s.Add1PGame(id, ngc.humanColour, &game)
	r := CreateResponse(game.GetGameState().NextMove, 0, 0, game.GetBoard())
	r.Add("id", id)

	return *r, nil
}
