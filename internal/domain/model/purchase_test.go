package model_test

import (
	"blue-discount/internal/domain/model"
	"time"

	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("purchase test", func() {
	Describe("NewPurchase()", func() {
		It("should get purchase", func() {
			now := time.Now()
			userID := uuid.Must(uuid.NewV4())
			user := model.NewUser(userID, &now)

			price := int64(3700)
			productID := uuid.Must(uuid.NewV4())
			product := model.NewProduct(productID, price)

			campaigns := []model.Campaign{
				model.NewCampaign(model.BirthdayCampaignName,
					birthdayPercent, true, &now),
			}

			discountCalc := model.NewDiscountCalculator(campaigns)

			purchase := model.NewPurchase(user, product, discountCalc)

			purchase.CalculateDiscount(maxDiscount)

			Ω(purchase.User).Should(Equal(user))
			Ω(purchase.Product).Should(Equal(product))
			Ω(purchase.Discount.Value).Should(Equal(int64(185)))
		})
	})
})
