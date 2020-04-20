package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/christopherriley/3dttt/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ServerTestPort = 8080

var _ = Describe("Game Server Tests", func() {

	var subject server.GameServer
	var command string
	var params map[string]string
	var requestBody []byte
	var response *http.Response
	var err error

	url := fmt.Sprintf("http://localhost:%d/api/v1/game", ServerTestPort)

	BeforeEach(func() {
		params = make(map[string]string)
		go func() {
			subject.Start(ServerTestPort)
		}()
	})

	JustBeforeEach(func() {
		requestBody, err = json.Marshal(map[string]interface{}{
			"command": command,
			"params":  params,
		})

		if err != nil {
			log.Fatalln(err)
		}

		response, err = http.Post(url, "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			log.Fatalln(err)
		}
	})

	Describe("new 1p game with red starting", func() {
		BeforeEach(func() {
			command = "newgame_1p"
			params["colour"] = "red"
			params["move_first"] = "TRUE"
		})
		It("succeeds", func() {
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})
