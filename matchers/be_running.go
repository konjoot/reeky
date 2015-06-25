package matchers

import (
	. "github.com/konjoot/reeky/mocks"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

func BeRunning() *beRunningMatcher {
	return &beRunningMatcher{}
}

type beRunningMatcher struct{}

func (m *beRunningMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.BeTrueMatcher{}).Match(actual.(*EngineMock).IsRunning())
}

func (m *beRunningMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be running")
}

func (m *beRunningMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be running")
}
