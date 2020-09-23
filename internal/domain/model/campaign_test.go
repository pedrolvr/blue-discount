package model_test

import (
	"blue-discount/internal/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("campaign test", func() {
	Describe("NewCampaign()", func() {
		It("should get campaign", func() {
			percent := int32(5)
			name := model.BirthdayCampaignName
			product := model.NewCampaign(name, percent, true, nil)

			Ω(product.Name).Should(Equal(name))
			Ω(product.Percent).Should(Equal(percent))
			Ω(product.Active).Should(Equal(true))
			Ω(product.AppliedAt).Should(BeNil())
		})
	})
})
