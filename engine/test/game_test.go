package engine_test

import (
	"github.com/christopherriley/3dttt/engine"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
})
