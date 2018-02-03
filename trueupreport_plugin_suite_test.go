package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTrueupreportPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TrueupreportPlugin Suite")
}
