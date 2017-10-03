package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Int Helpers", func() {
	Describe("MaxInt", func() {
		It("returns the bigger number", func() {
			Expect(MaxInt(1, 5)).To(Equal(5))
		})
	})
	Describe("MinInt", func() {
		It("returns the smaller number", func() {
			Expect(MinInt(1, 5)).To(Equal(1))
		})
	})
})
