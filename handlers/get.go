package handlers

import (
	"fmt"
	"net/http"

	"github.com/konjoot/reeky/errors"
	"github.com/labstack/echo"

	i "github.com/konjoot/reeky/interfaces"
)

func Get(c *echo.Context) (e error) {
	var (
		r  i.Finder
		v  i.Viewer
		ok bool
	)

	if r, ok = c.Get("resource").(i.Finder); !ok {
		return errors.NewEmptyResourceError()
	}

	fmt.Printf("%#v\n", c.P(0))
	if v, e = r.Find(c.P(0)); e != nil {
		return
	}

	if e = c.JSON(http.StatusOK, v.View()); e != nil {
		return
	}

	return
}
