package repository

import (
	"blue-discount/internal/domain/model"
	"errors"
	"fmt"

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

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRowNotFound
	}

	if err != nil {
		return m, fmt.Errorf("get by id %s: %w", id, err)
	}

	return m, nil
}
