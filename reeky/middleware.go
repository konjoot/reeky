package reeky

import mw "github.com/labstack/echo/middleware"

func (app *App) SetMiddleWare() bool {
	app.Engine.Use(mw.Logger())
	app.Engine.Use(mw.Recover())
	return true
}
