package command

import (
	"fmt"

	"github.com/christopherriley/3dttt/engine"
	"github.com/christopherriley/3dttt/server/state"
)

type MoveCommand struct {
	id  string
	peg engine.PegLabel
}

func (mc *MoveCommand) Create(p Params) error {
	var id, pegStr string
	var peg engine.PegLabel
	var err error

	if id, err = p.Get("id"); err != nil {
		return err
	}

	if pegStr, err = p.Get("peg"); err != nil {
		return err
	}

	if peg, err = engine.StringToPeg(pegStr); err != nil {
		return err
	}

	mc.id = id
	mc.peg = peg

	return nil
}

func (mc MoveCommand) Execute(s *state.GlobalState) (Response, error) {
	var game *engine.Game
	var err error

	if game, err = s.FetchGame(mc.id); err != nil {
		return Response{}, err
	}

	r := CreateResponse()

	if err = game.Move(mc.peg); err != nil {
		r.Add("move_status", "invalid")
	} else {
		r.Add("move_status", "accepted")
	}

	state := game.GetGameState()
	var stateStr string
	if stateStr, err = engine.BoardStateToString(state.BoardState); err != nil {
		return *r, err
	}
	r.Add("game_state", stateStr)
	r.Add("red_score", fmt.Sprintf("%d", state.RedLines))
	r.Add("blue_score", fmt.Sprintf("%d", state.BlueLines))

	return *r, nil
}
