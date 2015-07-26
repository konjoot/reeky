package matchers

import (
	. "github.com/konjoot/reeky/mocks"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeRunning() *baseMatcher {
	return Matcher(&beRunningMatcher{})
}

type beRunningMatcher struct{}

func (m *beRunningMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (m *beRunningMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*EngineMock).IsRunning()
}

func (m *beRunningMatcher) Format(actual interface{}) string {
	return actual.(*EngineMock).String()
}

func (_ *beRunningMatcher) Message() string {
	return "to be running"
}

func (_ *beRunningMatcher) String() (s string) {
	return
}
