package engine_test

import (
	"github.com/christopherriley/3dttt/engine"
	testsamples "github.com/christopherriley/3dttt/engine/test/samples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game Tests", func() {

	var subject engine.Game

	Describe("new game with red starting", func() {
		BeforeEach(func() {
			subject = engine.NewGame(engine.Red)
		})
		It("has zero blue lines", func() {
			Expect(subject.GetGameState().BlueLines).To(Equal(0))
		})
		It("has zero red lines", func() {
			Expect(subject.GetGameState().RedLines).To(Equal(0))
		})
		It("is red's turn", func() {
			Expect(subject.GetGameState().NextMove).To(Equal(engine.RedToMove))
		})
	})

	Describe("completed game with blue winning three to two", func() {
		BeforeEach(func() {
			subject = engine.NewGame(engine.Red)

			for _, move := range testsamples.BlueWinsThreeToTwo {
				subject.Move(move.Peg)
			}
		})
		It("has three blue lines", func() {
			Expect(subject.GetGameState().BlueLines).To(Equal(3))
		})
		It("has two red lines", func() {
			Expect(subject.GetGameState().RedLines).To(Equal(2))
		})
		It("has blue as winner", func() {
			Expect(subject.GetGameState().NextMove).To(Equal(engine.BlueWins))
		})
	})

	Describe("in-progress game with red winning two to zero", func() {
		BeforeEach(func() {
			subject = engine.NewGame(engine.Red)

			for _, move := range testsamples.RedWinningTwoToZero {
				subject.Move(move.Peg)
			}
		})
		It("has zero blue lines", func() {
			Expect(subject.GetGameState().BlueLines).To(Equal(0))
		})
		It("has two red lines", func() {
			Expect(subject.GetGameState().RedLines).To(Equal(2))
		})
		It("is blue's turn", func() {
			Expect(subject.GetGameState().NextMove).To(Equal(engine.BlueToMove))
		})
	})
})
