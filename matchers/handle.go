package matchers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	"strings"
)

func Handle(method string) *handleMatcher {
	return &handleMatcher{expected: echo.Route{Method: method}}
}

type handleMatcher struct {
	expected echo.Route
}

func (m *handleMatcher) On(path string) *handleMatcher {
	m.expected.Path = path
	return m
}

func (m *handleMatcher) By(handler string) *baseMatcher {
	m.expected.Handler = "github.com/konjoot/reeky/reeky." + handler
	return Matcher(m)
}

func (m *handleMatcher) Matcher() types.GomegaMatcher {
	return &matchers.ContainElementMatcher{Element: m.expected}
}

func (m *handleMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*echo.Echo).Routes()
}

func (_ *handleMatcher) Format(actual interface{}) string {
	s := make([]string, 1)

	for _, route := range actual.(*echo.Echo).Routes() {
		s = append(s, fmt.Sprintf("%#v", route))
	}

	return "[  " + strings.Join(s, "\n\t   ") + "  ]"
}

func (_ *handleMatcher) Message() string {
	return "to have route"
}

func (m *handleMatcher) String() (s string) {
	return fmt.Sprintf("%#v", m.expected)
}
