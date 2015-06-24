package reeky_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestReeky(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reeky Suite")
}
