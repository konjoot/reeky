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
		context  *echo.Context
		entity   *ResourceMock
		request  *http.Request
		response *httptest.ResponseRecorder
		form     map[string]string
	)

	BeforeEach(func() {
		form = map[string]string{"Name": "Test", "Desc": "TestBoard"}
		request = test.JsonRequest("POST", "/boards", form)
		response = test.Response()
	})

	Describe("Creator", func() {
		JustBeforeEach(func() {
			context = test.Context(request, response, entity)
			Creator(context)
		})

		Describe("positive case", func() {
			BeforeEach(func() {
				entity = &ResourseMock{}
			})

			It("should create entity and return right response", func() {
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).To(BeCreated())
				Expect(response).To(HaveHeader("Location").WithUrlFor(entity))
				Expect(response).To(HaveEmptyBody())
			})
		})

		Describe("negative case (conflict)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{Conflict: true}
			})

			It("should not create entity and return right error message", func() {
				Expect(form).To(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response).NotTo(HaveHeader("Location"))
				Expect(response).To(HaveEmptyBody())
			})
		})

		Describe("negative case (invalid params)", func() {
			BeforeEach(func() {
				entity = &ResourseMock{Invalid: true}
			})

			It("should not create entity and return right error message", func() {
				Expect(form).NotTo(BeBindedTo(entity))
				Expect(entity).NotTo(BeCreated())
				Expect(response).NotTo(HaveHeader("Location"))
				Expect(response).To(HaveEmptyBody())
			})
		})
	})
})
