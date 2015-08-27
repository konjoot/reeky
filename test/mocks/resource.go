package mocks

import (
	"fmt"
	"reflect"

	"github.com/konjoot/reeky/errors"

	i "github.com/konjoot/reeky/interfaces"
)

type ResourceMock struct {
	F        interface{}
	V        interface{}
	Invalid  bool
	NotFound bool

	created  bool
	findedBy string
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
	return fmt.Sprintf("ResourceMock{Invalid: %t, created: %t, findedBy: %#v, Form: %#v, View: %#v}", r.Invalid, r.created, r.findedBy, r.F, r.V)
}

func (r *ResourceMock) Url() string {
	return "some/url"
}

func (r *ResourceMock) Form() interface{} {
	return r.F
}

func (r *ResourceMock) Save() (e error) {
	if r.Invalid {
		e = errors.NewConflictError()
		return
	}

	r.created = true

	return
}

func (r *ResourceMock) FindedBy() string {
	return r.findedBy
}

func (r *ResourceMock) Finded() bool {
	return r.findedBy != ""
}

func (r *ResourceMock) Find(id string) (i.Viewer, error) {
	r.findedBy = id
	return r, nil
}

func (r *ResourceMock) View() interface{} {
	return r.V
}
