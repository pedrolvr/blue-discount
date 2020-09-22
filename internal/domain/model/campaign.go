package model

import "time"

type Campaign struct {
	Name      string     `json:"name"`
	Percent   int32      `json:"percent"`
	Active    bool       `json:"active"`
	AppliedAt *time.Time `json:"appliedAt"`
}

func (Campaign) TableName() string {
	return "campaign"
}
