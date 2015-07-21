package main

import (
	"github.com/konjoot/reeky/reeky"
	"github.com/labstack/echo"
)

func main() {
	app := &reeky.App{Engine: echo.New()}
	app.RunOn("8080")
}
