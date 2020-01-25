package command

import (
	"encoding/json"
	"fmt"

	"github.com/christopherriley/3dttt/server/state"
)

type CommandResponse struct {
	responseMap map[string]string
}

func CreateCommandResponse() *CommandResponse {
	var cr CommandResponse
	cr.responseMap = make(map[string]string)

	return &cr
}

func (cr *CommandResponse) Add(key, value string) {
	cr.responseMap[key] = value
}

func (cr CommandResponse) String() string {
	j, _ := json.Marshal(cr.responseMap)
	return string(j)
}

type Command interface {
	Execute(s *state.GlobalState) (CommandResponse, error)
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
