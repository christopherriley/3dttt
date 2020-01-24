package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func newGameHandler(params map[string]interface{}) error {
	var playerColourStr, playerFirstStr string
	var ok bool
	playerColourStr, ok = params["colour"].(string)
	if !ok {
		return fmt.Errorf("parameter 'colour' missing")
	}
	playerFirstStr, ok = params["move_first"].(string)
	if !ok {
		return fmt.Errorf("prameter 'move_first' missing")
	}

	fmt.Println("new game handler: colour:", playerColourStr, ", first:", playerFirstStr)

	return nil
}

func gamePostHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println("POST handler")
	var gr gameRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&gr)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	switch gr.CommandName {
	case "newgame":
		params := gr.CommandParams.(map[string]interface{})
		if err := newGameHandler(params); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}\n", err)))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"command '%s' not found\"}\n", gr.CommandName)))
	}

	w.Header().Set("Content-Type", "application/json")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/game", gamePostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
