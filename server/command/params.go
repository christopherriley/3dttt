package command

import (
	"fmt"
	"strings"
)

type Params struct {
	paramsMap map[string]string
}

func CreateParams(m map[string]interface{}) (Params, error) {
	var p Params
	p.paramsMap = make(map[string]string)
	for k, v := range m {
		valString, ok := v.(string)
		if !ok {
			return Params{}, fmt.Errorf("parameter '%s' with value '%s' could not be converted to string", k, v)
		}
		p.paramsMap[k] = valString
	}

	return p, nil
}

func (p Params) Get(k string) (string, error) {
	v, ok := p.paramsMap[k]
	if !ok {
		return "", fmt.Errorf("parameter '%s' missing", k)
	}
	return strings.ToUpper(v), nil
}
