package menu

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Menu", func() {
	var (
		menu Menu
	)
	BeforeEach(func() {
		menu = Menu{
			SubMenus: SubMenus{
				SubMenu{
					Title: "submenu1",
					Key:   rune('z'),
					subMenus: SubMenus{
						SubMenu{
							Title: "submenu1-submenu1",
							Key:   rune('a'),
						},
					},
				},
			},
		}
	})
	Describe("FindSubMenu", func() {
		Context("when there is a SubMenu", func() {
			Context("when there is one SubMenu", func() {
				It("returns the SubMenu", func() {
					subMenu, err := menu.FindSubMenu(rune('z'))
					Expect(err).To(BeNil())
					Expect(subMenu.Title).To(Equal("submenu1"))
				})
			})

			Context("when SubMenu is nested", func() {
				It("returns the nested SubMenu", func() {
					menu.Query = []byte{'z'}
					subMenu, err := menu.FindSubMenu(rune('a'))
					Expect(err).To(BeNil())
					Expect(subMenu.Title).To(Equal("submenu1-submenu1"))
				})
			})
		})

		Context("when there is no SubMenu", func() {
			Context("with no menu.Query", func() {
				It("returns an error and an empty SubMenu", func() {
					subMenu, err := menu.FindSubMenu(rune('a'))
					Expect(err).To(MatchError("SubMenuNotFound"))
					Expect(subMenu).To(BeEquivalentTo(SubMenu{}))
				})
			})
			Context("with a menu.Query", func() {
				It("returns an error and an empty SubMenu", func() {
					menu.Query = []byte{'z'}
					subMenu, err := menu.FindSubMenu(rune('z'))
					Expect(err).To(MatchError("SubMenuNotFound"))
					Expect(subMenu).To(BeEquivalentTo(SubMenu{}))
				})
			})
		})
	})

	Describe("FindLastSubMenu", func() {
		Context("when there is a SubMenu", func() {
			Context("when it is the first one", func() {
				It("returns the SubMenu", func() {
					menu.Query = []byte{'z'}
					subMenu, err := menu.FindLastSubMenu()
					Expect(err).To(BeNil())
					Expect(subMenu.Title).To(Equal("submenu1"))
				})

			})

			Context("when it is nested", func() {
				It("returns the SubMenu", func() {
					menu.Query = []byte{'z', 'a'}
					subMenu, err := menu.FindLastSubMenu()
					Expect(err).To(BeNil())
					Expect(subMenu.Title).To(Equal("submenu1-submenu1"))
				})
			})
		})
	})

	Context("when there is no Query", func() {
		It("returns an error and an empty SubMenu", func() {
			menu.Query = []byte{}
			subMenu, err := menu.FindLastSubMenu()
			Expect(err).To(MatchError("SubMenuNotFound"))
			Expect(subMenu).To(BeEquivalentTo(SubMenu{}))
		})
	})
})
