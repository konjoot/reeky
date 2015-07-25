package matchers

import (
	"github.com/labstack/echo"
	"reflect"
	"runtime"

	. "github.com/konjoot/reeky/mocks"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

func UseMiddleWare(midware echo.MiddlewareFunc) *useMiddleWareMatcher {
	name := runtime.FuncForPC(reflect.ValueOf(midware).Pointer()).Name()
	return &useMiddleWareMatcher{midware: name}
}

type useMiddleWareMatcher struct {
	midware string
}

func (m *useMiddleWareMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.ContainElementMatcher{Element: m.midware}).Match(actual.(*EngineMock).MiddleWares())
}

func (m *useMiddleWareMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*EngineMock).MiddleWares(), "to have middleware", m.midware)
}

func (m *useMiddleWareMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*EngineMock).MiddleWares(), "not to have middleware", m.midware)
}
