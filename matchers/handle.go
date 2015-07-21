package matchers

import (
	"github.com/labstack/echo"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

func Handle(method string) *handleMatcher {
	return &handleMatcher{expected: &echo.Route{Method: method}}
}

type handleMatcher struct {
	expected *echo.Route
}

func (m *handleMatcher) On(path string) *handleMatcher {
	m.expected.Path = path
	return m
}

func (m *handleMatcher) By(handler string) *handleMatcher {
	m.expected.Handler = "github.com/konjoot/reeky/reeky." + handler
	return m
}

func (m *handleMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.ContainElementMatcher{Element: m.expected}).Match(actual.(*echo.Echo).Routes())
}

func (m *handleMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*echo.Echo).Routes(), "to have route", m.expected)
}

func (m *handleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*echo.Echo).Routes(), "not to have route", m.expected)
}
