package model

import (
	"blue-discount/pkg/util"
)

const (
	BirthdayCampaignName    = "birthday"
	BlackFridayCampaignName = "black-friday"
)

type Discount struct {
	Value   int64
	Percent int32
}

func NewDiscount(v int64, perc int32) Discount {
	return Discount{
		Value:   v,
		Percent: perc,
	}
}

type DiscountCalculator struct {
	Campaigns []Campaign `json:"-"`
}

func NewDiscountCalculator(campaigns []Campaign) DiscountCalculator {
	return DiscountCalculator{
		Campaigns: campaigns,
	}
}

func (m *DiscountCalculator) Calculate(maxPercent int32, user User, product Product) Discount {
	var valueSum int64
	var percentSum int32
	var discount Discount

	for _, c := range m.Campaigns {
		if !c.Active {
			return discount
		}

		strategy := DiscountStrategyFactory(c)

		if strategy == nil {
			continue
		}

		valueSum = valueSum + strategy.Apply(user, product)

		if valueSum > 0 {
			percentSum = percentSum + c.Percent
		}

		if percentSum > maxPercent {
			break
		}

		discount.Value = valueSum
		discount.Percent = percentSum
	}

	return discount
}

func DiscountStrategyFactory(c Campaign) DiscountStrategy {
	if c.Name == BirthdayCampaignName {
		return &BirthdayDiscount{c}
	}

	if c.Name == BlackFridayCampaignName {
		return &OnDateDiscount{c}
	}

	return nil
}

type DiscountStrategy interface {
	Apply(User, Product) int64
}

type BirthdayDiscount struct {
	Campaign Campaign
}

func (s *BirthdayDiscount) Apply(user User, product Product) int64 {
	campaign := s.Campaign

	if user.BornAt != nil && util.IsBirthday(*user.BornAt) {
		return util.DiscountInCents(product.Price, campaign.Percent)
	}

	return 0
}

type OnDateDiscount struct {
	Campaign Campaign
}

func (s *OnDateDiscount) Apply(user User, product Product) int64 {
	campaign := s.Campaign

	if campaign.AppliedAt != nil && util.DateIsToday(*campaign.AppliedAt) {
		return util.DiscountInCents(product.Price, campaign.Percent)
	}

	return 0
}
