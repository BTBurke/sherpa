package sherpa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSherpa(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sherpa Suite")
}
