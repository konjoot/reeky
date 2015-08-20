package reeky_test

import (
	. "github.com/konjoot/reeky/matchers"
	. "github.com/konjoot/reeky/mocks"
	. "github.com/konjoot/reeky/reeky"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reeky", func() {
	var (
		app    *App
		engine *EngineMock
		port   string
	)

	BeforeEach(func() {
		port = "8080"
		engine = &EngineMock{}
		app = &App{Engine: engine}
	})

	Describe("RunOn", func() {
		It("should run engine on specified port", func() {
			Expect(engine).NotTo(BeRunning())
			Expect(app).NotTo(BeOk())
			Expect(engine.Port()).To(BeZero())

			app.RunOn(port)

			Expect(engine).To(BeRunning())
			Expect(app).To(BeOk())
			Expect(engine.Port()).To(Equal(":" + port))
		})
	})
})
