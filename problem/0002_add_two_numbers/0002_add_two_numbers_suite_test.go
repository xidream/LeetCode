package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test0002AddTwoNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "0002AddTwoNumbers Suite")
}
