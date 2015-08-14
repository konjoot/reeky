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
				Expect(response).To(HaveStatus("201"))
				Expect(response).To(HaveHeader("Location").WithUrlFor(entity))
				Expect(response).To(HaveEmptyBody())
			})
		})

		Describe("negative case (Conflict)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{Conflict: true}
			})

			It("should not create entity and set errors to context", func() {
				Expect(err).To(HaveType(ConflictError))
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response).NotTo(HaveStatus("201"))
				Expect(response).NotTo(HaveHeader("Location"))
				Expect(response).To(HaveEmptyBody())
			})
		})

		Describe("negative case (Unprocessable Entity)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{Invalid: true}
			})

			It("should not create entity and set errors to context", func() {
				Expect(err).To(HaveType(ValidationError))
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response).NotTo(HaveStatus("201"))
				Expect(response).NotTo(HaveHeader("Location"))
				Expect(response).To(HaveEmptyBody())
			})
		})

		Describe("negative case (Unsupported Media Type)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{}
			})

			It("should not create entity and set errors to context", func() {
				Expect(err).To(HaveType(echo.UnsupportedMediaType))
				Expect(form).NotTo(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response).NotTo(HaveStatus("201"))
				Expect(response).NotTo(HaveHeader("Location"))
				Expect(response).To(HaveEmptyBody())
				Expect(context).To(HaveErrors())
			})
		})

		Describe("negative case (Failed Dependency)", func() {
			It("should not create entity and set errors to context", func() {
				Expect(err).To(HaveType(EmptyResourceError))
				Expect(entity).To(BeNil())
				Expect(response).NotTo(HaveStatus("201"))
				Expect(response).NotTo(HaveHeader("Location"))
				Expect(response).To(HaveEmptyBody())
			})
		})
	})
})
