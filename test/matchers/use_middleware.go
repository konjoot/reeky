package matchers

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	. "github.com/konjoot/reeky/test/mocks"

	"github.com/labstack/echo"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
)

func UseMiddleWare(midware echo.MiddlewareFunc) *baseMatcher {
	name := runtime.FuncForPC(reflect.ValueOf(midware).Pointer()).Name()
	return Matcher(&useMiddleWareMatcher{midware: name})
}

type useMiddleWareMatcher struct {
	midware string
}

func (m *useMiddleWareMatcher) Matcher() types.GomegaMatcher {
	return &matchers.ContainElementMatcher{Element: m.midware}
}

func (_ *useMiddleWareMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*EngineMock).MiddleWares()
}

func (_ *useMiddleWareMatcher) Format(actual interface{}) string {
	return "[  " + strings.Join(actual.(*EngineMock).MiddleWares(), "\n\t   ") + "  ]"
}

func (_ *useMiddleWareMatcher) Message() string {
	return "to have middleware"
}

func (m *useMiddleWareMatcher) String() (s string) {
	return fmt.Sprintf("%#v", m.midware)
}
