package reeky

import h "github.com/konjoot/reeky/handlers"

func (app *App) SetRoutes() bool {
	app.Engine.Get("/boards/:id", h.Get)
	app.Engine.Get("/boards", h.List)
	app.Engine.Put("/boards/:id", h.Update)
	app.Engine.Post("/boards", h.Create)
	app.Engine.Delete("/boards/:id", h.Destroy)
	return true
}
