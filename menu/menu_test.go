package menu

import (
	"goo/common"
	"goo/editors"
	"goo/windows"

	"github.com/nsf/termbox-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Menu", func() {
	var (
		menu   Menu
		window *windows.Window
	)
	BeforeEach(func() {
		window = &windows.Window{}
		window.Dimensions.Cols = 20
		menu = Menu{
			TriggerKey: termbox.KeySpace,
			SubMenus: SubMenus{
				SubMenu{
					Title: "submenu1",
					Key:   rune('z'),
					subMenus: SubMenus{
						SubMenu{
							Title: "submenu1-submenu1",
							Key:   rune('a'),
							actions: Actions{
								Action{
									Key: rune('b'),
									Fn: func(e *editors.Editor) {
										e.InsertState()
									},
								},
							},
						},
					},
				},
				SubMenu{
					Title: "submenu2",
					Key:   rune('z'),
				},
			},
		}
	})

	Describe("TriggerListener", func() {
		Context("when trigger key", func() {
			It("sets trigger to true", func() {
				Expect(menu.triggered).To(BeFalse())
				menu.TriggerListener(window, termbox.KeySpace)
				Expect(menu.triggered).To(BeTrue())
			})

			It("sets the window content to the main submenus", func() {
				Expect(window.Content).To(BeEmpty())
				menu.TriggerListener(window, termbox.KeySpace)
				Expect(window.Content).To(BeEquivalentTo(common.Cells{
					common.RunesToCells([]rune("  submenu1➔ z")),
					common.RunesToCells([]rune("  submenu2➔ z")),
				}))
			})
		})

		Context("when non trigger key", func() {
			It("does not change triggered to true", func() {
				Expect(menu.triggered).To(BeFalse())
				menu.TriggerListener(window, termbox.KeyBackspace)
				Expect(menu.triggered).To(BeFalse())
			})

			It("does not write to the window content", func() {
				Expect(window.Content).To(BeEmpty())
				menu.TriggerListener(window, termbox.KeyBackspace)
				Expect(window.Content).To(BeEmpty())
			})
		})
	})

	Describe("IsTriggered", func() {
		Context("when triggered true", func() {
			It("returns true", func() {
				menu.triggered = true
				Expect(menu.IsTriggered()).To(BeTrue())
			})
		})
		Context("when triggered false", func() {
			It("returns false", func() {
				menu.triggered = false
				Expect(menu.IsTriggered()).To(BeFalse())
			})
		})
	})

	Describe("Process", func() {
		var (
			editor *editors.Editor
			window *windows.Window
		)

		BeforeEach(func() {
			editor = &editors.Editor{}
			window = &windows.Window{}
			window.Dimensions.Cols = 50
			window.Dimensions.Rows = 100
		})

		Context("on first enter to SubMenu", func() {
			It("renders the submenu tree", func() {
				menu.Process(window, editor, rune('z'))
				Expect(window.Content).To(BeEquivalentTo(common.Cells{
					common.RunesToCells([]rune("  submenu1-submenu1➔ a")),
				}))
			})
		})

		Context("when submenu can't be found", func() {
			It("returns an empty window.content", func() {
				menu.Process(window, editor, rune('g'))
				Expect(window.Content).To(BeEmpty())
			})

			It("sets triggered to false", func() {
				menu.triggered = true
				Expect(menu.triggered).To(BeTrue())
				menu.Process(window, editor, rune('g'))
				Expect(menu.triggered).To(BeFalse())
			})
		})

		Context("when action is called", func() {
			BeforeEach(func() {
				menu.Query = []rune{'z', 'a'}
				menu.triggered = true
			})

			It("executes the action", func() {
				Expect(editor.State).To(Equal(editors.StateNormal))
				menu.Process(window, editor, rune('b'))
				Expect(menu.triggered).To(BeFalse())
			})

			It("sets menu.triggered to false", func() {
				Expect(menu.triggered).To(BeTrue())
				menu.Process(window, editor, rune('b'))
				Expect(menu.triggered).To(BeFalse())
			})

			It("returns an empty window content and executes the action", func() {
				Expect(window.Content).To(BeEmpty())
				menu.Process(window, editor, rune('b'))
				Expect(window.Content).To(BeEmpty())
			})
		})
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
					menu.Query = []rune{'z'}
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
					menu.Query = []rune{'z'}
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
					menu.Query = []rune{'z'}
					subMenu, err := menu.FindLastSubMenu()
					Expect(err).To(BeNil())
					Expect(subMenu.Title).To(Equal("submenu1"))
				})

			})

			Context("when it is nested", func() {
				It("returns the SubMenu", func() {
					menu.Query = []rune{'z', 'a'}
					subMenu, err := menu.FindLastSubMenu()
					Expect(err).To(BeNil())
					Expect(subMenu.Title).To(Equal("submenu1-submenu1"))
				})
			})
		})
	})

	Context("when there is no Query", func() {
		It("returns an error and an empty SubMenu", func() {
			menu.Query = []rune{}
			subMenu, err := menu.FindLastSubMenu()
			Expect(err).To(MatchError("SubMenuNotFound"))
			Expect(subMenu).To(BeEquivalentTo(SubMenu{}))
		})
	})
})
