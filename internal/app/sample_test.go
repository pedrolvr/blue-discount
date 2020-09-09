package app_test

import (
	"blue-discount/internal/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sample", func() {
	It("anything", func() {
		Ω(app.Any()).Should(Equal("Anything"))
	})
})
