package reeky_test

import (
	. "github.com/konjoot/reeky/reeky"
	"github.com/labstack/echo"

	. "github.com/konjoot/reeky/matchers"
	. "github.com/konjoot/reeky/mocks"
	"github.com/konjoot/reeky/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Handlers", func() {
	var (
		err      error
		form     map[string]string
		context  *echo.Context
		entity   *ResourceMock
		response *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		form = map[string]string{"Name": "Test", "Desc": "TestBoard"}
		response = httptest.NewRecorder()
	})

	Describe("Creator", func() {
		JustBeforeEach(func() {
			request := http.NewRequest("POST", "/tests", test.NewJsonReader(form))
			context := test.Context(request, response, entity)
			err := Creator(context)
		})

		Describe("positive case", func() {
			BeforeEach(func() {
				entity = &ResourseMock{}
			})

			It("should create entity and return right response", func() {
				Expect(err).To(BeNil())
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).To(BeCreated())
				Expect(response.Code).To(Equal(201))
				Expect(response.Header().Get("Location")).To(Equal(entity.Url()))
				Expect(response.Body).To(BeEmpty())
			})
		})

		Describe("negative case (Conflict)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{Conflict: true}
			})

			It("should not create entity and set errors to context", func() {
				Expect(err).To(BeAssignableToTypeOf(ConflictError))
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeNil())
				Expect(response.Body).To(BeEmpty())
			})
		})

		Describe("negative case (Unprocessable Entity)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{Invalid: true}
			})

			It("should not create entity and set errors to context", func() {
				Expect(err).To(BeAssignableToTypeOf(ValidationError))
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeNil())
				Expect(response.Body).To(BeEmpty())
			})
		})

		Describe("negative case (Unsupported Media Type)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{}
			})

			It("should not create entity and set errors to context", func() {
				Expect(err).To(BeAssignableToTypeOf(echo.UnsupportedMediaType))
				Expect(form).NotTo(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeNil())
				Expect(response.Body).To(BeEmpty())
			})
		})

		Describe("negative case (Failed Dependency)", func() {
			It("should not create entity and set errors to context", func() {
				Expect(err).To(BeAssignableToTypeOf(EmptyResourceError))
				Expect(entity).To(BeNil())
				Expect(response.Code).NotTo(Equal(201))
				Expect(response.Header().Get("Location")).To(BeNil())
				Expect(response.Body).To(BeEmpty())
			})
		})
	})
})
