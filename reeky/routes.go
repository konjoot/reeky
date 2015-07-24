package reeky

import mw "github.com/labstack/echo/middleware"

func (app *App) Setup() (ok bool) {
	// Middleware
	app.Engine.Use(mw.Logger())
	app.Engine.Use(mw.Recover())

	// Routes
	app.Engine.Get("/boards/:id", Getter)
	app.Engine.Get("/boards", ListGetter)
	app.Engine.Put("/boards/:id", Updater)
	app.Engine.Post("/boards", Creator)
	app.Engine.Delete("/boards/:id", Destroyer)
	app.Ok, ok = true, true
	return
}
