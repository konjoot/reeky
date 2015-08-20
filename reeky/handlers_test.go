package reeky_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/konjoot/reeky/errors"
	. "github.com/konjoot/reeky/reeky"
	. "github.com/konjoot/reeky/test/matchers"
	. "github.com/konjoot/reeky/test/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/konjoot/reeky/test"
	"github.com/labstack/echo"
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

	BeforeEach(func() {
		fMap = map[string]string{"Name": "Test", "Desc": "TestDesc"}
		body = test.NewJsonReader(fMap)
		form = test.Form()
		response = httptest.NewRecorder()
	})

	Describe("Creator", func() {
		JustBeforeEach(func() {
			request, _ := http.NewRequest("POST", "/tests", body)
			context := test.Context(request, response, entity)
			err = Creator(context)
		})

		Describe("positive case", func() {
			BeforeEach(func() {
				entity = &ResourceMock{Form: form}
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

		Describe("negative case (Conflict)", func() {
			BeforeEach(func() {
				entity = &ResourceMock{Form: form, Conflict: true}
			})

			It("should not create entity and return ConflictError", func() {
				Expect(err).To(BeTypeOf(ConflictError{}))
				Expect(fMap).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Unprocessable Entity)", func() {
			BeforeEach(func() {
				entity = &ResourceMock{Form: form, Invalid: true}
			})

			It("should not create entity and return ValidationError", func() {
				Expect(err).To(BeTypeOf(ValidationError{}))
				Expect(fMap).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Unsupported Media Type)", func() {
			BeforeEach(func() {
				body = test.NewStringReader("bad request")
				entity = &ResourceMock{Form: form}
			})

			It("should not create entity and return UnsupportedMediaType error", func() {
				Expect(err).To(BeTypeOf(echo.UnsupportedMediaType))
				Expect(fMap).NotTo(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Failed Dependency)", func() {
			It("should not create entity and return EmptyResourceError", func() {
				Expect(err).To(BeTypeOf(EmptyResourceError{}))
				Expect(entity).To(BeNil())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeEmpty())
				Expect(response.Body.Len()).To(BeZero())
			})
		})
	})
})
