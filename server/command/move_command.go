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
	var state1P *state.GameState1P
	var err error

	if state1P, err = s.Fetch1PGame(mc.id); err != nil {
		return Response{}, err
	}

	gameState := state1P.Game.GetGameState()

	if gameState.NextMove == engine.Draw ||
		gameState.NextMove == engine.RedWins ||
		gameState.NextMove == engine.BlueWins {
		return Response{}, fmt.Errorf("invalid move - game is over")
	} else if gameState.NextMove == engine.RedToMove && state1P.PlayerColour == engine.Blue ||
		gameState.NextMove == engine.BlueToMove && state1P.PlayerColour == engine.Red {
		return Response{}, fmt.Errorf("invalid move - not human player turn")
	}

	moveStatus := "accepted"
	var r *Response
	if err = state1P.Game.Move(mc.peg); err != nil {
		moveStatus = "invalid"
	}

	r = CreateResponse(state1P.Game.GetGameState().NextMove,
		state1P.Game.GetGameState().RedLines,
		state1P.Game.GetGameState().BlueLines,
		state1P.Game.GetBoard())
	r.Add("move_status", moveStatus)

	return *r, nil
}
