package interfaces

import (
	ifaces "github.com/konjoot/reeky/interfaces"

	"github.com/onsi/gomega/types"
)

type MatcherIface interface {
	Matcher() types.GomegaMatcher
	Prepare(actual interface{}) interface{}
	Format(actual interface{}) string
	Message() string
	ifaces.Stringer
}
