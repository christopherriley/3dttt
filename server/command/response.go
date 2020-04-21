package command

import (
	"encoding/json"
	"fmt"

	engine "github.com/christopherriley/3dttt/engine"
)

type gameState struct {
	NextMove   string       `json:"next_move"`
	RedScore   int          `json:"red_score"`
	BlueScore  int          `json:"blue_score"`
	BoardState engine.Board `json:"board_state"`
}

type Response struct {
	responseMap map[string]interface{}
}

func CreateResponse(nextMove engine.NextMove, redScore, blueScore int, board engine.Board) *Response {
	var r Response
	r.responseMap = make(map[string]interface{})

	nextMoveStr, _ := engine.NextMoveToString(nextMove)

	r.Add("state", gameState{
		NextMove:   nextMoveStr,
		RedScore:   redScore,
		BlueScore:  blueScore,
		BoardState: board,
	})

	return &r
}

func (r *Response) Add(key string, value interface{}) {
	r.responseMap[key] = value
}

func (r Response) String() string {
	j, err := json.Marshal(r.responseMap)
	if err != nil {
		fmt.Printf("encountered error: %s", err)
		return "ERROR"
	}
	return string(j)
}
