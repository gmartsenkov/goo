package menu

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Action", func() {
	Describe("FindByKey", func() {
		var (
			actions Actions
		)
		Context("when action exists", func() {
			BeforeEach(func() {
				actions = Actions{
					Action{
						Key: rune('b'),
					},
				}
			})
			It("returns the action", func() {
				action, err := actions.FindByKey(rune('b'))
				Expect(err).To(BeNil())
				Expect(action).To(BeEquivalentTo(Action{Key: rune('b')}))
			})
		})

		Context("when action does not exist", func() {
			It("returns an empty action and an error", func() {
				action, err := actions.FindByKey(rune('a'))
				Expect(err).To(MatchError(errors.New("ActionNotFound")))
				Expect(action).To(BeEquivalentTo(Action{}))
			})
		})
	})
})
