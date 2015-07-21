package matchers

import (
	. "github.com/konjoot/reeky/reeky"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

func BeOk() *beOkMatcher {
	return &beOkMatcher{}
}

type beOkMatcher struct{}

func (m *beOkMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.BeTrueMatcher{}).Match(actual.(*App).Ok)
}

func (m *beOkMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have status Ok")
}

func (m *beOkMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have status Ok")
}
