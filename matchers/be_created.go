package matchers

import (
	. "github.com/konjoot/reeky/test/interfaces"

	"fmt"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeCreated() *baseMatcher {
	return Matcher(&beCreatedMatcher{})
}

type beCreatedMatcher struct{}

func (m *beCreatedMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (m *beCreatedMatcher) Prepare(actual interface{}) interface{} {
	return actual.(Creatable).Created()
}

func (m *beCreatedMatcher) Format(actual interface{}) string {
	return fmt.Sprintf("%v", actual)
}

func (_ *beCreatedMatcher) Message() string {
	return "to be created"
}

func (_ *beCreatedMatcher) String() (s string) {
	return
}
