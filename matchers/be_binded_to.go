package matchers

import (
	. "github.com/konjoot/reeky/mocks"
	. "github.com/konjoot/reeky/mocks/interfaces"

	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeBindedTo() *beBindedToMatcher {
	return Matcher(&beBindedToMatcher{})
}

type beBindedToMatcher struct{}

func (m *beBindedToMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (m *beBindedToMatcher) Prepare(actual interface{}) interface{} {
	return actual.(Bindable).Binded()
}

func (m *beBindedToMatcher) Format(actual interface{}) string {
	return actual.(Stringer).String()
}

func (_ *beBindedToMatcher) Message() string {
	return "to be binded"
}

func (_ *beBindedToMatcher) String() (s string) {
	return
}
