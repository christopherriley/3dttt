package command

import (
	"fmt"
	"net/http"

	"github.com/christopherriley/3dttt/server/state"
)

type Command interface {
	Execute(s *state.GlobalState, w http.ResponseWriter) error
}

func CreateCommand(name string, params map[string]interface{}) (Command, error) {
	fmt.Println("name:", name)
	switch name {
	case "newgame":
		var ngc NewGameCommand
		if err := ngc.Create(params); err != nil {
			return nil, err
		}
		return ngc, nil
	default:
		return nil, fmt.Errorf("unknown command: '%s'", name)
	}
}
