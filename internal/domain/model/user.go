package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID     uuid.UUID  `gorm:"type:uuid;primary_key;"`
	BornAt *time.Time `json:"bornAt"`
}

func (User) TableName() string {
	return "user"
}
