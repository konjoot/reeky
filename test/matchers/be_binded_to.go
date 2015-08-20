package matchers

import (
	"fmt"

	. "github.com/konjoot/reeky/interfaces"
	. "github.com/konjoot/reeky/test/interfaces"

	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

type modelIface interface {
	Bindable
	Stringer
}

func BeBindedTo(model modelIface) *baseMatcher {
	return Matcher(&beBindedToMatcher{model: model})
}

type beBindedToMatcher struct {
	model modelIface
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
