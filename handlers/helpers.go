package handlers

import (
	"github.com/konjoot/reeky/errors"
	"github.com/labstack/echo"

	i "github.com/konjoot/reeky/interfaces"
)

func Resource(c *echo.Context) (r i.ResourceIface, e error) {
	var ok bool

	if r, ok = c.Get("resource").(i.ResourceIface); !ok {
		e = errors.NewEmptyResourceError()
	}

	return
}
