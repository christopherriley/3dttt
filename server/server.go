package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/christopherriley/3dttt/server/command"
	"github.com/christopherriley/3dttt/server/state"
	"github.com/gorilla/mux"
)

type gameRequest struct {
	CommandName   string      `json:"command"`
	CommandParams interface{} `json:"params"`
}

type newGameCommand struct {
	PlayerColour    string `json:"colour"`
	PlayerMoveFirst string `json:"move_first"`
}

type GameServer struct {
	state state.GlobalState
}

func (gs GameServer) Start(listenPort int) {
	gs.state.Initialize()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/game", gs.gamePostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), r))
}

func (gs GameServer) gamePostHandler(w http.ResponseWriter, req *http.Request) {
	var gr gameRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&gr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"failed to decode request: %s\"}\n", err)))
		return
	}

	if len(gr.CommandName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"missing command name\"}\n")))
		return
	}

	var paramsMap map[string]interface{}
	var ok bool
	if paramsMap, ok = gr.CommandParams.(map[string]interface{}); !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"failed to extract command parameters\"}"))
	}
	params := command.CreateParams(paramsMap)
	var cmd command.Command
	if cmd, err = command.CreateCommand(gr.CommandName, params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"failed to create command: %s\"}\n", err)))
		return
	}

	var r command.Response
	if r, err = cmd.Execute(&gs.state); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"failed to execute command: %s\"}\n", err)))
	}

	w.Write([]byte(r.String()))
}
