package common_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFormatters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Formatters")
}
