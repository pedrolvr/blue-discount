package model_test

import (
	"blue-discount/internal/domain/model"
	"time"

	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	userID             = "e4174ef5-fe62-4f9e-9a55-377bd339e318"
	productID          = "15b69bd4-0357-45ed-96ca-62c308ed77e8"
	blackFridayPercent = int32(10)
	birthdayPercent    = int32(5)
	maxDiscount        = int32(10)
)

func newBirthdayCampaign(active bool) model.Campaign {
	return model.NewCampaign(model.BirthdayCampaignName, birthdayPercent, active, nil)
}

func newBlackFridayCampaign(active bool, appliedAt *time.Time) model.Campaign {
	return model.NewCampaign(model.BlackFridayCampaignName, blackFridayPercent, active, appliedAt)
}

func newNotImplementedCampaign(active bool) model.Campaign {
	return model.NewCampaign("not-implemented-campaign", 15, active, nil)
}

var _ = Describe("discount test", func() {
	userUID := uuid.Must(uuid.FromString(userID))
	productUID := uuid.Must(uuid.FromString(productID))

	var (
		now      time.Time
		tomorrow time.Time
	)

	BeforeEach(func() {
		now = time.Now()
		tomorrow = now.Add(24 * time.Hour)
	})

	Describe("DiscountStrategyFactory()", func() {
		Context("when campaign is birthday", func() {
			It("get birthday discount strategy", func() {
				campaign := newBirthdayCampaign(true)
				strategy := model.DiscountStrategyFactory(campaign)

				Ω(strategy.(*model.BirthdayDiscount).Campaign).Should(Equal(campaign))
			})
		})

		Context("when campaign is black friday", func() {
			It("get black friday discount strategy", func() {
				campaign := newBlackFridayCampaign(true, nil)
				strategy := model.DiscountStrategyFactory(campaign)

				Ω(strategy.(*model.OnDateDiscount).Campaign).Should(Equal(campaign))
			})
		})

		Context("when not implemented campaign", func() {
			It("get nil", func() {
				campaign := newNotImplementedCampaign(true)
				strategy := model.DiscountStrategyFactory(campaign)

				Ω(strategy).Should(BeNil())
			})
		})
	})

	Describe("Calculate()", func() {
		Context("when product price is 10000", func() {
			product := model.NewProduct(productUID, 10000)

			Context("and is user`s birthday", func() {
				Context("and the campaign is not enabled", func() {
					It("should get 0 of discount in cents", func() {
						campaigns := []model.Campaign{
							newBirthdayCampaign(false),
						}

						user := model.NewUser(userUID, &now)

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscount, user, product)

						Ω(discount.Value).Should(Equal(int64(0)))
					})
				})

				It("should get 500 of discount in cents", func() {
					campaigns := []model.Campaign{
						newBirthdayCampaign(true),
					}

					user := model.NewUser(userUID, &now)

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(500)))
				})
			})

			Context("and is not user`s birthday", func() {
				It("should get 0 of discount", func() {
					campaigns := []model.Campaign{
						newBirthdayCampaign(true),
					}

					user := model.NewUser(userUID, &tomorrow)

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})
		})

		Context("when product price is 10000", func() {
			product := model.NewProduct(productUID, 10000)

			Context("and is black friday", func() {
				Context("and the campaign is not enabled", func() {
					It("should get 0 of discount in cents", func() {
						user := model.NewUser(userUID, nil)

						campaigns := []model.Campaign{
							newBlackFridayCampaign(false, &now),
						}

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscount, user, product)

						Ω(discount.Value).Should(Equal(int64(0)))
						Ω(discount.Percentage).Should(Equal(int32(0)))
					})
				})

				It("should get 1000 of discount in cents", func() {
					user := model.NewUser(userUID, nil)

					campaigns := []model.Campaign{
						newBlackFridayCampaign(true, &now),
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(1000)))
					Ω(discount.Percentage).Should(Equal(blackFridayPercent))
				})
			})

			Context("and is not black friday", func() {
				It("should get 1000 of discount in cents", func() {
					user := model.NewUser(userUID, nil)

					campaigns := []model.Campaign{
						newBlackFridayCampaign(true, &tomorrow),
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})
		})

		Context("when product price is 10000", func() {
			product := model.NewProduct(productUID, 10000)

			Context("and is black friday and is user`s birthday", func() {
				It("should get 1000 of discount in cents", func() {
					campaigns := []model.Campaign{
						newBlackFridayCampaign(true, &now),
						newBirthdayCampaign(true),
					}

					user := model.NewUser(userUID, &now)

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(1000)))
					Ω(discount.Percentage).Should(Equal(blackFridayPercent))
				})

				Context("and max discount is 15%", func() {
					It("should get 1500 of discount in cents", func() {
						maxDiscountApplied := int32(15)

						campaigns := []model.Campaign{
							newBlackFridayCampaign(true, &now),
							newBirthdayCampaign(true),
						}

						user := model.NewUser(userUID, &now)

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscountApplied, user, product)

						Ω(discount.Value).Should(Equal(int64(1500)))
						Ω(discount.Percentage).Should(Equal(maxDiscountApplied))
					})
				})
			})

			Context("and is not black friday and is user`s birthday", func() {
				It("should get 500 of discount in cents", func() {
					now := time.Now()
					tomorrow := now.Add(24 * time.Hour)

					campaigns := []model.Campaign{
						newBlackFridayCampaign(true, &tomorrow),
						newBirthdayCampaign(true),
					}

					user := model.NewUser(userUID, &now)

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(500)))
					Ω(discount.Percentage).Should(Equal(birthdayPercent))
				})

				Context("and max discount is 15%", func() {
					It("should get 500 of discount in cents", func() {
						maxDiscountApplied := int32(15)

						campaigns := []model.Campaign{
							newBlackFridayCampaign(true, &tomorrow),
							newBirthdayCampaign(true),
						}

						user := model.NewUser(userUID, &now)

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscountApplied, user, product)

						Ω(discount.Value).Should(Equal(int64(500)))
						Ω(discount.Percentage).Should(Equal(birthdayPercent))
					})
				})
			})
		})

		Context("when product price is 10000", func() {
			product := model.Product{Price: 10000}

			Context("and campaign is not implemented", func() {
				It("should get 0 of discount in cents", func() {
					user := model.NewUser(userUID, nil)

					campaigns := []model.Campaign{
						newNotImplementedCampaign(true),
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})

			Context("and black friday campaign does not have date set", func() {
				It("should get 0 of discount in cents", func() {
					user := model.NewUser(userUID, nil)

					campaigns := []model.Campaign{
						newBlackFridayCampaign(true, &tomorrow),
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})
		})
	})
})
