package common_test

import (
	"goo/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Int Helpers", func() {
	Describe("MaxInt", func() {
		It("returns the bigger number", func() {
			Expect(common.MaxInt(1, 5)).To(Equal(5))
		})
	})
	Describe("MinInt", func() {
		It("returns the smaller number", func() {
			Expect(common.MinInt(1, 5)).To(Equal(1))
		})
	})
})
