package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cells", func() {
	var (
		cellArray []Cell
	)

	BeforeEach(func() {
		cellArray = []Cell{
			Cell{Ch: 't'},
			Cell{Ch: 'e'},
			Cell{Ch: 's'},
			Cell{Ch: 't'},
		}
	})
	Describe("AsRunes", func() {
		It("converts the cells to runes", func() {
			Expect(CellsAsRuneArray(cellArray)).To(Equal([]rune("test")))
		})
	})
	Describe("BytesToCells", func() {
		It("coverts the bytes array into cells array", func() {
			byteArray := []byte("test")
			Expect(BytesToCells(byteArray)).To(Equal(cellArray))
		})
	})
	Describe("RunesToCells", func() {
		It("coverts the runes array into cells array", func() {
			runeArray := []rune("test")
			Expect(RunesToCells(runeArray)).To(Equal(cellArray))
		})
	})
})
