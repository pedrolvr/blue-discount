package util_test

import (
	"blue-discount/pkg/util"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("moment", func() {
	tomorrow := time.Now().Add(24 * time.Hour)

	Describe("IsBirthday()", func() {
		Context("when birthday is today", func() {
			It("should return true", func() {
				Ω(util.IsBirthday(time.Now())).Should(Equal(true))
			})
		})

		Context("when birthday is not today", func() {
			It("should return false", func() {
				Ω(util.IsBirthday(tomorrow)).Should(Equal(false))
			})
		})
	})

	Describe("DateIsToday()", func() {
		Context("when date is today", func() {
			It("should return true", func() {
				Ω(util.DateIsToday(time.Now())).Should(Equal(true))
			})
		})

		Context("when date is not today", func() {
			It("should return false", func() {
				Ω(util.DateIsToday(tomorrow)).Should(Equal(false))
			})
		})
	})
})
