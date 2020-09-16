package model

import "time"

type User struct {
	ID     string    `json:"id"`
	BornAt time.Time `json:"bornAt"`
}
