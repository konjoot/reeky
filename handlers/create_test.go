package handlers_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/konjoot/reeky/test"

	. "github.com/konjoot/reeky/errors"
	. "github.com/konjoot/reeky/handlers"
	. "github.com/konjoot/reeky/test/matchers"
	. "github.com/konjoot/reeky/test/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {
	var (
		err      error
		body     io.Reader
		fMap     map[string]string
		form     interface{}
		entity   *ResourceMock
		response *httptest.ResponseRecorder
	)

	Describe("Create", func() {
		BeforeEach(func() {
			fMap = map[string]string{"Name": "Test", "Desc": "TestDesc"}
			body = test.NewJsonReader(fMap)
			form = &test.Form{}
			response = httptest.NewRecorder()
		})

		JustBeforeEach(func() {
			request, _ := http.NewRequest("POST", "/tests", body)
			context := test.Context(request, response, entity)
			err = Create(context)
		})

		Describe("positive case", func() {
			BeforeEach(func() {
				entity = &ResourceMock{F: form}
			})

			It("should create entity and return right response", func() {
				Expect(err).To(BeNil())
				Expect(fMap).To(BeBindedTo(entity))
				Expect(entity).To(BeCreated())
				Expect(response.Code).To(Equal(201))
				Expect(response.Header().Get("Location")).To(Equal(entity.Url()))
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Error while saving)", func() {
			BeforeEach(func() {
				entity = &ResourceMock{F: form, Invalid: true}
			})

			It("should not create entity and return ConflictError", func() {
				Expect(err).To(BeTypeOf(NewConflictError()))
				Expect(fMap).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Syntax error)", func() {
			BeforeEach(func() {
				body = test.NewStringReader("bad request")
				entity = &ResourceMock{F: form}
			})

			It("should not create entity and return SyntaxError error", func() {
				Expect(err).To(BeTypeOf(&json.SyntaxError{}))
				Expect(fMap).NotTo(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (UnsupportedMediaType error)", func() {
			JustBeforeEach(func() {
				response = httptest.NewRecorder()
				entity = &ResourceMock{F: form}
				request, _ := http.NewRequest("POST", "/tests", body)
				context := test.BadContext(request, response, entity)
				err = Create(context)
			})

			It("should not create entity and return UnsupportedMediaType error", func() {
				Expect(err).To(Equal(errors.New("echo ⇒ unsupported media type")))
				Expect(fMap).NotTo(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).NotTo(Equal(entity.Url()))
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Failed Dependency)", func() {
			JustBeforeEach(func() {
				response = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/tests", body)
				context := test.Context(request, response, nil)
				err = Create(context)
			})

			It("should not create entity and return EmptyResourceError", func() {
				Expect(err).To(BeTypeOf(NewEmptyResourceError()))
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})
	})
})
