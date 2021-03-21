package consumer

import (
	gin "github.com/onsi/ginkgo"
	gom "github.com/onsi/gomega"
	"testing"
)

func TestFirstConsumer(t *testing.T) {
	//When a Gomega assertion fails, Gomega  calls a GomegaFailHandler.
	//This is a function that you must provide using gomega.RegisterFailHandler().
	gom.RegisterFailHandler(gin.Fail)

	//RunSpecs(t *testing.T, suiteDescription string) tells Ginkgo to start the test suite.
	gin.RunSpecs(t, "First Consumer Suite")
}

var _ = gin.BeforeSuite(func() {
	// before tests call
})

var _ = gin.AfterSuite(func() {
	// after all test run
})
