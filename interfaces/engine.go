package interfaces

import (
	"github.com/labstack/echo"
)

type EngineIface interface {
	Run(addr string)
	Use(m ...echo.Middleware)
	Get(path string, h echo.Handler)
	Put(path string, h echo.Handler)
	Post(path string, h echo.Handler)
	Delete(path string, h echo.Handler)
}
