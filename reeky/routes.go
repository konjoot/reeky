package reeky

func (app *App) SetRoutes() bool {
	app.Engine.Get("/boards/:id", Getter)
	app.Engine.Get("/boards", ListGetter)
	app.Engine.Put("/boards/:id", Updater)
	app.Engine.Post("/boards", Creator)
	app.Engine.Delete("/boards/:id", Destroyer)
	return true
}
