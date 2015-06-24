package mocks

import (
	. "github.com/konjoot/reeky/interfaces"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

type EngineMock struct {
	EngineIface
	port      string
	isRunning bool
}

func (e *EngineMock) Run(addr string) (err error) {
	e.port, e.isRunning = addr, true
	return
}

func (e *EngineMock) Port() string {
	return e.port
}

func BeRunning() *beRunningMatcher {
	return &beRunningMatcher{}
}

type beRunningMatcher struct{}

func (m *beRunningMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.BeTrueMatcher{}).Match(actual.(*EngineMock).isRunning)
}

func (m *beRunningMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be running")
}

func (m *beRunningMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be running")
}
