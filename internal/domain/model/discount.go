package model

import (
	"blue-discount/internal/shared/util"
)

const (
	BirthdayCampaignName    = "birthday"
	BlackFridayCampaignName = "black-friday"
)

type Discount struct {
	Value      int64
	Percentage int32
	Campaigns  []Campaign `json:"-"`
}

func NewDiscount(campaigns []Campaign) Discount {
	return Discount{
		Campaigns: campaigns,
	}
}

func (m *Discount) Calculate(maxPercent int32, user User, product Product) {
	var value int64
	var percentApplied int32

	for _, c := range m.Campaigns {
		if !c.Enabled {
			return
		}

		strategy := DiscountStrategyFactory(c)

		if strategy == nil {
			return
		}

		percentApplied = percentApplied + c.Percent

		if percentApplied > maxPercent {
			return
		}

		value = value + strategy.Apply(user, product)

		if value > 0 {
			m.Value = value
			m.Percentage = percentApplied
		}
	}
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

	if util.IsBirthday(user.BornAt) {
		return util.DiscountInCents(product.Price, campaign.Percent)
	}

	return 0
}

type OnDateDiscount struct {
	Campaign Campaign
}

func (s *OnDateDiscount) Apply(user User, product Product) int64 {
	campaign := s.Campaign

	if campaign.AppliedAt == nil {
		return 0
	}

	if util.DateIsToday(*campaign.AppliedAt) {
		return util.DiscountInCents(product.Price, campaign.Percent)
	}

	return 0
}
