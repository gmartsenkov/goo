package windows

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window", func() {
	var (
		window *Window
	)
	BeforeEach(func() {
		window = &Window{}
	})

	Describe("SplitAndSetContent", func() {
		It("transforms the string into a two dimensional byte array", func() {
			Expect(window.Content).To(BeEmpty())
			content := []byte("1\n2\n3")
			window.SplitAndSetContent(content)
			Expect(window.Content).ToNot(BeEmpty())
			Expect(window.ContentAsRuneArray()).To(BeEquivalentTo([][]rune{
				[]rune("1"),
				[]rune("2"),
				[]rune("3"),
			}))
		})
	})
})
