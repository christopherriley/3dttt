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

type commandParams map[string]interface{}

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

type moveResponse struct {
	MoveStatus string    `json:"move_status"`
	State      gameState `json:"state"`
}

type cpuMoveResponse struct {
	CPUMove string    `json:"cpu_move"`
	State   gameState `json:"state"`
}

func pegLetterToIndex(s string) int {
	switch s {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	case "D":
		return 3
	case "E":
		return 4
	case "F":
		return 5
	case "G":
		return 6
	case "H":
		return 7
	}

	return -1
}

func post(url string, command string, params commandParams) (responseStatusCode int, responseBodyBytes []byte, err error) {
	requestBodyBytes, marshalErr := json.Marshal(commandParams{
		"command": command,
		"params":  params,
	})
	if marshalErr != nil {
		log.Fatalln(marshalErr)
	}

	response, postErr := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	if postErr != nil {
		err = postErr
		return
	}

	var readErr error
	responseBodyBytes, readErr = ioutil.ReadAll(response.Body)
	if readErr != nil {
		err = readErr
		return
	}

	responseStatusCode = response.StatusCode

	fmt.Println()
	fmt.Println("*****************************************************")
	fmt.Println("POST REQUEST BODY")
	fmt.Println(string(requestBodyBytes))
	fmt.Println()

	fmt.Println("POST RESPONSE BODY")
	fmt.Println(string(responseBodyBytes))
	fmt.Println()

	return
}

var _ = Describe("Game Server Tests", func() {
	var subject server.GameServer
	var command string
	var params commandParams
	var responseStatusCode int

	url := fmt.Sprintf("http://localhost:%d/api/v1/game", ServerTestPort)

	BeforeEach(func() {
		params = make(commandParams)
		go func() {
			subject.Start(ServerTestPort)
		}()
	})

	Describe("POST operations", func() {
		var bodyBytes []byte
		var err error

		JustBeforeEach(func() {
			if responseStatusCode, bodyBytes, err = post(url, command, params); err != nil {
				log.Fatalln(err)
			}
		})

		Describe("New Game command", func() {
			var response newGameResponse

			JustBeforeEach(func() {
				unmarshalErr := json.Unmarshal(bodyBytes, &response)
				if unmarshalErr != nil {
					log.Fatalln(unmarshalErr)
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
				It("returns a game ID", func() {
					Expect(response.Id).ToNot(BeEmpty())
				})
				It("sets next move to RedToMove", func() {
					Expect(response.State.NextMove).To(Equal("RedToMove"))
				})
				It("reports a score of 0-0", func() {
					Expect(response.State.RedScore).To(BeZero())
					Expect(response.State.BlueScore).To(BeZero())
				})
				It("reports an empty game board", func() {
					Expect(response.State.BoardState.Peg[0]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[1]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[2]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[3]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[4]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[5]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[6]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[7]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
				})
			})
		})

		Describe("Move command", func() {
			var response moveResponse

			BeforeEach(func() {
				command = "move"
			})

			JustBeforeEach(func() {
				unmarshalErr := json.Unmarshal(bodyBytes, &response)
				if unmarshalErr != nil {
					log.Fatalln(unmarshalErr)
				}
			})

			Describe("with new game, player to move", func() {
				BeforeEach(func() {
					newGameParams := commandParams{}
					newGameParams["colour"] = "red"
					newGameParams["move_first"] = "TRUE"
					statusCode, bodyBytes, err := post(url, "newgame_1p", newGameParams)
					if err != nil {
						log.Fatalln(err)
					}
					if statusCode != 200 {
						log.Fatalln("failed to create new game")
					}

					var response newGameResponse
					unmarshalErr := json.Unmarshal(bodyBytes, &response)
					if unmarshalErr != nil {
						log.Fatalln(unmarshalErr)
					}

					params["id"] = response.Id
					params["peg"] = "A"
				})

				It("succeeds", func() {
					Expect(responseStatusCode).To(Equal(200))
				})
				It("accepts the move", func() {
					Expect(response.MoveStatus).To(Equal("accepted"))
				})
				It("expects Blue to move next", func() {
					Expect(response.State.NextMove).To(Equal("BlueToMove"))
				})
				It("reports a score of 0-0", func() {
					Expect(response.State.RedScore).To(BeZero())
					Expect(response.State.BlueScore).To(BeZero())
				})
				It("reports a game board with one red peg", func() {
					Expect(response.State.BoardState.Peg[0]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.Red, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[1]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[2]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[3]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[4]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[5]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[6]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
					Expect(response.State.BoardState.Peg[7]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.None, engine.None, engine.None},
					}))
				})
			})
		})

		Describe("CPU Move command", func() {
			var response cpuMoveResponse

			BeforeEach(func() {
				command = "cpu_move"
			})

			JustBeforeEach(func() {
				unmarshalErr := json.Unmarshal(bodyBytes, &response)
				if unmarshalErr != nil {
					log.Fatalln(unmarshalErr)
				}
			})

			Describe("with new game, cpu to move", func() {
				BeforeEach(func() {
					newGameParams := commandParams{}
					newGameParams["colour"] = "red"
					newGameParams["move_first"] = "FALSE"
					statusCode, bodyBytes, err := post(url, "newgame_1p", newGameParams)
					if err != nil {
						log.Fatalln(err)
					}
					if statusCode != 200 {
						log.Fatalln("failed to create new game")
					}

					var response newGameResponse
					unmarshalErr := json.Unmarshal(bodyBytes, &response)
					if unmarshalErr != nil {
						log.Fatalln(unmarshalErr)
					}

					params["id"] = response.Id
				})

				It("succeeds", func() {
					Expect(responseStatusCode).To(Equal(200))
				})
				It("returns a CPU move", func() {
					Expect(response.CPUMove).ToNot(BeEmpty())
				})
				It("expects Blue to move next", func() {
					Expect(response.State.NextMove).To(Equal("RedToMove"))
				})
				It("reports a score of 0-0", func() {
					Expect(response.State.RedScore).To(BeZero())
					Expect(response.State.BlueScore).To(BeZero())
				})
				It("reports a game board with one red peg", func() {
					Expect(response.State.BoardState.Peg[pegLetterToIndex(response.CPUMove)]).To(Equal(engine.Peg{
						Slot: [3]engine.Colour{engine.Blue, engine.None, engine.None},
					}))
				})
			})
		})
	})
})
