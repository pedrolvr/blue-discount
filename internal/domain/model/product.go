package model

import "github.com/gofrs/uuid"

type Product struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Price int64     `json:"price"`
}

func NewProduct(ID uuid.UUID, price int64) Product {
	return Product{
		ID:    ID,
		Price: price,
	}
}

func (Product) TableName() string {
	return "product"
}
