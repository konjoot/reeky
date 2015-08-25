package reeky_test

import (
	"github.com/labstack/echo"

	. "github.com/konjoot/reeky/reeky"
	. "github.com/konjoot/reeky/test/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var (
		app    *App
		engine *echo.Echo
	)

	BeforeEach(func() {
		engine = echo.New()
		app = &App{Engine: engine}
		app.SetRoutes()
	})

	Describe("Routes", func() {
		It("/boards", func() {
			Expect(engine).To(Handle("GET").On("/boards/:id").By("Get"))
			Expect(engine).To(Handle("GET").On("/boards").By("List"))
			Expect(engine).To(Handle("PUT").On("/boards/:id").By("Update"))
			Expect(engine).To(Handle("POST").On("/boards").By("Create"))
			Expect(engine).To(Handle("DELETE").On("/boards/:id").By("Destroy"))
		})
	})
})
