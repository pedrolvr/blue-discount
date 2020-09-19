package model

import "github.com/gofrs/uuid"

type Product struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Price int64     `json:"price"`
}

func (Product) TableName() string {
	return "product"
}
