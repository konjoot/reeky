package matchers

import (
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"

	. "github.com/konjoot/reeky/interfaces"
	. "github.com/konjoot/reeky/test/interfaces"
)

func BeFinded() *baseMatcher {
	return Matcher(&beFindedMatcher{})
}

type beFindedMatcher struct{}

func (_ *beFindedMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (_ *beFindedMatcher) Prepare(actual interface{}) interface{} {
	return actual.(Findable).Finded()
}

func (_ *beFindedMatcher) Format(actual interface{}) string {
	return actual.(Stringer).String()
}

func (_ *beFindedMatcher) Message() string {
	return "to be finded"
}

func (m *beFindedMatcher) String() (s string) {
	return
}
