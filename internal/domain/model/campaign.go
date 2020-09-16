package model

import "time"

type Campaign struct {
	Name      string     `json:"name"`
	Percent   int32      `json:"percent"`
	Enabled   bool       `json:"enabled"`
	AppliedAt *time.Time `appliedAt:"id"`
}
