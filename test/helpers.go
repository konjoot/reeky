package test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

func Context(req *http.Request, res http.ResponseWriter, r interface{}) (c *echo.Context) {
	c = echo.NewContext(req, echo.NewResponse(res), echo.New())

	if r != nil {
		c.Set("Resource", r)
	}

	return
}

func NewJsonReader(form interface{}) io.Reader {
	jsForm, _ := json.Marshal(form)
	return bytes.NewReader(jsForm)
}
