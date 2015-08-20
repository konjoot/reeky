package matchers

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func BeTypeOf(ex interface{}) *baseMatcher {
	return Matcher(&beTypeOfMatcher{expected: ex})
}

type beTypeOfMatcher struct {
	expected interface{}
}

func (m *beTypeOfMatcher) Matcher() types.GomegaMatcher {
	return &matchers.EqualMatcher{Expected: reflect.TypeOf(m.expected)}
}

func (_ *beTypeOfMatcher) Prepare(actual interface{}) interface{} {
	return reflect.TypeOf(actual)
}

func (_ *beTypeOfMatcher) Format(actual interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(actual))
}

func (_ *beTypeOfMatcher) Message() string {
	return "to be type of"
}

func (m *beTypeOfMatcher) String() string {
	return fmt.Sprintf("%v", reflect.TypeOf(m.expected))
}
