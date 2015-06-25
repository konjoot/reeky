package reeky

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
}

func (app *App) RunOn(port string) {
	app.Engine.Run(":" + port)
}

func (app *App) SetRoutes() (ok bool) {
	return
}
