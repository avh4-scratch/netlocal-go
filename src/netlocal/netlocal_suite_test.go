package allup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAllup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Netlocal Suite")
}
