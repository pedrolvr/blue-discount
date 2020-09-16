package model_test

import (
	"blue-discount/internal/domain/model"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	blackFridayPercent = int32(10)
	birthdayPercent    = int32(5)
	maxDiscount        = int32(10)
)

var _ = Describe("discount test", func() {
	Describe("DiscountStrategyFactory()", func() {
		Context("when campaign is birthday", func() {
			It("get birthday discount strategy", func() {
				campaign := model.Campaign{
					Name:    model.BirthdayCampaignName,
					Percent: 5,
				}

				strategy := model.DiscountStrategyFactory(campaign)

				Ω(strategy.(*model.BirthdayDiscount).Campaign).Should(Equal(campaign))
			})
		})

		Context("when campaign is black friday", func() {
			It("get black friday discount strategy", func() {
				campaign := model.Campaign{
					Name:    model.BlackFridayCampaignName,
					Percent: 10,
				}

				strategy := model.DiscountStrategyFactory(campaign)

				Ω(strategy.(*model.OnDateDiscount).Campaign).Should(Equal(campaign))
			})
		})

		Context("when not implemented campaign", func() {
			It("get nil", func() {
				campaign := model.Campaign{
					Name:    "not-implemented-campaign",
					Percent: 10,
				}

				strategy := model.DiscountStrategyFactory(campaign)

				Ω(strategy).Should(BeNil())
			})
		})
	})

	Describe("Calculate()", func() {
		Context("when product price is 10000", func() {
			product := model.Product{Price: 10000}

			Context("and is user`s birthday", func() {
				Context("and the campaign is not enabled", func() {
					It("should get 0 of discount in cents", func() {
						campaigns := []model.Campaign{
							{Enabled: false},
						}

						user := model.User{BornAt: time.Now()}

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscount, user, product)

						Ω(discount.Value).Should(Equal(int64(0)))
					})
				})

				It("should get 500 of discount in cents", func() {
					campaigns := []model.Campaign{
						{
							Name:    model.BirthdayCampaignName,
							Percent: birthdayPercent,
							Enabled: true,
						},
					}

					user := model.User{BornAt: time.Now()}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(500)))
				})
			})

			Context("and is not user`s birthday", func() {
				It("should get 0 of discount", func() {
					campaigns := []model.Campaign{
						{
							Name:    model.BirthdayCampaignName,
							Percent: birthdayPercent,
							Enabled: true,
						},
					}

					tomorrow := time.Now().Add(24 * time.Hour)

					user := model.User{BornAt: tomorrow}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})
		})

		Context("when product price is 10000", func() {
			product := model.Product{Price: 10000}

			Context("and is black friday", func() {
				Context("and the campaign is not enabled", func() {
					It("should get 0 of discount in cents", func() {
						campaigns := []model.Campaign{
							{Enabled: false},
						}

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscount, model.User{}, product)

						Ω(discount.Value).Should(Equal(int64(0)))
						Ω(discount.Percentage).Should(Equal(int32(0)))
					})
				})

				It("should get 1000 of discount in cents", func() {
					now := time.Now()

					campaigns := []model.Campaign{
						{
							Name:      model.BlackFridayCampaignName,
							Percent:   blackFridayPercent,
							AppliedAt: &now,
							Enabled:   true,
						},
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, model.User{}, product)

					Ω(discount.Value).Should(Equal(int64(1000)))
					Ω(discount.Percentage).Should(Equal(blackFridayPercent))
				})
			})

			Context("and is not black friday", func() {
				It("should get 1000 of discount in cents", func() {
					tomorrow := time.Now().Add(24 * time.Hour)

					campaigns := []model.Campaign{
						{
							Name:      model.BlackFridayCampaignName,
							Percent:   blackFridayPercent,
							AppliedAt: &tomorrow,
							Enabled:   true,
						},
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, model.User{}, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})
		})

		Context("when product price is 10000", func() {
			product := model.Product{Price: 10000}

			Context("and is black friday and user`s birthday", func() {
				It("should get 1000 of discount in cents", func() {
					now := time.Now()

					campaigns := []model.Campaign{
						{
							Name:      model.BlackFridayCampaignName,
							Percent:   blackFridayPercent,
							AppliedAt: &now,
							Enabled:   true,
						},
						{
							Name:    model.BirthdayCampaignName,
							Percent: birthdayPercent,
							Enabled: true,
						},
					}

					user := model.User{BornAt: time.Now()}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, user, product)

					Ω(discount.Value).Should(Equal(int64(1000)))
					Ω(discount.Percentage).Should(Equal(blackFridayPercent))
				})

				Context("and max discount is 15%", func() {
					It("should get 1500 of discount in cents", func() {
						now := time.Now()

						maxDiscountApplied := int32(15)

						campaigns := []model.Campaign{
							{
								Name:      model.BlackFridayCampaignName,
								Percent:   blackFridayPercent,
								AppliedAt: &now,
								Enabled:   true,
							},
							{
								Name:    model.BirthdayCampaignName,
								Percent: birthdayPercent,
								Enabled: true,
							},
						}

						user := model.User{BornAt: time.Now()}

						discount := model.NewDiscount(campaigns)

						discount.Calculate(maxDiscountApplied, user, product)

						Ω(discount.Value).Should(Equal(int64(1500)))
						Ω(discount.Percentage).Should(Equal(maxDiscountApplied))
					})
				})
			})
		})

		Context("when product price is 10000", func() {
			product := model.Product{Price: 10000}

			Context("and campaign is not implemented", func() {
				It("should get 0 of discount in cents", func() {
					campaigns := []model.Campaign{
						{
							Name:    "not-implemented-campaign",
							Enabled: true,
						},
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, model.User{}, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})

			Context("and black friday campaign does not have date set", func() {
				It("should get 0 of discount in cents", func() {
					campaigns := []model.Campaign{
						{
							Name:    model.BlackFridayCampaignName,
							Percent: blackFridayPercent,
							Enabled: true,
						},
					}

					discount := model.NewDiscount(campaigns)

					discount.Calculate(maxDiscount, model.User{}, product)

					Ω(discount.Value).Should(Equal(int64(0)))
					Ω(discount.Percentage).Should(Equal(int32(0)))
				})
			})
		})
	})
})
