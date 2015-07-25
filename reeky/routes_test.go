package reeky_test

import (
	. "github.com/konjoot/reeky/reeky"
	"github.com/labstack/echo"

	. "github.com/konjoot/reeky/matchers"
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
			Expect(engine).To(Handle("GET").On("/boards/:id").By("Getter"))
			Expect(engine).To(Handle("GET").On("/boards").By("ListGetter"))
			Expect(engine).To(Handle("PUT").On("/boards/:id").By("Updater"))
			Expect(engine).To(Handle("POST").On("/boards").By("Creator"))
			Expect(engine).To(Handle("DELETE").On("/boards/:id").By("Destroyer"))
		})
	})
})
