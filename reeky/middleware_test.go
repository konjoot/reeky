package reeky_test

import (
	. "github.com/konjoot/reeky/reeky"
	. "github.com/konjoot/reeky/test/matchers"
	. "github.com/konjoot/reeky/test/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	mw "github.com/labstack/echo/middleware"
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
