package mocks

type EngineMock struct {
	port    string
	running bool
}

func (e *EngineMock) Run(addr string) {
	e.port, e.running = addr, true
}

func (e *EngineMock) Port() string {
	return e.port
}

func (e *EngineMock) IsRunning() bool {
	return e.running
}
