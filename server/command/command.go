package command

import (
	"fmt"

	"github.com/christopherriley/3dttt/server/state"
)

type Command interface {
	Execute(s *state.GlobalState) (Response, error)
}

func CreateCommand(name string, p Params) (Command, error) {
	switch name {
	case "newgame_1p":
		var ngc New1PGameCommand
		if err := ngc.Create(p); err != nil {
			return nil, err
		}
		return ngc, nil
	case "move":
		var mc MoveCommand
		if err := mc.Create(p); err != nil {
			return nil, err
		}
		return mc, nil
	default:
		return nil, fmt.Errorf("unknown command: '%s'", name)
	}
}
