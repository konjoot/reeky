package reeky

import (
	ifaces "github.com/konjoot/reeky/interfaces"
)

type App struct {
	Ok     bool
	Engine ifaces.EngineIface
}

func (app *App) RunOn(port string) {
	app.Setup()
	app.Engine.Run(":" + port)
}

func (app *App) Setup() (ok bool) {
	app.Ok, ok = true, true
	return
}
