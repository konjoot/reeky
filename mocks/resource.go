package mocks

import (
	"fmt"
	"reflect"
)

type ResourceMock struct {
	Form     interface{}
	Invalid  bool
	Conflict bool
	created  bool
}

type formMap map[string]string

func (r *ResourceMock) Created() bool {
	return r.created
}

func (r *ResourceMock) BindedWith(f interface{}) (binded bool) {
	var (
		name string
		val  interface{}
	)

	if _, ok := f.(formMap); !ok {
		return
	}

	rForm := reflect.ValueOf(r.Form).Elem()

	for name, val = range f.(formMap) {
		if field := rForm.FieldByName(name); field.IsValid() && field.Interface() == val {
			continue
		}
		return
	}

	return true
}

func (r *ResourceMock) String() string {
	return fmt.Sprintf("ResourceMock{Invalid: %t, Conflict: %t, created: %t, Form: %#v}", r.Invalid, r.Conflict, r.created, r.Form)
}

func (r *ResourceMock) Url() string {
	return "some/url"
}
