package reeky

import h "github.com/konjoot/reeky/handlers"

func (app *App) SetRoutes() bool {
	app.Engine.Get("/boards/:id", h.Getter)
	app.Engine.Get("/boards", h.ListGetter)
	app.Engine.Put("/boards/:id", h.Updater)
	app.Engine.Post("/boards", h.Creator)
	app.Engine.Delete("/boards/:id", h.Destroyer)
	return true
}
