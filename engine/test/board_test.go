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
})
