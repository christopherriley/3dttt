package engine_test

import (
	"github.com/christopherriley/3dttt/engine"
	testsamples "github.com/christopherriley/3dttt/engine/test/samples"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board Tests", func() {

	var subject engine.Board

	BeforeEach(func() {
		subject = engine.NewBoard()
	})

	Describe("with empty board", func() {
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
		JustBeforeEach(func() {
			subject = testsamples.CreateOneHorizontalRedRow()
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
		JustBeforeEach(func() {
			subject = testsamples.CreateOneVerticalBlueRow()
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
		JustBeforeEach(func() {
			subject = testsamples.CreateBlueWinsThreeToTwoBoard()
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

var _ = Describe("Game Tests", func() {
})
