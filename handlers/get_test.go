package handlers_test

import (
	"errors"
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
		id       string
		view     interface{}
		entity   *ResourceMock
		response *httptest.ResponseRecorder
	)

	Describe("Get", func() {
		BeforeEach(func() {
			id = "2"
			view = &test.Form{"Test", "TestDesc"}
			response = httptest.NewRecorder()
		})

		JustBeforeEach(func() {
			request, _ := http.NewRequest("GET", "/tests/"+id, nil)
			context := test.Context(request, response, entity)
			err = Get(context)
		})

		Describe("positive case", func() {
			BeforeEach(func() {
				entity = &ResourceMock{V: view}
			})

			It("should find and render entity", func() {
				Expect(err).To(BeNil())
				Expect(entity).To(BeFindedBy(id))
				Expect(response.Code).To(Equal(200))
				Expect(response.Body.String()).To(MatchJSON(`{"Name":"Test", "Desc":"TestDesc"}`))
			})
		})

		Describe("negative case (NotFound error)", func() {
			BeforeEach(func() {
				entity = &ResourceMock{NotFound: true}
			})

			It("should not render entity and return NotFoundError", func() {
				Expect(err).To(BeTypeOf(NewNotFoundError()))
				Expect(entity).NotTo(BeFinded())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (UnsupportedMediaType error)", func() {
			JustBeforeEach(func() {
				response = httptest.NewRecorder()
				entity = &ResourceMock{V: view}
				request, _ := http.NewRequest("GET", "/tests/"+id, nil)
				context := test.BadContext(request, response, entity)
				err = Get(context)
			})

			It("should not render entity and return UnsupportedMediaType error", func() {
				Expect(err).To(Equal(errors.New("echo ⇒ unsupported media type")))
				Expect(entity).NotTo(BeFinded())
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Describe("negative case (Failed Dependency)", func() {
			JustBeforeEach(func() {
				response = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/tests/"+id, nil)
				context := test.Context(request, response, nil)
				err = Get(context)
			})

			It("should not render entity and return EmptyResourceError", func() {
				Expect(err).To(BeTypeOf(NewEmptyResourceError()))
				Expect(entity).NotTo(BeFinded())
				Expect(response.Body.Len()).To(BeZero())
			})
		})
	})
})
