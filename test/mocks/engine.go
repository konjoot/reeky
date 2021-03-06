package mocks

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/labstack/echo"
)

type EngineMock struct {
	port     string
	running  bool
	midwares []string
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

func (e *EngineMock) Use(m ...echo.Middleware) {
	for _, h := range m {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		e.midwares = append(e.midwares, name)
	}
}

func (e *EngineMock) MiddleWares() []string {
	return e.midwares
}

func (e *EngineMock) String() string {
	return fmt.Sprintf("EngineMock{running: %t, port: \"%s\"}", e.running, e.port)
}

func (e *EngineMock) Get(path string, h echo.Handler)    {}
func (e *EngineMock) Put(path string, h echo.Handler)    {}
func (e *EngineMock) Post(path string, h echo.Handler)   {}
func (e *EngineMock) Delete(path string, h echo.Handler) {}
