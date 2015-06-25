package matchers

import (
	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

func Handle(method string) *handleMatcher {
	return &handleMatcher{expected: gin.RouteInfo{Method: method}}
}

type handleMatcher struct {
	expected gin.RouteInfo
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
	return (&matchers.ContainElementMatcher{Element: m.expected}).Match(actual.(*gin.Engine).Routes())
}

func (m *handleMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*gin.Engine).Routes(), "to have route", m.expected)
}

func (m *handleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*gin.Engine).Routes(), "not to have route", m.expected)
}
