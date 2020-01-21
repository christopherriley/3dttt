package engine_test

import (
	"github.com/christopherriley/3dttt/engine"
	testsamples "github.com/christopherriley/3dttt/engine/test/samples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createBlueWinsThreeToTwoGame() engine.Game {
	game := engine.NewGame(engine.Red)

	for _, move := range testsamples.BlueWinsThreeToTwo {
		game.Move(move.Peg)
	}

	return game
}

var _ = Describe("Game Tests", func() {

	var subject engine.Game

	Describe("new game with red starting", func() {
		BeforeEach(func() {
			subject = engine.NewGame(engine.Red)
		})
		It("is has zero blue lines", func() {
			Expect(subject.GetGameState().BlueLines).To(Equal(0))
		})
		It("it has zero red lines", func() {
			Expect(subject.GetGameState().RedLines).To(Equal(0))
		})
		It("is red's turn", func() {
			Expect(subject.GetGameState().BoardState).To(Equal(engine.RedToMove))
		})
	})

	Describe("completed game with blue winning three to two", func() {
		BeforeEach(func() {
			subject = createBlueWinsThreeToTwoGame()
		})
		It("is has three blue lines", func() {
			Expect(subject.GetGameState().BlueLines).To(Equal(3))
		})
		It("it has two red lines", func() {
			Expect(subject.GetGameState().RedLines).To(Equal(2))
		})
		It("has blue as winner", func() {
			Expect(subject.GetGameState().BoardState).To(Equal(engine.BlueWins))
		})
	})
})
