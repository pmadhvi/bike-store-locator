package external_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Setup the test suite for testing external api's
func TestExternal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "External Test Suite")
}
