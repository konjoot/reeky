package handlers

import (
	"net/http"

	"github.com/konjoot/reeky/errors"
	"github.com/labstack/echo"

	i "github.com/konjoot/reeky/interfaces"
)

func Create(c *echo.Context) (e error) {
	var (
		r  i.FormerSaverUrler
		ok bool
	)

	if r, ok = c.Get("resource").(i.FormerSaverUrler); !ok {
		return errors.NewEmptyResourceError()
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
