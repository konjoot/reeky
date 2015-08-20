package matchers

import (
	"fmt"

	ifaces "github.com/konjoot/reeky/test/interfaces"
)

func Matcher(m ifaces.MatcherIface) *baseMatcher {
	return &baseMatcher{m}
}

type baseMatcher struct{ ifaces.MatcherIface }

func (m *baseMatcher) Match(actual interface{}) (success bool, err error) {
	return m.Matcher().Match(m.Prepare(actual))
}

func (m *baseMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf(m.Template(false), m.Format(actual), m.Message())
}

func (m *baseMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf(m.Template(true), m.Format(actual), m.Message())
}

func (m *baseMatcher) Template(negate bool) (s string) {
	s = "Expected\n\t%s\n"

	if negate {
		s += "not "
	}

	s += "%s"

	if str := m.String(); len(str) > 0 {
		s += fmt.Sprintf("\n\t%s", str)
	}

	return
}
