package command

import "encoding/json"

type Response struct {
	responseMap map[string]string
}

func CreateResponse() *Response {
	var r Response
	r.responseMap = make(map[string]string)

	return &r
}

func (r *Response) Add(key, value string) {
	r.responseMap[key] = value
}

func (r Response) String() string {
	j, _ := json.Marshal(r.responseMap)
	return string(j)
}
