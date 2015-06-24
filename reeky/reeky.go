package reeky

import (
	. "github.com/konjoot/reeky/interfaces"
)

type App struct {
	Engine EngineIface
}

func (app *App) RunOn(port string) {
	app.Engine.Run(":" + port)
}
