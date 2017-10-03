package windows

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window Line Numbers", func() {
	var (
		window *Window
	)
	BeforeEach(func() {
		window = &Window{
			EnableLineNum: true,
		}
		window.AppendLineRuneArray([]rune("1"))
		window.AppendLineRuneArray([]rune("2"))
		window.SetCursor(0, 0)
	})
	Describe("LineNumerLen", func() {
		Context("when EnableLineNum is true", func() {
			It("returns the length of the last line number", func() {
				Expect(window.lineNumerLen()).To(Equal(1))
			})

			It("returns 2 when lines between 10-99", func() {
				Expect(window.lineNumerLen()).To(Equal(1))

				for i := 0; i <= 11; i++ {
					window.AppendLineRuneArray([]rune("a"))
				}
				Expect(window.lineNumerLen()).To(Equal(2))
			})
		})

		Context("when EnableLineNum is false", func() {
			It("returns 0", func() {
				window.EnableLineNum = false
				Expect(window.lineNumerLen()).To(Equal(0))
			})
		})
	})
})
