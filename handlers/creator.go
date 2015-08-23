package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	i "github.com/konjoot/reeky/interfaces"
)

func Creator(c *echo.Context) (e error) {
	var r i.ResourceIface

	if r, e = Resource(c); e != nil {
		return
	}

	if e = c.Bind(r.Form()); e != nil {
		return
	}

	if e = r.Save(); e != nil {
		return
	}

	c.Response().Header().Set("Location", r.Url())

	if e = c.NoContent(http.StatusCreated); e != nil {
		return
	}

	return
}
