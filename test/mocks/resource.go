package mocks

import (
	"fmt"
	"reflect"
)

type ResourceMock struct {
	F        interface{}
	Invalid  bool
	Conflict bool

	created bool
}

func (r *ResourceMock) Created() bool {
	return r.created
}

func (r *ResourceMock) BindedWith(f interface{}) (binded bool) {
	var (
		name string
		val  interface{}
	)

	if _, ok := f.(map[string]string); !ok {
		return
	}

	rForm := reflect.ValueOf(r.F).Elem()

	for name, val = range f.(map[string]string) {
		if field := rForm.FieldByName(name); field.IsValid() && field.Interface() == val {
			continue
		}
		return
	}

	return true
}

func (r *ResourceMock) String() string {
	return fmt.Sprintf("ResourceMock{Invalid: %t, Conflict: %t, created: %t, Form: %#v}", r.Invalid, r.Conflict, r.created, r.F)
}

func (r *ResourceMock) Url() string {
	return "some/url"
}

func (r *ResourceMock) Form() interface{} {
	return r.F
}

func (r *ResourceMock) Save() error {
	r.created = !(r.Conflict && r.Invalid)

	return nil
}
