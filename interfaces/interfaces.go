package interfaces

type EngineIface interface {
	Run(addr string) (err error)
}
