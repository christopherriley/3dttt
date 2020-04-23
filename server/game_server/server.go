package game_server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/christopherriley/3dttt/server/command"
	"github.com/christopherriley/3dttt/server/state"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS_3DTTT")
	if len(allowedOrigins) == 0 {
		fmt.Println("* warning: no CORS origins are set - clients may not be able to connect\n\nPlease set the ALLOWED_ORIGINS_3DTTT env var\n")
	}
	fmt.Println("allowed origins: ", allowedOrigins)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{allowedOrigins},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), handler))
}

func (gs GameServer) gamePostHandler(w http.ResponseWriter, req *http.Request) {
	var gr gameRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&gr)

	w.Header().Set("Content-Type", "application/json")

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
		return
	}

	params, paramErr := command.CreateParams(paramsMap)
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"failed to create command params: %s\"}\n", err)))
		return
	}

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
		return
	}

	w.Write([]byte(r.String()))
}
