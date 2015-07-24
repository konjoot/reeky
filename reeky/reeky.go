package reeky

import (
	// "fmt"
	. "github.com/konjoot/reeky/interfaces"
)

type App struct {
	Ok     bool
	Engine EngineIface
}

func (app *App) RunOn(port string) {
	app.Setup()
	// fmt.Printf("Reeky is running on port %s\nuse Ctrl-C for exit\n\n", port)
	app.Engine.Run(":" + port)
}
