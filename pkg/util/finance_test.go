package util_test

import (
	"blue-discount/pkg/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("finance", func() {
	Describe("DiscountInCents()", func() {
		Context("when value is 10000 in cents", func() {
			Context("and percentage is 5%", func() {
				It("should get 500 in cents", func() {
					Ω(util.DiscountInCents(10000, 5)).Should(Equal(int64(500)))
				})
			})

			Context("and percentage is 35%", func() {
				It("should get 3500 in cents", func() {
					Ω(util.DiscountInCents(10000, 35)).Should(Equal(int64(3500)))
				})
			})
		})

		Context("when value is 3500", func() {
			Context("and percentage is 7%", func() {
				It("should get 245 in cents", func() {
					Ω(util.DiscountInCents(3500, 7)).Should(Equal(int64(245)))
				})
			})
		})

		Context("when value is -4700 in cents", func() {
			Context("and percentage is 27%", func() {
				It("should get 1269 in cents", func() {
					Ω(util.DiscountInCents(4700, 27)).Should(Equal(int64(1269)))
				})
			})
		})
	})
})
