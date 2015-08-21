package interfaces

type ResourceIface interface {
	Form() interface{}
	Save() error
	Url() string
}
