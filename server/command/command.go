package command

import (
	"fmt"
)

type Command interface {
	Execute() error
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
