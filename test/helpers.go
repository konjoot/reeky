package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type testForm struct {
	Name string
	Desc string
}

func Context(req *http.Request, res http.ResponseWriter, r interface{}) (c *echo.Context) {
	req.Header.Set("Content-Type", "application/json")

	c = echo.NewContext(req, echo.NewResponse(res), echo.New())

	if r != nil {
		c.Set("resource", r)
	}

	return
}

func NewJsonReader(form interface{}) io.Reader {
	jsForm, _ := json.Marshal(form)
	return bytes.NewReader(jsForm)
}

func NewStringReader(s string) io.Reader {
	return strings.NewReader(s)
}

func Form() *testForm {
	return &testForm{}
}
