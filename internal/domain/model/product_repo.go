package model

type ProductRepo interface {
	GetByID(id string) (Product, error)
}
