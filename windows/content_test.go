package windows

import (
	"goo/common"

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

	Describe("AppendLineRuneArray", func() {
		It("appends a line at the end of the content from rune array", func() {
			Expect(window.Content).To(BeEmpty())
			window.AppendLineRuneArray([]rune("test"))
			Expect(len(window.Content)).To(Equal(1))
			Expect(window.Content).To(Equal(common.Cells{
				common.RunesToCells([]rune("test")),
			}))
		})
	})

	Describe("ContentAsRuneArray", func() {
		It("returns the content as a rune array", func() {
			Expect(window.ContentAsRuneArray()).To(BeEmpty())
			window.Content = common.Cells{
				common.RunesToCells([]rune("test")),
			}
			Expect(len(window.ContentAsRuneArray())).To(Equal(1))
			Expect(window.ContentAsRuneArray()[0]).To(Equal([]rune("test")))
		})
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
