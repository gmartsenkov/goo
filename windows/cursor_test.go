package windows

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window Cursor", func() {
	var (
		window *Window
	)
	BeforeEach(func() {
		window = &Window{}
		window.SetCursor(0, 0)
	})

	Describe("MoveCursorUp", func() {
		BeforeEach(func() {
			window.Content = [][]byte{
				[]byte("1"),
				[]byte("23"),
			}
		})
		It("moves the cursor up", func() {
			window.SetCursor(0, 1)
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 0,
				Y: 1,
			}))
			window.MoveCursorUp()
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 0,
				Y: 0,
			}))
		})
		Context("when above line is shorter", func() {
			It("corrects X as well", func() {
				window.SetCursor(1, 1)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 1,
					Y: 1,
				}))
				window.MoveCursorUp()
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
			})
		})
	})
})
