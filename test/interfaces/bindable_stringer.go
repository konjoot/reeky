package interfaces

import ifaces "github.com/konjoot/reeky/interfaces"

type BindableStringer interface {
	Bindable
	ifaces.Stringer
}
