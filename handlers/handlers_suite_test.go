package handlers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReeky(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers")
}
