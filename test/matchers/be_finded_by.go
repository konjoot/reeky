package matchers

import (
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"

	. "github.com/konjoot/reeky/interfaces"
	. "github.com/konjoot/reeky/test/interfaces"
)

func BeFindedBy(id string) *baseMatcher {
	return Matcher(&beFindedByMatcher{id: id})
}

type beFindedByMatcher struct {
	id string
}

func (m *beFindedByMatcher) Matcher() types.GomegaMatcher {
	return &matchers.EqualMatcher{Expected: m.id}
}

func (_ *beFindedByMatcher) Prepare(actual interface{}) interface{} {
	return actual.(FindedByI).FindedBy()
}

func (_ *beFindedByMatcher) Format(actual interface{}) string {
	return actual.(Stringer).String()
}

func (_ *beFindedByMatcher) Message() string {
	return "to be finded by"
}

func (m *beFindedByMatcher) String() string {
	return m.id
}
