package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/christopherriley/3dttt/engine"
	"github.com/christopherriley/3dttt/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ServerTestPort = 8080

type gameState struct {
	NextMove   string       `json:"next_move"`
	RedScore   int          `json:"red_score"`
	BlueScore  int          `json:"blue_score"`
	BoardState engine.Board `json:"board_state"`
}

type newGameResponse struct {
	Id    string    `json:"id"`
	State gameState `json:"state"`
}

var _ = Describe("Game Server Tests", func() {
	var subject server.GameServer
	var command string
	var params map[string]string
	var responseStatusCode int

	url := fmt.Sprintf("http://localhost:%d/api/v1/game", ServerTestPort)

	BeforeEach(func() {
		params = make(map[string]string)
		go func() {
			subject.Start(ServerTestPort)
		}()
	})

	Describe("POST operations", func() {
		var bodyBytes []byte

		JustBeforeEach(func() {
			requestBody, marshalErr := json.Marshal(map[string]interface{}{
				"command": command,
				"params":  params,
			})
			if marshalErr != nil {
				log.Fatalln(marshalErr)
			}

			response, postErr := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
			if postErr != nil {
				log.Fatalln(postErr)
			}

			responseStatusCode = response.StatusCode

			var readErr error
			bodyBytes, readErr = ioutil.ReadAll(response.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}
		})

		Describe("New Game command", func() {
			var response newGameResponse

			JustBeforeEach(func() {
				unmarshalErr := json.Unmarshal(bodyBytes, &response)
				if unmarshalErr != nil {
					log.Fatal(unmarshalErr)
				}
			})

			Describe("new 1p game with red starting", func() {
				BeforeEach(func() {
					command = "newgame_1p"
					params["colour"] = "red"
					params["move_first"] = "TRUE"
				})
				It("succeeds", func() {
					Expect(responseStatusCode).To(Equal(200))
				})
				It("sets next move to RedToMove", func() {
					Expect(response.State.NextMove).To(Equal("RedToMove"))
				})
			})
		})
	})
})
