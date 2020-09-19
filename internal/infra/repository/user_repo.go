package repository

import (
	"blue-discount/internal/domain/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) model.UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetByID(id string) (model.User, error) {
	var m model.User
	err := r.db.Where(`id = ?`, id).First(&m).Error
	return m, err
}
