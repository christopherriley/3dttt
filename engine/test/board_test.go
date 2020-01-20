package engine_test

import (
	"github.com/christopherriley/3dttt/engine"
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

		It("evaluates to zero", func() {
			Expect(subject.Evaluate()).To(BeZero())
		})
	})

	Describe("with one horizontal red row", func() {
		JustBeforeEach(func() {
			subject.Peg[engine.A].Add(engine.Red)
			subject.Peg[engine.B].Add(engine.Red)
			subject.Peg[engine.C].Add(engine.Red)
		})
		It("is not full", func() {
			Expect(subject.IsFull()).To(BeFalse())
		})

		It("evaluates to +1", func() {
			Expect(subject.Evaluate()).To(Equal(1))
		})
	})

	Describe("with one horizontal blue row", func() {
		JustBeforeEach(func() {
			subject.Peg[engine.A].Add(engine.Blue)
			subject.Peg[engine.B].Add(engine.Blue)
			subject.Peg[engine.C].Add(engine.Blue)
		})
		It("is not full", func() {
			Expect(subject.IsFull()).To(BeFalse())
		})

		It("evaluates to -1", func() {
			Expect(subject.Evaluate()).To(Equal(-1))
		})
	})

	Describe("with full board, blue winner", func() {

		// RED LINE 1: vertical on peg A
		// RED LINE 2: diagonal on peg A, B, C
		//
		// BLUE LINE 1: horizontal on peg F, G, H
		// BLUE LINE 2: diagonal on peg F, D, B
		// BLUE LINE 2: horizontal on peg F, D, B
		//
		// [R]   [B]   [R]
		// [R]   [R]   [B]
		// [R]   [B]   [B]
		//    [R]   [B]
		//    [B]   [R]
		//    [B]   [R]
		// [R]   [B]   [R]
		// [R]   [R]   [B]
		// [B]   [B]   [B]
		JustBeforeEach(func() {
			subject.Peg[engine.A].Add(engine.Red)
			subject.Peg[engine.A].Add(engine.Red)
			subject.Peg[engine.A].Add(engine.Red)

			subject.Peg[engine.B].Add(engine.Blue)
			subject.Peg[engine.B].Add(engine.Red)
			subject.Peg[engine.B].Add(engine.Blue)

			subject.Peg[engine.C].Add(engine.Red)
			subject.Peg[engine.C].Add(engine.Blue)
			subject.Peg[engine.C].Add(engine.Blue)

			subject.Peg[engine.D].Add(engine.Red)
			subject.Peg[engine.D].Add(engine.Blue)
			subject.Peg[engine.D].Add(engine.Blue)

			subject.Peg[engine.E].Add(engine.Blue)
			subject.Peg[engine.E].Add(engine.Red)
			subject.Peg[engine.E].Add(engine.Red)

			subject.Peg[engine.F].Add(engine.Red)
			subject.Peg[engine.F].Add(engine.Red)
			subject.Peg[engine.F].Add(engine.Blue)

			subject.Peg[engine.G].Add(engine.Blue)
			subject.Peg[engine.G].Add(engine.Red)
			subject.Peg[engine.G].Add(engine.Blue)

			subject.Peg[engine.H].Add(engine.Red)
			subject.Peg[engine.H].Add(engine.Blue)
			subject.Peg[engine.H].Add(engine.Blue)
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

		It("evaluates to blue win", func() {
			Expect(subject.Evaluate()).To(Equal(engine.BlueWinScore))
		})
	})
})
