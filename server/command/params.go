package command

import (
	"fmt"
	"strings"
)

type Params struct {
	paramsMap map[string]string
}

func CreateParams(m map[string]interface{}) Params {
	var p Params
	p.paramsMap = make(map[string]string)
	for k, v := range m {
		p.paramsMap[k] = v.(string)
	}

	return p
}

func (p Params) Get(k string) (string, error) {
	v, ok := p.paramsMap[k]
	if !ok {
		return "", fmt.Errorf("parameter '%s' missing", k)
	}
	return strings.ToUpper(v), nil
}
