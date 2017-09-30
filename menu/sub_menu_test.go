package menu

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SubMenu", func() {
	Describe("FindByKey", func() {
		var (
			subMenus SubMenus
		)
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
})
