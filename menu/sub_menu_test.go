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
					Key:   rune('t'),
				},
				SubMenu{
					Title: "Test 2",
					Key:   rune('t'),
				},
				SubMenu{
					Title: "Test 3",
					Key:   rune('t'),
				},
				SubMenu{
					Title: "Test 4",
					Key:   rune('t'),
				},
			}
		})

		Context("with width 20", func() {
			It("returns correct content", func() {
				content := subMenus.ContentForWindow(SubMenu{}, 30)
				Expect(content).To(Equal([][]rune{
					[]rune("  Test 1➔ t  Test 2➔ t"),
					[]rune("  Test 3➔ t  Test 4➔ t"),
				}))
			})
		})
		Context("with width 10", func() {
			It("returns correct content", func() {
				content := subMenus.ContentForWindow(SubMenu{}, 10)
				Expect(content).To(BeEquivalentTo([][]rune{
					[]rune{},
					[]rune("  Test 1➔ t"),
					[]rune("  Test 2➔ t"),
					[]rune("  Test 3➔ t"),
					[]rune("  Test 4➔ t"),
				}))
			})
		})
	})
})
