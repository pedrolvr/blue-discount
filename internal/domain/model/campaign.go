package model

import "time"

type Campaign struct {
	Name      string     `json:"name"`
	Percent   int32      `json:"percent"`
	Active    bool       `json:"active"`
	AppliedAt *time.Time `json:"appliedAt"`
	Order     int32      `json:"order"`
}

func NewCampaign(name string, perc int32, active bool, appliedAt *time.Time) Campaign {
	return Campaign{
		Name:      name,
		Percent:   perc,
		Active:    active,
		AppliedAt: appliedAt,
	}
}

func (Campaign) TableName() string {
	return "campaign"
}
