package reeky_test

import (
	. "github.com/konjoot/reeky/reeky"
	mw "github.com/labstack/echo/middleware"

	. "github.com/konjoot/reeky/matchers"
	. "github.com/konjoot/reeky/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var (
		app    *App
		engine *EngineMock
	)

	BeforeEach(func() {
		engine = &EngineMock{}
		app = &App{Engine: engine}
		app.SetMiddleWare()
	})

	It("should use expected middleware", func() {
		Expect(engine).To(UseMiddleWare(mw.Logger()))
		Expect(engine).To(UseMiddleWare(mw.Recover()))
	})
})
