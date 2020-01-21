package engine_test

import (
	"github.com/christopherriley/3dttt/engine"
	testsamples "github.com/christopherriley/3dttt/engine/test/samples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createBlueWinsThreeToTwoBoard() engine.Board {
	var board engine.Board

	for _, move := range testsamples.BlueWinsThreeToTwo {
		board.Peg[move.Peg].Add(move.Colour)
	}

	return board
}

func createOneHorizontalRedRow() engine.Board {
	var board engine.Board

	board.Peg[engine.A].Add(engine.Red)
	board.Peg[engine.B].Add(engine.Red)
	board.Peg[engine.C].Add(engine.Red)

	return board
}

func createOneVerticalBlueRow() engine.Board {
	var board engine.Board

	board.Peg[engine.A].Add(engine.Blue)
	board.Peg[engine.A].Add(engine.Blue)
	board.Peg[engine.A].Add(engine.Blue)

	return board
}

var _ = Describe("Board Tests", func() {

	var subject engine.Board

	Describe("with empty board", func() {
		BeforeEach(func() {
			subject = engine.NewBoard()
		})
		It("is not full", func() {
			Expect(subject.IsFull()).To(BeFalse())
		})

		It("has zero red rows", func() {
			Expect(subject.CountCompleteLines(engine.Red)).To(Equal(0))
		})
		It("has zero blue rows", func() {
			Expect(subject.CountCompleteLines(engine.Blue)).To(Equal(0))
		})
	})

	Describe("with one horizontal red row", func() {
		BeforeEach(func() {
			subject = createOneHorizontalRedRow()
		})
		It("is not full", func() {
			Expect(subject.IsFull()).To(BeFalse())
		})
		It("has one red row", func() {
			Expect(subject.CountCompleteLines(engine.Red)).To(Equal(1))
		})
		It("has zero blue rows", func() {
			Expect(subject.CountCompleteLines(engine.Blue)).To(Equal(0))
		})
	})

	Describe("with one vertical blue row", func() {
		BeforeEach(func() {
			subject = createOneVerticalBlueRow()
		})
		It("is not full", func() {
			Expect(subject.IsFull()).To(BeFalse())
		})
		It("has one blue row", func() {
			Expect(subject.CountCompleteLines(engine.Blue)).To(Equal(1))
		})
		It("has zero red rows", func() {
			Expect(subject.CountCompleteLines(engine.Red)).To(Equal(0))
		})
	})

	Describe("with full board, blue winner", func() {
		BeforeEach(func() {
			subject = createBlueWinsThreeToTwoBoard()
		})
		It("is full", func() {
			Expect(subject.IsFull()).To(BeTrue())
		})

		It("has 2 completed red lines", func() {
			Expect(subject.CountCompleteLines(engine.Red)).To(Equal(2))
		})

		It("has 3 completed blue lines", func() {
			Expect(subject.CountCompleteLines(engine.Blue)).To(Equal(3))
		})
	})
})
