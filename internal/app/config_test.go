package app_test

import (
	"blue-discount/internal/app"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configPath = "../../config/"

var _ = Describe("config", func() {
	Describe("ReadConfig()", func() {
		Context("when reading the config", func() {
			Context("and the path is correct", func() {
				It("should get a right config object", func() {
					c, err := app.ReadConfig("app", configPath)
					Ω(c.Discount.MaxApplied).Should(Equal(int32(10)))
					Ω(c.DB.Host).Should(Equal("localhost"))
					Ω(err).Should(BeNil())
				})
			})

			Context("and the config file is invalid", func() {
				It("should get a right config object", func() {
					_, err := app.ReadConfig("invalid-file", configPath)
					Ω(err.Error()).Should(ContainSubstring("error config file"))
				})
			})

			Context("and the config file is invalid", func() {
				It("should get a right config object", func() {
					os.Setenv("DISCOUNT.MAX_APPLIED", "false")

					_, err := app.ReadConfig("app", configPath)

					Ω(err).ShouldNot(BeNil())
					Ω(err.Error()).Should(ContainSubstring("unable to decode config"))
				})
			})
		})
	})
})
