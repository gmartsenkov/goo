package menu

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SubMenu", func() {
	var (
		subMenus SubMenus
	)
	Describe("FindByKey", func() {
		Context("when action exists", func() {
			BeforeEach(func() {
				subMenus = SubMenus{
					SubMenu{
						Key: rune('b'),
					},
				}
			})
			It("returns the SubMenu", func() {
				subMenu, err := subMenus.FindByKey(rune('b'))
				Expect(err).To(BeNil())
				Expect(subMenu).To(BeEquivalentTo(SubMenu{Key: rune('b')}))
			})
		})

		Context("when SubMenu does not exist", func() {
			It("returns an empty SubMenu and an error", func() {
				subMenu, err := subMenus.FindByKey(rune('a'))
				Expect(err).To(MatchError(errors.New("SubMenuNotFound")))
				Expect(subMenu).To(BeEquivalentTo(SubMenu{}))
			})
		})
	})

	Describe("ContentForWindow", func() {
		BeforeEach(func() {
			subMenus = SubMenus{
				SubMenu{
					Title: "Test 1",
				},
				SubMenu{
					Title: "Test 2",
				},
				SubMenu{
					Title: "Test 3",
				},
				SubMenu{
					Title: "Test 4",
				},
			}
		})

		Context("with width 20", func() {
			It("returns correct content", func() {
				content := subMenus.ContentForWindow(20)
				Expect(content).To(Equal([][]byte{
					[]byte(" Test 1  Test 2 "),
					[]byte(" Test 3  Test 4 "),
				}))
			})
		})
		Context("with width 10", func() {
			It("returns correct content", func() {
				content := subMenus.ContentForWindow(10)
				Expect(content).To(Equal([][]byte{
					[]byte(" Test 1 "),
					[]byte(" Test 2 "),
					[]byte(" Test 3 "),
					[]byte(" Test 4 "),
				}))
			})
		})
	})
})
