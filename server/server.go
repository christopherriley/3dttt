package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/christopherriley/3dttt/server/command"
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

	if len(gr.CommandName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"missing command name\"}\n")))
		return
	}

	var cmd command.Command
	if cmd, err = command.CreateCommand(gr.CommandName, gr.CommandParams.(map[string]interface{})); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"failed to create command: %s\"}\n", err)))
		return
	}

	cmd.Execute()

	w.Header().Set("Content-Type", "application/json")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/game", gamePostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
