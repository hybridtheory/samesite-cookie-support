package models

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test-reports/junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "SameSite cookie support suite", []Reporter{
		junitReporter,
	})
}
