package reeky

import (
	"fmt"
	. "github.com/konjoot/reeky/interfaces"
)

type App struct {
	routes  bool
	midware bool
	Engine  EngineIface
}

func (app *App) RunOn(port string) {
	app.Setup()
	// fmt.Printf("Reeky is running on port %s\nuse Ctrl-C for exit\n\n", port)
	app.Engine.Run(":" + port)
}

func (app *App) Setup() {
	app.midware = app.SetMiddleWare()
	app.routes = app.SetRoutes()
}

func (app *App) Ok() bool {
	return app.midware && app.routes
}

func (app *App) String() string {
	return fmt.Sprintf("App{Ok: \"%t\"}", app.Ok())
}
