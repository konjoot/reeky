package main

import (
	"github.com/gin-gonic/gin"
	"github.com/konjoot/reeky/reeky"
)

func main() {
	app := &reeky.App{Engine: gin.Default()}
	app.RunOn("8080")
}
