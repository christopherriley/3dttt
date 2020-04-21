package command

import (
	"fmt"
	"os"

	"github.com/christopherriley/3dttt/cpu_player"
	"github.com/christopherriley/3dttt/engine"
	"github.com/christopherriley/3dttt/server/state"
)

type CPUMoveCommand struct {
	id  string
	peg engine.PegLabel
}

func (mc *CPUMoveCommand) Create(p Params) error {
	var id string
	var err error

	if id, err = p.Get("id"); err != nil {
		return err
	}

	mc.id = id

	return nil
}

func (mc CPUMoveCommand) Execute(s *state.GlobalState) (Response, error) {
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
	} else if gameState.NextMove == engine.RedToMove && state1P.PlayerColour == engine.Red ||
		gameState.NextMove == engine.BlueToMove && state1P.PlayerColour == engine.Blue {
		return Response{}, fmt.Errorf("invalid move - not cpu player turn")
	}

	move := cpu_player.GetNextMove(state1P.Game.GetBoard(), state1P.CPUColour, 6)

	if err = state1P.Game.Move(move); err != nil {
		fmt.Printf("WARNING: CPU MADE A BAD MOVE: %s", err)
		os.Exit(1)
	}

	r := CreateResponse(state1P.Game.GetGameState().NextMove, 0, 0, state1P.Game.GetBoard())
	r.Add("cpu_move", engine.PegToString(move))

	return *r, nil
}
