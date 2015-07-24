package mocks

import "github.com/labstack/echo"

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

func (e *EngineMock) Get(path string, h echo.Handler)    {}
func (e *EngineMock) Put(path string, h echo.Handler)    {}
func (e *EngineMock) Post(path string, h echo.Handler)   {}
func (e *EngineMock) Delete(path string, h echo.Handler) {}
func (e *EngineMock) Use(m ...echo.Middleware)           {}
