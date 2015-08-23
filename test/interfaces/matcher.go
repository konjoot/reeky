package interfaces

import (
	"github.com/onsi/gomega/types"

	i "github.com/konjoot/reeky/interfaces"
)

type MatcherIface interface {
	Matcher() types.GomegaMatcher
	Prepare(actual interface{}) interface{}
	Format(actual interface{}) string
	Message() string
	i.Stringer
}
