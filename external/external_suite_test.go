package external_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Setup the test suite for testing Geocoding Api external request
func TestExternal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "External Suite")
}
