package matchers

import (
	. "github.com/konjoot/reeky/test/interfaces"

	"fmt"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeBindedTo(model BindableStringer) *baseMatcher {
	return Matcher(&beBindedToMatcher{model: model})
}

type beBindedToMatcher struct {
	model BindableStringer
}

func (_ *beBindedToMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (m *beBindedToMatcher) Prepare(actual interface{}) interface{} {
	return m.model.BindedWith(actual)
}

func (_ *beBindedToMatcher) Format(actual interface{}) string {
	return fmt.Sprintf("%v", actual)
}

func (_ *beBindedToMatcher) Message() string {
	return "to be binded to"
}

func (m *beBindedToMatcher) String() string {
	return fmt.Sprintf("%v", m.model)
}
