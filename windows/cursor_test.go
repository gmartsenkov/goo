package windows

import (
	"goo/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window Cursor", func() {
	var (
		window *Window
	)
	BeforeEach(func() {
		window = &Window{
			Dimensions: common.Dimensions{
				Rows: 2,
				Cols: 2,
			},
		}
		window.SetCursor(0, 0)
	})

	Describe("MoveCursorUp", func() {
		BeforeEach(func() {
			window.Content = [][]byte{
				[]byte("1"),
				[]byte("23"),
				[]byte("234"),
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

		Context("when there is no line above", func() {
			It("does not alter the cursor", func() {
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
				window.MoveCursorUp()
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
			})
		})

		Context("when line is outside the window bounds", func() {
			BeforeEach(func() {
				window.Dimensions.Rows = 1
			})
			It("decreases the vertical offset", func() {
				window.SetCursor(0, 0)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
				Expect(window.OffsetV).To(Equal(0))

				window.MoveCursorDown()
				Expect(window.OffsetV).To(Equal(0))
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))
				window.MoveCursorDown()

				Expect(window.OffsetV).To(Equal(1))
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))

				window.MoveCursorUp()
				Expect(window.OffsetV).To(Equal(1))
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))

				window.MoveCursorUp()
				Expect(window.OffsetV).To(Equal(0))
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
			})
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

	Describe("MoveCursorDown", func() {
		BeforeEach(func() {
			window.Content = [][]byte{
				[]byte("23"),
				[]byte("1"),
			}
		})
		It("moves the cursor down", func() {
			window.SetCursor(0, 0)
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 0,
				Y: 0,
			}))
			window.MoveCursorDown()
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 0,
				Y: 1,
			}))
		})
		Context("when there is no line below", func() {
			It("does not alter the cursor", func() {
				window.SetCursor(0, 1)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))
				window.MoveCursorDown()
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))
			})
		})

		Context("when line is outside the window bounds", func() {
			BeforeEach(func() {
				window.Dimensions.Rows = 1
				window.Content = [][]byte{
					[]byte("1"),
					[]byte("23"),
					[]byte("234"),
				}
			})
			It("increments the vertical offset", func() {
				window.SetCursor(0, 0)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
				Expect(window.OffsetV).To(Equal(0))

				window.MoveCursorDown()
				Expect(window.OffsetV).To(Equal(0))
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))
				window.MoveCursorDown()

				Expect(window.OffsetV).To(Equal(1))
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))
			})
		})

		Context("when below line is shorter", func() {
			It("corrects X", func() {
				window.SetCursor(1, 0)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 1,
					Y: 0,
				}))
				window.MoveCursorDown()
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 1,
				}))
			})
		})
	})

	Describe("MoveCursorLeft", func() {
		BeforeEach(func() {
			window.Content = [][]byte{
				[]byte("1234"),
			}
		})

		It("moves the cursor to the left", func() {
			window.SetCursor(3, 0)
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 3,
				Y: 0,
			}))

			window.MoveCursorLeft()
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 2,
				Y: 0,
			}))
		})

		Context("when cursor is on the first character", func() {
			It("doesn't move the cursor to the left", func() {
				window.SetCursor(0, 0)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))

				window.MoveCursorLeft()
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 0,
					Y: 0,
				}))
			})
		})
	})

	Describe("MoveCursorRight", func() {
		BeforeEach(func() {
			window.Content = [][]byte{
				[]byte("1234"),
			}
		})

		It("moves the cursor to the right", func() {
			window.SetCursor(0, 0)
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 0,
				Y: 0,
			}))

			window.MoveCursorRight()
			Expect(window.Cursor).To(BeEquivalentTo(Cursor{
				X: 1,
				Y: 0,
			}))
		})

		Context("when cursor is on the last character", func() {
			It("doesn't move the cursor any further", func() {
				window.SetCursor(3, 0)
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 3,
					Y: 0,
				}))

				window.MoveCursorRight()
				Expect(window.Cursor).To(BeEquivalentTo(Cursor{
					X: 3,
					Y: 0,
				}))
			})
		})
	})
})
