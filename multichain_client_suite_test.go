package multichain_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultichainClien(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultichainClient Suite")
}
