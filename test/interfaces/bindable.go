package interfaces

type Bindable interface {
	BindedWith(interface{}) bool
}
