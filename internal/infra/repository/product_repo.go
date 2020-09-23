package repository

import (
	"blue-discount/internal/domain/model"
	"errors"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) model.ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) GetByID(id string) (model.Product, error) {
	var m model.Product
	err := r.db.Where(`id = ?`, id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRowNotFound
	}
	return m, err
}
