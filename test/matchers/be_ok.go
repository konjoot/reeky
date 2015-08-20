package matchers

import (
	. "github.com/konjoot/reeky/reeky"

	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeOk() *baseMatcher {
	return Matcher(&beOkMatcher{})
}

type beOkMatcher struct{}

func (m *beOkMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (m *beOkMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*App).Ok()
}

func (m *beOkMatcher) Format(actual interface{}) string {
	return actual.(*App).String()
}

func (_ *beOkMatcher) Message() string {
	return "to be Ok"
}

func (_ *beOkMatcher) String() (s string) {
	return
}
